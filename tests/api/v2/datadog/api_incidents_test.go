// /*
//  * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
//  * This product includes software developed at Datadog (https://www.datadoghq.com/).
//  * Copyright 2019-Present Datadog, Inc.
//  */

package test

import (
	"context"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
)

func TestIncidentsLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	client := Client(ctx)

	setIncidentsUnstableOperations(client, true)
	defer setIncidentsUnstableOperations(client, false)

	time := tests.ClockFromContext(ctx).Now()
	testIncidentData := datadog.NewIncidentWithDefaults()
	testIncidentData.SetType("incidents")
	testIncidentData.SetAttributes(*datadog.NewIncidentAttributesWithDefaults())
	testIncidentData.Attributes.SetTitle("Test-Incident")
	testIncidentData.Attributes.SetState("active")
	testIncidentData.Attributes.SetCustomerImpacted(true)
	testIncidentData.Attributes.SetCustomerImpactScope("none")
	testIncidentData.Attributes.SetCustomerImpactStart(time)
	testIncidentData.Attributes.SetCustomerImpactEnd(time)
	testIncidentData.Attributes.SetDetected(time)

	testIncidentCreateData := datadog.NewIncidentCreateRequestWithInitialDataWithDefaults()

	data := datadog.NewIncidentCreateWithInitialData()
	data.SetAttributes(*datadog.NewIncidentCreateWithInitialDataAttributes())
	data.Attributes.UnsetCreationIdempotencyKey()
	data.Attributes.SetTitle(testIncidentData.Attributes.GetTitle())
	data.Attributes.SetState(testIncidentData.Attributes.GetState())
	data.Attributes.SetCustomerImpacted(testIncidentData.Attributes.GetCustomerImpacted())
	data.Attributes.SetCustomerImpactScope(testIncidentData.Attributes.GetCustomerImpactScope())
	data.Attributes.SetCustomerImpactStart(testIncidentData.Attributes.GetCustomerImpactStart())
	data.Attributes.SetCustomerImpactEnd(testIncidentData.Attributes.GetCustomerImpactEnd())
	data.Attributes.SetDetected(testIncidentData.Attributes.GetDetected())

	testIncidentCreateData.SetData(*data)

	// Create Incident
	incidentRsp, httpresp, err := client.IncidentsApi.CreateIncident(ctx).Body(*testIncidentCreateData).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error creating Incident %v: Response %s: %v", testIncidentData, err.Error(), bStr)
	}
	incident := incidentRsp.GetData()
	incidentAttrs := incident.GetAttributes()

	assert.Equal(httpresp.StatusCode, 201)

	// Lets ensure that we always delete the incident if we successfully created it, even if the update and
	// get requests fail
	defer func(test_client *datadog.APIClient, incident datadog.Incident) {
		//Delete Incident
		httpresp, err := test_client.IncidentsApi.DeleteIncident(ctx, incident.GetId()).Execute()
		if err != nil {
			bStr := err.(datadog.GenericOpenAPIError).Model()
			t.Fatalf("Error deleting Incident %v: Response %s: %v", incident, err.Error(), bStr)
		}
		assert.Equal(httpresp.StatusCode, 204)
	}(client, incident)

	assert.Equal(incident.GetType(), testIncidentData.GetType())
	assert.Equal(incidentAttrs.GetTitle(), testIncidentData.Attributes.GetTitle())
	assert.Equal(incidentAttrs.GetState(), testIncidentData.Attributes.GetState())
	assert.Equal(incidentAttrs.GetCustomerImpacted(), testIncidentData.Attributes.GetCustomerImpacted())
	assert.Equal(incidentAttrs.GetCustomerImpactScope(), testIncidentData.Attributes.GetCustomerImpactScope())

	// Edit Incident
	incident.Attributes.UnsetCreationIdempotencyKey()
	incident.Attributes.SetTitle("Test-Incident-Updated")
	incidentUpdatedRsp, httpresp, err := client.IncidentsApi.PatchIncident(ctx, incident.GetId()).Body(*datadog.NewIncidentRequest(incident)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error updating Incident %v: Response %s: %v", incident, err.Error(), bStr)
	}
	incidentUpdated := incidentUpdatedRsp.GetData()
	incidentAttrsUpdated := incidentUpdated.GetAttributes()
	assert.Equal(httpresp.StatusCode, 200)
	assert.Equal(incidentAttrsUpdated.GetTitle(), incidentAttrsUpdated.GetTitle())

	// Get Incident
	incidentGetResp, httpresp, err := client.IncidentsApi.GetIncident(ctx, incident.GetId()).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting Incident %v: Response %s: %v", incident, err.Error(), bStr)
	}
	assert.Equal(httpresp.StatusCode, 200)

	incidentGet := incidentGetResp.GetData()
	incidentAttrsGet := incidentGet.GetAttributes()
	assert.Equal(incidentGet.GetId(), incidentUpdated.GetId())
	assert.Equal(incidentAttrsGet.GetTitle(), incidentAttrsUpdated.GetTitle())

	// Get Incidents
	incidentsGetResp, httpresp, err := client.IncidentsApi.GetIncidents(ctx).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting Incidents %v: Response %s: %v", incident, err.Error(), bStr)
	}
	assert.Equal(httpresp.StatusCode, 200)
	assert.True(len(incidentsGetResp.GetData()) >= 1)

}

func createTestIncident(ctx context.Context, t *testing.T, client *datadog.APIClient) datadog.Incident {
	testIncidentData := datadog.NewIncidentCreateRequestWithInitialDataWithDefaults()

	data := datadog.NewIncidentCreateWithInitialData()
	data.SetAttributes(*datadog.NewIncidentCreateWithInitialDataAttributes())
	data.Attributes.SetTitle("Test-Incident")
	data.Attributes.SetState("unknown")
	data.Attributes.SetCustomerImpacted(false)

	testIncidentData.SetData(*data)
	incidentRsp, httpresp, err := client.IncidentsApi.CreateIncident(ctx).Body(*testIncidentData).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error creating Incident %v: Response %s: %v", testIncidentData, err.Error(), bStr)
	}
	assert := tests.Assert(ctx, t)
	assert.Equal(httpresp.StatusCode, 201)
	return incidentRsp.GetData()

}

func deleteTestIncident(ctx context.Context, t *testing.T, client *datadog.APIClient, id string) {
	httpresp, err := client.IncidentsApi.DeleteIncident(ctx, id).Execute()
	if err != nil && httpresp.StatusCode != 404 {
		t.Logf("Deleting Incident: %v failed with %v, Another test may have already deleted this Incident: %v",
			id, httpresp.StatusCode, err)
	}

}
