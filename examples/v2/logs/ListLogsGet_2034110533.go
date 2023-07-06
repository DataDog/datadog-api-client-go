// Get a quick list of logs returns "OK" response

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsApi(apiClient)
	resp, r, err := api.ListLogsGet(ctx, *datadogV2.NewListLogsGetOptionalParameters().WithFilterQuery("datadog-agent").WithFilterIndexes([]string{
		"main",
	}).WithFilterFrom(time.Date(2020, 9, 17, 11, 48, 36, 0, time.UTC)).WithFilterTo(time.Date(2020, 9, 17, 12, 48, 36, 0, time.UTC)).WithPageLimit(5))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.ListLogsGet`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsApi.ListLogsGet`:\n%s\n", responseContent)
}
