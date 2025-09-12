// Search flaky tests returns "OK" response with specific cursor

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
					Cursor: datadog.PtrString("Q29udGludW91cyBUZXN0aW5nLltETyBOT1QgREVMRVRFXVtPTi1ERU1BTkQgRlVOQ1RJT05BTElUSUVTXVtPVkVSUklERV0gRXh0cmFWYXJpYWJsZXM="),
					Limit:  datadog.PtrInt64(2),
				},
				Sort: datadogV2.FLAKYTESTSSEARCHSORT_FQN_ASCENDING.Ptr(),
			},
			Type: datadogV2.FLAKYTESTSSEARCHREQUESTDATATYPE_SEARCH_FLAKY_TESTS_REQUEST.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.SearchFlakyTests", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTestOptimizationApi(apiClient)
	resp, r, err := api.SearchFlakyTests(ctx, *datadogV2.NewSearchFlakyTestsOptionalParameters().WithBody(body))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestOptimizationApi.SearchFlakyTests`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TestOptimizationApi.SearchFlakyTests`:\n%s\n", responseContent)
}
