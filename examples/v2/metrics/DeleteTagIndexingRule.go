// Delete a tag indexing rule returns "No Content" response

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

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteTagIndexingRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMetricsApi(apiClient)
	r, err := api.DeleteTagIndexingRule(ctx, TagIndexingRuleDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.DeleteTagIndexingRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
