// List teams ownership rules returns "OK" response

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
	// there is a valid "teams_ownership_mapping" in the system
	TeamsOwnershipMappingDataAttributesViewName := os.Getenv("TEAMS_OWNERSHIP_MAPPING_DATA_ATTRIBUTES_VIEW_NAME")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumTeamsOwnershipApi(apiClient)
	resp, r, err := api.ListTeamsOwnershipRules(ctx, *datadogV2.NewListTeamsOwnershipRulesOptionalParameters().WithFilterViewName([]string{
		TeamsOwnershipMappingDataAttributesViewName,
	}))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumTeamsOwnershipApi.ListTeamsOwnershipRules`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumTeamsOwnershipApi.ListTeamsOwnershipRules`:\n%s\n", responseContent)
}
