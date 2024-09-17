// List all AWS Logs integrations returns "OK" response

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewAWSLogsIntegrationApi(apiClient)
	resp, r, err := api.ListAWSLogsIntegrations(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSLogsIntegrationApi.ListAWSLogsIntegrations`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSLogsIntegrationApi.ListAWSLogsIntegrations`:\n%s\n", responseContent)
}
