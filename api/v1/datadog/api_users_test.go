package datadog_test

import (
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

func generateUniqueUser(t *testing.T) datadog.User {
	prefix := fmt.Sprintf("%s-%d", t.Name(), time.Now().UnixNano())
	email := strings.ToLower(prefix) + "@integration-tests-accnt-for-sdk-ci.com"
	return datadog.User{
		Name:       datadog.PtrString(prefix),
		Email:      datadog.PtrString(email),
		Handle:     datadog.PtrString(email),
		AccessRole: datadog.PtrString("ro"),
	}
}

// [TODO] You can't update another user's email
// This is based on who owns the APP key that is making the changes
var UPDATEUSER = datadog.User{
	Name:       datadog.PtrString("test update user"),
	Disabled:   datadog.PtrBool(true),
	AccessRole: datadog.PtrString("st"),
}

func TestCreateUser(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testUser := generateUniqueUser(t)
	defer disableUser(testUser.GetHandle())

	// Assert User Created with proper fields
	userCreateResponse, httpresp, err := TESTAPICLIENT.UsersApi.CreateUser(TESTAUTH, testUser)
	if err != nil {
		t.Errorf("Error Creating User: %v", err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	user := userCreateResponse.GetUser()
	assert.Equal(t, user.GetName(), testUser.GetName())
	assert.Equal(t, user.GetHandle(), testUser.GetHandle())
	assert.Equal(t, user.GetAccessRole(), testUser.GetAccessRole())
	assert.Equal(t, user.GetEmail(), testUser.GetEmail())

	// Assert User Get with proper fields
	userGetResponse, httpresp, err := TESTAPICLIENT.UsersApi.GetUser(TESTAUTH, testUser.GetEmail())
	if err != nil {
		t.Errorf("Error Getting User %s: Response %s: %v", testUser.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	user = userGetResponse.GetUser()
	assert.Equal(t, user.GetName(), testUser.GetName())
	assert.Equal(t, user.GetHandle(), testUser.GetHandle())
	assert.Equal(t, user.GetAccessRole(), testUser.GetAccessRole())
	assert.Equal(t, user.GetEmail(), testUser.GetEmail())
}

func TestUpdateUser(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testUser := generateUniqueUser(t)
	defer disableUser(testUser.GetHandle())

	// Assert User Created with proper fields
	userCreateResponse, httpresp, err := TESTAPICLIENT.UsersApi.CreateUser(TESTAUTH, testUser)
	if err != nil {
		t.Errorf("Error creating User %v: Response %s: %v", testUser, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	userCreateResponse.GetUser()

	// Assert User Get with proper fields
	userUpdateResponse, httpresp, err := TESTAPICLIENT.UsersApi.UpdateUser(TESTAUTH, testUser.GetHandle(), UPDATEUSER)
	if err != nil {
		t.Errorf("Error getting User %s: Response %s: %v", testUser.GetHandle(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	user := userUpdateResponse.GetUser()
	// Test fields were updated
	assert.Equal(t, user.GetName(), UPDATEUSER.GetName())
	assert.Equal(t, user.GetDisabled(), UPDATEUSER.GetDisabled())
	assert.Equal(t, user.GetAccessRole(), UPDATEUSER.GetAccessRole())
	// Test unchanged field remains unchanged
	assert.Equal(t, user.GetHandle(), testUser.GetHandle())
}

func TestDisableUser(t *testing.T) {
	// We already test this in the disableUser cleanup function, but good to have an explicit test
	testUser := generateUniqueUser(t)
	defer disableUser(testUser.GetHandle())

	// Assert User Created with proper fields
	_, httpresp, err := TESTAPICLIENT.UsersApi.CreateUser(TESTAUTH, testUser)
	if err != nil {
		t.Errorf("Error Creating User: %v", err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	_, httpresp, err = TESTAPICLIENT.UsersApi.DisableUser(TESTAUTH, testUser.GetHandle())
	if err != nil {
		t.Errorf("Error disabling User %s: Response %s: %v", testUser.GetHandle(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
}

func disableUser(userID string) {
	_, httpresp, err := TESTAPICLIENT.UsersApi.DisableUser(TESTAUTH, userID)
	if httpresp.StatusCode != 200 || err != nil {
		log.Printf("Error disabling User: %v, Another test may have already disabled this user.", userID)
	}
}
