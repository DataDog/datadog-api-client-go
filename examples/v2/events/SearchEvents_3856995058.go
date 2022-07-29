// Search events returns "OK" response with pagination

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	body := datadog.EventsListRequest{
		Filter: &datadog.EventsQueryFilter{
			From: common.PtrString("now-15m"),
			To:   common.PtrString("now"),
		},
		Options: &datadog.EventsQueryOptions{
			Timezone: common.PtrString("GMT"),
		},
		Page: &datadog.EventsRequestPage{
			Limit: common.PtrInt32(2),
		},
		Sort: datadog.EVENTSSORT_TIMESTAMP_ASCENDING.Ptr(),
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.SearchEvents", true)
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewEventsApi(apiClient)
	resp, _, err := api.SearchEventsWithPagination(ctx, *datadog.NewSearchEventsOptionalParameters().WithBody(body))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.SearchEvents`: %v\n", err)
	}

	for item := range resp {
		responseContent, _ := json.MarshalIndent(item, "", "  ")
		fmt.Fprintf(os.Stdout, "%s\n", responseContent)
	}
}
