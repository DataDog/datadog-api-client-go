package api

import (
	"context"
	"strings"
	"testing"

	datadogV1 "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	datadogV2 "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
	testV1 "github.com/DataDog/datadog-api-client-go/tests/api/v1/datadog"
	testV2 "github.com/DataDog/datadog-api-client-go/tests/api/v2/datadog"
	"gopkg.in/h2non/gock.v1"
)

func TestDeserializationUnkownNestedOneOfInList(t *testing.T) {
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

	gock.New(URL).
		Get("synthetics/tests/api/public_id").
		Reply(299).
		AddHeader("Content-Type", "application/json").
		Body(strings.NewReader(responseBody))
	defer gock.Off()

	resp, httpresp, err := testV1.Client(ctx).SyntheticsApi.GetAPITest(ctx, "public_id")
	assert.Nil(err)
	assert.Equal(299, httpresp.StatusCode)
	// Root object deserializes correctly
	assert.Nil(resp.UnparsedObject)
	// Assertions object has the 3 expected assertions
	assert.Len(resp.Config.GetAssertions(), 3)
	// Unknown assertion is Unparsed
	assert.Equal("A non existent operator", resp.Config.GetAssertions()[2].UnparsedObject.(map[string]interface{})["operator"])
}

func TestDeserializationUnkownNestedEnumInList(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = testV1.WithClient(testV1.WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)

	responseBody := `
{
    "status": "live",
    "public_id": "2fx-64b-fb8",
    "tags": [
        "mini-website",
        "team:synthetics",
        "firefox",
        "synthetics-ci-browser",
        "edge",
        "chrome"
    ],
    "locations": [
        "aws:ap-northeast-1",
        "aws:eu-north-1",
        "aws:eu-west-3",
        "aws:eu-central-1"
    ],
    "message": "This mini-website check failed, please investigate why. @slack-synthetics-ops-worker",
    "name": "Mini Website - Click Trap",
    "monitor_id": 7647262,
    "type": "browser",
    "created_at": "2018-12-20T13:19:23.734004+00:00",
    "modified_at": "2021-06-30T15:46:49.387631+00:00",
    "config": {
        "variables": [],
        "setCookie": "",
        "request": {
            "url": "http://34.95.79.70/click-trap",
            "headers": {},
            "method": "GET"
        },
        "assertions": [],
        "configVariables": []
    },
    "options": {
        "ci": {
            "executionRule": "blocking"
        },
        "retry": {
            "count": 1,
            "interval": 1000
        },
        "min_location_failed": 1,
        "min_failure_duration": 0,
        "noScreenshot": false,
        "tick_every": 300,
        "forwardProxy": false,
        "disableCors": false,
        "device_ids": [
            "chrome.laptop_large",
            "firefox.laptop_large",
            "A non existent device ID"
        ],
        "monitor_options": {
            "renotify_interval": 360
        },
        "ignoreServerCertificateError": true
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

	resp, httpresp, err := testV1.Client(ctx).SyntheticsApi.GetBrowserTest(ctx, "public_id")
	assert.Nil(err)
	assert.Equal(299, httpresp.StatusCode)
	// Root object deserializes correctly
	assert.Nil(resp.UnparsedObject)
	// Options object has the 3 expected device IDs
	assert.Len(resp.Options.GetDeviceIds(), 3)
	assert.Equal(datadogV1.SyntheticsDeviceID("A non existent device ID"), resp.Options.GetDeviceIds()[2])
}

func TestDeserializationUnkownTopLevelEnum(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = testV1.WithClient(testV1.WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)

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

	resp, httpresp, err := testV1.Client(ctx).SyntheticsApi.GetBrowserTest(ctx, "public_id")
	assert.Nil(err)
	assert.Equal(299, httpresp.StatusCode)
	// Root object is unparsed
	assert.NotNil(resp.UnparsedObject)
	assert.Equal("A non existent test type", resp.UnparsedObject["type"])
	assert.Equal("Check on www.10.0.0.1.xip.io", resp.UnparsedObject["name"])
}

func TestDeserializationUnkownNestedEnum(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = testV1.WithClient(testV1.WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)

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
            "method": "A non existent method",
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
	URL, err := testV1.Client(ctx).GetConfig().ServerURLWithContext(ctx, "SyntheticsApiService.GetAPITest")
	assert.NoError(err)

	gock.New(URL).
		Get("synthetics/tests/api/public_id").
		Reply(299).
		AddHeader("Content-Type", "application/json").
		Body(strings.NewReader(responseBody))
	defer gock.Off()

	resp, httpresp, err := testV1.Client(ctx).SyntheticsApi.GetAPITest(ctx, "public_id")
	assert.Nil(err)
	assert.Equal(299, httpresp.StatusCode)
	// Root object is properly deserialized
	assert.Nil(resp.UnparsedObject)
	assert.Nil(resp.Config.UnparsedObject)
	// Direct parent of invalid enum is unparsed
	assert.NotNil(resp.Config.Request.UnparsedObject)
	assert.Equal("A non existent method", resp.Config.Request.UnparsedObject["method"])
	assert.Equal(float64(30), resp.Config.Request.UnparsedObject["timeout"])
}

func TestDeserializationUnkownNestedOneOf(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = testV2.WithClient(testV2.WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)

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

	resp, httpresp, err := testV2.Client(ctx).LogsArchivesApi.CreateLogsArchive(ctx, datadogV2.LogsArchiveCreateRequest{})
	assert.Nil(err)
	assert.Equal(299, httpresp.StatusCode)
	// Root object is properly deserialized
	assert.Nil(resp.UnparsedObject)
	assert.Nil(resp.Data.Attributes.UnparsedObject)
	// OneOf is unparsed
	assert.NotNil(resp.Data.Attributes.Destination.Get().UnparsedObject)
	assert.Equal("A non existent destination", resp.Data.Attributes.Destination.Get().UnparsedObject.(map[string]interface{})["type"])
}
