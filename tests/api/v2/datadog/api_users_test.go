package test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
)

func testingUserCreateAttributes(ctx context.Context, t *testing.T) *datadog.UserCreateAttributes {
	uca := datadog.NewUserCreateAttributes()
	name := *tests.UniqueEntityName(ctx, t)
	uca.SetEmail(fmt.Sprintf("%s@datadoghq.com", name))
	uca.SetName(name)
	uca.SetTitle("Big boss")
	return uca
}

func disableUser(ctx context.Context, userID string) {
	_, err := Client(ctx).UsersApi.DisableUser(ctx, userID).Execute()
	if err == nil {
		return
	}
	log.Printf("Error disabling User: %v, Another test may have already disabled this user: %s", userID, err.Error())
}

func TestUserLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// first, test creating a user
	uca := testingUserCreateAttributes(ctx, t)
	ucd := datadog.NewUserCreateData()
	ucd.SetAttributes(*uca)
	ucp := datadog.NewUserCreatePayload()
	ucp.SetData(*ucd)
	ur, httpresp, err := Client(ctx).UsersApi.CreateUser(ctx).Body(*ucp).Execute()
	if err != nil {
		t.Fatalf("Error creating User %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(httpresp.StatusCode, 201)
	urData := ur.GetData()
	uid := urData.GetId()
	defer disableUser(ctx, uid)

	urAttributes := urData.GetAttributes()
	assert.Equal(urAttributes.GetEmail(), uca.GetEmail())
	assert.Equal(urAttributes.GetName(), uca.GetName())
	assert.Equal(urAttributes.GetTitle(), uca.GetTitle())

	// now, test updating it
	uua := datadog.NewUserUpdateAttributes()
	uua.SetDisabled(false)
	uua.SetName(uca.GetName() + "-updated")
	uud := datadog.NewUserUpdateData()
	uud.SetAttributes(*uua)
	uud.SetId(uid)
	uup := datadog.NewUserUpdatePayload()
	uup.SetData(*uud)
	// no response payload
	httpresp, err = Client(ctx).UsersApi.UpdateUser(ctx, uid).Body(*uup).Execute()
	if err != nil {
		t.Fatalf("Error updating User %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(httpresp.StatusCode, 204)

	// now, test getting it
	urp, httpresp, err := Client(ctx).UsersApi.GetUser(ctx, uid).Execute()
	if err != nil {
		t.Fatalf("Error getting User %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(httpresp.StatusCode, 200)
	urData = urp.GetData()
	urAttributes = urData.GetAttributes()
	assert.Equal(urAttributes.GetEmail(), uca.GetEmail())
	assert.Equal(urAttributes.GetName(), uca.GetName() + "-updated")
	assert.Equal(urAttributes.GetTitle(), uca.GetTitle())
	assert.Equal(urAttributes.GetDisabled(), false)

	// now, test disabling it
	// no response payload
	httpresp, err = Client(ctx).UsersApi.DisableUser(ctx, uid).Execute()
	if err != nil {
		t.Fatalf("Error disabling User %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(httpresp.StatusCode, 204)

	// now, test filtering for it in the list call
	usrp, httpresp, err := Client(ctx).UsersApi.
		ListUsers(ctx).
		Filter(uca.GetEmail()).
		PageSize(1).
		PageNumber(0).
		SortDir(datadog.QUERYSORTORDER_ASC).
		Execute()
	if err != nil {
		t.Fatalf("Error listing users %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(httpresp.StatusCode, 200)
	assert.Equal(len(usrp.GetData()), 1)
	urAttributes = usrp.GetData()[0].GetAttributes()
	assert.Equal(urAttributes.GetEmail(), uca.GetEmail())
	urMetaAttributes := usrp.GetMeta()
	page := urMetaAttributes.GetPage()
	assert.GreaterOrEqual(page.GetTotalCount(), int64(1))
	assert.GreaterOrEqual(page.GetTotalFilteredCount(), int64(1))

	// NOTE: to test getting a user organization, we'd need to have a "whoami" API endpoint
	// to get the UUID of the current user, but there's no such stable endpoint right now
	// (a user can only get organization for itself, never for a different user)
}

func TestUserInvitation(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	uca := testingUserCreateAttributes(ctx, t)
	ucd := datadog.NewUserCreateData()
	ucd.SetAttributes(*uca)
	ucp := datadog.NewUserCreatePayload()
	ucp.SetData(*ucd)
	ur, httpresp, err := Client(ctx).UsersApi.CreateUser(ctx).Body(*ucp).Execute()
	if err != nil {
		t.Fatalf("Error creating User %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(httpresp.StatusCode, 201)
	urData := ur.GetData()
	id := urData.GetId()
	defer disableUser(ctx, id)

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

	resp, httpresp, err := Client(ctx).UsersApi.SendInvitations(ctx).Body(*uip).Execute()
	if err != nil {
		t.Fatalf("Error sending invitation for %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(httpresp.StatusCode, 201)
	respID := resp.GetData()[0].GetId()

	// now, test getting the invitation
	oneresp, httpresp, err := Client(ctx).UsersApi.GetInvitation(ctx, respID).Execute()
	if err != nil {
		t.Fatalf("Error getting invitation %s: Response %s: %v", respID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(httpresp.StatusCode, 200)
	data := oneresp.GetData()
	assert.Equal(data.GetId(), respID)
}
