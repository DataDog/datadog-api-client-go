/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"fmt"
	"gopkg.in/h2non/gock.v1"
	"net/http"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/DataDog/datadog-api-client-go/v2/tests"
)

func TestApiKeyFunctions(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	if tests.GetRecording() == tests.ModeReplaying {
		t.Skip("This test case does not support reply from recording")
	}

	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewKeyManagementApi(Client(ctx))

	// TODO remove when API keys support non secret primary identifiers
	// NOTE do not forget to replace time -> c.Clock
	Client(ctx).GetConfig().HTTPClient = http.DefaultClient

	// Create API Key
	// ----------------------------------
	testAPIKeyName := *tests.UniqueEntityName(ctx, t)
	apiKeyData, httpresp, err := api.CreateAPIKey(ctx, datadogV1.ApiKey{Name: &testAPIKeyName})
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
	apiKeyData, httpresp, err = api.GetAPIKey(ctx, createAPIKeyValue)
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
	respListData, httpresp, err := api.ListAPIKeys(ctx)
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
	apiKeyData, httpresp, err = api.UpdateAPIKey(ctx, createAPIKeyValue, datadogV1.ApiKey{Name: &newAPIKeyName})
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
	apiKeyData, httpresp, err = api.DeleteAPIKey(ctx, createAPIKeyValue)
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
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadogV1.ApiKey
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadogV1.ApiKey{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadogV1.ApiKey{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewKeyManagementApi(Client(ctx))

			_, httpresp, err := api.CreateAPIKey(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAPIKeysMgmtGetErrors(t *testing.T) {
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
			api := datadogV1.NewKeyManagementApi(Client(ctx))

			_, httpresp, err := api.GetAPIKey(ctx, "whatever")
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAPIKeysMgmtUpdateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	name := "nonexistent key"
	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadogV1.ApiKey
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadogV1.ApiKey{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadogV1.ApiKey{}, 403},
		"404 Not Found":   {WithTestAuth, datadogV1.ApiKey{Name: &name}, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewKeyManagementApi(Client(ctx))

			_, httpresp, err := api.UpdateAPIKey(ctx, "whatever", tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAPIKeysMgmtDelete400Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewKeyManagementApi(Client(ctx))

	res, err := tests.ReadFixture("fixtures/key-mgmt/invalid_number_of_keys_400.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because we need 0 API keys to trigger the 400
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "TODO")
	assert.NoError(err)
	gock.New(URL).Delete("/api/v1/api_key/whatever").Reply(400).JSON(res)
	defer gock.Off()

	_, httpresp, err := api.DeleteAPIKey(ctx, "whatever")
	assert.Equal(400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestAPIKeysMgmtDeleteErrors(t *testing.T) {
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
			api := datadogV1.NewKeyManagementApi(Client(ctx))

			_, httpresp, err := api.DeleteAPIKey(ctx, "whatever")
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAppKeysMgmtListErrors(t *testing.T) {
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
			api := datadogV1.NewKeyManagementApi(Client(ctx))

			_, httpresp, err := api.ListApplicationKeys(ctx)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAppKeysMgmtCreateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadogV1.ApplicationKey
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadogV1.ApplicationKey{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadogV1.ApplicationKey{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewKeyManagementApi(Client(ctx))

			_, httpresp, err := api.CreateApplicationKey(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAppKeysMgmtGetErrors(t *testing.T) {
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
			api := datadogV1.NewKeyManagementApi(Client(ctx))

			_, httpresp, err := api.GetApplicationKey(ctx, "whatever")
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAppKeysMgmtUpdateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	name := "nonexistent key"
	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadogV1.ApplicationKey
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadogV1.ApplicationKey{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadogV1.ApplicationKey{}, 403},
		"404 Not Found":   {WithTestAuth, datadogV1.ApplicationKey{Name: &name}, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewKeyManagementApi(Client(ctx))

			_, httpresp, err := api.UpdateApplicationKey(ctx, "whatever", tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAppKeysMgmtDeleteErrors(t *testing.T) {
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
			api := datadogV1.NewKeyManagementApi(Client(ctx))

			_, httpresp, err := api.DeleteApplicationKey(ctx, "whatever")
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func deleteAPIKey(ctx context.Context, t *testing.T, apiKeyValue string) {
	api := datadogV1.NewKeyManagementApi(Client(ctx))

	_, httpresp, err := api.DeleteAPIKey(ctx, apiKeyValue)
	if httpresp.StatusCode != 200 || err != nil {
		t.Logf("Deleting api key: %v failed with %v, Another test may have already deleted this api key.", apiKeyValue[len(apiKeyValue)-4:], httpresp.StatusCode)
	}
}
