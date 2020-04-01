/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"fmt"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/stretchr/testify/assert"
)

func TestTeamsLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	SetUnstableIncidentsAPIs(true)
	defer SetUnstableIncidentsAPIs(false)

	testTeamData := datadog.NewTeam()
	testTeamData.SetType("teams")
	testTeamData.SetAttributes(*datadog.NewTeamAttributesWithDefaults())
	testTeamData.Attributes.SetName("Test-Team")
	testTeamAttributes := testTeamData.GetAttributes()
	// Create Team
	teamRsp, httpresp, err := TESTAPICLIENT.TeamsApi.CreateTeam(TESTAUTH).Body(*datadog.NewTeamPayload(*testTeamData)).
		Execute()
	if err != nil {
		t.Fatalf("Error creating Team %v: Response %s: %v", testTeamData, err.Error(), err)
	}
	team := teamRsp.GetData()
	teamAttrs := team.GetAttributes()
	assert.Equal(t, httpresp.StatusCode, 201)

	// Lets ensure that we always delete the team if we successfully created it, even if the update and
	// get requests fail
	defer func(test_client *datadog.APIClient, team datadog.Team) {
		//Delete Team
		httpresp, err := test_client.TeamsApi.DeleteTeam(TESTAUTH, team.GetId()).Execute()
		if err != nil {
			t.Fatalf("Error deleting Team %s: Response %s: %v", team.GetId(), err.Error(), err)
		}
		assert.Equal(t, httpresp.StatusCode, 204)
	}(TESTAPICLIENT, team)

	assert.Equal(t, team.GetType(), testTeamData.GetType())
	assert.Equal(t, teamAttrs.GetName(), testTeamAttributes.GetName())

	// Edit Team
	team.Attributes.SetName("Test-Team-Updated")
	teamUpdatedRsp, httpresp, err := TESTAPICLIENT.TeamsApi.PatchTeam(TESTAUTH, team.GetId()).Body(*datadog.NewTeamPayload(team)).Execute()
	if err != nil {
		t.Fatalf("Error updating Team %v: Response %s: %v", team, err.Error(), err)
	}
	teamUpdated := teamUpdatedRsp.GetData()
	teamAttrsUpdated := teamUpdated.GetAttributes()
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, teamAttrsUpdated.GetName(), teamAttrsUpdated.GetName())

	// Get Team
	teamGetResp, httpresp, err := TESTAPICLIENT.TeamsApi.GetTeam(TESTAUTH, team.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error getting Team %s: Response %s: %v", team.GetId(), err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	teamGet := teamGetResp.GetData()
	teamAttrsGet := teamGet.GetAttributes()
	assert.Equal(t, teamGet.GetId(), teamUpdated.GetId())
	assert.Equal(t, teamAttrsGet.GetName(), teamAttrsUpdated.GetName())

	// Get Teams

	teamsGetResp, httpresp, err := TESTAPICLIENT.TeamsApi.GetTeams(TESTAUTH).Execute()
	if err != nil {
		t.Fatalf("Error getting Team %s: Response %s: %v", team.GetId(), err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.True(t, len(teamsGetResp.GetData()) >= 1)

}

func createTestTeam(t *testing.T) datadog.Team {
	testTeamData := datadog.NewTeam()
	testTeamData.SetType("teams")
	testTeamData.SetAttributes(*datadog.NewTeamAttributesWithDefaults())
	testTeamData.Attributes.SetName(fmt.Sprintf("TestTeam%s", generateUniqueString(32)))
	// Create Team
	teamRsp, httpresp, err := TESTAPICLIENT.TeamsApi.CreateTeam(TESTAUTH).Body(*datadog.NewTeamPayload(*testTeamData)).Execute()
	if err != nil {
		t.Fatalf("Error creating Team %v: Response %s: %v", testTeamData, err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 201)
	return teamRsp.GetData()
}

func deleteTestTeam(t *testing.T, id string) {
	httpresp, err := TESTAPICLIENT.TeamsApi.DeleteTeam(TESTAUTH, id).Execute()
	if err != nil && httpresp.StatusCode != 404 {
		t.Logf("Deleting Team: %v failed with %v, Another test may have already deleted this team: %v",
			id, httpresp.StatusCode, err)
	}

}
