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
	"net/http"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"

	"gopkg.in/h2non/gock.v1"
)

func TestApiKeyFunctions(t *testing.T) {
	if !tests.IsRecording() {
		t.Skip("This test case does not support reply from recording")
	}

	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// TODO remove when API keys support non secret primary identifiers
	// NOTE do not forget to replace time -> c.Clock
	Client(ctx).GetConfig().HTTPClient = http.DefaultClient

	// Create API Key
	// ----------------------------------
	testAPIKeyName := fmt.Sprintf("%s:%d", t.Name(), time.Now().UnixNano())
	apiKeyData, httpresp, err := Client(ctx).KeyManagementApi.CreateAPIKey(ctx).Body(datadog.ApiKey{Name: &testAPIKeyName}).Execute()
	if err != nil {
		t.Errorf("Error creating api key %v: Response %s: %v", testAPIKeyName, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteAPIKey(ctx, apiKeyData.ApiKey.GetKey())
	assert.Equal(200, httpresp.StatusCode)

	createAPIKeyReturned := apiKeyData.GetApiKey()
	createAPIKeyName := createAPIKeyReturned.GetName()
	createAPIKeyCreated := createAPIKeyReturned.GetCreated()
	createAPIKeyCreatedBy := createAPIKeyReturned.GetCreatedBy()
	createAPIKeyValue := createAPIKeyReturned.GetKey()

	// none of these values should be empty
	assert.NotEmpty(createAPIKeyName)
	assert.NotEmpty(createAPIKeyCreated)
	assert.NotEmpty(createAPIKeyCreatedBy)
	assert.NotEmpty(createAPIKeyValue)
	assert.Equal(testAPIKeyName, createAPIKeyName)

	// Get API Key
	// ----------------------------------
	apiKeyData, httpresp, err = Client(ctx).KeyManagementApi.GetAPIKey(ctx, createAPIKeyValue).Execute()
	if err != nil {
		t.Errorf("Error getting api key %v: Response %s: %v", createAPIKeyValue, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	getAPIKeyReturned := apiKeyData.GetApiKey()
	getAPIKeyName := getAPIKeyReturned.GetName()
	getAPIKeyCreated := getAPIKeyReturned.GetCreated()
	getAPIKeyCreatedBy := getAPIKeyReturned.GetCreatedBy()
	getAPIKeyValue := getAPIKeyReturned.GetKey()

	// should be the same as the create
	assert.Equal(createAPIKeyName, getAPIKeyName)
	assert.Equal(createAPIKeyCreated, getAPIKeyCreated)
	assert.Equal(createAPIKeyCreatedBy, getAPIKeyCreatedBy)
	assert.Equal(createAPIKeyValue, getAPIKeyValue)

	// Get All API Keys
	// ----------------------------------
	respListData, httpresp, err := Client(ctx).KeyManagementApi.ListAPIKeys(ctx).Execute()
	if err != nil {
		t.Errorf("Error getting all api keys: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	getAllAPIKeyReturned := respListData.GetApiKeys()

	// should have more than 1 at least
	assert.True(len(getAllAPIKeyReturned) > 1)

	// Edit API Key
	// ----------------------------------
	newAPIKeyName := fmt.Sprintf("new-%s", testAPIKeyName)
	apiKeyData, httpresp, err = Client(ctx).KeyManagementApi.UpdateAPIKey(ctx, createAPIKeyValue).Body(datadog.ApiKey{Name: &newAPIKeyName}).Execute()
	if err != nil {
		t.Errorf("Error editing api key %v: Response %s: %v", createAPIKeyValue, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	editAPIKeyReturned := apiKeyData.GetApiKey()
	editAPIKeyName := editAPIKeyReturned.GetName()
	editAPIKeyCreated := editAPIKeyReturned.GetCreated()
	editAPIKeyCreatedBy := editAPIKeyReturned.GetCreatedBy()
	editAPIKeyValue := editAPIKeyReturned.GetKey()

	// everything should be the same except the name
	assert.NotEqual(getAPIKeyName, editAPIKeyName)
	assert.Equal(getAPIKeyCreated, editAPIKeyCreated)
	assert.Equal(getAPIKeyCreatedBy, editAPIKeyCreatedBy)
	assert.Equal(getAPIKeyValue, editAPIKeyValue)

	// Delete API Key
	// ----------------------------------
	apiKeyData, httpresp, err = Client(ctx).KeyManagementApi.DeleteAPIKey(ctx, createAPIKeyValue).Execute()
	if err != nil {
		t.Errorf("Error deleting api key %v: Response %s: %v", createAPIKeyValue, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	deleteAPIKeyReturned := apiKeyData.GetApiKey()
	deleteAPIKeyName := deleteAPIKeyReturned.GetName()
	deleteAPIKeyCreated := deleteAPIKeyReturned.GetCreated()
	deleteAPIKeyCreatedBy := deleteAPIKeyReturned.GetCreatedBy()
	deleteAPIKeyValue := deleteAPIKeyReturned.GetKey()

	// should return the key thats been deleted
	assert.Equal(editAPIKeyName, deleteAPIKeyName)
	assert.Equal(editAPIKeyCreated, deleteAPIKeyCreated)
	assert.Equal(editAPIKeyCreatedBy, deleteAPIKeyCreatedBy)
	assert.Equal(editAPIKeyValue, deleteAPIKeyValue)
}

func TestApplicationKeyFunctions(t *testing.T) {
	if !tests.IsRecording() {
		t.Skip("This test case does not support reply from recording")
	}

	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// TODO remove when API keys support non secret primary identifiers
	// NOTE do not forget to replace time -> c.Clock
	Client(ctx).GetConfig().HTTPClient = http.DefaultClient

	// Create Application Key
	// ----------------------------------
	testAppKeyName := fmt.Sprintf("%s:%d", t.Name(), time.Now().UnixNano())
	appKeyData, httpresp, err := Client(ctx).KeyManagementApi.CreateApplicationKey(ctx).Body(datadog.ApplicationKey{Name: &testAppKeyName}).Execute()
	if err != nil {
		t.Errorf("Error creating api key %v: Response %s: %v", testAppKeyName, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteAppKey(ctx, appKeyData.ApplicationKey.GetHash())
	assert.Equal(200, httpresp.StatusCode)

	createAppKeyReturned := appKeyData.GetApplicationKey()
	createAppKeyOwner := createAppKeyReturned.GetOwner()
	createAppKeyHash := createAppKeyReturned.GetHash()
	createAppKeyName := createAppKeyReturned.GetName()

	// all values should not be nil
	assert.NotEmpty(createAppKeyOwner)
	assert.NotEmpty(createAppKeyHash)
	assert.NotEmpty(createAppKeyName)
	assert.Equal(testAppKeyName, createAppKeyName)

	// Get Application Key
	// ----------------------------------
	appKeyData, httpresp, err = Client(ctx).KeyManagementApi.GetApplicationKey(ctx, createAppKeyHash).Execute()
	if err != nil {
		t.Errorf("Error getting app key %v: Response %s: %v", createAppKeyHash, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	getAppKeyReturned := appKeyData.GetApplicationKey()
	getAppKeyOwner := getAppKeyReturned.GetOwner()
	getAppKeyHash := getAppKeyReturned.GetHash()
	getAppKeyName := getAppKeyReturned.GetName()

	// should be the same as the create
	assert.Equal(createAppKeyOwner, getAppKeyOwner)
	assert.Equal(createAppKeyHash, getAppKeyHash)
	assert.Equal(createAppKeyName, getAppKeyName)

	// Get All Application Keys
	// ----------------------------------
	respListData, httpresp, err := Client(ctx).KeyManagementApi.ListApplicationKeys(ctx).Execute()
	if err != nil {
		t.Errorf("Error getting all app keys: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	getAllAppKeyReturned := respListData.GetApplicationKeys()

	// should have more than one at least
	assert.True(len(getAllAppKeyReturned) > 1)

	// Edit Application Key
	// ----------------------------------
	newAppKeyName := fmt.Sprintf("New %s", testAppKeyName)
	appKeyData, httpresp, err = Client(ctx).KeyManagementApi.UpdateApplicationKey(ctx, getAppKeyHash).Body(datadog.ApplicationKey{Name: &newAppKeyName}).Execute()
	if err != nil {
		t.Errorf("Error editing app key %v: Response %s: %v", getAppKeyHash, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	editAppKeyReturned := appKeyData.GetApplicationKey()
	editAppKeyOwner := editAppKeyReturned.GetOwner()
	editAppKeyHash := editAppKeyReturned.GetHash()
	editAppKeyName := editAppKeyReturned.GetName()

	// everything should be the same except the name
	assert.NotEqual(getAppKeyName, editAppKeyName)
	assert.Equal(getAppKeyOwner, editAppKeyOwner)
	assert.Equal(getAppKeyHash, editAppKeyHash)

	// Delete Application Key
	// ----------------------------------
	appKeyData, httpresp, err = Client(ctx).KeyManagementApi.DeleteApplicationKey(ctx, createAppKeyHash).Execute()
	if err != nil {
		t.Errorf("Error deleting app key %v: Response %s: %v", createAppKeyHash, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	deleteAppKeyReturned := appKeyData.GetApplicationKey()
	deleteAppKeyOwner := deleteAppKeyReturned.GetOwner()
	deleteAppKeyHash := deleteAppKeyReturned.GetHash()
	deleteAppKeyName := deleteAppKeyReturned.GetName()

	// should return the deleted app key
	assert.Equal(editAppKeyOwner, deleteAppKeyOwner)
	assert.Equal(editAppKeyHash, deleteAppKeyHash)
	assert.Equal(editAppKeyName, deleteAppKeyName)
}

func TestAPIKeysMgmtListErrors(t *testing.T) {
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

			_, httpresp, err := Client(ctx).KeyManagementApi.ListAPIKeys(ctx).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAPIKeysMgmtCreateErrors(t *testing.T) {

	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.ApiKey
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.ApiKey{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.ApiKey{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).KeyManagementApi.CreateAPIKey(ctx).Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAPIKeysMgmtGetErrors(t *testing.T) {

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

			_, httpresp, err := Client(ctx).KeyManagementApi.GetAPIKey(ctx, "whatever").Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAPIKeysMgmtUpdateErrors(t *testing.T) {

	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	name := "nonexistent key"
	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.ApiKey
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.ApiKey{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.ApiKey{}, 403},
		"404 Not Found":   {WithTestAuth, datadog.ApiKey{Name: &name}, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).KeyManagementApi.UpdateAPIKey(ctx, "whatever").Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAPIKeysMgmtDelete400Error(t *testing.T) {

	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	res, err := tests.ReadFixture("fixtures/key-mgmt/invalid_number_of_keys_400.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because we need 0 API keys to trigger the 400
	gock.New("https://api.datadoghq.com").Delete("/api/v1/api_key/whatever").Reply(400).JSON(res)
	defer gock.Off()

	_, httpresp, err := Client(ctx).KeyManagementApi.DeleteAPIKey(ctx, "whatever").Execute()
	assert.Equal(400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestAPIKeysMgmtDeleteErrors(t *testing.T) {
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

			_, httpresp, err := Client(ctx).KeyManagementApi.DeleteAPIKey(ctx, "whatever").Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAppKeysMgmtListErrors(t *testing.T) {
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

			_, httpresp, err := Client(ctx).KeyManagementApi.ListApplicationKeys(ctx).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAppKeysMgmtCreateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.ApplicationKey
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.ApplicationKey{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.ApplicationKey{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).KeyManagementApi.CreateApplicationKey(ctx).Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAppKeysMgmtCreate409Error(t *testing.T) {
	if !tests.IsRecording() {
		t.Skip("This test case does not support reply from recording")
	}

	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Create an app key to trigger the 409 conflict
	testAPPKeyName := fmt.Sprintf("%s:%d", t.Name(), time.Now().UnixNano())
	appKeyData, httpresp, err := Client(ctx).KeyManagementApi.CreateApplicationKey(ctx).Body(datadog.ApplicationKey{Name: &testAPPKeyName}).Execute()
	if err != nil {
		t.Fatalf("Error creating api key %v: Response %s: %v", testAPPKeyName, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteAppKey(ctx, appKeyData.ApplicationKey.GetHash())
	assert.Equal(200, httpresp.StatusCode)

	_, httpresp, err = Client(ctx).KeyManagementApi.CreateApplicationKey(ctx).Body(datadog.ApplicationKey{Name: &testAPPKeyName}).Execute()
	assert.Equal(409, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestAppKeysMgmtGetErrors(t *testing.T) {
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

			_, httpresp, err := Client(ctx).KeyManagementApi.GetApplicationKey(ctx, "whatever").Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAppKeysMgmtUpdateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	name := "nonexistent key"
	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.ApplicationKey
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.ApplicationKey{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.ApplicationKey{}, 403},
		"404 Not Found":   {WithTestAuth, datadog.ApplicationKey{Name: &name}, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).KeyManagementApi.UpdateApplicationKey(ctx, "whatever").Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAppKeysMgmtUpdate409Error(t *testing.T) {
	if !tests.IsRecording() {
		t.Skip("This test case does not support reply from recording")
	}

	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Create two app keys to trigger the 409 conflict
	testAPPKeyName1 := fmt.Sprintf("%s:%d", t.Name(), time.Now().UnixNano())
	appKeyData1, httpresp, err := Client(ctx).KeyManagementApi.CreateApplicationKey(ctx).Body(datadog.ApplicationKey{Name: &testAPPKeyName1}).Execute()
	if err != nil {
		t.Errorf("Error creating app key %v: Response %s: %v", testAPPKeyName1, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteAppKey(ctx, appKeyData1.ApplicationKey.GetHash())
	assert.Equal(200, httpresp.StatusCode)

	testAPPKeyName2 := fmt.Sprintf("%s:%d2", t.Name(), time.Now().UnixNano())
	appKeyData2, httpresp, err := Client(ctx).KeyManagementApi.CreateApplicationKey(ctx).Body(datadog.ApplicationKey{Name: &testAPPKeyName2}).Execute()
	if err != nil {
		t.Errorf("Error creating app key %v: Response %s: %v", testAPPKeyName2, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteAppKey(ctx, appKeyData2.ApplicationKey.GetHash())
	assert.Equal(200, httpresp.StatusCode)

	_, httpresp, err = Client(ctx).KeyManagementApi.UpdateApplicationKey(ctx, appKeyData1.ApplicationKey.GetHash()).Body(datadog.ApplicationKey{Name: &testAPPKeyName2}).Execute()
	assert.Equal(409, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestAppKeysMgmtDeleteErrors(t *testing.T) {
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

			_, httpresp, err := Client(ctx).KeyManagementApi.DeleteApplicationKey(ctx, "whatever").Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func deleteAPIKey(ctx context.Context, apiKeyValue string) {
	_, httpresp, err := Client(ctx).KeyManagementApi.DeleteAPIKey(ctx, apiKeyValue).Execute()
	if httpresp.StatusCode != 200 || err != nil {
		log.Printf("Deleting api key: %v failed with %v, Another test may have already deleted this api key.", apiKeyValue[len(apiKeyValue)-4:], httpresp.StatusCode)
	}
}

func deleteAppKey(ctx context.Context, appKeyHash string) {
	_, httpresp, err := Client(ctx).KeyManagementApi.DeleteApplicationKey(ctx, appKeyHash).Execute()
	if httpresp.StatusCode != 200 || err != nil {
		log.Printf("Deleting app key: %v failed with %v, Another test may have already deleted this app key.", appKeyHash[len(appKeyHash)-4:], httpresp.StatusCode)
	}
}
