// Remove a team hierarchy link returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "team_hierarchy_link" in the system
	TeamHierarchyLinkDataID := os.Getenv("TEAM_HIERARCHY_LINK_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTeamsApi(apiClient)
	r, err := api.RemoveTeamHierarchyLink(ctx, TeamHierarchyLinkDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.RemoveTeamHierarchyLink`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
