// Search pipelines events returns "OK" response with pagination

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
			From: datadog.PtrString("now-30s"),
			To:   datadog.PtrString("now"),
		},
		Options: &datadogV2.CIAppQueryOptions{
			Timezone: datadog.PtrString("GMT"),
		},
		Page: &datadogV2.CIAppQueryPageOptions{
			Limit: datadog.PtrInt32(2),
		},
		Sort: datadogV2.CIAPPSORT_TIMESTAMP_ASCENDING.Ptr(),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCIVisibilityPipelinesApi(apiClient)
	resp, _, err := api.SearchCIAppPipelineEventsWithPagination(ctx, *datadogV2.NewSearchCIAppPipelineEventsOptionalParameters().WithBody(body))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CIVisibilityPipelinesApi.SearchCIAppPipelineEvents`: %v\n", err)
	}

	for item := range resp {
		responseContent, _ := json.MarshalIndent(item, "", "  ")
		fmt.Fprintf(os.Stdout, "%s\n", responseContent)
	}
}
