package api

import (
	"context"
	"strings"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/DataDog/datadog-api-client-go/v2/tests"
	testV1 "github.com/DataDog/datadog-api-client-go/v2/tests/api/datadogV1"
	testV2 "github.com/DataDog/datadog-api-client-go/v2/tests/api/datadogV2"
	"gopkg.in/h2non/gock.v1"
)

func TestDeserializationUnknownNestedOneOfInList(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = testV1.WithClient(testV1.WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)

	responseBody := `
{
    "status": "paused",
    "public_id": "jv7-wfd-kvt",
    "tags": [],
    "locations": [
        "pl:pl-kevin-y-6382df0d72d4588e1817f090b131541f"
    ],
    "message": "",
    "name": "Test on www.example.com",
    "monitor_id": 28558768,
    "type": "api",
    "created_at": "2021-01-12T10:11:40.802074+00:00",
    "modified_at": "2021-01-22T16:42:10.520384+00:00",
    "subtype": "http",
    "config": {
        "request": {
            "url": "https://www.example.com",
            "method": "GET",
            "timeout": 30
        },
        "assertions": [
            {
                "operator": "lessThan",
                "type": "responseTime",
                "target": 1000
            },
            {
                "operator": "is",
                "type": "statusCode",
                "target": 200
            },
            {
                "operator": "A non existent operator",
                "type": "body",
                "target": {
                    "xPath": "//html/head/title",
                    "operator": "contains",
                    "targetValue": "Example"
                }
            }
        ],
        "configVariables": []
    },
    "options": {
        "monitor_options": {
            "notify_audit": false,
            "locked": false,
            "include_tags": true,
            "new_host_delay": 300,
            "notify_no_data": false,
            "renotify_interval": 0
        },
        "retry": {
            "count": 0,
            "interval": 300
        },
        "min_location_failed": 1,
        "min_failure_duration": 0,
        "tick_every": 60
    }
}
`
	// Mock the synthetics API.
	URL, err := testV1.Client(ctx).GetConfig().ServerURLWithContext(ctx, "SyntheticsApiService.GetAPITest")
	assert.NoError(err)
	api := datadogV1.NewSyntheticsApi(testV1.Client(ctx))

	gock.New(URL).
		Get("synthetics/tests/api/public_id").
		Reply(299).
		AddHeader("Content-Type", "application/json").
		Body(strings.NewReader(responseBody))
	defer gock.Off()

	resp, httpresp, err := api.GetAPITest(ctx, "public_id")
	assert.Nil(err)
	assert.Equal(299, httpresp.StatusCode)
	// Root object deserializes correctly
	assert.Nil(resp.UnparsedObject)
	// Assertions object has the 3 expected assertions
	assert.Len(resp.Config.GetAssertions(), 3)
	// Unknown assertion is Unparsed
	assert.Equal("A non existent operator", resp.Config.GetAssertions()[2].UnparsedObject.(map[string]interface{})["operator"])
	assert.True(datadog.ContainsUnparsedObject(resp))
}

func TestDeserializationUnknownNestedEnumInList(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = testV2.WithClient(testV2.WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	api := datadogV2.NewDowntimesApi(testV2.Client(ctx))

	responseBody := `
{
    "data": {
        "type": "downtime",
        "attributes": {
            "mute_first_recovery_notification": false,
            "schedule": {
                "end": null,
                "start": "2024-11-10T03:12:35.223223+00:00"
            },
            "notify_end_types": [
                "expired"
            ],
            "status": "active",
            "monitor_identifier": {
                "monitor_tags": [
                    "*"
                ]
            },
            "display_timezone": "UTC",
            "notify_end_states": [
                "warn",
                "alert",
                "not an end state"
            ],
            "created": "2024-11-10T03:12:35.241213+00:00",
            "modified": "2024-11-10T03:12:35.241213+00:00",
            "canceled": null,
            "message": null,
            "scope": "host:java-hostsMuteErrorsTest-local-1731208355"
        },
        "id": "a5546ef7-fea3-4a1b-b82e-04f8067f655a"
    }
}
`
	// Mock the synthetics API.
	URL, err := testV2.Client(ctx).GetConfig().ServerURLWithContext(ctx, "DowntimeApiService.GetDowntime")
	assert.NoError(err)

	gock.New(URL).
		Get("downtime/downtime_id").
		Reply(299).
		AddHeader("Content-Type", "application/json").
		Body(strings.NewReader(responseBody))
	defer gock.Off()

	resp, httpresp, err := api.GetDowntime(ctx, "downtime_id")
	assert.Nil(err)
	assert.Equal(299, httpresp.StatusCode)
	// Root object deserializes correctly
	assert.Nil(resp.UnparsedObject)
	// Options object has the 3 expected device IDs
	assert.Len(resp.Data.Attributes.GetNotifyEndStates(), 3)
	assert.Equal(datadogV2.DowntimeNotifyEndStateTypes("not an end state"), resp.Data.Attributes.GetNotifyEndStates()[2])
	assert.True(datadog.ContainsUnparsedObject(resp))
}

func TestDeserializationUnkownTopLevelEnum(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = testV1.WithClient(testV1.WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewSyntheticsApi(testV1.Client(ctx))

	responseBody := `
{
    "status": "live",
    "public_id": "g6d-gcm-pdq",
    "tags": [],
    "locations": [
        "aws:eu-central-1",
        "aws:ap-northeast-1"
    ],
    "message": "",
    "name": "Check on www.10.0.0.1.xip.io",
    "monitor_id": 7464050,
    "type": "A non existent test type",
    "created_at": "2018-12-07T17:30:49.785089+00:00",
    "modified_at": "2019-09-04T17:01:09.921070+00:00",
    "subtype": "http",
    "config": {
        "request": {
            "url": "https://www.10.0.0.1.xip.io",
            "method": "GET",
            "timeout": 30
        },
        "assertions": [
            {
                "operator": "is",
                "type": "statusCode",
                "target": 200
            }
        ]
    },
    "options": {
        "tick_every": 60
    }
}
`
	// Mock the synthetics API.
	URL, err := testV1.Client(ctx).GetConfig().ServerURLWithContext(ctx, "SyntheticsApiService.GetBrowserTest")
	assert.NoError(err)

	gock.New(URL).
		Get("synthetics/tests/browser/public_id").
		Reply(299).
		AddHeader("Content-Type", "application/json").
		Body(strings.NewReader(responseBody))
	defer gock.Off()

	resp, httpresp, err := api.GetBrowserTest(ctx, "public_id")
	assert.Nil(err)
	assert.Equal(299, httpresp.StatusCode)
	// Root object is unparsed
	assert.NotNil(resp.UnparsedObject)
	assert.Equal("A non existent test type", resp.UnparsedObject["type"])
	assert.Equal("Check on www.10.0.0.1.xip.io", resp.UnparsedObject["name"])
	assert.True(datadog.ContainsUnparsedObject(resp))
}

func TestDeserializationUnkownNestedEnum(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = testV1.WithClient(testV1.WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewSyntheticsApi(testV1.Client(ctx))

	responseBody := `
{
    "status": "live",
    "public_id": "g6d-gcm-pdq",
    "tags": [],
    "locations": [
        "aws:eu-central-1",
        "aws:ap-northeast-1"
    ],
    "message": "",
    "name": "Check on www.10.0.0.1.xip.io",
    "monitor_id": 7464050,
    "type": "api",
    "created_at": "2018-12-07T17:30:49.785089+00:00",
    "modified_at": "2019-09-04T17:01:09.921070+00:00",
    "subtype": "http",
    "config": {
        "request": {
            "url": "https://www.10.0.0.1.xip.io",
            "method": "GET",
            "timeout": 30
        },
        "assertions": [
            {
                "operator": "not-an-operator",
                "type": "statusCode",
                "target": 200
            }
        ]
    },
    "options": {
        "tick_every": 60
    }
}
`
	// Mock the synthetics API.
	URL, err := testV1.Client(ctx).GetConfig().ServerURLWithContext(ctx, "SyntheticsApiService.GetAPITest")
	assert.NoError(err)

	gock.New(URL).
		Get("synthetics/tests/api/public_id").
		Reply(299).
		AddHeader("Content-Type", "application/json").
		Body(strings.NewReader(responseBody))
	defer gock.Off()

	resp, httpresp, err := api.GetAPITest(ctx, "public_id")
	assert.Nil(err)
	assert.Equal(299, httpresp.StatusCode)
	// Direct parent of invalid enum is unparsed
	assert.NotNil(resp.Config.Assertions[0].UnparsedObject)
	assert.Equal("not-an-operator", resp.Config.Assertions[0].UnparsedObject.(map[string]interface{})["operator"])
	assert.Equal(float64(200), resp.Config.Assertions[0].UnparsedObject.(map[string]interface{})["target"])
	assert.True(datadog.ContainsUnparsedObject(resp))
}

func TestDeserializationUnknownNestedOneOf(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = testV2.WithClient(testV2.WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	api := datadogV2.NewLogsArchivesApi(testV2.Client(ctx))

	responseBody := `
{
    "data": {
        "type": "archives",
        "id": "n_XDSxVpScepiBnyhysj_A",
        "attributes": {
            "name": "my first azure archive",
            "query": "service:toto",
            "state": "UNKNOWN",
            "destination": {
                "container": "my-container",
                "storage_account": "storageaccount",
                "path": "/path/blou",
                "type": "A non existent destination",
                "integration": {
                    "tenant_id": "tf-TestAccDatadogLogsArchiveAzure_basic-local-1624981538",
                    "client_id": "testc7f6-1234-5678-9101-3fcbf464test"
                }
            },
            "rehydration_tags": [],
            "include_tags": false
        }
    }
}
`
	// Mock the synthetics API.
	URL, err := testV2.Client(ctx).GetConfig().ServerURLWithContext(ctx, "LogsArchivesApiService.CreateLogsArchive")
	assert.NoError(err)

	gock.New(URL).
		Post("logs/config/archives").
		Reply(299).
		AddHeader("Content-Type", "application/json").
		Body(strings.NewReader(responseBody))
	defer gock.Off()

	resp, httpresp, err := api.CreateLogsArchive(ctx, datadogV2.LogsArchiveCreateRequest{})
	assert.Nil(err)
	assert.Equal(299, httpresp.StatusCode)
	// Root object is properly deserialized
	assert.Nil(resp.UnparsedObject)
	assert.Nil(resp.Data.Attributes.UnparsedObject)
	// OneOf is unparsed
	assert.NotNil(resp.Data.Attributes.Destination.Get().UnparsedObject)
	assert.Equal("A non existent destination", resp.Data.Attributes.Destination.Get().UnparsedObject.(map[string]interface{})["type"])
	assert.True(datadog.ContainsUnparsedObject(resp))
}
