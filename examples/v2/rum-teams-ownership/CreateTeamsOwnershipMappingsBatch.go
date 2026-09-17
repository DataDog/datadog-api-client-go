// Bulk create and remove teams ownership mappings returns "OK" response

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
	body := datadogV2.TeamsOwnershipMappingBatchRequest{
		AtomicOperations: []datadogV2.TeamsOwnershipMappingBatchOperation{
			{
				Op: datadogV2.TEAMSOWNERSHIPMAPPINGBATCHOPERATIONOP_ADD,
				Data: &datadogV2.TeamsOwnershipMappingBatchOperationData{
					Type: datadogV2.TEAMSOWNERSHIPMAPPINGTYPE_TEAMS_OWNERSHIP_MAPPINGS,
					Attributes: datadogV2.TeamsOwnershipMappingBatchOperationDataAttributes{
						TeamHandle: datadog.PtrString("team-rum"),
						ViewName:   datadog.PtrString("/checkout-examplerumteamsownership"),
						Service:    datadog.PtrString("web-checkout-examplerumteamsownership"),
						MatchType:  datadogV2.TEAMSOWNERSHIPMATCHTYPE_EXACT.Ptr(),
					},
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumTeamsOwnershipApi(apiClient)
	resp, r, err := api.CreateTeamsOwnershipMappingsBatch(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumTeamsOwnershipApi.CreateTeamsOwnershipMappingsBatch`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumTeamsOwnershipApi.CreateTeamsOwnershipMappingsBatch`:\n%s\n", responseContent)
}
