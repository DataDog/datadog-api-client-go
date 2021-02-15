/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"

	"gopkg.in/h2non/gock.v1"
)

func TestApiKeyFunctions(t *testing.T) {
	if tests.GetRecording() == tests.ModeReplaying {
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
	testAPIKeyName := *tests.UniqueEntityName(ctx, t)
	apiKeyData, httpresp, err := Client(ctx).KeyManagementApi.CreateAPIKey(ctx).Body(datadog.ApiKey{Name: &testAPIKeyName}).Execute()
	if err != nil {
		t.Errorf("Error creating api key %v: Response %s: %v", testAPIKeyName, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteAPIKey(ctx, t, apiKeyData.ApiKey.GetKey())
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
	newAPIKeyName := fmt.Sprintf("%s-new", testAPIKeyName)
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
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "TODO")
	assert.NoError(err)
	gock.New(URL).Delete("/api/v1/api_key/whatever").Reply(400).JSON(res)
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

func deleteAPIKey(ctx context.Context, t *testing.T, apiKeyValue string) {
	_, httpresp, err := Client(ctx).KeyManagementApi.DeleteAPIKey(ctx, apiKeyValue).Execute()
	if httpresp.StatusCode != 200 || err != nil {
		t.Logf("Deleting api key: %v failed with %v, Another test may have already deleted this api key.", apiKeyValue[len(apiKeyValue)-4:], httpresp.StatusCode)
	}
}

func deleteAppKey(ctx context.Context, t *testing.T, appKeyHash string) {
	_, httpresp, err := Client(ctx).KeyManagementApi.DeleteApplicationKey(ctx, appKeyHash).Execute()
	if httpresp.StatusCode != 200 || err != nil {
		t.Logf("Deleting app key: %v failed with %v, Another test may have already deleted this app key.", appKeyHash[len(appKeyHash)-4:], httpresp.StatusCode)
	}
}
