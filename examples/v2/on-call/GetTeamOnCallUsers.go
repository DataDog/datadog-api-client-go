// Get team on-call users returns "OK" response

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
	// there are valid "routing_rules" in the system
	RoutingRulesDataID := os.Getenv("ROUTING_RULES_DATA_ID")


	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOnCallApi(apiClient)
	resp, r, err := api.GetTeamOnCallUsers(ctx, RoutingRulesDataID, *datadogV2.NewGetTeamOnCallUsersOptionalParameters().WithInclude("responders,escalations.responders"), )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnCallApi.GetTeamOnCallUsers`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OnCallApi.GetTeamOnCallUsers`:\n%s\n", responseContent)
}