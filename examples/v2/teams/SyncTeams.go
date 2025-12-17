// Link Teams with GitHub Teams returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.TeamSyncRequest{
		Data: datadogV2.TeamSyncData{
			Attributes: datadogV2.TeamSyncAttributes{
				Source: datadogV2.TEAMSYNCATTRIBUTESSOURCE_GITHUB,
				Type:   datadogV2.TEAMSYNCATTRIBUTESTYPE_LINK,
			},
			Type: datadogV2.TEAMSYNCBULKTYPE_TEAM_SYNC_BULK,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTeamsApi(apiClient)
	r, err := api.SyncTeams(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.SyncTeams`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
