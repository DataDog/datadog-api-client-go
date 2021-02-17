package test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/DataDog/datadog-api-client-go/tests"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gopkg.in/h2non/gock.v1"
	is "gotest.tools/assert/cmp"
)

const (
	staticTeamName    = "slack_integration_test_team"
	staticChannelName = "#test-channel"
)

func TestSlackIntegrationGetAllChannelsMocked(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	defer gock.Off()

	fixturePath, err := filepath.Abs("fixtures/slack-integration/get_all_channels.json")
	if err != nil {
		t.Errorf("Failed to get fixture file path: %s", err)
	}
	data, err := ioutil.ReadFile(fixturePath)
	if err != nil {
		t.Errorf("Failed to open fixture file: %s", err)
	}

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "SlackIntegrationApi.GetSlackIntegrationChannels")
	assert.NoError(err)
	gock.New(URL).
		Get(fmt.Sprintf("/api/v1/integration/slack/configuration/accounts/%s/channels", staticTeamName)).
		Reply(200).
		JSON(data)

	var expected []datadog.SlackIntegrationChannel
	json.Unmarshal([]byte(data), &expected)

	api := Client(ctx).SlackIntegrationApi
	slackChannelsResp, httpResp, err := api.GetSlackIntegrationChannels(ctx, staticTeamName).Execute()
	if err != nil {
		t.Errorf("Failed to get slack integration all channels: %v", err)
	}
	assert.Equal(200, httpResp.StatusCode)
	assert.Equal(3, len(slackChannelsResp))
	assert.True(is.DeepEqual(expected, slackChannelsResp)().Success())
}

func TestSlackIntegrationGetAllChannelsErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			uniqueTeamName := *tests.UniqueEntityName(ctx, t)
			_, httpresp, err := Client(ctx).SlackIntegrationApi.GetSlackIntegrationChannels(ctx, uniqueTeamName).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSlackIntegrationCreateChannelMocked(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	defer gock.Off()

	fixturePath, err := filepath.Abs("fixtures/slack-integration/create_channel.json")
	if err != nil {
		t.Errorf("Failed to get fixture file path: %s", err)
	}
	data, err := ioutil.ReadFile(fixturePath)
	if err != nil {
		t.Errorf("Failed to open fixture file: %s", err)
	}

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "SlackIntegrationApi.CreateSlackIntegrationChannel")
	assert.NoError(err)
	gock.New(URL).
		Post(fmt.Sprintf("/api/v1/integration/slack/configuration/accounts/%s/channels", staticTeamName)).
		Reply(200).
		JSON(data)

	var expected datadog.SlackIntegrationChannel
	json.Unmarshal([]byte(data), &expected)

	api := Client(ctx).SlackIntegrationApi

	channelPayload := datadog.NewSlackIntegrationChannel()
	channelPayload.SetName(staticChannelName)
	channelPayload.Display = &datadog.SlackIntegrationChannelDisplay{
		Message:  datadog.PtrBool(false),
		Notified: datadog.PtrBool(true),
		Snapshot: datadog.PtrBool(true),
		Tags:     datadog.PtrBool(false),
	}

	slackChannelsResp, httpResp, err := api.CreateSlackIntegrationChannel(ctx, staticTeamName).Body(*channelPayload).Execute()
	if err != nil {
		t.Errorf("Failed to create slack integration channel: %v", err)
	}
	assert.Equal(200, httpResp.StatusCode)
	assert.True(is.DeepEqual(expected, slackChannelsResp)().Success())
}

func TestSlackIntegrationCreateChannelErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.SlackIntegrationChannel
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.SlackIntegrationChannel{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.SlackIntegrationChannel{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			uniqueTeamName := *tests.UniqueEntityName(ctx, t)
			_, httpresp, err := Client(ctx).SlackIntegrationApi.CreateSlackIntegrationChannel(ctx, uniqueTeamName).Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSlackIntegrationCreateChannel404Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)

	data, err := tests.ReadFixture("fixtures/slack-integration/error_404.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because an existing slack account is needed for 404 response
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "SlackIntegrationApi.CreateSlackIntegrationChannel")
	assert.NoError(err)
	gock.New(URL).
		Post(fmt.Sprintf("/api/v1/integration/slack/configuration/accounts/%s/channels", staticTeamName)).
		Reply(404).JSON(data)
	defer gock.Off()

	_, httpresp, err := Client(ctx).SlackIntegrationApi.CreateSlackIntegrationChannel(ctx, staticTeamName).Body(datadog.SlackIntegrationChannel{}).Execute()
	assert.Equal(404, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestSlackIntegrationGetChannelMocked(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	defer gock.Off()

	fixturePath, err := filepath.Abs("fixtures/slack-integration/get_channel.json")
	if err != nil {
		t.Errorf("Failed to get fixture file path: %s", err)
	}
	data, err := ioutil.ReadFile(fixturePath)
	if err != nil {
		t.Errorf("Failed to open fixture file: %s", err)
	}

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "SlackIntegrationApi.GetSlackIntegrationChannel")
	assert.NoError(err)
	gock.New(URL).
		Get(fmt.Sprintf("/api/v1/integration/slack/configuration/accounts/%s/channels/%s", staticTeamName, staticChannelName)).
		Reply(200).
		JSON(data)

	var expected datadog.SlackIntegrationChannel
	json.Unmarshal([]byte(data), &expected)

	api := Client(ctx).SlackIntegrationApi
	slackChannelsResp, httpResp, err := api.GetSlackIntegrationChannel(ctx, staticTeamName, staticChannelName).Execute()
	if err != nil {
		t.Errorf("Failed to get slack integration channel: %v", err)
	}
	assert.Equal(200, httpResp.StatusCode)
	assert.True(is.DeepEqual(expected, slackChannelsResp)().Success())
}

func TestSlackIntegrationGetChannelErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			uniqueTeamName := *tests.UniqueEntityName(ctx, t)
			uniqueChannelName := *tests.UniqueEntityName(ctx, t)
			_, httpresp, err := Client(ctx).SlackIntegrationApi.GetSlackIntegrationChannel(ctx, uniqueTeamName, uniqueChannelName).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSlackIntegrationGetChannel404Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)

	data, err := tests.ReadFixture("fixtures/slack-integration/error_404.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because an existing slack account is needed for 404 response
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "SlackIntegrationApi.GetSlackIntegrationChannel")
	assert.NoError(err)
	gock.New(URL).
		Get(fmt.Sprintf("/api/v1/integration/slack/configuration/accounts/%s/channels/%s", staticTeamName, staticChannelName)).
		Reply(404).JSON(data)
	defer gock.Off()

	_, httpresp, err := Client(ctx).SlackIntegrationApi.GetSlackIntegrationChannel(ctx, staticTeamName, staticChannelName).Execute()
	assert.Equal(404, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestSlackIntegrationUpdateChannelMocked(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	defer gock.Off()

	fixturePath, err := filepath.Abs("fixtures/slack-integration/update_channel.json")
	if err != nil {
		t.Errorf("Failed to get fixture file path: %s", err)
	}
	data, err := ioutil.ReadFile(fixturePath)
	if err != nil {
		t.Errorf("Failed to open fixture file: %s", err)
	}

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "SlackIntegrationApi.UpdateSlackIntegrationChannel")
	assert.NoError(err)
	gock.New(URL).
		Patch(fmt.Sprintf("/api/v1/integration/slack/configuration/accounts/%s/channels/%s", staticTeamName, staticChannelName)).
		Reply(200).
		JSON(data)

	var expected datadog.SlackIntegrationChannel
	json.Unmarshal([]byte(data), &expected)

	api := Client(ctx).SlackIntegrationApi
	channelPayload := datadog.NewSlackIntegrationChannel()
	channelPayload.SetName(staticChannelName)
	channelPayload.Display = &datadog.SlackIntegrationChannelDisplay{
		Message:  datadog.PtrBool(false),
		Notified: datadog.PtrBool(false),
		Snapshot: datadog.PtrBool(false),
		Tags:     datadog.PtrBool(false),
	}

	slackChannelsResp, httpResp, err := api.UpdateSlackIntegrationChannel(ctx, staticTeamName, staticChannelName).Body(*channelPayload).Execute()
	if err != nil {
		t.Errorf("Failed to update slack integration channel: %v", err)
	}
	assert.Equal(200, httpResp.StatusCode)
	assert.True(is.DeepEqual(expected, slackChannelsResp)().Success())
}

func TestSlackIntegrationUpdateChannelErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.SlackIntegrationChannel
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.SlackIntegrationChannel{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.SlackIntegrationChannel{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			uniqueTeamName := *tests.UniqueEntityName(ctx, t)
			uniqueChannelName := *tests.UniqueEntityName(ctx, t)
			_, httpresp, err := Client(ctx).SlackIntegrationApi.UpdateSlackIntegrationChannel(ctx, uniqueTeamName, uniqueChannelName).Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSlackIntegrationUpdateChannel404Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)

	data, err := tests.ReadFixture("fixtures/slack-integration/error_404.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because an existing slack account is needed for 404 response
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "SlackIntegrationApi.UpdateSlackIntegrationChannel")
	assert.NoError(err)
	gock.New(URL).
		Patch(fmt.Sprintf("/api/v1/integration/slack/configuration/accounts/%s/channels/%s", staticTeamName, staticChannelName)).
		Reply(404).JSON(data)
	defer gock.Off()

	_, httpresp, err := Client(ctx).SlackIntegrationApi.UpdateSlackIntegrationChannel(ctx, staticTeamName, staticChannelName).Body(datadog.SlackIntegrationChannel{}).Execute()
	assert.Equal(404, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestSlackIntegrationRemoveChannelMocked(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	defer gock.Off()

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "SlackIntegrationApi.RemoveSlackIntegrationChannel")
	assert.NoError(err)
	gock.New(URL).
		Delete(fmt.Sprintf("/api/v1/integration/slack/configuration/accounts/%s/channels/%s", staticTeamName, staticChannelName)).
		Reply(204)

	api := Client(ctx).SlackIntegrationApi
	httpResp, err := api.RemoveSlackIntegrationChannel(ctx, staticTeamName, staticChannelName).Execute()
	if err != nil {
		t.Errorf("Failed to remove slack integration channel: %v", err)
	}
	assert.Equal(204, httpResp.StatusCode)
}

func TestSlackIntegrationRemoveChannelErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			uniqueTeamName := *tests.UniqueEntityName(ctx, t)
			uniqueChannelName := *tests.UniqueEntityName(ctx, t)
			httpresp, err := Client(ctx).SlackIntegrationApi.RemoveSlackIntegrationChannel(ctx, uniqueTeamName, uniqueChannelName).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError := err.(datadog.GenericOpenAPIError)
			assert.NotEmpty(apiError)
		})
	}
}
