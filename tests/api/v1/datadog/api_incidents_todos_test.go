/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func TestIncidentTodosLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	INCIDENT := createTestIncident(t)
	defer deleteTestIncident(t, INCIDENT.GetId())

	time := TESTCLOCK.Now()
	formattedClock := *datadog.NewNullableTime(&time)
	testIncidentTodoData := datadog.NewIncidentTodoWithDefaults()
	testIncidentTodoData.SetType("incident_todos")
	testIncidentTodoData.SetAttributes(*datadog.NewIncidentTodoAttributesWithDefaults())
	testIncidentTodoData.Attributes.SetContent("TestTodoContent")

	// Create IncidentTodo
	incidentTodoRsp, httpresp, err := TESTAPICLIENT.IncidentsApi.CreateIncidentTodo(TESTAUTH, INCIDENT.GetId()).
		Body(*datadog.NewIncidentTodoPayload(*testIncidentTodoData)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error creating IncidentTodo %v: Response %s: %v", testIncidentTodoData, err.Error(), bStr)
	}
	incidentTodo := incidentTodoRsp.GetData()
	incidentTodoAttrs := incidentTodo.GetAttributes()
	assert.Equal(t, httpresp.StatusCode, 201)

	// Lets ensure that we always delete the incidentTodo if we successfully created it, even if the update and
	// get requests fail
	defer func(test_client *datadog.APIClient, incidentTodo datadog.IncidentTodo) {
		//Delete IncidentTodo
		httpresp, err := test_client.IncidentsApi.DeleteIncidentTodo(TESTAUTH, INCIDENT.GetId(), incidentTodo.GetId()).Execute()
		if err != nil {
			bStr := err.(datadog.GenericOpenAPIError).Model()
			t.Fatalf("Error deleting IncidentTodo %v: Response %s: %v", incidentTodo, err.Error(), bStr)
		}
		assert.Equal(t, httpresp.StatusCode, 204)
	}(TESTAPICLIENT, incidentTodo)

	assert.Equal(t, incidentTodo.GetType(), testIncidentTodoData.GetType())
	assert.Equal(t, incidentTodoAttrs.GetContent(), testIncidentTodoData.Attributes.GetContent())
	assert.Nil(t, incidentTodoAttrs.GetCompleted().Get())
	// Edit IncidentTodo
	incidentTodo.Attributes.SetCompleted(formattedClock)
	incidentTodoUpdatedRsp, httpresp, err := TESTAPICLIENT.IncidentsApi.
		PatchIncidentTodo(TESTAUTH, INCIDENT.GetId(), incidentTodo.GetId()).
		Body(*datadog.NewIncidentTodoPayload(incidentTodo)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error updating IncidentTodo %v: Response %s: %v", incidentTodo, err.Error(), bStr)
	}
	incidentTodoUpdated := incidentTodoUpdatedRsp.GetData()
	incidentTodoAttrsUpdated := incidentTodoUpdated.GetAttributes()
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.NotNil(t, incidentTodoAttrsUpdated.GetCompleted().Get())

	// Get IncidentTodo
	incidentTodoGetResp, httpresp, err := TESTAPICLIENT.IncidentsApi.
		GetIncidentTodo(TESTAUTH, INCIDENT.GetId(), incidentTodo.GetId()).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting IncidentTodo %v: Response %s: %v", incidentTodo, err.Error(), bStr)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	incidentTodoGet := incidentTodoGetResp.GetData()
	incidentTodoAttrsGet := incidentTodoGet.GetAttributes()
	assert.Equal(t, incidentTodoGet.GetId(), incidentTodoUpdated.GetId())
	assert.Equal(t, incidentTodoAttrsGet.GetCompleted(), incidentTodoAttrsUpdated.GetCompleted())

	// Get IncidentTodos
	incidentTodosGetResp, httpresp, err := TESTAPICLIENT.IncidentsApi.GetIncidentTodos(TESTAUTH, INCIDENT.GetId()).
		Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting IncidentTodos %v: Response %s: %v", incidentTodo, err.Error(), bStr)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.True(t, len(incidentTodosGetResp.GetData()) >= 1)

}
