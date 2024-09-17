// Get team memberships returns "Represents a user's association to a team" response with pagination

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
	resp, _ := api.GetTeamMembershipsWithPagination(ctx, "2e06bf2c-193b-41d4-b3c2-afccc080458f", *datadogV2.NewGetTeamMembershipsOptionalParameters().WithPageSize(2))

	for paginationResult := range resp {
		if paginationResult.Error != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.GetTeamMemberships`: %v\n", paginationResult.Error)
		}
		responseContent, _ := json.MarshalIndent(paginationResult.Item, "", "  ")
		fmt.Fprintf(os.Stdout, "%s\n", responseContent)
	}
}
