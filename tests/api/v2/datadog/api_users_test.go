package test

import (
	"fmt"
	"log"
	"testing"

	"gotest.tools/assert"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func testingUserCreateAttributes(c Client) *datadog.UserCreateAttributes {
	uca := datadog.NewUserCreateAttributes()
	uca.SetEmail(fmt.Sprintf("test-datadog-client-go-%d@datadoghq.com", c.Clock.Now().UnixNano()))
	uca.SetName("Test Datadog Client Go")
	uca.SetTitle("Big boss")
	return uca
}

func disableUser(c Client, userID string) {
	_, err := c.Client.UsersApi.DisableUser(c.Ctx, userID).Execute()
	if err == nil {
		return
	}
	log.Printf("Error disabling User: %v, Another test may have already disabled this user: %s", userID, err.Error())
}

func TestUserLifecycle(t *testing.T) {
	c := NewClientWithRecording(t)
	defer c.Close()

	// first, test creating a user
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
	httpresp, err = c.Client.UsersApi.UpdateUser(c.Ctx, uid).Body(*uup).Execute()
	if err != nil {
		t.Fatalf("Error updating User %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 204)

	// now, test getting it
	urp, httpresp, err := c.Client.UsersApi.GetUser(c.Ctx, uid).Execute()
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
	httpresp, err = c.Client.UsersApi.DisableUser(c.Ctx, uid).Execute()
	if err != nil {
		t.Fatalf("Error disabling User %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 204)

	// now, test filtering for it in the list call
	usrp, httpresp, err := c.Client.UsersApi.
		ListUsers(c.Ctx).
		Filter(uca.GetEmail()).
		PageSize(1).
		PageNumber(0).
		SortDir(datadog.QUERYSORTORDER_ASC).
		Execute()
	if err != nil {
		t.Fatalf("Error listing users %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, len(usrp.GetData()), 1)
	urAttributes = usrp.GetData()[0].GetAttributes()
	assert.Equal(t, urAttributes.GetEmail(), uca.GetEmail())
	urMetaAttributes := usrp.GetMeta()
	page := urMetaAttributes.GetPage()
	assert.Assert(t, page.GetTotalCount() >= 1)
	assert.Assert(t, page.GetTotalFilteredCount() >= 1)

	// NOTE: to test getting a user organization, we'd need to have a "whoami" API endpoint
	// to get the UUID of the current user, but there's no such stable endpoint right now
	// (a user can only get organization for itself, never for a different user)
}

func TestUserInvitation(t *testing.T) {
	c := NewClientWithRecording(t)
	defer c.Close()

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
	id := urData.GetId()
	defer disableUser(c, id)

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

	resp, httpresp, err := c.Client.UsersApi.SendInvitations(c.Ctx).Body(*uip).Execute()
	if err != nil {
		t.Fatalf("Error sending invitation for %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 201)
	respID := resp.GetData()[0].GetId()

	// now, test getting the invitation
	oneresp, httpresp, err := c.Client.UsersApi.GetInvitation(c.Ctx, respID).Execute()
	if err != nil {
		t.Fatalf("Error getting invitation %s: Response %s: %v", respID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	data := oneresp.GetData()
	assert.Equal(t, data.GetId(), respID)
}
