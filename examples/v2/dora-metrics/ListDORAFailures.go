// Get a list of incident events returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.DORAListFailuresRequest{
		Data: datadogV2.DORAListFailuresRequestData{
			Attributes: datadogV2.DORAListFailuresRequestAttributes{
				From:  datadog.PtrTime(time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)),
				Limit: datadog.PtrInt32(100),
				Query: datadog.PtrString("severity:(SEV-1 OR SEV-2) env:production team:backend"),
				Sort:  datadog.PtrString("-started_at"),
				To:    datadog.PtrTime(time.Date(2025, 1, 31, 23, 59, 59, 0, time.UTC)),
			},
			Type: datadogV2.DORALISTFAILURESREQUESTDATATYPE_DORA_FAILURES_LIST_REQUEST.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDORAMetricsApi(apiClient)
	resp, r, err := api.ListDORAFailures(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DORAMetricsApi.ListDORAFailures`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DORAMetricsApi.ListDORAFailures`:\n%s\n", responseContent)
}
