// Create a teams ownership mapping returns "Created" response

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
	body := datadogV2.TeamsOwnershipMappingCreateRequest{
		Data: datadogV2.TeamsOwnershipMappingCreateData{
			Attributes: datadogV2.TeamsOwnershipMappingCreateDataAttributes{
				ApplicationId: datadog.PtrUUID(uuid.MustParse("11111111-2222-3333-4444-555555555555")),
				MatchType:     datadogV2.TEAMSOWNERSHIPMATCHTYPE_EXACT.Ptr(),
				Service:       datadog.PtrString("web-checkout"),
				TeamHandle:    "team-rum",
				ViewName:      "/checkout",
			},
			Type: datadogV2.TEAMSOWNERSHIPMAPPINGTYPE_TEAMS_OWNERSHIP_MAPPINGS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateTeamsOwnershipMapping", true)
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
