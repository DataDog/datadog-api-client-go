package datadog_test

import (
	"log"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

var TESTUSER = datadog.User{
	Name:       datadog.PtrString("test user"),
	Email:      datadog.PtrString("test@example.com"),
	Handle:     datadog.PtrString("test@example.com"),
	AccessRole: datadog.PtrString("ro"),
}

// [TODO] You can't update another user's email
// This is based on who owns the APP key that is making the changes
var UPDATEUSER = datadog.User{
	Name:       datadog.PtrString("test update user"),
	Disabled:   datadog.PtrBool(true),
	AccessRole: datadog.PtrString("st"),
}

var UPDATEENABLEUSER = datadog.User{
	Disabled: datadog.PtrBool(false),
}

func TestCreateUser(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	defer disableUser(TESTUSER.GetHandle())

	// Assert User Created with proper fields
	userCreateResponse, httpresp, err := TESTAPICLIENT.UsersApi.CreateUser(TESTAUTH, TESTUSER)
	if err != nil {
		t.Errorf("Error Creating User: %v", err)
	}
	user := userCreateResponse.GetUser()
	assert.Equal(t, user.GetName(), TESTUSER.GetName())
	assert.Equal(t, user.GetHandle(), TESTUSER.GetHandle())
	assert.Equal(t, user.GetAccessRole(), TESTUSER.GetAccessRole())
	assert.Equal(t, user.GetEmail(), TESTUSER.GetEmail())

	// Assert User Get with proper fields
	userGetResponse, httpresp, err := TESTAPICLIENT.UsersApi.GetUser(TESTAUTH, TESTUSER.GetEmail())
	if httpresp.StatusCode != 200 || err != nil {
		t.Errorf("Error Getting User: %v: %v", httpresp, err)
	}
	user = userGetResponse.GetUser()
	assert.Equal(t, user.GetName(), TESTUSER.GetName())
	assert.Equal(t, user.GetHandle(), TESTUSER.GetHandle())
	assert.Equal(t, user.GetAccessRole(), TESTUSER.GetAccessRole())
	assert.Equal(t, user.GetEmail(), TESTUSER.GetEmail())
}

func TestUpdateUser(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	defer disableUser(TESTUSER.GetHandle())

	// Assert User Created with proper fields
	userCreateResponse, httpresp, err := TESTAPICLIENT.UsersApi.CreateUser(TESTAUTH, TESTUSER)
	if err != nil {
		t.Errorf("Error Creating User: %v", err)
	}
	userCreateResponse.GetUser()

	// Assert User Get with proper fields
	userUpdateResponse, httpresp, err := TESTAPICLIENT.UsersApi.UpdateUser(TESTAUTH, TESTUSER.GetHandle(), UPDATEUSER)
	if httpresp.StatusCode != 200 || err != nil {
		t.Errorf("Error Getting User: %v: %v", httpresp, err)
	}
	user := userUpdateResponse.GetUser()
	// Test fields were updated
	assert.Equal(t, user.GetName(), UPDATEUSER.GetName())
	assert.Equal(t, user.GetDisabled(), UPDATEUSER.GetDisabled())
	assert.Equal(t, user.GetAccessRole(), UPDATEUSER.GetAccessRole())
	// Test unchanged field remains unchanged
	assert.Equal(t, user.GetHandle(), TESTUSER.GetHandle())
}

func TestDisableUser(t *testing.T) {
	// We already test this in the disableUser cleanup function, but good to have an explicit test

	// First ensure that the user isn't currently disabled
	TESTAPICLIENT.UsersApi.UpdateUser(TESTAUTH, TESTUSER.GetHandle(), UPDATEENABLEUSER)

	_, httpresp, err := TESTAPICLIENT.UsersApi.DisableUser(TESTAUTH, TESTUSER.GetHandle())
	if httpresp.StatusCode != 200 || err != nil {
		t.Errorf("Error disabling User: %v: %v", httpresp, err)
	}
}

func disableUser(userID string) {
	_, httpresp, err := TESTAPICLIENT.UsersApi.DisableUser(TESTAUTH, userID)
	if httpresp.StatusCode != 200 || err != nil {
		log.Printf("Error disabling User: %v, Another test may have already disabled this user.", userID)
	}
}
