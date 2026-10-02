// Send CI job logs returns "Request accepted for processing" response

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
	body := []datadogV2.CILogItem{
		{
			Ddtags:           datadog.PtrString("runner:linux,architecture:amd64"),
			JobId:            "job-456",
			LineNumber:       datadog.PtrInt64(812),
			Message:          "Running go test ./...",
			PipelineUniqueId: "3eacb6f3-ff04-4e10-8a9c-46e6d054024a",
			ProviderName:     datadog.PtrString("example-provider"),
			SectionName:      datadog.PtrString("tests"),
			Status:           datadog.PtrString("warn"),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCIVisibilityLogsApi(apiClient)
	resp, r, err := api.SubmitCILog(ctx, body, *datadogV2.NewSubmitCILogOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CIVisibilityLogsApi.SubmitCILog`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CIVisibilityLogsApi.SubmitCILog`:\n%s\n", responseContent)
}
