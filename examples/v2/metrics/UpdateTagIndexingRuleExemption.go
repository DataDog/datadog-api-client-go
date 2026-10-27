// Update a tag indexing rule exemption returns "OK" response

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
	body := datadogV2.TagIndexingRuleExemptionUpdateRequest{
		Data: datadogV2.TagIndexingRuleExemptionUpdateData{
			Attributes: datadogV2.TagIndexingRuleExemptionUpdateAttributes{
				Reason: "This metric must remain exempt because of its query requirements.",
			},
			Type: datadogV2.TAGINDEXINGRULEEXEMPTIONTYPE_TAG_INDEXING_RULE_EXEMPTIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateTagIndexingRuleExemption", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMetricsApi(apiClient)
	resp, r, err := api.UpdateTagIndexingRuleExemption(ctx, "metric_name", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.UpdateTagIndexingRuleExemption`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.UpdateTagIndexingRuleExemption`:\n%s\n", responseContent)
}
