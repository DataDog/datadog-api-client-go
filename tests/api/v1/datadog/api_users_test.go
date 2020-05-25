/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"log"
	"strings"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"

	"gopkg.in/h2non/gock.v1"
)

func generateUniqueUser(ctx context.Context, t *testing.T) datadog.User {
	prefix := *tests.UniqueEntityName(ctx, t)
	email := strings.ToLower(prefix) + "@integration-tests-accnt-for-sdk-ci.com"
	return datadog.User{
		Name:       datadog.PtrString(prefix),
		Email:      datadog.PtrString(email),
		Handle:     datadog.PtrString(email),
		AccessRole: datadog.ACCESSROLE_READ_ONLY.Ptr(),
	}
}

func getUpdateUser(ctx context.Context, t *testing.T) datadog.User {
	// [TODO] You can't update another user's email
	// This is based on who owns the APP key that is making the changes
	return datadog.User{
		Name:       tests.UniqueEntityName(ctx, t),
		Disabled:   datadog.PtrBool(true),
		AccessRole: datadog.ACCESSROLE_STANDARD.Ptr(),
	}
}

func TestCreateUser(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	testUser := generateUniqueUser(ctx, t)
	defer disableUser(ctx, testUser.GetHandle())

	// Assert User Created with proper fields
	userCreateResponse, httpresp, err := Client(ctx).UsersApi.CreateUser(ctx).Body(testUser).Execute()
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
	userGetResponse, httpresp, err := Client(ctx).UsersApi.GetUser(ctx, testUser.GetEmail()).Execute()
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
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	testUser := generateUniqueUser(ctx, t)
	defer disableUser(ctx, testUser.GetHandle())

	// Assert User Created with proper fields
	userCreateResponse, httpresp, err := Client(ctx).UsersApi.CreateUser(ctx).Body(testUser).Execute()
	if err != nil {
		t.Fatalf("Error creating User %v: Response %s: %v", testUser, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	userCreateResponse.GetUser()

	// Assert User Get with proper fields
	updateUser := getUpdateUser(ctx, t)
	userUpdateResponse, httpresp, err := Client(ctx).UsersApi.UpdateUser(ctx, testUser.GetHandle()).Body(updateUser).Execute()
	if err != nil {
		t.Fatalf("Error getting User %s: Response %s: %v", testUser.GetHandle(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	user := userUpdateResponse.GetUser()
	// Test fields were updated
	assert.Equal(updateUser.GetName(), user.GetName())
	assert.Equal(updateUser.GetDisabled(), user.GetDisabled())
	assert.Equal(updateUser.GetAccessRole(), user.GetAccessRole())
	// Test unchanged field remains unchanged
	assert.Equal(testUser.GetHandle(), user.GetHandle())
}

func TestDisableUser(t *testing.T) {
	// We already test this in the disableUser cleanup function, but good to have an explicit test
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	testUser := generateUniqueUser(ctx, t)
	defer disableUser(ctx, testUser.GetHandle())

	// Assert User Created with proper fields
	_, httpresp, err := Client(ctx).UsersApi.CreateUser(ctx).Body(testUser).Execute()
	if err != nil {
		t.Fatalf("Error creating User %v: Response %s: %v", testUser, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	_, httpresp, err = Client(ctx).UsersApi.DisableUser(ctx, testUser.GetHandle()).Execute()
	if err != nil {
		t.Fatalf("Error disabling User %s: Response %s: %v", testUser.GetHandle(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
}

func TestListUsers(t *testing.T) {
	if tests.GetRecording() == tests.ModeReplaying {
		t.Skip("This test case does not support reply from recording")
	}

	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Assert User Created with proper fields
	userListResponse, httpresp, err := Client(ctx).UsersApi.ListUsers(ctx).Execute()
	if err != nil {
		t.Fatalf("Error listing Users. Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	// Just assert the user list isn't empty and contains a user object
	assert.NotEmpty(userListResponse.GetUsers())
	assert.NotEmpty(userListResponse.GetUsers()[0].GetName())
}

func TestUserCreateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.User
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.User{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.User{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).UsersApi.CreateUser(ctx).Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUserCreate409Error(t *testing.T) {
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	res, err := tests.ReadFixture("fixtures/user/error_409.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because it needs the user to be activated
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "UsersApiService.CreateUser")
	assert.NoError(err)
	gock.New(URL).Post("/api/v1/user").Reply(409).JSON(res)
	defer gock.Off()

	_, httpresp, err := Client(ctx).UsersApi.CreateUser(ctx).Body(datadog.User{}).Execute()
	assert.Equal(409, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestUserListErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

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

			_, httpresp, err := Client(ctx).UsersApi.ListUsers(ctx).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUserGetErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

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

			_, httpresp, err := Client(ctx).UsersApi.GetUser(ctx, "notahandle").Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUserUpdateErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()

	// Create user
	testUser := generateUniqueUser(ctx, t)
	_, _, err := Client(ctx).UsersApi.CreateUser(ctx).Body(testUser).Execute()
	if err != nil {
		t.Fatalf("Error creating User %v: Response %s: %v", testUser, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer disableUser(ctx, testUser.GetHandle())

	badUser := *datadog.NewUserWithDefaults()
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

			_, httpresp, err := Client(ctx).UsersApi.UpdateUser(ctx, tc.ID).Body(badUser).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUserDisableErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()

	// Create user and disable it to trigger 400
	testUser := generateUniqueUser(ctx, t)
	_, _, err := Client(ctx).UsersApi.CreateUser(ctx).Body(testUser).Execute()
	if err != nil {
		t.Fatalf("Error creating User %v: Response %s: %v", testUser, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	disableUser(ctx, testUser.GetHandle())

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

			_, httpresp, err := Client(ctx).UsersApi.DisableUser(ctx, tc.ID).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func disableUser(ctx context.Context, userID string) {
	_, httpresp, err := Client(ctx).UsersApi.DisableUser(ctx, userID).Execute()
	if httpresp.StatusCode != 200 || err != nil {
		log.Printf("Error disabling User: %v, Another test may have already disabled this user.", userID)
	}
}
