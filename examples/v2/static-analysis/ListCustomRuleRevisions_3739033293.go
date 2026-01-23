// List Custom Rule Revisions returns "Successful response" response with pagination

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
	configuration.SetUnstableOperationEnabled("v2.ListCustomRuleRevisions", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStaticAnalysisApi(apiClient)
	resp, _ := api.ListCustomRuleRevisionsWithPagination(ctx, "ruleset_name", "rule_name", *datadogV2.NewListCustomRuleRevisionsOptionalParameters())

	for paginationResult := range resp {
		if paginationResult.Error != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `StaticAnalysisApi.ListCustomRuleRevisions`: %v\n", paginationResult.Error)
		}
		responseContent, _ := json.MarshalIndent(paginationResult.Item, "", "  ")
		fmt.Fprintf(os.Stdout, "%s\n", responseContent)
	}
}
