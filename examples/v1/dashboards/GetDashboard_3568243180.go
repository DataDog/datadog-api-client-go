// Get a dashboard with five team tags and two AI tags

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	// there is a valid "dashboard_with_team_and_ai_tags" in the system
	DashboardWithTeamAndAiTagsID := os.Getenv("DASHBOARD_WITH_TEAM_AND_AI_TAGS_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewDashboardsApi(apiClient)
	resp, r, err := api.GetDashboard(ctx, DashboardWithTeamAndAiTagsID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.GetDashboard`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.GetDashboard`:\n%s\n", responseContent)
}
