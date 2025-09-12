// Search flaky tests returns "OK" response with cursor pagination

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
	body := datadogV2.FlakyTestsSearchRequest{
		Data: &datadogV2.FlakyTestsSearchRequestData{
			Attributes: &datadogV2.FlakyTestsSearchRequestAttributes{
				Filter: &datadogV2.FlakyTestsSearchFilter{
					Query: datadog.PtrString("*"),
				},
				Page: &datadogV2.FlakyTestsSearchPageOptions{
					Limit: datadog.PtrInt64(2),
				},
				Sort: datadogV2.FLAKYTESTSSEARCHSORT_FQN_ASCENDING.Ptr(),
			},
			Type: datadogV2.FLAKYTESTSSEARCHREQUESTDATATYPE_SEARCH_FLAKY_TESTS_REQUEST.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTestOptimizationApi(apiClient)
	resp, _ := api.SearchFlakyTestsWithPagination(ctx, *datadogV2.NewSearchFlakyTestsOptionalParameters().WithBody(body))

	for paginationResult := range resp {
		if paginationResult.Error != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `TestOptimizationApi.SearchFlakyTests`: %v\n", paginationResult.Error)
		}
		responseContent, _ := json.MarshalIndent(paginationResult.Item, "", "  ")
		fmt.Fprintf(os.Stdout, "%s\n", responseContent)
	}
}
