// Search RUM events returns "OK" response

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
	body := datadog.RUMSearchEventsRequest{
		Filter: &datadog.RUMQueryFilter{
			From:  datadog.PtrString("now-15m"),
			Query: datadog.PtrString("@type:session AND @session.type:user"),
			To:    datadog.PtrString("now"),
		},
		Options: &datadog.RUMQueryOptions{
			TimeOffset: datadog.PtrInt64(0),
			Timezone:   datadog.PtrString("GMT"),
		},
		Page: &datadog.RUMQueryPageOptions{
			Limit: datadog.PtrInt32(25),
		},
		Sort: datadog.RUMSORT_TIMESTAMP_ASCENDING.Ptr(),
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.RUMApi(apiClient)
	resp, r, err := api.SearchRUMEvents(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMApi.SearchRUMEvents`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMApi.SearchRUMEvents`:\n%s\n", responseContent)
}
