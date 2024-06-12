/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"strings"
	"testing"

	"gopkg.in/h2non/gock.v1"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/DataDog/datadog-api-client-go/v2/tests"
)

func generateUniqueUser(ctx context.Context, t *testing.T) datadogV1.User {
	prefix := *tests.UniqueEntityName(ctx, t)
	email := strings.ToLower(prefix) + "@integration-tests-accnt-for-sdk-ci.com"
	return datadogV1.User{
		Name:       datadog.PtrString(prefix),
		Email:      datadog.PtrString(email),
		Handle:     datadog.PtrString(email),
		AccessRole: *datadogV1.NewNullableAccessRole(datadogV1.ACCESSROLE_READ_ONLY.Ptr()),
	}
}

func getUpdateUser(ctx context.Context, t *testing.T) datadogV1.User {
	// [TODO] You can't update another user's email
	// This is based on who owns the APP key that is making the changes
	return datadogV1.User{
		Name:     tests.UniqueEntityName(ctx, t),
		Disabled: datadog.PtrBool(true),
	}
}

func TestCreateUser(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsersApi(Client(ctx))

	testUser := generateUniqueUser(ctx, t)
	defer disableUser(ctx, t, testUser.GetHandle())

	// Assert User Created with proper fields
	userCreateResponse, httpresp, err := api.CreateUser(ctx, testUser)
	if err != nil {
		t.Fatalf("Error creating User %s: Response %s: %v", testUser.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	user := userCreateResponse.GetUser()
	assert.Equal(testUser.GetName(), user.GetName())
	assert.Equal(testUser.GetHandle(), user.GetHandle())
	assert.Equal(testUser.GetAccessRole(), user.GetAccessRole())
	assert.Equal(testUser.GetEmail(), user.GetEmail())

	// Assert User Get with proper fields
	userGetResponse, httpresp, err := api.GetUser(ctx, testUser.GetEmail())
	if err != nil {
		t.Fatalf("Error Getting User %s: Response %s: %v", testUser.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	user = userGetResponse.GetUser()
	assert.Equal(testUser.GetName(), user.GetName())
	assert.Equal(testUser.GetHandle(), user.GetHandle())
	assert.Equal(testUser.GetAccessRole(), user.GetAccessRole())
	assert.Equal(testUser.GetEmail(), user.GetEmail())
}

func TestUpdateUser(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsersApi(Client(ctx))

	testUser := generateUniqueUser(ctx, t)
	defer disableUser(ctx, t, testUser.GetHandle())

	// Assert User Created with proper fields
	userCreateResponse, httpresp, err := api.CreateUser(ctx, testUser)
	if err != nil {
		t.Fatalf("Error creating User %v: Response %s: %v", testUser, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	userCreateResponse.GetUser()

	// Assert User Get with proper fields
	updateUser := getUpdateUser(ctx, t)
	userUpdateResponse, httpresp, err := api.UpdateUser(ctx, testUser.GetHandle(), updateUser)
	if err != nil {
		t.Fatalf("Error getting User %s: Response %s: %v", testUser.GetHandle(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	user := userUpdateResponse.GetUser()
	// Test fields were updated
	assert.Equal(updateUser.GetName(), user.GetName())
	assert.Equal(updateUser.GetDisabled(), user.GetDisabled())
	// Test unchanged field remains unchanged
	assert.Equal(testUser.GetHandle(), user.GetHandle())
}

func TestDisableUser(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// We already test this in the disableUser cleanup function, but good to have an explicit test
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsersApi(Client(ctx))

	testUser := generateUniqueUser(ctx, t)
	defer disableUser(ctx, t, testUser.GetHandle())

	// Assert User Created with proper fields
	_, httpresp, err := api.CreateUser(ctx, testUser)
	if err != nil {
		t.Fatalf("Error creating User %v: Response %s: %v", testUser, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	_, httpresp, err = api.DisableUser(ctx, testUser.GetHandle())
	if err != nil {
		t.Fatalf("Error disabling User %s: Response %s: %v", testUser.GetHandle(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
}

func TestListUsers(t *testing.T) {
	t.Skip("Skipping until account is cleaned up")
	// if tests.GetRecording() != tests.ModeIgnore {
	// 	t.Skip("This test case does not support reply from recording")
	// }

	// ctx, finish := tests.WithTestSpan(context.Background(), t)
	// defer finish()
	// // Setup the Client we'll use to interact with the Test account
	// ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	// defer finish()
	// assert := tests.Assert(ctx, t)
	// api := datadogV1.NewUsersApi(Client(ctx))

	// // Assert User Created with proper fields
	// userListResponse, httpresp, err := api.ListUsers(ctx)
	// if err != nil {
	// 	t.Fatalf("Error listing Users. Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	// }
	// assert.Equal(200, httpresp.StatusCode)
	// // Just assert the user list isn't empty and contains a user object
	// assert.NotEmpty(userListResponse.GetUsers())
	// assert.NotEmpty(userListResponse.GetUsers()[0].GetName())
}

func TestUserCreateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadogV1.User
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadogV1.User{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadogV1.User{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewUsersApi(Client(ctx))

			_, httpresp, err := api.CreateUser(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUserCreate409Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsersApi(Client(ctx))

	res, err := tests.ReadFixture("fixtures/user/error_409.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because it needs the user to be activated
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "UsersApiService.CreateUser")
	assert.NoError(err)
	gock.New(URL).Post("/api/v1/user").Reply(409).JSON(res)
	defer gock.Off()

	_, httpresp, err := api.CreateUser(ctx, datadogV1.User{})
	assert.Equal(409, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestUserListErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewUsersApi(Client(ctx))

			_, httpresp, err := api.ListUsers(ctx)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUserGetErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
		"404 Not Found": {WithTestAuth, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewUsersApi(Client(ctx))

			_, httpresp, err := api.GetUser(ctx, "notahandle")
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUserUpdateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	api := datadogV1.NewUsersApi(Client(ctx))

	// Create user
	testUser := generateUniqueUser(ctx, t)
	_, _, err := api.CreateUser(ctx, testUser)
	if err != nil {
		t.Fatalf("Error creating User %v: Response %s: %v", testUser, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer disableUser(ctx, t, testUser.GetHandle())

	badUser := *datadogV1.NewUserWithDefaults()
	badUser.SetEmail("notanemail")

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ID                 string
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, testUser.GetHandle(), 400},
		"403 Forbidden":   {WithFakeAuth, "notahandle", 403},
		"404 Not Found":   {WithTestAuth, "notahandle", 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewUsersApi(Client(ctx))

			_, httpresp, err := api.UpdateUser(ctx, tc.ID, badUser)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUserDisableErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	api := datadogV1.NewUsersApi(Client(ctx))

	// Create user and disable it to trigger 400
	testUser := generateUniqueUser(ctx, t)
	_, _, err := api.CreateUser(ctx, testUser)
	if err != nil {
		t.Fatalf("Error creating User %v: Response %s: %v", testUser, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	disableUser(ctx, t, testUser.GetHandle())

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ID                 string
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, testUser.GetHandle(), 400},
		"403 Forbidden":   {WithFakeAuth, "notahandle", 403},
		"404 Not Found":   {WithTestAuth, "notahandle", 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewUsersApi(Client(ctx))

			_, httpresp, err := api.DisableUser(ctx, tc.ID)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func disableUser(ctx context.Context, t *testing.T, userID string) {
	api := datadogV1.NewUsersApi(Client(ctx))

	_, httpresp, err := api.DisableUser(ctx, userID)
	if httpresp.StatusCode != 200 || err != nil {
		t.Logf("Error disabling User: %v, Another test may have already disabled this user.", userID)
	}
}
