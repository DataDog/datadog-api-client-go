package test

/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

import (
	"context"
	"fmt"
	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"reflect"
	"testing"
)

// This test uses mocking because: 1) it relies on private data. 2) It relies on external services

func TestLogsArchivesCreateS3(t *testing.T) {
	archiveType := "s3"
	action := "create"
	testLogArchivesCreate(t, archiveType, action)
}

func TestLogsArchivesCreateAzure(t *testing.T) {
	archiveType := "azure"
	action := "create"
	testLogArchivesCreate(t, archiveType, action)
}

func TestLogsArchivesCreateGCS(t *testing.T) {
	archiveType := "gcs"
	action := "create"
	testLogArchivesCreate(t, archiveType, action)
}

func testLogArchivesCreate(t *testing.T, archiveType string, action string) {
	if tests.GetRecording() == tests.ModeReplaying {
		t.Skip("This test case does not support reply from recording")
	}
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	client := Client(ctx)
	assert := tests.Assert(ctx, t)
	defer finish()
	inputArchive := datadog.NullableLogsArchiveCreateRequest{}
	inputArchive.UnmarshalJSON([]byte(readFixture(t, fmt.Sprintf("fixtures/logs/archives/%s/in/%s.json", archiveType, action))))
	outputArchive := datadog.NullableLogsArchive{}
	outputArchive.UnmarshalJSON([]byte(readFixture(t, fmt.Sprintf("fixtures/logs/archives/%s/out/%s.json", archiveType, action))))
	URL, err := client.GetConfig().ServerURLWithContext(ctx, fmt.Sprintf("LogsArchive.%s.%s", action, archiveType))
	assert.NoError(err)
	gock.New(URL).Post("/api/v2/logs/config/archives").JSON(inputArchive).Reply(200).JSON(outputArchive)
	archivesAreEqual(t, *inputArchive.Get().Data.Attributes, *outputArchive.Get().Data.Attributes)
	assert.Equal(outputArchive.Get().Data.Type, "archives")
	r := reflect.ValueOf(outputArchive.Get().Data.Attributes.Destination.Get().GetActualInstance())
	assert.Equal(reflect.Indirect(r).FieldByName("Type").String(), archiveType)
	result, httpresp, err := client.LogsArchivesApi.CreateLogsArchive(ctx).Body(*inputArchive.Get()).Execute()
	assert.Equal(httpresp.StatusCode, 200)
	assert.Equal(result, *outputArchive.Get())
	defer gock.Off()
}

func TestLogsArchivesGetByID(t *testing.T) {
	if tests.GetRecording() == tests.ModeReplaying {
		t.Skip("This test case does not support reply from recording")
	}
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	client := Client(ctx)
	id := "XVlBzgbaiC"
	action := "getbyid"
	archiveType := "s3"
	outputArchive := datadog.NullableLogsArchive{}
	outputArchive.UnmarshalJSON([]byte(readFixture(t, fmt.Sprintf("fixtures/logs/archives/%s/out/%s.json", archiveType, action))))
	URL, err := client.GetConfig().ServerURLWithContext(ctx, fmt.Sprintf("LogsArchive.%s.%s", action, archiveType))
	assert.NoError(err)
	gock.New(URL).Get(fmt.Sprintf("/api/v2/logs/config/archives/%s", id)).Reply(200).JSON(outputArchive)
	result, httpresp, err := client.LogsArchivesApi.GetLogsArchive(ctx, id).Execute()
	assert.Equal(httpresp.StatusCode, 200)
	checkS3Archive(t, *outputArchive.Get().Data)
	assert.Equal(result, *outputArchive.Get())
	defer gock.Off()
}

func TestLogsArchivesDelete(t *testing.T) {
	if tests.GetRecording() == tests.ModeReplaying {
		t.Skip("This test case does not support reply from recording")
	}
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	id := "XVlBzgbaiC"
	action := "deleteById"
	archiveType := "s3"
	client := Client(ctx)
	URL, err := client.GetConfig().ServerURLWithContext(ctx, fmt.Sprintf("LogsArchive.%s.%s", action, archiveType))
	assert.NoError(err)
	gock.New(URL).Delete(fmt.Sprintf("/api/v2/logs/config/archives/%s", id)).Reply(204)
	httpresp, err := client.LogsArchivesApi.DeleteLogsArchive(ctx, id).Execute()
	assert.NoError(err)
	assert.Equal(httpresp.StatusCode, 204)
	defer gock.Off()
}

func TestLogsArchivesGetAll(t *testing.T) {
	if tests.GetRecording() == tests.ModeReplaying {
		t.Skip("This test case does not support reply from recording")
	}
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	action := "getall"
	archiveType := "s3"
	outputArchives := datadog.NullableLogsArchives{}
	outputArchives.UnmarshalJSON([]byte(readFixture(t, fmt.Sprintf("fixtures/logs/archives/%s/out/%s.json", archiveType, action))))
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, fmt.Sprintf("LogsArchive.%s.%s", action, archiveType))
	assert.NoError(err)
	gock.New(URL).Get("/api/v2/logs/config/archives/").Reply(200).JSON(outputArchives)
	assert.True(len(*outputArchives.Get().Data) > 0)
	checkS3Archive(t, (*outputArchives.Get().Data)[0])
	defer gock.Off()
}

func checkS3Archive(t *testing.T, outputArchive datadog.LogsArchiveDefinition) {
	assert.Equal(t, outputArchive.Type, "archives")
	destination := outputArchive.Attributes.Destination.Get().LogsArchiveDestinationS3
	assert.Equal(t, destination.Type, datadog.LOGSARCHIVEDESTINATIONS3TYPE_S3)
	assert.Equal(t, destination.Integration.AccountId, "711111111111")
	assert.Equal(t, destination.Integration.RoleName, "DatadogGoClientTestIntegrationRole")
	assert.Equal(t, destination.Path, datadog.PtrString("/path/blou"))
	assert.Equal(t, destination.Bucket, "dd-logs-test-datadog-api-client-go")
	assert.Equal(t, outputArchive.Attributes.Name, "datadog-api-client-go Tests Archive")
	assert.Equal(t, outputArchive.Attributes.Query, "source:tata")
	assert.Equal(t, outputArchive.Id, datadog.PtrString("XVlBzgbaiC"))
}

func archivesAreEqual(t *testing.T, archiveRequest datadog.LogsArchiveCreateRequestAttributes, createdArchive datadog.LogsArchiveAttributes) {
	assert.Equal(t, archiveRequest.Name, createdArchive.Name)
	assert.Equal(t, archiveRequest.Query, createdArchive.Query)
	s3DestinationReq := archiveRequest.Destination.LogsArchiveDestinationS3
	s3DestinationCreated := createdArchive.Destination.Get().LogsArchiveDestinationS3
	assert.Equal(t, s3DestinationReq, s3DestinationCreated)
	azureDestinationReq := archiveRequest.Destination.LogsArchiveDestinationAzure
	azureDestinationCreated := createdArchive.Destination.Get().LogsArchiveDestinationAzure
	assert.Equal(t, azureDestinationReq, azureDestinationCreated)
	gcsDestinationReq := archiveRequest.Destination.LogsArchiveDestinationGCS
	gcsDestinationCreated := createdArchive.Destination.Get().LogsArchiveDestinationGCS
	assert.Equal(t, gcsDestinationReq, gcsDestinationCreated)
}

func readFixture(t *testing.T, path string) string {
	res, err := tests.ReadFixture(path)
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	return res
}
