// Search pipelines events returns "OK" response

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
	body := datadogV2.CIAppPipelineEventsRequest{
		Filter: &datadogV2.CIAppPipelinesQueryFilter{
			From:  datadog.PtrString("now-15m"),
			Query: datadog.PtrString("@ci.provider.name:github AND @ci.status:error"),
			To:    datadog.PtrString("now"),
		},
		Options: &datadogV2.CIAppQueryOptions{
			Timezone: datadog.PtrString("GMT"),
		},
		Page: &datadogV2.CIAppQueryPageOptions{
			Limit: datadog.PtrInt32(5),
		},
		Sort: datadogV2.CIAPPSORT_TIMESTAMP_ASCENDING.Ptr(),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCIVisibilityPipelinesApi(apiClient)
	resp, r, err := api.SearchCIAppPipelineEvents(ctx, *datadogV2.NewSearchCIAppPipelineEventsOptionalParameters().WithBody(body))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CIVisibilityPipelinesApi.SearchCIAppPipelineEvents`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CIVisibilityPipelinesApi.SearchCIAppPipelineEvents`:\n%s\n", responseContent)
}
