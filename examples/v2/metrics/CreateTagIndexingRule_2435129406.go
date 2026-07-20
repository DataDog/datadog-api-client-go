// Create a tag indexing rule with exclude-mode tag usage fields returns "Created" response

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
	body := datadogV2.TagIndexingRuleCreateRequest{
		Data: datadogV2.TagIndexingRuleCreateData{
			Attributes: datadogV2.TagIndexingRuleCreateAttributes{
				ExcludeTagsMode:          datadog.PtrBool(true),
				IgnoredMetricNameMatches: []string{},
				MetricNameMatches: []string{
					"dd.test.*",
				},
				Name: "my-indexing-rule",
				Options: &datadogV2.TagIndexingRuleOptions{
					Data: &datadogV2.TagIndexingRuleOptionsData{
						DynamicTags: &datadogV2.TagIndexingRuleDynamicTags{
							ExcludeNotQueriedWindowSeconds: datadog.PtrInt64(3600),
							ExcludeNotUsedInAssets:         datadog.PtrBool(true),
						},
						ManagePreexistingMetrics: datadog.PtrBool(true),
						OverridePreviousRules:    datadog.PtrBool(false),
					},
					Version: datadog.PtrInt64(1),
				},
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
	configuration.SetUnstableOperationEnabled("v2.CreateTagIndexingRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMetricsApi(apiClient)
	resp, r, err := api.CreateTagIndexingRule(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.CreateTagIndexingRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.CreateTagIndexingRule`:\n%s\n", responseContent)
}
