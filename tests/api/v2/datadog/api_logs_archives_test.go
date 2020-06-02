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
	"gopkg.in/h2non/gock.v1"
	"reflect"
	"testing"
)

// This test uses mocking because: 1) it relies on private data. 2) It relies on external services

func TestLogsArchivesCreateS3(t *testing.T) {
	cases := []struct {
		archiveType string
		action      string
	}{
		{
			archiveType: "s3",
			action:      "create",
		},
		{
			archiveType: "azure",
			action:      "create",
		},
		{
			archiveType: "gcs",
			action:      "create",
		},
	}
	if tests.GetRecording() == tests.ModeReplaying {
		t.Skip("This test case does not support reply from recording")
	}
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	defer gock.Off()
	client := Client(ctx)
	assert := tests.Assert(ctx, t)
	for _, c := range cases {
		inputArchive := datadog.NullableLogsArchiveCreateRequest{}
		inputArchive.UnmarshalJSON([]byte(readFixture(t, fmt.Sprintf("fixtures/logs/archives/%s/in/%s.json", c.archiveType, c.action))))
		outputArchive := datadog.NullableLogsArchive{}
		outputArchive.UnmarshalJSON([]byte(readFixture(t, fmt.Sprintf("fixtures/logs/archives/%s/out/%s.json", c.archiveType, c.action))))
		URL, err := client.GetConfig().ServerURLWithContext(ctx, fmt.Sprintf("LogsArchive.%s.%s", c.action, c.archiveType))
		assert.NoError(err)
		gock.New(URL).Post("/api/v2/logs/config/archives").JSON(inputArchive).Reply(200).JSON(outputArchive)
		archivesAreEqual(t, ctx, *inputArchive.Get().Data.Attributes, *outputArchive.Get().Data.Attributes)
		assert.Equal(outputArchive.Get().Data.Type, "archives")
		r := reflect.ValueOf(outputArchive.Get().Data.Attributes.Destination.Get().GetActualInstance())
		assert.Equal(reflect.Indirect(r).FieldByName("Type").String(), c.archiveType)
		result, httpresp, err := client.LogsArchivesApi.CreateLogsArchive(ctx).Body(*inputArchive.Get()).Execute()
		assert.Equal(httpresp.StatusCode, 200)
		assert.Equal(result, *outputArchive.Get())
	}
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
	checkS3Archive(t, ctx, *outputArchive.Get().Data)
	assert.Equal(result, *outputArchive.Get())
	defer gock.Off()
}

func TestLogsArchivesDelete(t *testing.T) {
	if tests.GetRecording() == tests.ModeReplaying {
		t.Skip("This test case does not support reply from recording")
	}
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	defer gock.Off()
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
}

func TestLogsArchivesGetAll(t *testing.T) {
	if tests.GetRecording() == tests.ModeReplaying {
		t.Skip("This test case does not support reply from recording")
	}
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	defer gock.Off()
	assert := tests.Assert(ctx, t)
	action := "getall"
	archiveType := "s3"
	outputArchives := datadog.NullableLogsArchives{}
	outputArchives.UnmarshalJSON([]byte(readFixture(t, fmt.Sprintf("fixtures/logs/archives/%s/out/%s.json", archiveType, action))))
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, fmt.Sprintf("LogsArchive.%s.%s", action, archiveType))
	assert.NoError(err)
	gock.New(URL).Get("/api/v2/logs/config/archives/").Reply(200).JSON(outputArchives)
	assert.True(len(*outputArchives.Get().Data) > 0)
	checkS3Archive(t, ctx, (*outputArchives.Get().Data)[0])
}

func checkS3Archive(t *testing.T, ctx context.Context, outputArchive datadog.LogsArchiveDefinition) {
	assert := tests.Assert(ctx, t)
	assert.Equal(outputArchive.Type, "archives")
	destination := outputArchive.Attributes.Destination.Get().LogsArchiveDestinationS3
	assert.Equal(destination.Type, datadog.LOGSARCHIVEDESTINATIONS3TYPE_S3)
	assert.Equal(destination.Integration.AccountId, "711111111111")
	assert.Equal(destination.Integration.RoleName, "DatadogGoClientTestIntegrationRole")
	assert.Equal(destination.Path, datadog.PtrString("/path/blou"))
	assert.Equal(destination.Bucket, "dd-logs-test-datadog-api-client-go")
	assert.Equal(outputArchive.Attributes.Name, "datadog-api-client-go Tests Archive")
	assert.Equal(outputArchive.Attributes.Query, "source:tata")
	assert.Equal(outputArchive.Id, datadog.PtrString("XVlBzgbaiC"))
}

func archivesAreEqual(t *testing.T, ctx context.Context, archiveRequest datadog.LogsArchiveCreateRequestAttributes, createdArchive datadog.LogsArchiveAttributes) {
	assert := tests.Assert(ctx, t)
	assert.Equal(archiveRequest.Name, createdArchive.Name)
	assert.Equal(archiveRequest.Query, createdArchive.Query)
	s3DestinationReq := archiveRequest.Destination.LogsArchiveDestinationS3
	s3DestinationCreated := createdArchive.Destination.Get().LogsArchiveDestinationS3
	assert.Equal(s3DestinationReq, s3DestinationCreated)
	azureDestinationReq := archiveRequest.Destination.LogsArchiveDestinationAzure
	azureDestinationCreated := createdArchive.Destination.Get().LogsArchiveDestinationAzure
	assert.Equal(azureDestinationReq, azureDestinationCreated)
	gcsDestinationReq := archiveRequest.Destination.LogsArchiveDestinationGCS
	gcsDestinationCreated := createdArchive.Destination.Get().LogsArchiveDestinationGCS
	assert.Equal(gcsDestinationReq, gcsDestinationCreated)
}

func readFixture(t *testing.T, path string) string {
	t.Helper()
	res, err := tests.ReadFixture(path)
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	return res
}
