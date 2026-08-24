// Create teams ownership mapping returns "Created" response

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
	body := datadogV2.TeamsOwnershipMappingCreateRequest{
		Data: datadogV2.TeamsOwnershipMappingCreateData{
			Type: datadogV2.TEAMSOWNERSHIPMAPPINGTYPE_TEAMS_OWNERSHIP_MAPPINGS,
			Attributes: datadogV2.TeamsOwnershipMappingCreateDataAttributes{
				TeamHandle: "team-rum",
				ViewName:   "/checkout-examplerumteamsownership",
				Service:    datadog.PtrString("web-checkout-examplerumteamsownership"),
				MatchType:  datadogV2.TEAMSOWNERSHIPMATCHTYPE_EXACT.Ptr(),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumTeamsOwnershipApi(apiClient)
	resp, r, err := api.CreateTeamsOwnershipMapping(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumTeamsOwnershipApi.CreateTeamsOwnershipMapping`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumTeamsOwnershipApi.CreateTeamsOwnershipMapping`:\n%s\n", responseContent)
}
