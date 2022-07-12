// Delete an existing incident team returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "team" in the system
	TeamDataID := os.Getenv("TEAM_DATA_ID")

	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteIncidentTeam", true)
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewIncidentTeamsApi(apiClient)
	r, err := api.DeleteIncidentTeam(ctx, TeamDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentTeamsApi.DeleteIncidentTeam`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
