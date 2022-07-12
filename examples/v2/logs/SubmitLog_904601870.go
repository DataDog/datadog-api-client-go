// Send deflate logs returns "Request accepted for processing (always 202 empty JSON)." response

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
	body := []datadog.HTTPLogItem{
		{
			Ddsource: datadog.PtrString("nginx"),
			Ddtags:   datadog.PtrString("env:staging,version:5.1"),
			Hostname: datadog.PtrString("i-012345678"),
			Message:  "2019-11-19T14:37:58,995 INFO [process.name][20081] Hello World",
			Service:  datadog.PtrString("payment"),
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewLogsApi(apiClient)
	resp, r, err := api.SubmitLog(ctx, body, *datadog.NewSubmitLogOptionalParameters().WithContentEncoding(datadog.CONTENTENCODING_DEFLATE))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.SubmitLog`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsApi.SubmitLog`:\n%s\n", responseContent)
}
