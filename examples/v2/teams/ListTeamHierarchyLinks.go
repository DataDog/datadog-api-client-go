// Get team hierarchy links returns "OK" response

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
	// there is a valid "team_hierarchy_link" in the system
	TeamHierarchyLinkDataRelationshipsParentTeamDataID := os.Getenv("TEAM_HIERARCHY_LINK_DATA_RELATIONSHIPS_PARENT_TEAM_DATA_ID")
	TeamHierarchyLinkDataRelationshipsSubTeamDataID := os.Getenv("TEAM_HIERARCHY_LINK_DATA_RELATIONSHIPS_SUB_TEAM_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTeamsApi(apiClient)
	resp, r, err := api.ListTeamHierarchyLinks(ctx, *datadogV2.NewListTeamHierarchyLinksOptionalParameters().WithFilterParentTeam(TeamHierarchyLinkDataRelationshipsParentTeamDataID).WithFilterSubTeam(TeamHierarchyLinkDataRelationshipsSubTeamDataID).WithPageNumber(0).WithPageSize(100))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.ListTeamHierarchyLinks`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TeamsApi.ListTeamHierarchyLinks`:\n%s\n", responseContent)
}
