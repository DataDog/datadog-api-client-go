package test

import (
	"fmt"
	"log"
	"testing"

	"gotest.tools/assert"

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
		// doesn't exist any more => no need to disable
		return
	}
	if httpresp.StatusCode != 204 || err != nil {
		log.Printf("Error disabling Role: %v, Another test may have already disabled this user: %s", roleID, err.Error())
	}
}

func TestRoleLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// first, test creating a user
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
	assert.Equal(t, 1, len(rsrp.GetData()))
	rrAttributes = rsrp.GetData()[0].GetAttributes()
	assert.Equal(t, rua.GetName(), rrAttributes.GetName())
	rrMetaAttributes := rsrp.GetMeta()
	page := rrMetaAttributes.GetPage()
	assert.Assert(t, page.GetTotalCount() >= 1)
	assert.Assert(t, page.GetTotalFilteredCount() >= 1)

	// now, test deleting it
	// no response payload
	httpresp, err = TestAPIClient.RolesApi.DeleteRole(TestAuth, rid).Execute()
	if err != nil {
		t.Fatalf("Error deleting Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 204, httpresp.StatusCode)
}

/*
func xxestRoleInvitation(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	rca := testingRoleCreateAttributes()
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	ucp := datadog.NewRoleCreatePayload()
	ucp.SetData(*rcd)
	ur, httpresp, err := TestAPIClient.RolesApi.CreateRole(TestAuth).Body(*ucp).Execute()
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 201)
	rrData := rr.GetData()
	id := rrData.GetId()
	defer disableRole(id)

	// first, create the user invitation
	rtud := datadog.NewRelationshipToRoleData()
	rtud.SetId(id)
	rtu := datadog.NewRelationshipToRole()
	rtu.SetData(*rtud)
	uir := datadog.NewRoleInvitationRelationships()
	uir.SetRole(*rtu)
	rid := datadog.NewRoleInvitationData()
	rid.SetRelationships(*uir)
	uip := datadog.NewRoleInvitationPayload()
	uip.SetData([]datadog.RoleInvitationData{*rid})

	resp, httpresp, err := TestAPIClient.RolesApi.SendInvitations(TestAuth).Body(*uip).Execute()
	if err != nil {
		t.Fatalf("Error sending invitation for %s: Response %s: %v", rca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 201)
	respID := resp.GetData()[0].GetId()

	// now, test getting the invitation
	oneresp, httpresp, err := TestAPIClient.RolesApi.GetInvitation(TestAuth, respID).Execute()
	if err != nil {
		t.Fatalf("Error getting invitation %s: Response %s: %v", respID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	data := oneresp.GetData()
	assert.Equal(t, data.GetId(), respID)
}
*/
