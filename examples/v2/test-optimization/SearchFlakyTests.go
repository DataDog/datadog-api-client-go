// Search flaky tests returns "OK" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.FlakyTestsSearchRequest{
Data: &datadogV2.FlakyTestsSearchRequestData{
Attributes: &datadogV2.FlakyTestsSearchRequestAttributes{
Filter: &datadogV2.FlakyTestsSearchFilter{
Query: datadog.PtrString(`flaky_test_state:active @git.repository.id_v2:"github.com/datadog/shopist"`),
},
Page: &datadogV2.FlakyTestsSearchPageOptions{
Cursor: datadog.PtrString("eyJzdGFydEF0IjoiQVFBQUFYS2tMS3pPbm40NGV3QUFBQUJCV0V0clRFdDZVbG8zY3pCRmNsbHJiVmxDWlEifQ=="),
Limit: datadog.PtrInt64(25),
},
Sort: datadogV2.FLAKYTESTSSEARCHSORT_FAILURE_RATE_ASCENDING.Ptr(),
},
Type: datadogV2.FLAKYTESTSSEARCHREQUESTDATATYPE_SEARCH_FLAKY_TESTS_REQUEST.Ptr(),
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.SearchFlakyTests", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTestOptimizationApi(apiClient)
	resp, r, err := api.SearchFlakyTests(ctx, *datadogV2.NewSearchFlakyTestsOptionalParameters().WithBody(body), )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestOptimizationApi.SearchFlakyTests`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TestOptimizationApi.SearchFlakyTests`:\n%s\n", responseContent)
}