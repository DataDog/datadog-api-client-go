// Update a tag indexing rule returns "OK" response

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
	// there is a valid "tag_indexing_rule" in the system
	TagIndexingRuleDataID := os.Getenv("TAG_INDEXING_RULE_DATA_ID")

	body := datadogV2.TagIndexingRuleUpdateRequest{
		Data: datadogV2.TagIndexingRuleUpdateData{
			Attributes: &datadogV2.TagIndexingRuleUpdateAttributes{
				IgnoredMetricNameMatches: []string{},
				MetricNameMatches: []string{
					"dd.test.*",
				},
				Name: datadog.PtrString("my-indexing-rule"),
				Options: &datadogV2.TagIndexingRuleOptions{
					Data: &datadogV2.TagIndexingRuleOptionsData{
						DynamicTags: &datadogV2.TagIndexingRuleDynamicTags{
							QueriedTagsWindowSeconds: datadog.PtrInt64(3600),
							RelatedAssetTags:         datadog.PtrBool(false),
						},
						ManagePreexistingMetrics: datadog.PtrBool(true),
						MetricMatch: &datadogV2.TagIndexingRuleMetricMatch{
							QueriedWindowSeconds: datadog.PtrInt64(3600),
						},
						OverridePreviousRules: datadog.PtrBool(false),
					},
					Version: datadog.PtrInt64(1),
				},
				RuleOrder: datadog.PtrInt64(2),
				Tags: []string{
					"env",
					"service",
				},
			},
			Type: datadogV2.TAGINDEXINGRULETYPE_TAG_INDEXING_RULES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMetricsApi(apiClient)
	resp, r, err := api.UpdateTagIndexingRule(ctx, TagIndexingRuleDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.UpdateTagIndexingRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.UpdateTagIndexingRule`:\n%s\n", responseContent)
}
