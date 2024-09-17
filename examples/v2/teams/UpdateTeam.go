// Update a team returns "OK" response

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
	// there is a valid "dd_team" in the system
	DdTeamDataAttributesHandle := os.Getenv("DD_TEAM_DATA_ATTRIBUTES_HANDLE")
	DdTeamDataID := os.Getenv("DD_TEAM_DATA_ID")

	body := datadogV2.TeamUpdateRequest{
		Data: datadogV2.TeamUpdate{
			Attributes: datadogV2.TeamUpdateAttributes{
				Handle: DdTeamDataAttributesHandle,
				Name:   "Example Team updated",
				Avatar: *datadog.NewNullableString(datadog.PtrString("🥑")),
				Banner: *datadog.NewNullableInt64(datadog.PtrInt64(7)),
				HiddenModules: []string{
					"m3",
				},
				VisibleModules: []string{
					"m1",
					"m2",
				},
			},
			Type: datadogV2.TEAMTYPE_TEAM,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTeamsApi(apiClient)
	resp, r, err := api.UpdateTeam(ctx, DdTeamDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.UpdateTeam`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TeamsApi.UpdateTeam`:\n%s\n", responseContent)
}
