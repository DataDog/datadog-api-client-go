package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func testingRoleCreateAttributes() *datadog.RoleCreateAttributes {
	rca := datadog.NewRoleCreateAttributes()
	rca.SetName(fmt.Sprintf("test-role-datadog-client-go-%d", TestClock.Now().UnixNano()))
	return rca
}

func deleteRole(roleID string) {
	httpresp, err := TestAPIClient.RolesApi.DeleteRole(TestAuth, roleID).Execute()
	if httpresp.StatusCode == 404 {
		// doesn't exist any more => no need to delete
		return
	}
	if httpresp.StatusCode != 204 || err != nil {
		log.Printf("Error disabling Role: %v, Another test may have already deleted this role: %s", roleID, err.Error())
	}
}

func TestRoleLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// first, test creating a role
	rca := testingRoleCreateAttributes()
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := TestAPIClient.RolesApi.CreateRole(TestAuth).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(rid)

	rrAttributes := rrData.GetAttributes()
	assert.Equal(t, rca.GetName(), rrAttributes.GetName())

	// now, test updating it
	updatedRoleName := "updated-" + rca.GetName()
	rua := datadog.NewRoleUpdateAttributes()
	rua.SetName(updatedRoleName)
	rud := datadog.NewRoleUpdateData()
	rud.SetAttributes(*rua)
	rud.SetId(rid)
	rup := datadog.NewRoleUpdatePayload()
	rup.SetData(*rud)

	urp, httpresp, err := TestAPIClient.RolesApi.UpdateRole(TestAuth, rid).Body(*rup).Execute()
	if err != nil {
		t.Fatalf("Error updating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	urData := urp.GetData()
	urAttributes := urData.GetAttributes()
	assert.Equal(t, rua.GetName(), urAttributes.GetName())

	// now, test getting it
	rrp, httpresp, err := TestAPIClient.RolesApi.GetRole(TestAuth, rid).Execute()
	if err != nil {
		t.Fatalf("Error getting Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData = rrp.GetData()
	rrAttributes = rrData.GetAttributes()
	assert.Equal(t, updatedRoleName, rrAttributes.GetName())

	// now, test filtering for it in the list call
	rsrp, httpresp, err := TestAPIClient.RolesApi.
		ListRoles(TestAuth).
		Filter(updatedRoleName).
		PageSize(1).
		PageNumber(0).
		Sort(datadog.ROLESSORT_MODIFIED_AT_DESCENDING).
		Execute()
	if err != nil {
		t.Fatalf("Error listing roles %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.NotEmpty(t, rsrp.GetData())
	rrAttributes = rsrp.GetData()[0].GetAttributes()
	assert.Equal(t, rua.GetName(), rrAttributes.GetName())
	rrMetaAttributes := rsrp.GetMeta()
	page := rrMetaAttributes.GetPage()
	assert.GreaterOrEqual(t, page.GetTotalCount(), int64(1))
	assert.GreaterOrEqual(t, page.GetTotalFilteredCount(), int64(1))

	// now, test deleting it
	// no response payload
	httpresp, err = TestAPIClient.RolesApi.DeleteRole(TestAuth, rid).Execute()
	if err != nil {
		t.Fatalf("Error deleting Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 204, httpresp.StatusCode)
}

func TestRolePermissionsLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// first, create a role
	rca := testingRoleCreateAttributes()
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := TestAPIClient.RolesApi.CreateRole(TestAuth).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(rid)

	// find a permission
	permissions, httpresp, err := TestAPIClient.RolesApi.ListPermissions(TestAuth).Execute()
	if err != nil {
		t.Fatalf("Error listing permissions: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, permissions.HasData())
	assert.NotEmpty(t, permissions.GetData())

	permission := permissions.GetData()[0]
	pid := permission.GetId()

	// add a permission to the role
	rtp := datadog.NewRelationshipToPermissionWithDefaults()
	rtpd := datadog.NewRelationshipToPermissionDataWithDefaults()
	rtpd.SetId(pid)
	rtp.SetData(*rtpd)

	crrtps, httpresp, err := TestAPIClient.RolesApi.AddPermissionToRole(TestAuth, rid).Body(*rtp).Execute()
	if err != nil {
		t.Fatalf("Error creating permission relation: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Contains(t, crrtps.GetData(), permission)

	// get all permissions for the role
	lrrtps, httpresp, err := TestAPIClient.RolesApi.ListRolePermissions(TestAuth, rid).Execute()
	if err != nil {
		t.Fatalf("Error listing permission relations: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Contains(t, lrrtps.GetData(), permission)

	// remove the permission from the role
	drrtps, httpresp, err := TestAPIClient.RolesApi.RemovePermissionFromRole(TestAuth, rid).Body(*rtp).Execute()
	if err != nil {
		t.Fatalf("Error remove permission relation: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.NotContains(t, drrtps.GetData(), permission)
}

func TestRoleUsersLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// first, create a role
	rca := testingRoleCreateAttributes()
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := TestAPIClient.RolesApi.CreateRole(TestAuth).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(rid)

	// create a user
	uca := testingUserCreateAttributes()
	ucd := datadog.NewUserCreateData()
	ucd.SetAttributes(*uca)
	ucp := datadog.NewUserCreatePayload()
	ucp.SetData(*ucd)
	ur, httpresp, err := TestAPIClient.UsersApi.CreateUser(TestAuth).Body(*ucp).Execute()
	if err != nil {
		t.Fatalf("Error creating User %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 201)
	urData := ur.GetData()
	uid := urData.GetId()
	defer disableUser(uid)

	// add a user to the role
	rtu := datadog.NewRelationshipToUserWithDefaults()
	rtud := datadog.NewRelationshipToUserDataWithDefaults()
	rtud.SetId(uid)
	rtu.SetData(*rtud)

	crrtus, httpresp, err := TestAPIClient.RolesApi.AddUserToRole(TestAuth, rid).Body(*rtu).Execute()
	if err != nil {
		t.Fatalf("Error creating user relation: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)

	found := false
	for _, ua := range crrtus.GetData() {
		if uid == ua.GetId() {
			found = true
			break
		}
	}
	assert.True(t, found, "user %s not found in relation to role %s", uid, rid)

	// get all users for the role
	lrrtus, httpresp, err := TestAPIClient.RolesApi.ListRoleUsers(TestAuth, rid).Execute()
	if err != nil {
		t.Fatalf("Error listing permission relations: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)

	found = false
	for _, ua := range lrrtus.GetData() {
		if uid == ua.GetId() {
			found = true
			break
		}
	}
	assert.True(t, found, "user %s not found in relation to role %s", uid, rid)

	// remove the permission from the role
	drrtus, httpresp, err := TestAPIClient.RolesApi.RemoveUserFromRole(TestAuth, rid).Body(*rtu).Execute()
	if err != nil {
		t.Fatalf("Error remove permission relation: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	for _, ua := range drrtus.GetData() {
		if uid == ua.GetId() {
			t.Fatalf("User %s must not exist in the relation %s", uid, rid)
		}
	}
}
