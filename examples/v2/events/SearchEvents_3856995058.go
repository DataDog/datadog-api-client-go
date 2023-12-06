// Search events returns "OK" response with pagination

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
	body := datadogV2.EventsListRequest{
		Filter: &datadogV2.EventsQueryFilter{
			From: datadog.PtrString("now-15m"),
			To:   datadog.PtrString("now"),
		},
		Options: &datadogV2.EventsQueryOptions{
			Timezone: datadog.PtrString("GMT"),
		},
		Page: &datadogV2.EventsRequestPage{
			Limit: datadog.PtrInt32(2),
		},
		Sort: datadogV2.EVENTSSORT_TIMESTAMP_ASCENDING.Ptr(),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewEventsApi(apiClient)
	resp, _ := api.SearchEventsWithPagination(ctx, *datadogV2.NewSearchEventsOptionalParameters().WithBody(body))

	for paginationResult := range resp {
		if paginationResult.Error != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.SearchEvents`: %v\n", paginationResult.Error)
		}
		responseContent, _ := json.MarshalIndent(paginationResult.Item, "", "  ")
		fmt.Fprintf(os.Stdout, "%s\n", responseContent)
	}
}
