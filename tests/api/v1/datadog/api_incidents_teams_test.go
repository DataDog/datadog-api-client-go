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

func TestIncidentTeamsLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	SetUnstableIncidentsAPIs(true)
	defer SetUnstableIncidentsAPIs(false)

	INCIDENT := createTestIncident(t)
	defer deleteTestIncident(t, INCIDENT.GetId())
	TEAM_1 := createTestTeam(t)
	defer deleteTestTeam(t, TEAM_1.GetId())

	// Create IncidentTeam
	incidentTeamRsp, httpresp, err := TESTAPICLIENT.IncidentsApi.AddTeamToIncident(TESTAUTH, INCIDENT.GetId()).
		Body(*datadog.NewTeamPayload(TEAM_1)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error adding team to incident %v: Response %s: %v", TEAM_1, err.Error(), bStr)
	}
	assert.Equal(t, httpresp.StatusCode, 201)

	// Lets ensure that we always delete the incidentTeam if we successfully created it, even if the update and
	// get requests fail
	defer func(test_client *datadog.APIClient, incidentTeam datadog.Team) {
		//Delete IncidentTeam
		httpresp, err := test_client.IncidentsApi.
			RemoveTeamFromIncident(TESTAUTH, INCIDENT.GetId(), incidentTeam.GetId()).Execute()
		if err != nil {
			bStr := err.(datadog.GenericOpenAPIError).Model()
			t.Fatalf("Error removeing team for incident %v: Response %s: %v", incidentTeam, err.Error(), bStr)
		}
		assert.Equal(t, httpresp.StatusCode, 204)
	}(TESTAPICLIENT, TEAM_1)

	assert.True(t, len(incidentTeamRsp.GetData()) >= 1)
	assert.NotNil(t, incidentTeamRsp.GetData()[0].GetId())
	assert.Equal(t, incidentTeamRsp.GetData()[0].GetType(), TEAM_1.GetType())
	assert.Equal(t, incidentTeamRsp.GetData()[0].GetAttributes().Name, TEAM_1.GetAttributes().Name)

	// Get IncidentTeams
	incidentTeamsGetResp, httpresp, err := TESTAPICLIENT.IncidentsApi.GetTeamsForIncident(TESTAUTH, INCIDENT.GetId()).
		Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting Teams for Incident %v: Response %s: %v", TEAM_1, err.Error(), bStr)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.True(t, len(incidentTeamsGetResp.GetData()) >= 1)
	assert.NotNil(t, incidentTeamRsp.GetData()[0].GetId())
	assert.Equal(t, incidentTeamRsp.GetData()[0].GetType(), TEAM_1.GetType())
	assert.Equal(t, incidentTeamRsp.GetData()[0].GetAttributes().Name, TEAM_1.GetAttributes().Name)

}
