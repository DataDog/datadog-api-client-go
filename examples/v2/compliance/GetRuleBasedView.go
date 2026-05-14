// Get the rule-based view of compliance findings returns "OK" response

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetRuleBasedView", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewComplianceApi(apiClient)
	resp, r, err := api.GetRuleBasedView(ctx, 1739982278000, *datadogV2.NewGetRuleBasedViewOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceApi.GetRuleBasedView`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ComplianceApi.GetRuleBasedView`:\n%s\n", responseContent)
}
