// Send gzip logs returns "Response from server (always 200 empty JSON)." response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := []datadogV1.HTTPLogItem{
		{
			Message: "Example-Send_gzip_logs_returns_Response_from_server_always_200_empty_JSON_response",
			Ddtags:  datadog.PtrString("host:ExampleSendgziplogsreturnsResponsefromserveralways200emptyJSONresponse"),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewLogsApi(apiClient)
	resp, r, err := api.SubmitLog(ctx, body, *datadogV1.NewSubmitLogOptionalParameters().WithContentEncoding(datadogV1.CONTENTENCODING_GZIP))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.SubmitLog`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsApi.SubmitLog`:\n%s\n", responseContent)
}
