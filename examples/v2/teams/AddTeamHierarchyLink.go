// Create a team hierarchy link returns "CREATED" response

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
	body := datadogV2.TeamHierarchyLinkCreateRequest{
		Data: datadogV2.TeamHierarchyLinkCreate{
			Relationships: datadogV2.TeamHierarchyLinkCreateRelationships{
				ParentTeam: datadogV2.TeamHierarchyLinkCreateTeamRelationship{
					Data: datadogV2.TeamHierarchyLinkCreateTeam{
						Id:   "692e8073-12c4-4c71-8408-5090bd44c9c8",
						Type: datadogV2.TEAMTYPE_TEAM,
					},
				},
				SubTeam: datadogV2.TeamHierarchyLinkCreateTeamRelationship{
					Data: datadogV2.TeamHierarchyLinkCreateTeam{
						Id:   "692e8073-12c4-4c71-8408-5090bd44c9c8",
						Type: datadogV2.TEAMTYPE_TEAM,
					},
				},
			},
			Type: datadogV2.TEAMHIERARCHYLINKTYPE_TEAM_HIERARCHY_LINKS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTeamsApi(apiClient)
	resp, r, err := api.AddTeamHierarchyLink(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.AddTeamHierarchyLink`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TeamsApi.AddTeamHierarchyLink`:\n%s\n", responseContent)
}
