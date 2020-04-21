package test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func testingRoleCreateAttributes(c *Client) *datadog.RoleCreateAttributes {
	rca := datadog.NewRoleCreateAttributes()
	rca.SetName(fmt.Sprintf("test-role-datadog-client-go-%d", c.Clock.Now().UnixNano()))
	return rca
}

func deleteRole(c *Client, roleID string) {
	_, err := c.Client.RolesApi.DeleteRole(c.Ctx, roleID).Execute()
	if err == nil {
		return
	}
	log.Printf("Error disabling Role: %v, Another test may have already deleted this role: %s", roleID, err.Error())
}

func TestRoleLifecycle(t *testing.T) {
	c := NewClientWithRecording(t)
	defer c.Close()

	// first, test creating a role
	rca := testingRoleCreateAttributes(c)
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := c.Client.RolesApi.CreateRole(c.Ctx).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(c, rid)

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

	urr, httpresp, err := c.Client.RolesApi.UpdateRole(c.Ctx, rid).Body(*rup).Execute()
	if err != nil {
		t.Fatalf("Error updating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	urData := urr.GetData()
	urAttributes := urData.GetAttributes()
	assert.Equal(t, rua.GetName(), urAttributes.GetName())

	// now, test getting it
	grr, httpresp, err := c.Client.RolesApi.GetRole(c.Ctx, rid).Execute()
	if err != nil {
		t.Fatalf("Error getting Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData = grr.GetData()
	rrAttributes = rrData.GetAttributes()
	assert.Equal(t, updatedRoleName, rrAttributes.GetName())

	// now, test filtering for it in the list call
	rsr, httpresp, err := c.Client.RolesApi.
		ListRoles(c.Ctx).
		Filter(updatedRoleName).
		PageSize(1).
		PageNumber(0).
		Sort(datadog.ROLESSORT_MODIFIED_AT_DESCENDING).
		Execute()
	if err != nil {
		t.Fatalf("Error listing roles %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.NotEmpty(t, rsr.GetData())
	rrAttributes = rsr.GetData()[0].GetAttributes()
	assert.Equal(t, rua.GetName(), rrAttributes.GetName())
	rrMetaAttributes := rsr.GetMeta()
	page := rrMetaAttributes.GetPage()
	assert.GreaterOrEqual(t, page.GetTotalCount(), int64(1))
	assert.GreaterOrEqual(t, page.GetTotalFilteredCount(), int64(1))

	// now, test deleting it
	// no response payload
	httpresp, err = c.Client.RolesApi.DeleteRole(c.Ctx, rid).Execute()
	if err != nil {
		t.Fatalf("Error deleting Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 204, httpresp.StatusCode)
}

func TestRolePermissionsLifecycle(t *testing.T) {
	c := NewClientWithRecording(t)
	defer c.Close()

	// first, create a role
	rca := testingRoleCreateAttributes(c)
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := c.Client.RolesApi.CreateRole(c.Ctx).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(c, rid)

	// find a permission
	permissions, httpresp, err := c.Client.RolesApi.ListPermissions(c.Ctx).Execute()
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

	crrtps, httpresp, err := c.Client.RolesApi.AddPermissionToRole(c.Ctx, rid).Body(*rtp).Execute()
	if err != nil {
		t.Fatalf("Error creating permission relation: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Contains(t, crrtps.GetData(), permission)

	// get all permissions for the role
	lrrtps, httpresp, err := c.Client.RolesApi.ListRolePermissions(c.Ctx, rid).Execute()
	if err != nil {
		t.Fatalf("Error listing permission relations: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Contains(t, lrrtps.GetData(), permission)

	// remove the permission from the role
	drrtps, httpresp, err := c.Client.RolesApi.RemovePermissionFromRole(c.Ctx, rid).Body(*rtp).Execute()
	if err != nil {
		t.Fatalf("Error remove permission relation: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.NotContains(t, drrtps.GetData(), permission)
}

func TestRoleUsersLifecycle(t *testing.T) {
	c := NewClientWithRecording(t)
	defer c.Close()

	// first, create a role
	rca := testingRoleCreateAttributes(c)
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := c.Client.RolesApi.CreateRole(c.Ctx).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(c, rid)

	// create a user
	uca := testingUserCreateAttributes(c)
	ucd := datadog.NewUserCreateData()
	ucd.SetAttributes(*uca)
	ucp := datadog.NewUserCreatePayload()
	ucp.SetData(*ucd)
	ur, httpresp, err := c.Client.UsersApi.CreateUser(c.Ctx).Body(*ucp).Execute()
	if err != nil {
		t.Fatalf("Error creating User %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 201)
	urData := ur.GetData()
	uid := urData.GetId()
	defer disableUser(c, uid)

	// add a user to the role
	rtu := datadog.NewRelationshipToUserWithDefaults()
	rtud := datadog.NewRelationshipToUserDataWithDefaults()
	rtud.SetId(uid)
	rtu.SetData(*rtud)

	crrtus, httpresp, err := c.Client.RolesApi.AddUserToRole(c.Ctx, rid).Body(*rtu).Execute()
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
	lrrtus, httpresp, err := c.Client.RolesApi.ListRoleUsers(c.Ctx, rid).Execute()
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
	drrtus, httpresp, err := c.Client.RolesApi.RemoveUserFromRole(c.Ctx, rid).Body(*rtu).Execute()
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

func TestListRolesErrors(t *testing.T) {
	testCases := map[string]struct {
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {FakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			c := NewClientWithRecording(t)
			defer c.Close()

			_, httpresp, err := c.Client.RolesApi.ListRoles(tc.Ctx).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestCreateRoleErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(t)
	defer c.Close()

	// first, test creating a role
	rca := testingRoleCreateAttributes(c)
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)

	// invalid role without data
	rcp400 := datadog.NewRoleCreatePayload()

	testCases := map[string]struct {
		Ctx                context.Context
		ExpectedStatusCode int
		Body               *datadog.RoleCreatePayload
	}{
		"400 Bad Request": {c.Ctx, 400, rcp400},
		"403 Forbidden":   {FakeAuth, 403, rcp},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			c := NewClientWithRecording(t)
			defer c.Close()

			rr, httpresp, err := c.Client.RolesApi.CreateRole(tc.Ctx).Body(*tc.Body).Execute()
			// make sure that we clean everything on error
			if 200 == httpresp.StatusCode {
				rrData := rr.GetData()
				rid := rrData.GetId()
				defer deleteRole(c, rid)
			}
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestGetRoleErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(t)
	defer c.Close()

	// first, test creating a role
	rca := testingRoleCreateAttributes(c)
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := c.Client.RolesApi.CreateRole(c.Ctx).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(c, rid)

	// bad role id
	rid404 := "00000000-dead-beef-dead-ffffffffffff"
	_, httpresp, err = c.Client.RolesApi.GetRole(c.Ctx, rid404).Execute()
	assert.Equal(t, 404, httpresp.StatusCode)

	testCases := map[string]struct {
		Ctx                context.Context
		ExpectedStatusCode int
		RoleID             string
	}{

		// "40X Bad Request": {c.Ctx, 404 /* FIXME invalid uuid should be 400 rather then 404 */, "this-is-an-invalid-role-id-for-get-endpoint-because-it-is-too-long-to-be-a-uuid"},
		// "403 Forbidden": {FakeAuth, 403, rid},
		"404 Not found": {c.Ctx, 404, rid404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			c := NewClientWithRecording(t)
			defer c.Close()

			_, httpresp, err := c.Client.RolesApi.GetRole(tc.Ctx, tc.RoleID).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestUpdateRoleErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(t)
	defer c.Close()

	// first, create a role
	rca := testingRoleCreateAttributes(c)
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := c.Client.RolesApi.CreateRole(c.Ctx).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(c, rid)

	// bad role id
	rid404 := "00000000-dead-beef-dead-ffffffffffff"
	_, httpresp, err = c.Client.RolesApi.GetRole(c.Ctx, rid404).Execute()
	assert.Equal(t, 404, httpresp.StatusCode)

	// working update payload
	updatedRoleName := "updated-" + rca.GetName()
	rua := datadog.NewRoleUpdateAttributesWithDefaults()
	rua.SetName(updatedRoleName)
	rud := datadog.NewRoleUpdateDataWithDefaults()
	rud.SetAttributes(*rua)
	rud.SetId(rid)
	rup := datadog.NewRoleUpdatePayloadWithDefaults()
	rup.SetData(*rud)

	// invalid role ID in the payload
	/*
		rud400 := datadog.NewRoleUpdateDataWithDefaults()
		rud400.SetAttributes(*rua)
		rud400.SetId(rid404)
		rup400 := datadog.NewRoleUpdatePayloadWithDefaults()
		rup400.SetData(*rud400)
	*/

	testCases := map[string]struct {
		Ctx                context.Context
		ExpectedStatusCode int
		RoleID             string
		Body               *datadog.RoleUpdatePayload
	}{
		"400 Bad Request":         {c.Ctx, 400, rid, datadog.NewRoleUpdatePayloadWithDefaults()},
		"403 Forbidden":           {FakeAuth, 403, rid, rup},
		"404 Bad Role ID in Path": {c.Ctx, 404, rid404, rup},
		// FIXME should be 400 "400 Bad Role ID in Request": {c.Ctx, 400, rid, rup400},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			c := NewClientWithRecording(t)
			defer c.Close()

			_, httpresp, err := c.Client.RolesApi.UpdateRole(tc.Ctx, tc.RoleID).Body(*tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestDeleteRoleErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(t)
	defer c.Close()

	// first, test creating a role
	rca := testingRoleCreateAttributes(c)
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := c.Client.RolesApi.CreateRole(c.Ctx).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(c, rid)

	// bad role id
	rid404 := "00000000-dead-beef-dead-ffffffffffff"
	_, httpresp, err = c.Client.RolesApi.GetRole(c.Ctx, rid404).Execute()
	assert.Equal(t, 404, httpresp.StatusCode)

	testCases := map[string]struct {
		Ctx                context.Context
		ExpectedStatusCode int
		RoleID             string
	}{
		"403 Forbidden": {FakeAuth, 403, rid},
		"404 Not found": {c.Ctx, 404, rid404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			c := NewClientWithRecording(t)
			defer c.Close()

			httpresp, err := c.Client.RolesApi.DeleteRole(tc.Ctx, tc.RoleID).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestListRolePermissionsErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(t)
	defer c.Close()

	// first, test creating a role
	rca := testingRoleCreateAttributes(c)
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := c.Client.RolesApi.CreateRole(c.Ctx).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(c, rid)

	testCases := map[string]struct {
		Ctx                context.Context
		ExpectedStatusCode int
		RoleID             string
	}{
		"403 Forbidden": {FakeAuth, 403, rid},
		"404 Not found": {c.Ctx, 404, "00000000-dead-beef-dead-ffffffffffff"},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			c := NewClientWithRecording(t)
			defer c.Close()

			_, httpresp, err := c.Client.RolesApi.ListRolePermissions(tc.Ctx, tc.RoleID).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestAddPermissionToRoleErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(t)
	defer c.Close()

	// first, test creating a role
	rca := testingRoleCreateAttributes(c)
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := c.Client.RolesApi.CreateRole(c.Ctx).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(c, rid)

	rtp := datadog.NewRelationshipToPermissionWithDefaults()
	rtpd := datadog.NewRelationshipToPermissionDataWithDefaults()
	rtpd.SetId("00000000-dead-beef-dead-ffffffffffff")
	rtp.SetData(*rtpd)

	testCases := map[string]struct {
		Ctx                context.Context
		ExpectedStatusCode int
		Body               *datadog.RelationshipToPermission
	}{
		"400 Bad Request": {c.Ctx, 400, datadog.NewRelationshipToPermissionWithDefaults()},
		"403 Forbidden":   {FakeAuth, 403, rtp},
		"404 Not found":   {c.Ctx, 404, rtp},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			c := NewClientWithRecording(t)
			defer c.Close()

			_, httpresp, err := c.Client.RolesApi.AddPermissionToRole(tc.Ctx, rid).Body(*tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestRemovePermissionFromRoleErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(t)
	defer c.Close()

	// first, test creating a role
	rca := testingRoleCreateAttributes(c)
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := c.Client.RolesApi.CreateRole(c.Ctx).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(c, rid)

	// find a permission
	permissions, httpresp, err := c.Client.RolesApi.ListPermissions(c.Ctx).Execute()
	if err != nil {
		t.Fatalf("Error listing permissions: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, permissions.HasData())
	assert.NotEmpty(t, permissions.GetData())

	permission := permissions.GetData()[0]
	pid := permission.GetId()

	rtp := datadog.NewRelationshipToPermissionWithDefaults()
	rtpd := datadog.NewRelationshipToPermissionDataWithDefaults()
	rtpd.SetId(pid)
	rtp.SetData(*rtpd)

	testCases := map[string]struct {
		Ctx                context.Context
		ExpectedStatusCode int
		RoleID             string
		Body               *datadog.RelationshipToPermission
	}{
		"400 Bad Request": {c.Ctx, 400, rid, datadog.NewRelationshipToPermissionWithDefaults()},
		"403 Forbidden":   {FakeAuth, 403, rid, rtp},
		// TODO this should be probably 404
		"40X Bad role": {c.Ctx, 400, "00000000-dead-beef-dead-ffffffffffff", rtp},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			c := NewClientWithRecording(t)
			defer c.Close()

			_, httpresp, err := c.Client.RolesApi.RemovePermissionFromRole(tc.Ctx, tc.RoleID).Body(*tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestListRoleUsersErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(t)
	defer c.Close()

	// first, create a role
	rca := testingRoleCreateAttributes(c)
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := c.Client.RolesApi.CreateRole(c.Ctx).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(c, rid)

	testCases := map[string]struct {
		Ctx                context.Context
		ExpectedStatusCode int
		RoleID             string
	}{
		"403 Forbidden": {FakeAuth, 403, rid},
		// !FIXME "404 Not found": {c.Ctx, 404, "10000000-dead-beef-dead-ffffffffffff"},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			c := NewClientWithRecording(t)
			defer c.Close()

			_, httpresp, err := c.Client.RolesApi.ListRoleUsers(tc.Ctx, tc.RoleID).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestAddUserToRoleErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(t)
	defer c.Close()

	// first, test creating a role
	rca := testingRoleCreateAttributes(c)
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := c.Client.RolesApi.CreateRole(c.Ctx).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(c, rid)

	rtu := datadog.NewRelationshipToUserWithDefaults()
	rtud := datadog.NewRelationshipToUserDataWithDefaults()
	rtud.SetId("00000000-dead-beef-dead-ffffffffffff")
	rtu.SetData(*rtud)

	testCases := map[string]struct {
		Ctx                context.Context
		ExpectedStatusCode int
		Body               *datadog.RelationshipToUser
	}{
		"400 Bad Request": {c.Ctx, 400, datadog.NewRelationshipToUserWithDefaults()},
		"403 Forbidden":   {FakeAuth, 403, rtu},
		"404 Not found":   {c.Ctx, 404, rtu},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			c := NewClientWithRecording(t)
			defer c.Close()

			_, httpresp, err := c.Client.RolesApi.AddUserToRole(tc.Ctx, rid).Body(*tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestRemoveUserFromRoleErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(t)
	defer c.Close()

	// first, test creating a role
	rca := testingRoleCreateAttributes(c)
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := c.Client.RolesApi.CreateRole(c.Ctx).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(c, rid)

	// bad role id
	rid404 := "00000000-dead-beef-dead-ffffffffffff"
	_, httpresp, err = c.Client.RolesApi.GetRole(c.Ctx, rid404).Execute()
	assert.Equal(t, 404, httpresp.StatusCode)

	// create a user
	uca := testingUserCreateAttributes(c)
	ucd := datadog.NewUserCreateData()
	ucd.SetAttributes(*uca)
	ucp := datadog.NewUserCreatePayload()
	ucp.SetData(*ucd)
	ur, httpresp, err := c.Client.UsersApi.CreateUser(c.Ctx).Body(*ucp).Execute()
	if err != nil {
		t.Fatalf("Error creating User %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 201)
	urData := ur.GetData()
	uid := urData.GetId()
	defer disableUser(c, uid)

	rtu := datadog.NewRelationshipToUserWithDefaults()
	rtud := datadog.NewRelationshipToUserDataWithDefaults()
	rtud.SetId(uid)
	rtu.SetData(*rtud)

	testCases := map[string]struct {
		Ctx                context.Context
		ExpectedStatusCode int
		RoleID             string
		Body               *datadog.RelationshipToUser
	}{
		"400 Bad Request": {c.Ctx, 400, rid, datadog.NewRelationshipToUserWithDefaults()},
		"403 Forbidden":   {FakeAuth, 403, rid, rtu},
		"404 Bad role":    {c.Ctx, 404, rid404, rtu},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			c := NewClientWithRecording(t)
			defer c.Close()

			_, httpresp, err := c.Client.RolesApi.RemoveUserFromRole(tc.Ctx, tc.RoleID).Body(*tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}
