// Get On-Call escalation policy returns "OK" response

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
	// there is a valid "escalation_policy" in the system
	EscalationPolicyDataID := os.Getenv("ESCALATION_POLICY_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOnCallApi(apiClient)
	resp, r, err := api.GetOnCallEscalationPolicy(ctx, EscalationPolicyDataID, *datadogV2.NewGetOnCallEscalationPolicyOptionalParameters().WithInclude("steps.targets"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnCallApi.GetOnCallEscalationPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OnCallApi.GetOnCallEscalationPolicy`:\n%s\n", responseContent)
}
