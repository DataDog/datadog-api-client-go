/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"fmt"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"gopkg.in/h2non/gock.v1"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/DataDog/datadog-api-client-go/v2/tests"
)

// This test uses mocking because: 1) it relies on private data. 2) It relies on external services

func TestLogsArchivesCreate(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	includeTags := true
	testCases := []struct {
		archiveType string
		archive     datadogV2.LogsArchiveCreateRequest
	}{
		{
			archiveType: "s3",
			archive: datadogV2.LogsArchiveCreateRequest{
				Data: &datadogV2.LogsArchiveCreateRequestDefinition{
					Attributes: &datadogV2.LogsArchiveCreateRequestAttributes{
						Destination: datadogV2.LogsArchiveCreateRequestDestination{
							LogsArchiveDestinationS3: &datadogV2.LogsArchiveDestinationS3{
								Bucket: "dd-logs-test-datadog-api-client-go",
								Integration: datadogV2.LogsArchiveIntegrationS3{
									AccountId: "711111111111",
									RoleName:  "DatadogGoClientTestIntegrationRole",
								},
								Path: datadog.PtrString("/path/blou"),
								Type: "s3",
							},
						},
						Name:            "datadog-api-client-go Tests Archive",
						Query:           "service:toto",
						RehydrationTags: []string{"team:intake", "team:app"},
						IncludeTags:     &includeTags,
					},
					Type: "archives",
				},
			},
		},
		{
			archiveType: "azure",
			archive: datadogV2.LogsArchiveCreateRequest{
				Data: &datadogV2.LogsArchiveCreateRequestDefinition{
					Attributes: &datadogV2.LogsArchiveCreateRequestAttributes{
						Destination: datadogV2.LogsArchiveCreateRequestDestination{
							LogsArchiveDestinationAzure: &datadogV2.LogsArchiveDestinationAzure{
								Container: "my-container",
								Integration: datadogV2.LogsArchiveIntegrationAzure{
									ClientId: "aaaaaaaa-1a1a-1a1a-1a1a-aaaaaaaaaaaa",
									TenantId: "aaaaaaaa-1a1a-1a1a-1a1a-aaaaaaaaaaaa",
								},
								Path:           datadog.PtrString("/path/blou"),
								Region:         datadog.PtrString("my-region"),
								StorageAccount: "storageAccount",
								Type:           "azure",
							},
						},
						Name:            "datadog-api-client-go Tests Archive",
						Query:           "service:toto",
						RehydrationTags: []string{"team:intake", "team:app"},
						IncludeTags:     &includeTags,
					},
					Type: "archives",
				},
			},
		},
		{
			archiveType: "gcs",
			archive: datadogV2.LogsArchiveCreateRequest{
				Data: &datadogV2.LogsArchiveCreateRequestDefinition{
					Attributes: &datadogV2.LogsArchiveCreateRequestAttributes{
						Destination: datadogV2.LogsArchiveCreateRequestDestination{
							LogsArchiveDestinationGCS: &datadogV2.LogsArchiveDestinationGCS{
								Bucket: "dd-logs-test-datadog-api-client-go",
								Integration: datadogV2.LogsArchiveIntegrationGCS{
									ClientEmail: "email@email.com",
									ProjectId:   "aaaaaaaa-1a1a-1a1a-1a1a-aaaaaaaaaaaa",
								},
								Path: datadog.PtrString("/path/blou"),
								Type: "gcs",
							},
						},
						Name:            "datadog-api-client-go Tests Archive",
						Query:           "service:toto",
						RehydrationTags: []string{"team:intake", "team:app"},
						IncludeTags:     &includeTags,
					},
					Type: "archives",
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.archiveType, func(t *testing.T) {
			ctx := WithClient(WithFakeAuth(ctx))

			client := Client(ctx)
			assert := tests.Assert(ctx, t)
			api := datadogV2.NewLogsArchivesApi(client)

			outputArchiveStr := readFixture(t, fmt.Sprintf("fixtures/logs/archives/%s/out/%s.json", tc.archiveType, "create"))
			outputArchive := datadogV2.LogsArchive{}
			outputArchive.UnmarshalJSON([]byte(outputArchiveStr))

			URL, err := client.GetConfig().ServerURLWithContext(ctx, "LogsArchivesApiService.CreateLogsArchive")
			assert.NoError(err)

			gock.New(URL).Post("/api/v2/logs/config/archives").MatchType("json").JSON(tc.archive).Reply(200).Type("json").BodyString(outputArchiveStr)
			defer gock.Off()

			result, httpresp, err := api.CreateLogsArchive(ctx, tc.archive)
			assert.NoError(err)
			assert.Equal(httpresp.StatusCode, 200)
			assert.Equal(result, outputArchive)
		})
	}
}

func TestLogsArchivesUpdate(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	client := Client(ctx)
	api := datadogV2.NewLogsArchivesApi(client)
	archiveType := "s3"
	action := "update"
	inputArchive := datadogV2.LogsArchiveCreateRequest{
		Data: &datadogV2.LogsArchiveCreateRequestDefinition{
			Attributes: &datadogV2.LogsArchiveCreateRequestAttributes{
				Destination: datadogV2.LogsArchiveCreateRequestDestination{
					LogsArchiveDestinationS3: &datadogV2.LogsArchiveDestinationS3{
						Bucket: "dd-logs-test-datadog-api-client-go",
						Integration: datadogV2.LogsArchiveIntegrationS3{
							AccountId: "711111111111",
							RoleName:  "DatadogGoClientTestIntegrationRole",
						},
						Path: datadog.PtrString("/path/blou"),
						Type: "s3",
					},
				},
				Name:            "datadog-api-client-go Tests Archive",
				Query:           "service:toto",
				RehydrationTags: []string{"team:intake", "team:app"},
			},
			Type: "archives",
		},
	}
	outputArchiveStr := readFixture(t, fmt.Sprintf("fixtures/logs/archives/%s/out/%s.json", archiveType, action))
	outputArchive := datadogV2.LogsArchive{}
	outputArchive.UnmarshalJSON([]byte(outputArchiveStr))
	URL, err := client.GetConfig().ServerURLWithContext(ctx, "LogsArchivesApiService.UpdateLogsArchive")
	assert.NoError(err)
	id := "FOObartotO"
	gock.New(URL).Put(fmt.Sprintf("/api/v2/logs/config/archives/%s", id)).MatchType("json").JSON(inputArchive).Reply(200).Type("json").BodyString(outputArchiveStr)
	defer gock.Off()
	result, httpresp, _ := api.UpdateLogsArchive(ctx, id, inputArchive)
	assert.Equal(httpresp.StatusCode, 200)
	assert.Equal(result, outputArchive)
}

func TestLogsArchivesGetByID(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	client := Client(ctx)
	api := datadogV2.NewLogsArchivesApi(client)
	id := "FOObartotO"
	action := "getbyid"
	archiveType := "s3"
	outputArchiveStr := readFixture(t, fmt.Sprintf("fixtures/logs/archives/%s/out/%s.json", archiveType, action))
	outputArchive := datadogV2.LogsArchive{}
	outputArchive.UnmarshalJSON([]byte(outputArchiveStr))
	URL, err := client.GetConfig().ServerURLWithContext(ctx, "LogsArchivesApiService.GetLogsArchive")
	assert.NoError(err)
	gock.New(URL).Get(fmt.Sprintf("/api/v2/logs/config/archives/%s", id)).Reply(200).Type("json").BodyString(outputArchiveStr)
	defer gock.Off()
	result, httpresp, _ := api.GetLogsArchive(ctx, id)
	assert.Equal(httpresp.StatusCode, 200)
	assert.Equal(result, outputArchive)
}

func TestLogsArchivesDelete(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	id := "FOObartotO"
	client := Client(ctx)
	api := datadogV2.NewLogsArchivesApi(client)
	URL, err := client.GetConfig().ServerURLWithContext(ctx, "LogsArchivesApiService.DeleteLogsArchive")
	assert.NoError(err)
	gock.New(URL).Delete(fmt.Sprintf("/api/v2/logs/config/archives/%s", id)).Reply(204)
	defer gock.Off()
	httpresp, err := api.DeleteLogsArchive(ctx, id)
	assert.NoError(err)
	assert.Equal(httpresp.StatusCode, 204)
}

func TestLogsArchivesGetAll(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	client := Client(ctx)
	assert := tests.Assert(ctx, t)
	api := datadogV2.NewLogsArchivesApi(client)
	action := "getall"
	archiveType := "s3"
	outputArchivesStr := readFixture(t, fmt.Sprintf("fixtures/logs/archives/%s/out/%s.json", archiveType, action))
	outputArchives := datadogV2.LogsArchives{}
	outputArchives.UnmarshalJSON([]byte(outputArchivesStr))
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "LogsArchivesApiService.ListLogsArchives")
	assert.NoError(err)
	gock.New(URL).Get("/api/v2/logs/config/archives").Reply(200).Type("json").JSON(outputArchivesStr)
	defer gock.Off()
	result, httpresp, err := api.ListLogsArchives(ctx)
	assert.NoError(err)
	assert.Equal(httpresp.StatusCode, 200)
	assert.True(len(result.Data) > 0)
	assert.Equal(outputArchives, result)
}

func TestGetLogsArchiveOrder(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	client := Client(ctx)
	api := datadogV2.NewLogsArchivesApi(client)
	assert := tests.Assert(ctx, t)

	fileName := "default"
	outputArchiveOrderStr := readFixture(t, fmt.Sprintf("fixtures/logs/archives/archive_order/out/%s.json", fileName))
	outputArchiveOrder := datadogV2.LogsArchiveOrder{}
	_ = outputArchiveOrder.UnmarshalJSON([]byte(outputArchiveOrderStr))

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "LogsArchivesApiService.GetLogsArchiveOrder")
	assert.NoError(err)
	gock.New(URL).Get("/api/v2/logs/config/archive-order").Reply(200).Type("json").JSON(outputArchiveOrderStr)
	defer gock.Off()
	result, httpresp, err := api.GetLogsArchiveOrder(ctx)
	assert.NoError(err)
	assert.Equal(httpresp.StatusCode, 200)
	assert.Equal(outputArchiveOrder, result)

}

func TestUpdateLogsArchiveOrder(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	client := Client(ctx)
	api := datadogV2.NewLogsArchivesApi(client)
	assert := tests.Assert(ctx, t)
	input := createUpdatedLogsArchiveOrder(t)

	fileName := "default"
	outputArchiveOrderStr := readFixture(t, fmt.Sprintf("fixtures/logs/archives/archive_order/out/%s.json", fileName))
	outputArchiveOrder := datadogV2.LogsArchiveOrder{}
	_ = outputArchiveOrder.UnmarshalJSON([]byte(outputArchiveOrderStr))

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "LogsArchivesApiService.GetLogsArchiveOrder")
	assert.NoError(err)
	gock.New(URL).Put("/api/v2/logs/config/archive-order").Reply(200).Type("json").JSON(outputArchiveOrderStr)
	defer gock.Off()
	result, httpresp, err := api.UpdateLogsArchiveOrder(ctx, *input)
	assert.NoError(err)
	assert.Equal(httpresp.StatusCode, 200)
	assert.Equal(outputArchiveOrder, result)
}

func createUpdatedLogsArchiveOrder(t *testing.T) *datadogV2.LogsArchiveOrder {
	fileName := "updated"
	oldArchiveOrderStr := readFixture(t, fmt.Sprintf("fixtures/logs/archives/archive_order/out/%s.json", fileName))
	oldArchiveOrder := datadogV2.LogsArchiveOrder{}
	_ = oldArchiveOrder.UnmarshalJSON([]byte(oldArchiveOrderStr))
	data := oldArchiveOrder.GetData()
	attributes := data.GetAttributes()
	archiveIds := attributes.GetArchiveIds()
	archiveIds = append(archiveIds, archiveIds[0])
	archiveIds = append(archiveIds[:1], archiveIds[2:]...)

	archiveOrderAttribute := datadogV2.NewLogsArchiveOrderAttributesWithDefaults()
	archiveOrderAttribute.SetArchiveIds(archiveIds)
	archiveOrderDefinition := datadogV2.NewLogsArchiveOrderDefinitionWithDefaults()
	archiveOrderDefinition.SetAttributes(*archiveOrderAttribute)
	newArchiveOrder := datadogV2.NewLogsArchiveOrderWithDefaults()
	newArchiveOrder.SetData(*archiveOrderDefinition)
	return newArchiveOrder
}

func readFixture(t *testing.T, path string) string {
	t.Helper()
	res, err := tests.ReadFixture(path)
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	return res
}
