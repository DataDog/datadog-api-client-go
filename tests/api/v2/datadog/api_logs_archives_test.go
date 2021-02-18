/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
	"gopkg.in/h2non/gock.v1"
)

// This test uses mocking because: 1) it relies on private data. 2) It relies on external services

func TestLogsArchivesCreate(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	includeTags := true
	testCases := []struct {
		archiveType string
		archive     datadog.LogsArchiveCreateRequest
	}{
		{
			archiveType: "s3",
			archive: datadog.LogsArchiveCreateRequest{
				Data: &datadog.LogsArchiveCreateRequestDefinition{
					Attributes: &datadog.LogsArchiveCreateRequestAttributes{
						Destination: datadog.LogsArchiveCreateRequestDestination{
							LogsArchiveDestinationS3: &datadog.LogsArchiveDestinationS3{
								Bucket: "dd-logs-test-datadog-api-client-go",
								Integration: datadog.LogsArchiveIntegrationS3{
									AccountId: "711111111111",
									RoleName:  "DatadogGoClientTestIntegrationRole",
								},
								Path: datadog.PtrString("/path/blou"),
								Type: "s3",
							},
						},
						Name:            "datadog-api-client-go Tests Archive",
						Query:           "service:toto",
						RehydrationTags: &[]string{"team:intake", "team:app"},
						IncludeTags:     &includeTags,
					},
					Type: "archives",
				},
			},
		},
		{
			archiveType: "azure",
			archive: datadog.LogsArchiveCreateRequest{
				Data: &datadog.LogsArchiveCreateRequestDefinition{
					Attributes: &datadog.LogsArchiveCreateRequestAttributes{
						Destination: datadog.LogsArchiveCreateRequestDestination{
							LogsArchiveDestinationAzure: &datadog.LogsArchiveDestinationAzure{
								Container: "my-container",
								Integration: datadog.LogsArchiveIntegrationAzure{
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
						RehydrationTags: &[]string{"team:intake", "team:app"},
						IncludeTags:     &includeTags,
					},
					Type: "archives",
				},
			},
		},
		{
			archiveType: "gcs",
			archive: datadog.LogsArchiveCreateRequest{
				Data: &datadog.LogsArchiveCreateRequestDefinition{
					Attributes: &datadog.LogsArchiveCreateRequestAttributes{
						Destination: datadog.LogsArchiveCreateRequestDestination{
							LogsArchiveDestinationGCS: &datadog.LogsArchiveDestinationGCS{
								Bucket: "dd-logs-test-datadog-api-client-go",
								Integration: datadog.LogsArchiveIntegrationGCS{
									ClientEmail: "email@email.com",
									ProjectId:   "aaaaaaaa-1a1a-1a1a-1a1a-aaaaaaaaaaaa",
								},
								Path: datadog.PtrString("/path/blou"),
								Type: "gcs",
							},
						},
						Name:            "datadog-api-client-go Tests Archive",
						Query:           "service:toto",
						RehydrationTags: &[]string{"team:intake", "team:app"},
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

			outputArchiveStr := readFixture(t, fmt.Sprintf("fixtures/logs/archives/%s/out/%s.json", tc.archiveType, "create"))
			outputArchive := datadog.NullableLogsArchive{}
			outputArchive.UnmarshalJSON([]byte(outputArchiveStr))

			URL, err := client.GetConfig().ServerURLWithContext(ctx, "LogsArchivesApiService.CreateLogsArchive")
			assert.NoError(err)

			gock.New(URL).Post("/api/v2/logs/config/archives").MatchType("json").JSON(tc.archive).Reply(200).Type("json").BodyString(outputArchiveStr)
			defer gock.Off()

			result, httpresp, err := client.LogsArchivesApi.CreateLogsArchive(ctx).Body(tc.archive).Execute()
			assert.NoError(err)
			assert.Equal(httpresp.StatusCode, 200)
			assert.Equal(result, *outputArchive.Get())
		})
	}
}

func TestLogsArchivesUpdate(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	client := Client(ctx)
	archiveType := "s3"
	action := "update"
	inputArchive := datadog.LogsArchiveCreateRequest{
		Data: &datadog.LogsArchiveCreateRequestDefinition{
			Attributes: &datadog.LogsArchiveCreateRequestAttributes{
				Destination: datadog.LogsArchiveCreateRequestDestination{
					LogsArchiveDestinationS3: &datadog.LogsArchiveDestinationS3{
						Bucket: "dd-logs-test-datadog-api-client-go",
						Integration: datadog.LogsArchiveIntegrationS3{
							AccountId: "711111111111",
							RoleName:  "DatadogGoClientTestIntegrationRole",
						},
						Path: datadog.PtrString("/path/blou"),
						Type: "s3",
					},
				},
				Name:            "datadog-api-client-go Tests Archive",
				Query:           "service:toto",
				RehydrationTags: &[]string{"team:intake", "team:app"},
			},
			Type: "archives",
		},
	}
	outputArchiveStr := readFixture(t, fmt.Sprintf("fixtures/logs/archives/%s/out/%s.json", archiveType, action))
	outputArchive := datadog.NullableLogsArchive{}
	outputArchive.UnmarshalJSON([]byte(outputArchiveStr))
	URL, err := client.GetConfig().ServerURLWithContext(ctx, "LogsArchivesApiService.UpdateLogsArchive")
	assert.NoError(err)
	id := "FOObartotO"
	gock.New(URL).Put(fmt.Sprintf("/api/v2/logs/config/archives/%s", id)).MatchType("json").JSON(inputArchive).Reply(200).Type("json").BodyString(outputArchiveStr)
	defer gock.Off()
	result, httpresp, err := client.LogsArchivesApi.UpdateLogsArchive(ctx, id).Body(inputArchive).Execute()
	assert.Equal(httpresp.StatusCode, 200)
	assert.Equal(result, *outputArchive.Get())
}

func TestLogsArchivesGetByID(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	client := Client(ctx)
	id := "FOObartotO"
	action := "getbyid"
	archiveType := "s3"
	outputArchiveStr := readFixture(t, fmt.Sprintf("fixtures/logs/archives/%s/out/%s.json", archiveType, action))
	outputArchive := datadog.NullableLogsArchive{}
	outputArchive.UnmarshalJSON([]byte(outputArchiveStr))
	URL, err := client.GetConfig().ServerURLWithContext(ctx, "LogsArchivesApiService.GetLogsArchive")
	assert.NoError(err)
	gock.New(URL).Get(fmt.Sprintf("/api/v2/logs/config/archives/%s", id)).Reply(200).Type("json").BodyString(outputArchiveStr)
	defer gock.Off()
	result, httpresp, err := client.LogsArchivesApi.GetLogsArchive(ctx, id).Execute()
	assert.Equal(httpresp.StatusCode, 200)
	assert.Equal(result, *outputArchive.Get())
}

func TestLogsArchivesDelete(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	id := "FOObartotO"
	client := Client(ctx)
	URL, err := client.GetConfig().ServerURLWithContext(ctx, "LogsArchivesApiService.DeleteLogsArchive")
	assert.NoError(err)
	gock.New(URL).Delete(fmt.Sprintf("/api/v2/logs/config/archives/%s", id)).Reply(204)
	defer gock.Off()
	httpresp, err := client.LogsArchivesApi.DeleteLogsArchive(ctx, id).Execute()
	assert.NoError(err)
	assert.Equal(httpresp.StatusCode, 204)
}

func TestLogsArchivesGetAll(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	client := Client(ctx)
	assert := tests.Assert(ctx, t)
	action := "getall"
	archiveType := "s3"
	outputArchivesStr := readFixture(t, fmt.Sprintf("fixtures/logs/archives/%s/out/%s.json", archiveType, action))
	outputArchives := datadog.NullableLogsArchives{}
	outputArchives.UnmarshalJSON([]byte(outputArchivesStr))
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "LogsArchivesApiService.ListLogsArchives")
	assert.NoError(err)
	gock.New(URL).Get("/api/v2/logs/config/archives").Reply(200).Type("json").JSON(outputArchivesStr)
	defer gock.Off()
	result, httpresp, err := client.LogsArchivesApi.ListLogsArchives(ctx).Execute()
	assert.NoError(err)
	assert.Equal(httpresp.StatusCode, 200)
	assert.True(len(*result.Data) > 0)
	assert.Equal(*outputArchives.Get(), result)
}

func TestGetLogsArchiveOrder(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	client := Client(ctx)
	assert := tests.Assert(ctx, t)

	fileName := "default"
	outputArchiveOrderStr := readFixture(t, fmt.Sprintf("fixtures/logs/archives/archive_order/out/%s.json", fileName))
	outputArchiveOrder := datadog.NullableLogsArchiveOrder{}
	_ = outputArchiveOrder.UnmarshalJSON([]byte(outputArchiveOrderStr))

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "LogsArchivesApiService.GetLogsArchiveOrder")
	assert.NoError(err)
	gock.New(URL).Get("/api/v2/logs/config/archive-order").Reply(200).Type("json").JSON(outputArchiveOrderStr)
	defer gock.Off()
	result, httpresp, err := client.LogsArchivesApi.GetLogsArchiveOrder(ctx).Execute()
	assert.NoError(err)
	assert.Equal(httpresp.StatusCode, 200)
	assert.Equal(*outputArchiveOrder.Get(), result)

}

func TestUpdateLogsArchiveOrder(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	client := Client(ctx)
	assert := tests.Assert(ctx, t)
	input := createUpdatedLogsArchiveOrder(t)

	fileName := "default"
	outputArchiveOrderStr := readFixture(t, fmt.Sprintf("fixtures/logs/archives/archive_order/out/%s.json", fileName))
	outputArchiveOrder := datadog.NullableLogsArchiveOrder{}
	_ = outputArchiveOrder.UnmarshalJSON([]byte(outputArchiveOrderStr))

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "LogsArchivesApiService.GetLogsArchiveOrder")
	assert.NoError(err)
	gock.New(URL).Put("/api/v2/logs/config/archive-order").Reply(200).Type("json").JSON(outputArchiveOrderStr)
	defer gock.Off()
	result, httpresp, err := client.LogsArchivesApi.UpdateLogsArchiveOrder(ctx).Body(*input).Execute()
	assert.NoError(err)
	assert.Equal(httpresp.StatusCode, 200)
	assert.Equal(*outputArchiveOrder.Get(), result)
}

func createUpdatedLogsArchiveOrder(t *testing.T) *datadog.LogsArchiveOrder {
	fileName := "updated"
	oldArchiveOrderStr := readFixture(t, fmt.Sprintf("fixtures/logs/archives/archive_order/out/%s.json", fileName))
	oldArchiveOrder := datadog.NullableLogsArchiveOrder{}
	_ = oldArchiveOrder.UnmarshalJSON([]byte(oldArchiveOrderStr))
	data := oldArchiveOrder.Get().GetData()
	attributes := data.GetAttributes()
	archiveIds := attributes.GetArchiveIds()
	archiveIds = append(archiveIds, archiveIds[0])
	archiveIds = append(archiveIds[:1], archiveIds[2:]...)

	archiveOrderAttribute := datadog.NewLogsArchiveOrderAttributesWithDefaults()
	archiveOrderAttribute.SetArchiveIds(archiveIds)
	archiveOrderDefinition := datadog.NewLogsArchiveOrderDefinitionWithDefaults()
	archiveOrderDefinition.SetAttributes(*archiveOrderAttribute)
	newArchiveOrder := datadog.NewLogsArchiveOrderWithDefaults()
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
