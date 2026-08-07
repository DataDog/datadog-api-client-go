// Get a tag rule returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.GetTagRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTagRulesApi(apiClient)
	resp, r, err := api.GetTagRule(ctx, "policy_id", *datadogV2.NewGetTagRuleOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagRulesApi.GetTagRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TagRulesApi.GetTagRule`:\n%s\n", responseContent)
}
