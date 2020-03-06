/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
)

func TestTeamsLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	client := Client(ctx)

	setIncidentsUnstableOperations(client, true)
	defer setIncidentsUnstableOperations(client, false)

	testTeamData := datadog.NewTeam()
	testTeamData.SetType("teams")
	testTeamData.SetAttributes(*datadog.NewTeamAttributesWithDefaults())
	testTeamData.Attributes.SetName("Test-Team")
	testTeamAttributes := testTeamData.GetAttributes()
	// Create Team
	teamRsp, httpresp, err := client.TeamsApi.CreateTeam(ctx).Body(*datadog.NewTeamPayload(*testTeamData)).
		Execute()
	if err != nil {
		t.Fatalf("Error creating Team %v: Response %s: %v", testTeamData, err.Error(), err)
	}
	team := teamRsp.GetData()
	teamAttrs := team.GetAttributes()
	assert.Equal(httpresp.StatusCode, 201)

	// Lets ensure that we always delete the team if we successfully created it, even if the update and
	// get requests fail
	defer func(test_client *datadog.APIClient, team datadog.Team) {
		//Delete Team
		httpresp, err := test_client.TeamsApi.DeleteTeam(ctx, team.GetId()).Execute()
		if err != nil {
			t.Fatalf("Error deleting Team %s: Response %s: %v", team.GetId(), err.Error(), err)
		}
		assert.Equal(httpresp.StatusCode, 204)
	}(client, team)

	assert.Equal(team.GetType(), testTeamData.GetType())
	assert.Equal(teamAttrs.GetName(), testTeamAttributes.GetName())

	// Edit Team
	team.Attributes.SetName("Test-Team-Updated")
	teamUpdatedRsp, httpresp, err := client.TeamsApi.PatchTeam(ctx, team.GetId()).Body(*datadog.NewTeamPayload(team)).Execute()
	if err != nil {
		t.Fatalf("Error updating Team %v: Response %s: %v", team, err.Error(), err)
	}
	teamUpdated := teamUpdatedRsp.GetData()
	teamAttrsUpdated := teamUpdated.GetAttributes()
	assert.Equal(httpresp.StatusCode, 200)
	assert.Equal(teamAttrsUpdated.GetName(), teamAttrsUpdated.GetName())

	// Get Team
	teamGetResp, httpresp, err := client.TeamsApi.GetTeam(ctx, team.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error getting Team %s: Response %s: %v", team.GetId(), err.Error(), err)
	}
	assert.Equal(httpresp.StatusCode, 200)

	teamGet := teamGetResp.GetData()
	teamAttrsGet := teamGet.GetAttributes()
	assert.Equal(teamGet.GetId(), teamUpdated.GetId())
	assert.Equal(teamAttrsGet.GetName(), teamAttrsUpdated.GetName())

	// Get Teams

	teamsGetResp, httpresp, err := client.TeamsApi.GetTeams(ctx).Execute()
	if err != nil {
		t.Fatalf("Error getting Team %s: Response %s: %v", team.GetId(), err.Error(), err)
	}
	assert.Equal(httpresp.StatusCode, 200)
	assert.True(len(teamsGetResp.GetData()) >= 1)

}

var teamNameCounter int64 = time.Now().UnixNano()

func createTestTeam(t *testing.T, client *datadog.APIClient, ctx context.Context) datadog.Team {
	atomic.AddInt64(&teamNameCounter, 1)
	testTeamData := datadog.NewTeam()
	testTeamData.SetType("teams")
	testTeamData.SetAttributes(*datadog.NewTeamAttributesWithDefaults())
	testTeamData.Attributes.SetName(fmt.Sprintf("createTestTeam-%s", RandomString(16)))
	// Create Team
	teamRsp, httpresp, err := client.TeamsApi.CreateTeam(ctx).Body(*datadog.NewTeamPayload(*testTeamData)).Execute()
	if err != nil {
		t.Fatalf("Error creating Team %v: Response %s: %v", testTeamData, err.Error(), err)
	}
	assert := tests.Assert(ctx, t)
	assert.Equal(httpresp.StatusCode, 201)
	return teamRsp.GetData()
}

func deleteTestTeam(t *testing.T, client *datadog.APIClient, ctx context.Context, id string) {
	httpresp, err := client.TeamsApi.DeleteTeam(ctx, id).Execute()
	if err != nil && httpresp.StatusCode != 404 {
		t.Logf("Deleting Team: %v failed with %v, Another test may have already deleted this team: %v",
			id, httpresp.StatusCode, err)
	}

}
