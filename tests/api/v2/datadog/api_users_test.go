package test

import (
	"fmt"
	"log"
	"testing"

	"gotest.tools/assert"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func testingUserCreateAttributes() *datadog.UserCreateAttributes {
	uca := datadog.NewUserCreateAttributes()
	uca.SetEmail(fmt.Sprintf("testinguser-%d@datadoghq.com", TestClock.Now().UnixNano()))
	uca.SetName("Testign User")
	uca.SetTitle("Big boss")
	return uca
}

func disableUser(userID string) {
	httpresp, err := TestAPIClient.UsersApi.DisableUser(TestAuth, userID).Execute()
	if httpresp.StatusCode == 404 {
		// doesn't exist any more => no need to disable
		return
	}
	if httpresp.StatusCode != 204 || err != nil {
		log.Printf("Error disabling User: %v, Another test may have already disabled this user: %s", userID, err.Error())
	}
}

func TestUserLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// first, test creating a user
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

	urAttributes := urData.GetAttributes()
	assert.Equal(t, urAttributes.GetEmail(), uca.GetEmail())
	assert.Equal(t, urAttributes.GetName(), uca.GetName())
	assert.Equal(t, urAttributes.GetTitle(), uca.GetTitle())

	// now, test updating it
	uua := datadog.NewUserUpdateAttributes()
	uua.SetDisabled(false)
	uua.SetName("Joe Doe")
	uud := datadog.NewUserUpdateData()
	uud.SetAttributes(*uua)
	uud.SetId(uid)
	uup := datadog.NewUserUpdatePayload()
	uup.SetData(*uud)
	// no response payload
	httpresp, err = TestAPIClient.UsersApi.UpdateUser(TestAuth, uid).Body(*uup).Execute()
	if err != nil {
		t.Fatalf("Error updating User %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 204)

	// now, test getting it
	urp, httpresp, err := TestAPIClient.UsersApi.GetUser(TestAuth, uid).Execute()
	if err != nil {
		t.Fatalf("Error getting User %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	urData = urp.GetData()
	urAttributes = urData.GetAttributes()
	assert.Equal(t, urAttributes.GetEmail(), uca.GetEmail())
	assert.Equal(t, urAttributes.GetName(), "Joe Doe")
	assert.Equal(t, urAttributes.GetTitle(), uca.GetTitle())
	assert.Equal(t, urAttributes.GetDisabled(), false)

	// now, test disabling it
	// no response payload
	httpresp, err = TestAPIClient.UsersApi.DisableUser(TestAuth, uid).Execute()
	if err != nil {
		t.Fatalf("Error disabling User %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 204)

	// now, test filtering for it in the list call
	usrp, httpresp, err := TestAPIClient.UsersApi.ListUsers(TestAuth).Filter(uca.GetEmail()).PageSize(1).PageNumber(0).Execute()
	if err != nil {
		t.Fatalf("Error listing users %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, len(usrp.GetData()), 1)
	urAttributes = usrp.GetData()[0].GetAttributes()
	assert.Equal(t, urAttributes.GetEmail(), uca.GetEmail())

	// now, test getting a user organization
	// TODO: it seems that this can only be called for the current user, but how do we get current user UUID?
}

func TestUserInvitation(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// TODO: setting login_method
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
	id := urData.GetId()
	defer disableUser(id)

	// first, create the user invitation
	rtud := datadog.NewRelationshipToUserData()
	rtud.SetId(id)
	rtu := datadog.NewRelationshipToUser()
	rtu.SetData(*rtud)
	uir := datadog.NewUserInvitationRelationships()
	uir.SetUser(*rtu)
	uid := datadog.NewUserInvitationData()
	uid.SetRelationships(*uir)
	uip := datadog.NewUserInvitationPayload()
	uip.SetData([]datadog.UserInvitationData{*uid})

	resp, httpresp, err := TestAPIClient.UsersApi.SendInvitations(TestAuth).Body(*uip).Execute()
	if err != nil {
		t.Fatalf("Error sending invitation for %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 201)
	respId := resp.GetData()[0].GetId()

	// now, test getting the invitation
	oneresp, httpresp, err := TestAPIClient.UsersApi.GetInvitation(TestAuth, respId).Execute()
	if err != nil {
		t.Fatalf("Error getting invitation %s: Response %s: %v", respId, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	data := oneresp.GetData()
	assert.Equal(t, data.GetId(), respId)
}