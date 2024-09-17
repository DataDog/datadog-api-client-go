// Search tests events returns "OK" response

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
	body := datadogV2.CIAppTestEventsRequest{
		Filter: &datadogV2.CIAppTestsQueryFilter{
			From:  datadog.PtrString("now-15m"),
			Query: datadog.PtrString("@test.service:web-ui-tests AND @test.status:skip"),
			To:    datadog.PtrString("now"),
		},
		Options: &datadogV2.CIAppQueryOptions{
			Timezone: datadog.PtrString("GMT"),
		},
		Page: &datadogV2.CIAppQueryPageOptions{
			Limit: datadog.PtrInt32(25),
		},
		Sort: datadogV2.CIAPPSORT_TIMESTAMP_ASCENDING.Ptr(),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCIVisibilityTestsApi(apiClient)
	resp, r, err := api.SearchCIAppTestEvents(ctx, *datadogV2.NewSearchCIAppTestEventsOptionalParameters().WithBody(body))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CIVisibilityTestsApi.SearchCIAppTestEvents`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CIVisibilityTestsApi.SearchCIAppTestEvents`:\n%s\n", responseContent)
}
