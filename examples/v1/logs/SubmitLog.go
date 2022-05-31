// Send logs returns "Response from server (always 200 empty JSON)." response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := []datadog.HTTPLogItem{
		{
			Message: "Example-Send_logs_returns_Response_from_server_always_200_empty_JSON_response",
			Ddtags:  datadog.PtrString("host:ExampleSendlogsreturnsResponsefromserveralways200emptyJSONresponse"),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsApi.SubmitLog(ctx, body, *datadog.NewSubmitLogOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.SubmitLog`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsApi.SubmitLog`:\n%s\n", responseContent)
}
