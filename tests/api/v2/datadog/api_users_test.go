package test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
)

func testingUserCreateAttributes(ctx context.Context, t *testing.T) *datadog.UserCreateAttributes {
	uca := datadog.NewUserCreateAttributesWithDefaults()
	// the API lowercases returned emails, so let's send a lowercase value in the first place
	name := strings.ToLower(*tests.UniqueEntityName(ctx, t))
	uca.SetEmail(fmt.Sprintf("%s@datadoghq.com", name))
	uca.SetName(name)
	uca.SetTitle("Big boss")
	return uca
}

func disableUser(ctx context.Context, t *testing.T, userID string) {
	_, err := Client(ctx).UsersApi.DisableUser(ctx, userID)
	if err == nil {
		return
	}
	t.Logf("Error disabling User: %v, Another test may have already disabled this user: %s", userID, err.Error())
}

func TestUserLifecycle(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// first, test creating a user
	uca := testingUserCreateAttributes(ctx, t)
	ucd := datadog.NewUserCreateDataWithDefaults()
	ucd.SetAttributes(*uca)
	ucr := datadog.NewUserCreateRequestWithDefaults()
	ucr.SetData(*ucd)
	ur, httpresp, err := Client(ctx).UsersApi.CreateUser(ctx, *ucr)
	if err != nil {
		t.Fatalf("Error creating User %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(httpresp.StatusCode, 201)
	urData := ur.GetData()
	uid := urData.GetId()
	defer disableUser(ctx, t, uid)

	urAttributes := urData.GetAttributes()
	assert.Equal(urAttributes.GetEmail(), uca.GetEmail())
	assert.Equal(urAttributes.GetName(), uca.GetName())
	assert.Equal(urAttributes.GetTitle(), uca.GetTitle())

	// now, test updating it
	uua := datadog.NewUserUpdateAttributes()
	uua.SetDisabled(false)
	uua.SetName(uca.GetName() + "-updated")
	uud := datadog.NewUserUpdateDataWithDefaults()
	uud.SetAttributes(*uua)
	uud.SetId(uid)
	uur := datadog.NewUserUpdateRequestWithDefaults()
	uur.SetData(*uud)
	// no response payload
	_, httpresp, err = Client(ctx).UsersApi.UpdateUser(ctx, uid, *uur)
	if err != nil {
		t.Fatalf("Error updating User %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(httpresp.StatusCode, 200)

	// now, test getting it
	urp, httpresp, err := Client(ctx).UsersApi.GetUser(ctx, uid)
	if err != nil {
		t.Fatalf("Error getting User %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(httpresp.StatusCode, 200)
	urData = urp.GetData()
	urAttributes = urData.GetAttributes()
	assert.Equal(urAttributes.GetEmail(), uca.GetEmail())
	assert.Equal(urAttributes.GetName(), uca.GetName()+"-updated")
	assert.Equal(urAttributes.GetTitle(), uca.GetTitle())
	assert.Equal(urAttributes.GetDisabled(), false)

	// now, test disabling it
	// no response payload
	httpresp, err = Client(ctx).UsersApi.DisableUser(ctx, uid)
	if err != nil {
		t.Fatalf("Error disabling User %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(httpresp.StatusCode, 204)

	// now, test filtering for it in the list call
	usrp, httpresp, err := Client(ctx).UsersApi.ListUsers(ctx, *datadog.NewListUsersOptionalParameters().
		WithFilter(uca.GetEmail()).
		WithPageSize(1).
		WithPageNumber(0).
		WithSortDir(datadog.QUERYSORTORDER_ASC))
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

func TestUpdateUserErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()

	// Build update payload
	uid := "00000000-dead-beef-dead-ffffffffffff"
	uua := datadog.NewUserUpdateAttributes()
	uua.SetDisabled(false)
	uua.SetName("Joe Doe")
	uud := datadog.NewUserUpdateDataWithDefaults()
	uud.SetAttributes(*uua)
	uud.SetId(uid)
	uur := datadog.NewUserUpdateRequestWithDefaults()
	uur.SetData(*uud)

	// 422 needs a mismatched id, shallow copy is fine here
	uud422 := uud
	uud422.SetId("00000000-mismatch-body-id-ffffffffffff")
	uur422 := datadog.NewUserUpdateRequestWithDefaults()
	uur422.SetData(*uud)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
		UserID             string
		Body               *datadog.UserUpdateRequest
	}{
		"400 Bad Request":            {WithTestAuth, 400, uid, &datadog.UserUpdateRequest{}},
		"403 Forbidden":              {WithFakeAuth, 403, uid, uur},
		"404 Bad User ID in Path":    {WithTestAuth, 404, uid, uur},
		"422 Bad User ID in Request": {WithTestAuth, 422, uid, uur422},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).UsersApi.UpdateUser(ctx, tc.UserID, *tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUserInvitation(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	uca := testingUserCreateAttributes(ctx, t)
	ucd := datadog.NewUserCreateDataWithDefaults()
	ucd.SetAttributes(*uca)
	ucr := datadog.NewUserCreateRequestWithDefaults()
	ucr.SetData(*ucd)
	ur, httpresp, err := Client(ctx).UsersApi.CreateUser(ctx, *ucr)
	if err != nil {
		t.Fatalf("Error creating User %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(httpresp.StatusCode, 201)
	urData := ur.GetData()
	id := urData.GetId()
	defer disableUser(ctx, t, id)

	// first, create the user invitation
	rtud := datadog.NewRelationshipToUserDataWithDefaults()
	rtud.SetId(id)
	rtu := datadog.NewRelationshipToUserWithDefaults()
	rtu.SetData(*rtud)
	uir := datadog.NewUserInvitationRelationshipsWithDefaults()
	uir.SetUser(*rtu)
	uid := datadog.NewUserInvitationDataWithDefaults()
	uid.SetRelationships(*uir)
	uireq := datadog.NewUserInvitationsRequestWithDefaults()
	uireq.Data = append(uireq.Data, *uid)

	resp, httpresp, err := Client(ctx).UsersApi.SendInvitations(ctx, *uireq)
	if err != nil {
		assert.IsType(datadog.GenericOpenAPIError{}, err, "%v", err)
		t.Fatalf("Error sending invitation for %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(httpresp.StatusCode, 201)
	respID := resp.GetData()[0].GetId()

	// now, test getting the invitation
	oneresp, httpresp, err := Client(ctx).UsersApi.GetInvitation(ctx, respID)
	if err != nil {
		assert.IsType(datadog.GenericOpenAPIError{}, err, "%v", err)
		t.Fatalf("Error getting invitation %s: Response %s: %v", respID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(httpresp.StatusCode, 200)
	data := oneresp.GetData()
	assert.Equal(data.GetId(), respID)
}
