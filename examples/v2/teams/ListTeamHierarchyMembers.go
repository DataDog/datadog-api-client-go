// Get team hierarchy members returns "OK" response

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
	DdTeamDataID := os.Getenv("DD_TEAM_DATA_ID")

	// there is a valid "user" in the system
	UserDataAttributesEmail := os.Getenv("USER_DATA_ATTRIBUTES_EMAIL")

	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ListTeamHierarchyMembers", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTeamsApi(apiClient)
	resp, r, err := api.ListTeamHierarchyMembers(ctx, DdTeamDataID, *datadogV2.NewListTeamHierarchyMembersOptionalParameters().WithFilterKeyword(UserDataAttributesEmail).WithInclude("user,team"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.ListTeamHierarchyMembers`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TeamsApi.ListTeamHierarchyMembers`:\n%s\n", responseContent)
}
