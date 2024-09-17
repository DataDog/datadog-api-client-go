// Send logs returns "Request accepted for processing (always 202 empty JSON)." response

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
	body := []datadogV2.HTTPLogItem{
		{
			Ddsource: datadog.PtrString("nginx"),
			Ddtags:   datadog.PtrString("env:staging,version:5.1"),
			Hostname: datadog.PtrString("i-012345678"),
			Message:  "2019-11-19T14:37:58,995 INFO [process.name][20081] Hello World",
			Service:  datadog.PtrString("payment"),
			AdditionalProperties: map[string]string{
				"status": "info",
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsApi(apiClient)
	resp, r, err := api.SubmitLog(ctx, body, *datadogV2.NewSubmitLogOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.SubmitLog`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsApi.SubmitLog`:\n%s\n", responseContent)
}
