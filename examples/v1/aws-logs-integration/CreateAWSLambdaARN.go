// Add AWS Log Lambda ARN returns "OK" response

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
	body := datadog.AWSAccountAndLambdaRequest{
		AccountId: "1234567",
		LambdaArn: "arn:aws:lambda:us-east-1:1234567:function:LogsCollectionAPITest",
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.AWSLogsIntegrationApi(apiClient)
	resp, r, err := api.CreateAWSLambdaARN(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSLogsIntegrationApi.CreateAWSLambdaARN`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSLogsIntegrationApi.CreateAWSLambdaARN`:\n%s\n", responseContent)
}
