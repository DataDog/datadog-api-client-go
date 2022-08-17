// Get a list of all incident teams returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "team" in the system
	TeamDataAttributesName := os.Getenv("TEAM_DATA_ATTRIBUTES_NAME")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ListIncidentTeams", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentTeamsApi(apiClient)
	resp, r, err := api.ListIncidentTeams(ctx, *datadogV2.NewListIncidentTeamsOptionalParameters().WithFilter(TeamDataAttributesName))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentTeamsApi.ListIncidentTeams`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentTeamsApi.ListIncidentTeams`:\n%s\n", responseContent)
}
