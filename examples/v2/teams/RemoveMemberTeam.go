// Remove a member team returns "No Content" response

package main


import (
	"context"
	"encoding/json"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.RemoveMemberTeam", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTeamsApi(apiClient)
	r, err := api.RemoveMemberTeam(ctx, "super_team_id", "member_team_id", )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.RemoveMemberTeam`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}