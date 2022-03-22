// Enable an AWS Logs integration returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.AWSLogsServicesRequest{
		AccountId: "1234567",
		Services: []string{
			"s3",
			"elb",
			"elbv2",
			"cloudfront",
			"redshift",
			"lambda",
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.AWSLogsIntegrationApi.EnableAWSLogServices(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSLogsIntegrationApi.EnableAWSLogServices`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSLogsIntegrationApi.EnableAWSLogServices`:\n%s\n", responseContent)
}
