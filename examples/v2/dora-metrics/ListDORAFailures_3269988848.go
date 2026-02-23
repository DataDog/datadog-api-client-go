// Get a list of failure events returns "OK" response

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
				From:  datadog.PtrTime(time.Date(2025, 3, 23, 0, 0, 0, 0, time.UTC)),
				Limit: datadog.PtrInt32(1),
				To:    datadog.PtrTime(time.Date(2025, 3, 24, 0, 0, 0, 0, time.UTC)),
			},
			Type: datadogV2.DORALISTFAILURESREQUESTDATATYPE_DORA_FAILURES_LIST_REQUEST.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
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
