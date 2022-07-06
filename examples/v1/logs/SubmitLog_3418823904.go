// Send deflate logs returns "Response from server (always 200 empty JSON)." response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := []datadog.HTTPLogItem{
		{
			Message: "Example-Send_deflate_logs_returns_Response_from_server_always_200_empty_JSON_response",
			Ddtags:  datadog.PtrString("host:ExampleSenddeflatelogsreturnsResponsefromserveralways200emptyJSONresponse"),
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.LogsApi(apiClient)
	resp, r, err := api.SubmitLog(ctx, body, *datadog.NewSubmitLogOptionalParameters().WithContentEncoding(datadog.CONTENTENCODING_DEFLATE))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.SubmitLog`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsApi.SubmitLog`:\n%s\n", responseContent)
}
