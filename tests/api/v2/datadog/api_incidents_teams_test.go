/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"testing"

	"github.com/DataDog/datadog-api-client-go/tests"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func TestIncidentTeamsLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	client := Client(ctx)

	setIncidentsUnstableOperations(client, true)
	defer setIncidentsUnstableOperations(client, false)

	incident := createTestIncident(t, client, ctx)
	defer deleteTestIncident(t, client, ctx, incident.GetId())

	team1 := createTestTeam(t, client, ctx)
	defer deleteTestTeam(t, client, ctx, team1.GetId())

	// Create IncidentTeam
	incidentTeamRsp, httpresp, err := client.IncidentsApi.AddTeamToIncident(ctx, incident.GetId()).
		Body(*datadog.NewTeamPayload(team1)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error adding team to incident %v: Response %s: %v", team1, err.Error(), bStr)
	}
	assert.Equal(httpresp.StatusCode, 201)

	// Lets ensure that we always delete the incidentTeam if we successfully created it, even if the update and
	// get requests fail
	defer func(test_client *datadog.APIClient, incidentTeam datadog.Team) {
		//Delete IncidentTeam
		httpresp, err := test_client.IncidentsApi.
			RemoveTeamFromIncident(ctx, incident.GetId(), incidentTeam.GetId()).Execute()
		if err != nil {
			bStr := err.(datadog.GenericOpenAPIError).Model()
			t.Fatalf("Error removeing team for incident %v: Response %s: %v", incidentTeam, err.Error(), bStr)
		}
		assert.Equal(httpresp.StatusCode, 204)
	}(client, team1)

	assert.True(len(incidentTeamRsp.GetData()) >= 1)
	assert.NotNil(incidentTeamRsp.GetData()[0].GetId())
	assert.Equal(incidentTeamRsp.GetData()[0].GetType(), team1.GetType())
	assert.Equal(incidentTeamRsp.GetData()[0].GetAttributes().Name, team1.GetAttributes().Name)

	// Get IncidentTeams
	incidentTeamsGetResp, httpresp, err := client.IncidentsApi.GetTeamsForIncident(ctx, incident.GetId()).
		Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting Teams for Incident %v: Response %s: %v", team1, err.Error(), bStr)
	}
	assert.Equal(httpresp.StatusCode, 200)
	assert.True(len(incidentTeamsGetResp.GetData()) >= 1)
	assert.NotNil(incidentTeamRsp.GetData()[0].GetId())
	assert.Equal(incidentTeamRsp.GetData()[0].GetType(), team1.GetType())
	assert.Equal(incidentTeamRsp.GetData()[0].GetAttributes().Name, team1.GetAttributes().Name)

}
