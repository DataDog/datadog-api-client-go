// Search spans returns "OK" response with pagination

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
	body := datadogV2.SpansListRequest{
		Data: &datadogV2.SpansListRequestData{
			Attributes: &datadogV2.SpansListRequestAttributes{
				Filter: &datadogV2.SpansQueryFilter{
					From:  datadog.PtrString("now-15m"),
					Query: datadog.PtrString("service:python*"),
					To:    datadog.PtrString("now"),
				},
				Options: &datadogV2.SpansQueryOptions{
					Timezone: datadog.PtrString("GMT"),
				},
				Page: &datadogV2.SpansListRequestPage{
					Limit: datadog.PtrInt32(2),
				},
				Sort: datadogV2.SPANSSORT_TIMESTAMP_ASCENDING.Ptr(),
			},
			Type: datadogV2.SPANSLISTREQUESTTYPE_SEARCH_REQUEST.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSpansApi(apiClient)
	resp, _ := api.ListSpansWithPagination(ctx, body)

	for paginationResult := range resp {
		if paginationResult.Error != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `SpansApi.ListSpans`: %v\n", paginationResult.Error)
		}
		responseContent, _ := json.MarshalIndent(paginationResult.Item, "", "  ")
		fmt.Fprintf(os.Stdout, "%s\n", responseContent)
	}
}
