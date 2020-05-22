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
	"math/rand"
	_nethttp "net/http"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/stretchr/testify/assert"
)

const archiveName string = "datadog-api-client-go Tests Archive"
const archiveBucket string = "dd-logs-test-datadog-api-client-go"

var archiveIntegration = datadog.LogsArchiveIntegrationS3{
	AccountId: "711111111111",
	RoleName:  "DatadogGoClientTestIntegrationRole",
}
var archivePath = "/path/blou"

const archiveQuery string = "service:toto"

func TestLogsArchivesLifecycle(t *testing.T) {
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	client := Client(ctx)
	config := client.GetConfig()
	URL, err := config.ServerURLWithContext(ctx, "LogsArchivesTest.Lifecycle")
	assert.NoError(t, err)
	apiClient := apiClientTest{
		ctx:    ctx,
		client: client,
	}
	testClient := apiDummyClient{
		archives: map[string]datadog.LogsArchive{},
		client:   apiClient,
		URL:      URL,
	}
	defer gock.Off()

	archive := createArchive()
	// 1. Create
	createdArchive, httpresp, err := testClient.create(archive)
	assertNoError(t, err, httpresp)
	assert.Equal(t, 200, httpresp.StatusCode)
	archivesAreEqual(t, *archive.Data.Attributes, *createdArchive.Data.Attributes)

	// 2. Delete at the end
	defer deleteLogsArchive(t, &testClient, *createdArchive.Data.Id)

	// 3. Update by Id
	archive.Data.Attributes.Query = "source:tata"
	archiveUpdated, httpresp, err := testClient.updateById(*createdArchive.Data.Id, archive)
	assert.NoError(t, err)
	assert.Equal(t, "source:tata", archiveUpdated.Data.Attributes.Query)
	archivesAreEqual(t, *archive.Data.Attributes, *archiveUpdated.Data.Attributes)

	// 4. Get by Id
	archiveGot, _, err := testClient.getById(*createdArchive.Data.Id)
	assertNoError(t, err, httpresp)
	archivesAreEqual(t, *archive.Data.Attributes, *archiveGot.Data.Attributes)
	// 5. Get all
	archives, _, err := testClient.getAll()
	assertNoError(t, err, httpresp)
	assert.True(t, len(*archives.Data) > 0)
	lastArchive := (*archives.Data)[len(*archives.Data)-1]
	archivesAreEqual(t, *archive.Data.Attributes, *lastArchive.Attributes)
}

func TestLogsArchivesNotFound(t *testing.T) {
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	client := Client(ctx)
	config := client.GetConfig()
	URL, err := config.ServerURLWithContext(ctx, "LogsArchivesTest.NotFound")
	assert.NoError(t, err)
	apiClient := apiClientTest{
		ctx:    ctx,
		client: client,
	}
	testClient := apiDummyClient{
		archives: map[string]datadog.LogsArchive{},
		client:   apiClient,
		URL:      URL,
	}
	//defer gock.Off()

	archive := createArchive()
	// 1. Create
	createdArchive, httpresp, err := testClient.create(archive)
	if err != nil {
		t.Fatalf("Error creating logs archive: Response: %v, Http Resp: %v", err, httpresp)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	archivesAreEqual(t, *archive.Data.Attributes, *createdArchive.Data.Attributes)

	// 2. Delete
	deleteLogsArchive(t, &testClient, *createdArchive.Data.Id)

	// 3. Get by Id
	_, httpresp, err = testClient.getById(*createdArchive.Data.Id)
	assert.True(t, err != nil)
	assert.Equal(t, 404, httpresp.StatusCode)

	// 4. Delete by Id
	httpresp, err = testClient.deleteById(*createdArchive.Data.Id)
	assert.True(t, err != nil)
	assert.Equal(t, 404, httpresp.StatusCode)

	// 5. Update by Id
	_, httpresp, err = testClient.updateById(*createdArchive.Data.Id, archive)
	assert.True(t, err != nil)
	assert.Equal(t, 404, httpresp.StatusCode)
}

func archivesAreEqual(t *testing.T, archiveRequest datadog.LogsArchiveCreateRequestAttributes, createdArchive datadog.LogsArchiveAttributes) {
	assert.Equal(t, archiveRequest.Name, createdArchive.Name)
	assert.Equal(t, archiveRequest.Query, createdArchive.Query)
	s3DestinationReq := archiveRequest.Destination.LogsArchiveDestinationS3
	s3DestinationCreated := createdArchive.Destination.Get().LogsArchiveDestinationS3
	assert.Equal(t, s3DestinationReq.Bucket, s3DestinationCreated.Bucket)
	assert.Equal(t, s3DestinationReq.Path, s3DestinationCreated.Path)
	assert.Equal(t, s3DestinationReq.Integration, s3DestinationCreated.Integration)
	assert.Equal(t, s3DestinationReq.Type, s3DestinationCreated.Type)
}

func deleteLogsArchive(t *testing.T, testClient logsArchiveTestClient, id string) {
	httpresp, err := testClient.deleteById(id)
	assertNoError(t, err, httpresp)
	assert.Equal(t, 204, httpresp.StatusCode)
}

func createArchive() datadog.LogsArchiveCreateRequest {
	destinationS3 := datadog.LogsArchiveDestinationS3{
		Bucket:      archiveBucket,
		Integration: archiveIntegration,
		Path:        &archivePath,
		Type:        datadog.LOGSARCHIVEDESTINATIONS3TYPE_S3,
	}
	return datadog.LogsArchiveCreateRequest{
		Data: &datadog.LogsArchiveCreateRequestDefinition{
			Attributes: &datadog.LogsArchiveCreateRequestAttributes{
				Destination: datadog.LogsArchiveDestinationS3AsLogsArchiveCreateRequestDestination(&destinationS3),
				Name:        archiveName,
				Query:       archiveQuery,
			},
			Type: "archives",
		},
	}
}

type logsArchiveTestClient interface {
	getAll() (datadog.LogsArchives, *_nethttp.Response, error)
	create(archive datadog.LogsArchiveCreateRequest) (datadog.LogsArchive, *_nethttp.Response, error)
	getById(archiveId string) (datadog.LogsArchive, *_nethttp.Response, error)
	deleteById(archiveId string) (*_nethttp.Response, error)
	updateById(archiveId string, archive datadog.LogsArchiveCreateRequest) (datadog.LogsArchive, *_nethttp.Response, error)
}

type apiClientTest struct {
	ctx    context.Context
	client *datadog.APIClient
}

func (r *apiClientTest) getAll() (datadog.LogsArchives, *_nethttp.Response, error) {
	return r.client.LogsArchivesApi.ListLogsArchives(r.ctx).Execute()
}

func (r *apiClientTest) create(archive datadog.LogsArchiveCreateRequest) (datadog.LogsArchive, *_nethttp.Response, error) {
	return r.client.LogsArchivesApi.CreateLogsArchive(r.ctx).Body(archive).Execute()
}

func (r *apiClientTest) updateById(archiveId string, archive datadog.LogsArchiveCreateRequest) (datadog.LogsArchive, *_nethttp.Response, error) {
	return r.client.LogsArchivesApi.UpdateLogsArchive(r.ctx, archiveId).Body(archive).Execute()
}

func (r *apiClientTest) getById(archiveId string) (datadog.LogsArchive, *_nethttp.Response, error) {
	return r.client.LogsArchivesApi.GetLogsArchive(r.ctx, archiveId).Execute()
}

func (r *apiClientTest) deleteById(archiveId string) (*_nethttp.Response, error) {
	return r.client.LogsArchivesApi.DeleteLogsArchive(r.ctx, archiveId).Execute()
}

// Test Lifecycle

// 1. create(archive) -> (archive with id, httresp with code 200, nil)
// 2. updateById(archiveId, archive with Query = "source") -> (archive, httresp with code 200, nil)
// 3. getById(archiveId) -> (archive with id, httresp with code 200, nil)
// 4. getAll() -> ([archive with id], httresp with code 200, nil)

type apiDummyClient struct {
	archives map[string]datadog.LogsArchive
	client   apiClientTest
	URL      string
}

func (r *apiDummyClient) getAll() (datadog.LogsArchives, *_nethttp.Response, error) {
	archiveDefs := make([]datadog.LogsArchiveDefinition, len(r.archives))
	for _, archive := range r.archives {
		archiveDefs = append(archiveDefs, *archive.Data)
	}
	result := datadog.NewLogsArchives()
	result.Data = &archiveDefs
	gock.New(r.URL).Get("api/v2/logs/config/archives").Reply(200).JSON(result)
	defer gock.Off()
	return r.client.getAll()
}

func (r *apiDummyClient) create(archive datadog.LogsArchiveCreateRequest) (datadog.LogsArchive, *_nethttp.Response, error) {
	id := RandString()
	newArchive := logsArchiveCreateRequestToLogArchive(id, archive)
	r.archives[id] = newArchive
	gock.New(r.URL).Post("/api/v2/logs/config/archives").MatchType("json").JSON(archive).Reply(200).JSON(newArchive)
	defer gock.Off()
	return r.client.create(archive)
}

func (r *apiDummyClient) updateById(archiveId string, archive datadog.LogsArchiveCreateRequest) (datadog.LogsArchive, *_nethttp.Response, error) {
	if _, exists := r.archives[archiveId]; exists {
		newArchive := logsArchiveCreateRequestToLogArchive(archiveId, archive)
		r.archives[archiveId] = newArchive
		gock.New(r.URL).Put(fmt.Sprintf("api/v2/logs/config/archives/%s", archiveId)).MatchType("json").JSON(archive).Reply(200).JSON(newArchive)
	} else {
		newArchive := datadog.NewLogsArchive()
		gock.New(r.URL).Put(fmt.Sprintf("api/v2/logs/config/archives/%s", archiveId)).MatchType("json").JSON(archive).Reply(404).JSON(newArchive)
	}
	defer gock.Off()
	return r.client.updateById(archiveId, archive)
}

func (r *apiDummyClient) getById(archiveId string) (datadog.LogsArchive, *_nethttp.Response, error) {
	if value, exists := r.archives[archiveId]; exists {
		gock.New(r.URL).Get("api/v2/logs/config/archives").Reply(200).JSON(value)
	} else {
		gock.New(r.URL).Get("api/v2/logs/config/archives").Reply(404).JSON(datadog.NewLogsArchive())
	}
	defer gock.Off()
	return r.client.getById(archiveId)
}

func (r *apiDummyClient) deleteById(archiveId string) (*_nethttp.Response, error) {
	if _, exists := r.archives[archiveId]; exists {
		delete(r.archives, archiveId)
		gock.New(r.URL).Delete("api/v2/logs/config/archives").Reply(204)
	} else {
		gock.New(r.URL).Delete("api/v2/logs/config/archives").Reply(404)
	}
	return r.client.deleteById(archiveId)
}

func logsArchiveCreateRequestToLogArchive(id string, archive datadog.LogsArchiveCreateRequest) datadog.LogsArchive {
	return datadog.LogsArchive{
		Data: &datadog.LogsArchiveDefinition{
			Attributes: &datadog.LogsArchiveAttributes{
				Destination: *datadog.NewNullableLogsArchiveDestination(&datadog.LogsArchiveDestination{
					LogsArchiveDestinationS3:    archive.Data.Attributes.Destination.LogsArchiveDestinationS3,
					LogsArchiveDestinationAzure: archive.Data.Attributes.Destination.LogsArchiveDestinationAzure,
					LogsArchiveDestinationGCS:   archive.Data.Attributes.Destination.LogsArchiveDestinationGCS,
				}),
				Name:  archive.Data.Attributes.Name,
				Query: archive.Data.Attributes.Query,
			},
			Type: archive.Data.Type,
			Id:   datadog.PtrString(id),
		},
	}
}

func assertNoError(t *testing.T, err error, httpresp *_nethttp.Response) {
	if err != nil {
		if openApiError, ok := err.(datadog.GenericOpenAPIError); ok {
			t.Fatalf("Error creating logs archive: Response: %s, Http Resp: %v", openApiError.Body(), httpresp)
		} else {
			t.Fatalf("Error creating logs archive: Response: %s, Http Resp: %v", err, httpresp)
		}
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandString() string {
	n := 10
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
