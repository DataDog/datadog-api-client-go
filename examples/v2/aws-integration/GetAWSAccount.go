// Get an AWS integration by config ID returns "AWS Account object" response

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
	// there is a valid "aws_account_v2" in the system
	AwsAccountV2DataID := os.Getenv("AWS_ACCOUNT_V2_DATA_ID")


	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAWSIntegrationApi(apiClient)
	resp, r, err := api.GetAWSAccount(ctx, AwsAccountV2DataID, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.GetAWSAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSIntegrationApi.GetAWSAccount`:\n%s\n", responseContent)
}