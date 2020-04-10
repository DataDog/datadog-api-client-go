/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func generateUniqueUser(t *testing.T) datadog.User {
	prefix := fmt.Sprintf("%s-%d", t.Name(), TESTCLOCK.Now().UnixNano())
	email := strings.ToLower(prefix) + "@integration-tests-accnt-for-sdk-ci.com"
	return datadog.User{
		Name:       datadog.PtrString(prefix),
		Email:      datadog.PtrString(email),
		Handle:     datadog.PtrString(email),
		AccessRole: datadog.ACCESSROLE_READ_ONLY.Ptr(),
	}
}

// [TODO] You can't update another user's email
// This is based on who owns the APP key that is making the changes
var UPDATEUSER = datadog.User{
	Name:       datadog.PtrString("test update user"),
	Disabled:   datadog.PtrBool(true),
	AccessRole: datadog.ACCESSROLE_STANDARD.Ptr(),
}

func TestCreateUser(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testUser := generateUniqueUser(t)
	defer disableUser(testUser.GetHandle())

	// Assert User Created with proper fields
	userCreateResponse, httpresp, err := TESTAPICLIENT.UsersApi.CreateUser(TESTAUTH).Body(testUser).Execute()
	if err != nil {
		t.Fatalf("Error creating User %s: Response %s: %v", testUser.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)

	user := userCreateResponse.GetUser()
	assert.Equal(t, testUser.GetName(), user.GetName())
	assert.Equal(t, testUser.GetHandle(), user.GetHandle())
	assert.Equal(t, testUser.GetAccessRole(), user.GetAccessRole())
	assert.Equal(t, testUser.GetEmail(), user.GetEmail())

	// Assert User Get with proper fields
	userGetResponse, httpresp, err := TESTAPICLIENT.UsersApi.GetUser(TESTAUTH, testUser.GetEmail()).Execute()
	if err != nil {
		t.Fatalf("Error Getting User %s: Response %s: %v", testUser.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)

	user = userGetResponse.GetUser()
	assert.Equal(t, testUser.GetName(), user.GetName())
	assert.Equal(t, testUser.GetHandle(), user.GetHandle())
	assert.Equal(t, testUser.GetAccessRole(), user.GetAccessRole())
	assert.Equal(t, testUser.GetEmail(), user.GetEmail())
}

func TestUpdateUser(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testUser := generateUniqueUser(t)
	defer disableUser(testUser.GetHandle())

	// Assert User Created with proper fields
	userCreateResponse, httpresp, err := TESTAPICLIENT.UsersApi.CreateUser(TESTAUTH).Body(testUser).Execute()
	if err != nil {
		t.Fatalf("Error creating User %v: Response %s: %v", testUser, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	userCreateResponse.GetUser()

	// Assert User Get with proper fields
	userUpdateResponse, httpresp, err := TESTAPICLIENT.UsersApi.UpdateUser(TESTAUTH, testUser.GetHandle()).Body(UPDATEUSER).Execute()
	if err != nil {
		t.Fatalf("Error getting User %s: Response %s: %v", testUser.GetHandle(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)

	user := userUpdateResponse.GetUser()
	// Test fields were updated
	assert.Equal(t, UPDATEUSER.GetName(), user.GetName())
	assert.Equal(t, UPDATEUSER.GetDisabled(), user.GetDisabled())
	assert.Equal(t, UPDATEUSER.GetAccessRole(), user.GetAccessRole())
	// Test unchanged field remains unchanged
	assert.Equal(t, testUser.GetHandle(), user.GetHandle())
}

func TestDisableUser(t *testing.T) {
	// We already test this in the disableUser cleanup function, but good to have an explicit test
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testUser := generateUniqueUser(t)
	defer disableUser(testUser.GetHandle())

	// Assert User Created with proper fields
	_, httpresp, err := TESTAPICLIENT.UsersApi.CreateUser(TESTAUTH).Body(testUser).Execute()
	if err != nil {
		t.Fatalf("Error creating User %v: Response %s: %v", testUser, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)

	_, httpresp, err = TESTAPICLIENT.UsersApi.DisableUser(TESTAUTH, testUser.GetHandle()).Execute()
	if err != nil {
		t.Fatalf("Error disabling User %s: Response %s: %v", testUser.GetHandle(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
}

func TestListUsers(t *testing.T) {
	if !tests.IsRecording() {
		t.Skip("This test case does not support reply from recording")
	}

	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Assert User Created with proper fields
	userListResponse, httpresp, err := TESTAPICLIENT.UsersApi.ListUsers(TESTAUTH).Execute()
	if err != nil {
		t.Fatalf("Error listing Users. Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	// Just assert the user list isn't empty and contains a user object
	assert.NotEmpty(t, userListResponse.GetUsers())
	assert.NotEmpty(t, userListResponse.GetUsers()[0].GetName())
}

func TestUserCreateErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Create user to trigger 409
	testUser := generateUniqueUser(t)
	_, _, err := TESTAPICLIENT.UsersApi.CreateUser(TESTAUTH).Body(testUser).Execute()
	if err != nil {
		t.Fatalf("Error creating User %v: Response %s: %v", testUser, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer disableUser(testUser.GetHandle())

	testCases := []struct {
		Name               string
		Ctx                context.Context
		Body               datadog.User
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, datadog.User{}, 400},
		{"403 Forbidden", fake_auth, testUser, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.UsersApi.CreateUser(tc.Ctx).Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestUserCreate409Error(t *testing.T) {
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)

	res, err := tests.ReadFixture("fixtures/user/error_409.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because it needs the user to be activated
	gock.New("https://api.datadoghq.com").Post("/api/v1/user").Reply(409).JSON(res)
	defer gock.Off()

	_, httpresp, err := TESTAPICLIENT.UsersApi.CreateUser(TESTAUTH).Body(datadog.User{}).Execute()
	assert.Equal(t, 409, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}

func TestUserListErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.UsersApi.ListUsers(tc.Ctx).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestUserGetErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"403 Forbidden", fake_auth, 403},
		{"404 Not Found", TESTAUTH, 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.UsersApi.GetUser(tc.Ctx, "notahandle").Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestUserUpdateErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Create user
	testUser := generateUniqueUser(t)
	_, _, err := TESTAPICLIENT.UsersApi.CreateUser(TESTAUTH).Body(testUser).Execute()
	if err != nil {
		t.Fatalf("Error creating User %v: Response %s: %v", testUser, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer disableUser(testUser.GetHandle())

	badUser := *datadog.NewUserWithDefaults()
	badUser.SetEmail("notanemail")

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ID                 string
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, testUser.GetHandle(), 400},
		{"403 Forbidden", fake_auth, "notahandle", 403},
		{"404 Not Found", TESTAUTH, "notahandle", 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.UsersApi.UpdateUser(tc.Ctx, tc.ID).Body(badUser).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestUserDisableErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Create user and disable it to trigger 400
	testUser := generateUniqueUser(t)
	_, _, err := TESTAPICLIENT.UsersApi.CreateUser(TESTAUTH).Body(testUser).Execute()
	if err != nil {
		t.Fatalf("Error creating User %v: Response %s: %v", testUser, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	disableUser(testUser.GetHandle())

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ID                 string
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, testUser.GetHandle(), 400},
		{"403 Forbidden", fake_auth, "notahandle", 403},
		{"404 Not Found", TESTAUTH, "notahandle", 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.UsersApi.DisableUser(tc.Ctx, tc.ID).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func disableUser(userID string) {
	_, httpresp, err := TESTAPICLIENT.UsersApi.DisableUser(TESTAUTH, userID).Execute()
	if httpresp.StatusCode != 200 || err != nil {
		log.Printf("Error disabling User: %v, Another test may have already disabled this user.", userID)
	}
}
