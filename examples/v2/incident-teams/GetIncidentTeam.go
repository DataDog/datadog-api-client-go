// Get details of an incident team returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "team" in the system
	TeamDataID := os.Getenv("TEAM_DATA_ID")

	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	configuration.SetUnstableOperationEnabled("GetIncidentTeam", true)
	apiClient := common.NewAPIClient(configuration)
	api := datadog.IncidentTeamsApi(apiClient)
	resp, r, err := api.GetIncidentTeam(ctx, TeamDataID, *datadog.NewGetIncidentTeamOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentTeamsApi.GetIncidentTeam`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentTeamsApi.GetIncidentTeam`:\n%s\n", responseContent)
}
