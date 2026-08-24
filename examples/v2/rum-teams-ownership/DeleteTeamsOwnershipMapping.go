// Delete a teams ownership mapping returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "teams_ownership_mapping" in the system
	TeamsOwnershipMappingDataID := os.Getenv("TEAMS_OWNERSHIP_MAPPING_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumTeamsOwnershipApi(apiClient)
	r, err := api.DeleteTeamsOwnershipMapping(ctx, TeamsOwnershipMappingDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumTeamsOwnershipApi.DeleteTeamsOwnershipMapping`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
