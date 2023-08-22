// Create a team returns "CREATED" response

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
	body := datadogV2.TeamCreateRequest{
		Data: datadogV2.TeamCreate{
			Attributes: datadogV2.TeamCreateAttributes{
				Handle: "test-handle-a0fc0297eb519635",
				Name:   "test-name-a0fc0297eb519635",
			},
			Relationships: &datadogV2.TeamCreateRelationships{
				Users: &datadogV2.RelationshipToUsers{
					Data: []datadogV2.RelationshipToUserData{},
				},
			},
			Type: datadogV2.TEAMTYPE_TEAM,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTeamsApi(apiClient)
	resp, r, err := api.CreateTeam(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.CreateTeam`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TeamsApi.CreateTeam`:\n%s\n", responseContent)
}
