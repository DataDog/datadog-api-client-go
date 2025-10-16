// Get all member teams returns "OK" response with pagination

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ListMemberTeams", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTeamsApi(apiClient)
	resp, _ := api.ListMemberTeamsWithPagination(ctx, "super_team_id", *datadogV2.NewListMemberTeamsOptionalParameters(), )

        for paginationResult := range resp {
            if paginationResult.Error != nil {
                fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.ListMemberTeams`: %v\n", paginationResult.Error)
            }
            responseContent, _ := json.MarshalIndent(paginationResult.Item, "", "  ")
            fmt.Fprintf(os.Stdout, "%s\n", responseContent)
	}
}