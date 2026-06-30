// Create a tag indexing rule exemption returns "Created" response

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
	body := datadogV2.TagIndexingRuleExemptionCreateRequest{
		Data: datadogV2.TagIndexingRuleExemptionCreateData{
			Attributes: datadogV2.TagIndexingRuleExemptionCreateAttributes{
				Reason: "This metric has a pre-existing tag configuration.",
			},
			Type: datadogV2.TAGINDEXINGRULEEXEMPTIONTYPE_TAG_INDEXING_RULE_EXEMPTIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateTagIndexingRuleExemption", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMetricsApi(apiClient)
	resp, r, err := api.CreateTagIndexingRuleExemption(ctx, "metric_name", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.CreateTagIndexingRuleExemption`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.CreateTagIndexingRuleExemption`:\n%s\n", responseContent)
}
