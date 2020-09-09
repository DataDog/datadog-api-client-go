/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
)

func TestIncidentTodosLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	client := Client(ctx)

	setIncidentsUnstableOperations(client, true)
	defer setIncidentsUnstableOperations(client, false)

	incident := createTestIncident(ctx, t, client)
	defer deleteTestIncident(ctx, t, client, incident.GetId())

	time := tests.ClockFromContext(ctx).Now()
	testIncidentTodoData := datadog.NewIncidentTodoWithDefaults()
	testIncidentTodoData.SetType("incident_todos")
	testIncidentTodoData.SetAttributes(*datadog.NewIncidentTodoAttributesWithDefaults())
	testIncidentTodoData.Attributes.SetContent("TestTodoContent")

	// Create IncidentTodo
	incidentTodoRsp, httpresp, err := client.IncidentsApi.CreateIncidentTodo(ctx, incident.GetId()).
		Body(*datadog.NewIncidentTodoRequest(*testIncidentTodoData)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error creating IncidentTodo %v: Response %s: %v", testIncidentTodoData, err.Error(), bStr)
	}
	incidentTodo := incidentTodoRsp.GetData()
	incidentTodoAttrs := incidentTodo.GetAttributes()
	assert.Equal(httpresp.StatusCode, 201)

	// Lets ensure that we always delete the incidentTodo if we successfully created it, even if the update and
	// get requests fail
	defer func(test_client *datadog.APIClient, incidentTodo datadog.IncidentTodo) {
		//Delete IncidentTodo
		httpresp, err := test_client.IncidentsApi.DeleteIncidentTodo(ctx, incident.GetId(), incidentTodo.GetId()).Execute()
		if err != nil {
			bStr := err.(datadog.GenericOpenAPIError).Model()
			t.Fatalf("Error deleting IncidentTodo %v: Response %s: %v", incidentTodo, err.Error(), bStr)
		}
		assert.Equal(httpresp.StatusCode, 204)
	}(client, incidentTodo)

	assert.Equal(incidentTodo.GetType(), testIncidentTodoData.GetType())
	assert.Equal(incidentTodoAttrs.GetContent(), testIncidentTodoData.Attributes.GetContent())
	completed, _ := incidentTodoAttrs.GetCompletedOk()
	assert.Nil(completed)

	// Edit IncidentTodo
	incidentTodo.Attributes.SetCompleted(time)
	incidentTodoUpdatedRsp, httpresp, err := client.IncidentsApi.
		PatchIncidentTodo(ctx, incident.GetId(), incidentTodo.GetId()).
		Body(*datadog.NewIncidentTodoRequest(incidentTodo)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error updating IncidentTodo %v: Response %s: %v", incidentTodo, err.Error(), bStr)
	}
	incidentTodoUpdated := incidentTodoUpdatedRsp.GetData()
	incidentTodoAttrsUpdated := incidentTodoUpdated.GetAttributes()
	assert.Equal(httpresp.StatusCode, 200)
	completed, _ = incidentTodoAttrsUpdated.GetCompletedOk()
	assert.NotNil(completed)

	// Get IncidentTodo
	incidentTodoGetResp, httpresp, err := client.IncidentsApi.
		GetIncidentTodo(ctx, incident.GetId(), incidentTodo.GetId()).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting IncidentTodo %v: Response %s: %v", incidentTodo, err.Error(), bStr)
	}
	assert.Equal(httpresp.StatusCode, 200)

	incidentTodoGet := incidentTodoGetResp.GetData()
	incidentTodoAttrsGet := incidentTodoGet.GetAttributes()
	assert.Equal(incidentTodoGet.GetId(), incidentTodoUpdated.GetId())
	assert.Equal(incidentTodoAttrsGet.GetCompleted(), incidentTodoAttrsUpdated.GetCompleted())

	// Get IncidentTodos
	incidentTodosGetResp, httpresp, err := client.IncidentsApi.GetIncidentTodos(ctx, incident.GetId()).
		Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting IncidentTodos %v: Response %s: %v", incidentTodo, err.Error(), bStr)
	}
	assert.Equal(httpresp.StatusCode, 200)
	assert.True(len(incidentTodosGetResp.GetData()) >= 1)

}
