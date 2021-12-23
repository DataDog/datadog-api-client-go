// Get a list of all incident teams returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "team" in the system
	TeamDataAttributesName := os.Getenv("TEAM_DATA_ATTRIBUTES_NAME")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("ListIncidentTeams", true)
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentTeamsApi.ListIncidentTeams(ctx, *datadog.NewListIncidentTeamsOptionalParameters().WithFilter(TeamDataAttributesName))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentTeamsApi.ListIncidentTeams`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentTeamsApi.ListIncidentTeams`:\n%s\n", responseContent)
}
