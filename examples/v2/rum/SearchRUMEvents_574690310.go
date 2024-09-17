// Search RUM events returns "OK" response with pagination

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
	body := datadogV2.RUMSearchEventsRequest{
		Filter: &datadogV2.RUMQueryFilter{
			From:  datadog.PtrString("now-15m"),
			Query: datadog.PtrString("@type:session AND @session.type:user"),
			To:    datadog.PtrString("now"),
		},
		Options: &datadogV2.RUMQueryOptions{
			TimeOffset: datadog.PtrInt64(0),
			Timezone:   datadog.PtrString("GMT"),
		},
		Page: &datadogV2.RUMQueryPageOptions{
			Limit: datadog.PtrInt32(2),
		},
		Sort: datadogV2.RUMSORT_TIMESTAMP_ASCENDING.Ptr(),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMApi(apiClient)
	resp, _ := api.SearchRUMEventsWithPagination(ctx, body)

	for paginationResult := range resp {
		if paginationResult.Error != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `RUMApi.SearchRUMEvents`: %v\n", paginationResult.Error)
		}
		responseContent, _ := json.MarshalIndent(paginationResult.Item, "", "  ")
		fmt.Fprintf(os.Stdout, "%s\n", responseContent)
	}
}
