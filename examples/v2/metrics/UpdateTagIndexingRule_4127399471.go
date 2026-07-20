// Update a tag indexing rule with exclude-mode tag usage fields returns "OK" response

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
	// there is a valid "tag_indexing_rule_exclude_mode" in the system
	TagIndexingRuleExcludeModeDataID := os.Getenv("TAG_INDEXING_RULE_EXCLUDE_MODE_DATA_ID")

	body := datadogV2.TagIndexingRuleUpdateRequest{
		Data: datadogV2.TagIndexingRuleUpdateData{
			Attributes: &datadogV2.TagIndexingRuleUpdateAttributes{
				ExcludeTagsMode:          datadog.PtrBool(true),
				IgnoredMetricNameMatches: []string{},
				MetricNameMatches: []string{
					"dd.test.*",
				},
				Name: datadog.PtrString("my-indexing-rule"),
				Options: &datadogV2.TagIndexingRuleOptions{
					Data: &datadogV2.TagIndexingRuleOptionsData{
						DynamicTags: &datadogV2.TagIndexingRuleDynamicTags{
							ExcludeNotQueriedWindowSeconds: datadog.PtrInt64(7200),
							ExcludeNotUsedInAssets:         datadog.PtrBool(true),
						},
						ManagePreexistingMetrics: datadog.PtrBool(true),
						OverridePreviousRules:    datadog.PtrBool(false),
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
	configuration.SetUnstableOperationEnabled("v2.UpdateTagIndexingRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMetricsApi(apiClient)
	resp, r, err := api.UpdateTagIndexingRule(ctx, TagIndexingRuleExcludeModeDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.UpdateTagIndexingRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.UpdateTagIndexingRule`:\n%s\n", responseContent)
}
