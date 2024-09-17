// Delete an AWS Logs integration returns "OK" response

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
	body := datadogV1.AWSAccountAndLambdaRequest{
		AccountId: "1234567",
		LambdaArn: "arn:aws:lambda:us-east-1:1234567:function:LogsCollectionAPITest",
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewAWSLogsIntegrationApi(apiClient)
	resp, r, err := api.DeleteAWSLambdaARN(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSLogsIntegrationApi.DeleteAWSLambdaARN`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSLogsIntegrationApi.DeleteAWSLambdaARN`:\n%s\n", responseContent)
}
