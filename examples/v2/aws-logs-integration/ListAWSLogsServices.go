// Get list of AWS log ready services returns "AWS Logs Services List object" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAWSLogsIntegrationApi(apiClient)
	resp, r, err := api.ListAWSLogsServices(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSLogsIntegrationApi.ListAWSLogsServices`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSLogsIntegrationApi.ListAWSLogsServices`:\n%s\n", responseContent)
}