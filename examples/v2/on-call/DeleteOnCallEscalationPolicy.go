// Delete On-Call escalation policy returns "No Content" response

package main

import (
	"context"
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
	r, err := api.DeleteOnCallEscalationPolicy(ctx, EscalationPolicyDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnCallApi.DeleteOnCallEscalationPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
