// Get a list of logs returns "OK" response

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
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewLogsApi(apiClient)
	resp, r, err := api.ListLogsGet(ctx, *datadog.NewListLogsGetOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.ListLogsGet`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsApi.ListLogsGet`:\n%s\n", responseContent)
}
