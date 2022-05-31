// Send gzip logs returns "Request accepted for processing (always 202 empty JSON)." response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	body := []datadog.HTTPLogItem{
		{
			Ddsource: datadog.PtrString("nginx"),
			Ddtags:   datadog.PtrString("env:staging,version:5.1"),
			Hostname: datadog.PtrString("i-012345678"),
			Message:  "2019-11-19T14:37:58,995 INFO [process.name][20081] Hello World",
			Service:  datadog.PtrString("payment"),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsApi.SubmitLog(ctx, body, *datadog.NewSubmitLogOptionalParameters().WithContentEncoding(datadog.CONTENTENCODING_GZIP))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.SubmitLog`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsApi.SubmitLog`:\n%s\n", responseContent)
}
