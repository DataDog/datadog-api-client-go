package test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
)

func testingRoleCreateAttributes(ctx context.Context) *datadog.RoleCreateAttributes {
	rca := datadog.NewRoleCreateAttributes()
	rca.SetName(fmt.Sprintf("test-role-datadog-client-go-%d", tests.ClockFromContext(ctx).Now().UnixNano()))
	return rca
}

func deleteRole(ctx context.Context, roleID string) {
	_, err := Client(ctx).RolesApi.DeleteRole(ctx, roleID).Execute()
	if err == nil {
		return
	}
	log.Printf("Error disabling Role: %v, Another test may have already deleted this role: %s", roleID, err.Error())
}

func TestRoleLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()

	// first, test creating a role
	rca := testingRoleCreateAttributes(ctx)
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := Client(ctx).RolesApi.CreateRole(ctx).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(ctx, rid)

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

	urr, httpresp, err := Client(ctx).RolesApi.UpdateRole(ctx, rid).Body(*rup).Execute()
	if err != nil {
		t.Fatalf("Error updating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	urData := urr.GetData()
	urAttributes := urData.GetAttributes()
	assert.Equal(t, rua.GetName(), urAttributes.GetName())

	// now, test getting it
	grr, httpresp, err := Client(ctx).RolesApi.GetRole(ctx, rid).Execute()
	if err != nil {
		t.Fatalf("Error getting Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData = grr.GetData()
	rrAttributes = rrData.GetAttributes()
	assert.Equal(t, updatedRoleName, rrAttributes.GetName())

	// now, test filtering for it in the list call
	rsr, httpresp, err := Client(ctx).RolesApi.
		ListRoles(ctx).
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
	httpresp, err = Client(ctx).RolesApi.DeleteRole(ctx, rid).Execute()
	if err != nil {
		t.Fatalf("Error deleting Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 204, httpresp.StatusCode)
}

func TestRolePermissionsLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()

	// first, create a role
	rca := testingRoleCreateAttributes(ctx)
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := Client(ctx).RolesApi.CreateRole(ctx).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(ctx, rid)

	// find a permission
	permissions, httpresp, err := Client(ctx).RolesApi.ListPermissions(ctx).Execute()
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

	crrtps, httpresp, err := Client(ctx).RolesApi.AddPermissionToRole(ctx, rid).Body(*rtp).Execute()
	if err != nil {
		t.Fatalf("Error creating permission relation: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Contains(t, crrtps.GetData(), permission)

	// get all permissions for the role
	lrrtps, httpresp, err := Client(ctx).RolesApi.ListRolePermissions(ctx, rid).Execute()
	if err != nil {
		t.Fatalf("Error listing permission relations: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Contains(t, lrrtps.GetData(), permission)

	// remove the permission from the role
	drrtps, httpresp, err := Client(ctx).RolesApi.RemovePermissionFromRole(ctx, rid).Body(*rtp).Execute()
	if err != nil {
		t.Fatalf("Error remove permission relation: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.NotContains(t, drrtps.GetData(), permission)
}

func TestRoleUsersLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()

	// first, create a role
	rca := testingRoleCreateAttributes(ctx)
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := Client(ctx).RolesApi.CreateRole(ctx).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(ctx, rid)

	// create a user
	uca := testingUserCreateAttributes(ctx)
	ucd := datadog.NewUserCreateData()
	ucd.SetAttributes(*uca)
	ucp := datadog.NewUserCreatePayload()
	ucp.SetData(*ucd)
	ur, httpresp, err := Client(ctx).UsersApi.CreateUser(ctx).Body(*ucp).Execute()
	if err != nil {
		t.Fatalf("Error creating User %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 201)
	urData := ur.GetData()
	uid := urData.GetId()
	defer disableUser(ctx, uid)

	// add a user to the role
	rtu := datadog.NewRelationshipToUserWithDefaults()
	rtud := datadog.NewRelationshipToUserDataWithDefaults()
	rtud.SetId(uid)
	rtu.SetData(*rtud)

	crrtus, httpresp, err := Client(ctx).RolesApi.AddUserToRole(ctx, rid).Body(*rtu).Execute()
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
	lrrtus, httpresp, err := Client(ctx).RolesApi.ListRoleUsers(ctx, rid).Execute()
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
	drrtus, httpresp, err := Client(ctx).RolesApi.RemoveUserFromRole(ctx, rid).Body(*rtu).Execute()
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
	ctx, finish := WithRecorder(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			_, httpresp, err := Client(ctx).RolesApi.ListRoles(ctx).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestCreateRoleErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()

	// first, test creating a role
	rca := testingRoleCreateAttributes(ctx)
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)

	// invalid role without data
	rcp400 := datadog.NewRoleCreatePayload()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
		Body               *datadog.RoleCreatePayload
	}{
		"400 Bad Request": {WithTestAuth, 400, rcp400},
		"403 Forbidden":   {WithFakeAuth, 403, rcp},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			rr, httpresp, err := Client(ctx).RolesApi.CreateRole(ctx).Body(*tc.Body).Execute()
			// make sure that we clean everything on error
			if 200 == httpresp.StatusCode {
				rrData := rr.GetData()
				rid := rrData.GetId()
				defer deleteRole(ctx, rid)
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
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()

	// valid role ID
	rca := testingRoleCreateAttributes(ctx)
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := Client(ctx).RolesApi.CreateRole(ctx).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(ctx, rid)

	// bad role ID
	rid404 := "00000000-dead-beef-dead-ffffffffffff"
	_, httpresp, err = Client(ctx).RolesApi.GetRole(ctx, rid404).Execute()
	assert.Equal(t, 404, httpresp.StatusCode)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
		RoleID             string
	}{
		"403 Forbidden": {WithFakeAuth, 403, rid},
		"404 Not found": {WithTestAuth, 404, rid404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			_, httpresp, err := Client(ctx).RolesApi.GetRole(ctx, tc.RoleID).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestUpdateRoleErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()

	// valid role ID
	rca := testingRoleCreateAttributes(ctx)
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := Client(ctx).RolesApi.CreateRole(ctx).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(ctx, rid)

	// bad role ID
	rid404 := "00000000-dead-beef-dead-ffffffffffff"
	_, httpresp, err = Client(ctx).RolesApi.GetRole(ctx, rid404).Execute()
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
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
		RoleID             string
		Body               *datadog.RoleUpdatePayload
	}{
		"400 Bad Request":         {WithTestAuth, 400, rid, datadog.NewRoleUpdatePayloadWithDefaults()},
		"403 Forbidden":           {WithFakeAuth, 403, rid, rup},
		"404 Bad Role ID in Path": {WithTestAuth, 404, rid404, rup},
		// FIXME AAA-1540: should be 400 "400 Bad Role ID in Request": {WithTestAuth, 400, rid, rup400},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			_, httpresp, err := Client(ctx).RolesApi.UpdateRole(ctx, tc.RoleID).Body(*tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestDeleteRoleErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()

	// valid role ID
	rca := testingRoleCreateAttributes(ctx)
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := Client(ctx).RolesApi.CreateRole(ctx).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(ctx, rid)

	// bad role ID
	rid404 := "00000000-dead-beef-dead-ffffffffffff"
	_, httpresp, err = Client(ctx).RolesApi.GetRole(ctx, rid404).Execute()
	assert.Equal(t, 404, httpresp.StatusCode)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
		RoleID             string
	}{
		"403 Forbidden": {WithFakeAuth, 403, rid},
		"404 Not found": {WithTestAuth, 404, rid404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			httpresp, err := Client(ctx).RolesApi.DeleteRole(ctx, tc.RoleID).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestListRolePermissionsErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()

	// valid role ID
	rca := testingRoleCreateAttributes(ctx)
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := Client(ctx).RolesApi.CreateRole(ctx).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(ctx, rid)

	// bad role ID
	rid404 := "00000000-dead-beef-dead-ffffffffffff"
	_, httpresp, err = Client(ctx).RolesApi.GetRole(ctx, rid404).Execute()
	assert.Equal(t, 404, httpresp.StatusCode)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
		RoleID             string
	}{
		"403 Forbidden": {WithFakeAuth, 403, rid},
		"404 Not found": {WithTestAuth, 404, rid404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			_, httpresp, err := Client(ctx).RolesApi.ListRolePermissions(ctx, tc.RoleID).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestAddPermissionToRoleErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()

	// valid role ID
	rca := testingRoleCreateAttributes(ctx)
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := Client(ctx).RolesApi.CreateRole(ctx).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(ctx, rid)

	// bad role ID
	rid404 := "00000000-dead-beef-dead-ffffffffffff"
	_, httpresp, err = Client(ctx).RolesApi.GetRole(ctx, rid404).Execute()
	assert.Equal(t, 404, httpresp.StatusCode)

	rtp := datadog.NewRelationshipToPermissionWithDefaults()
	rtpd := datadog.NewRelationshipToPermissionDataWithDefaults()
	rtpd.SetId(rid404)
	rtp.SetData(*rtpd)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
		Body               *datadog.RelationshipToPermission
	}{
		"400 Bad Request": {WithTestAuth, 400, datadog.NewRelationshipToPermissionWithDefaults()},
		"403 Forbidden":   {WithFakeAuth, 403, rtp},
		"404 Not found":   {WithTestAuth, 404, rtp},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			_, httpresp, err := Client(ctx).RolesApi.AddPermissionToRole(ctx, rid).Body(*tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestRemovePermissionFromRoleErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()

	// valid role ID
	rca := testingRoleCreateAttributes(ctx)
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := Client(ctx).RolesApi.CreateRole(ctx).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(ctx, rid)

	// bad role ID
	rid404 := "00000000-dead-beef-dead-ffffffffffff"
	_, httpresp, err = Client(ctx).RolesApi.GetRole(ctx, rid404).Execute()
	assert.Equal(t, 404, httpresp.StatusCode)

	// find a permission
	permissions, httpresp, err := Client(ctx).RolesApi.ListPermissions(ctx).Execute()
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

	// invalid permission data
	rtp400 := datadog.NewRelationshipToPermissionWithDefaults()
	rtpd400 := datadog.NewRelationshipToPermissionDataWithDefaults()
	rtpd400.SetId("11111111-dead-beef-dead-ffffffffffff")
	rtp400.SetData(*rtpd400)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
		RoleID             string
		Body               *datadog.RelationshipToPermission
	}{
		"400 Bad Request":        {WithTestAuth, 400, rid, datadog.NewRelationshipToPermissionWithDefaults()},
		"400 Invalid Permission": {WithTestAuth, 400, rid, rtp400},
		"403 Forbidden":          {WithFakeAuth, 403, rid, rtp},
		"404 Role Not Found":     {WithTestAuth, 404, rid404, rtp},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			_, httpresp, err := Client(ctx).RolesApi.RemovePermissionFromRole(ctx, tc.RoleID).Body(*tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestListRoleUsersErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()

	// valid role ID
	rca := testingRoleCreateAttributes(ctx)
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := Client(ctx).RolesApi.CreateRole(ctx).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(ctx, rid)

	// bad role ID
	rid404 := "00000000-dead-beef-dead-ffffffffffff"
	_, httpresp, err = Client(ctx).RolesApi.GetRole(ctx, rid404).Execute()
	assert.Equal(t, 404, httpresp.StatusCode)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
		RoleID             string
	}{
		"403 Forbidden": {WithFakeAuth, 403, rid},
		"404 Not found": {WithTestAuth, 404, rid404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			_, httpresp, err := Client(ctx).RolesApi.ListRoleUsers(ctx, tc.RoleID).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestAddUserToRoleErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()

	// valid role ID
	rca := testingRoleCreateAttributes(ctx)
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := Client(ctx).RolesApi.CreateRole(ctx).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(ctx, rid)

	// bad role ID
	rid404 := "00000000-dead-beef-dead-ffffffffffff"
	_, httpresp, err = Client(ctx).RolesApi.GetRole(ctx, rid404).Execute()
	assert.Equal(t, 404, httpresp.StatusCode)

	rtu := datadog.NewRelationshipToUserWithDefaults()
	rtud := datadog.NewRelationshipToUserDataWithDefaults()
	rtud.SetId(rid404)
	rtu.SetData(*rtud)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
		Body               *datadog.RelationshipToUser
	}{
		"400 Bad Request": {WithTestAuth, 400, datadog.NewRelationshipToUserWithDefaults()},
		"403 Forbidden":   {WithFakeAuth, 403, rtu},
		"404 Not found":   {WithTestAuth, 404, rtu},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			_, httpresp, err := Client(ctx).RolesApi.AddUserToRole(ctx, rid).Body(*tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestRemoveUserFromRoleErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()

	// valid role ID
	rca := testingRoleCreateAttributes(ctx)
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreatePayload()
	rcp.SetData(*rcd)
	rr, httpresp, err := Client(ctx).RolesApi.CreateRole(ctx).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(ctx, rid)

	// bad role ID
	rid404 := "00000000-dead-beef-dead-ffffffffffff"
	_, httpresp, err = Client(ctx).RolesApi.GetRole(ctx, rid404).Execute()
	assert.Equal(t, 404, httpresp.StatusCode)

	// create a user
	uca := testingUserCreateAttributes(ctx)
	ucd := datadog.NewUserCreateData()
	ucd.SetAttributes(*uca)
	ucp := datadog.NewUserCreatePayload()
	ucp.SetData(*ucd)
	ur, httpresp, err := Client(ctx).UsersApi.CreateUser(ctx).Body(*ucp).Execute()
	if err != nil {
		t.Fatalf("Error creating User %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 201)
	urData := ur.GetData()
	uid := urData.GetId()
	defer disableUser(ctx, uid)

	rtu := datadog.NewRelationshipToUserWithDefaults()
	rtud := datadog.NewRelationshipToUserDataWithDefaults()
	rtud.SetId(uid)
	rtu.SetData(*rtud)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
		RoleID             string
		Body               *datadog.RelationshipToUser
	}{
		"400 Bad Request": {WithTestAuth, 400, rid, datadog.NewRelationshipToUserWithDefaults()},
		"403 Forbidden":   {WithFakeAuth, 403, rid, rtu},
		"404 Bad role":    {WithTestAuth, 404, rid404, rtu},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			_, httpresp, err := Client(ctx).RolesApi.RemoveUserFromRole(ctx, tc.RoleID).Body(*tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}
