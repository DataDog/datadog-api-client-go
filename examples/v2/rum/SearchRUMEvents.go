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
			From:  common.PtrString("now-15m"),
			Query: common.PtrString("@type:session AND @session.type:user"),
			To:    common.PtrString("now"),
		},
		Options: &datadog.RUMQueryOptions{
			TimeOffset: common.PtrInt64(0),
			Timezone:   common.PtrString("GMT"),
		},
		Page: &datadog.RUMQueryPageOptions{
			Limit: common.PtrInt32(25),
		},
		Sort: datadog.RUMSORT_TIMESTAMP_ASCENDING.Ptr(),
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewRUMApi(apiClient)
	resp, r, err := api.SearchRUMEvents(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMApi.SearchRUMEvents`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMApi.SearchRUMEvents`:\n%s\n", responseContent)
}
