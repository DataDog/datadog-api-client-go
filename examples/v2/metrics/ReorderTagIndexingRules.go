// Reorder tag indexing rules returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "tag_indexing_rule" in the system
	TagIndexingRuleDataID := os.Getenv("TAG_INDEXING_RULE_DATA_ID")

	body := datadogV2.TagIndexingRuleOrderRequest{
		Data: datadogV2.TagIndexingRuleOrderData{
			Attributes: datadogV2.TagIndexingRuleOrderAttributes{
				RuleIds: []string{
					TagIndexingRuleDataID,
				},
			},
			Type: datadogV2.TAGINDEXINGRULETYPE_TAG_INDEXING_RULES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ReorderTagIndexingRules", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMetricsApi(apiClient)
	r, err := api.ReorderTagIndexingRules(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.ReorderTagIndexingRules`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
