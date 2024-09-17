// Get all teams with fields_team parameter returns "OK" response

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTeamsApi(apiClient)
	resp, r, err := api.ListTeams(ctx, *datadogV2.NewListTeamsOptionalParameters().WithFieldsTeam([]datadogV2.TeamsField{
		datadogV2.TEAMSFIELD_ID,
		datadogV2.TEAMSFIELD_NAME,
		datadogV2.TEAMSFIELD_HANDLE,
	}))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.ListTeams`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TeamsApi.ListTeams`:\n%s\n", responseContent)
}
