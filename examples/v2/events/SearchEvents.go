// Search events returns "OK" response

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
			Query: common.PtrString("datadog-agent"),
			From:  common.PtrString("2020-09-17T11:48:36+01:00"),
			To:    common.PtrString("2020-09-17T12:48:36+01:00"),
		},
		Sort: datadog.EVENTSSORT_TIMESTAMP_ASCENDING.Ptr(),
		Page: &datadog.EventsRequestPage{
			Limit: common.PtrInt32(5),
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.SearchEvents", true)
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewEventsApi(apiClient)
	resp, r, err := api.SearchEvents(ctx, *datadog.NewSearchEventsOptionalParameters().WithBody(body))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.SearchEvents`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `EventsApi.SearchEvents`:\n%s\n", responseContent)
}
