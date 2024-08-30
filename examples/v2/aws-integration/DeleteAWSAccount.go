// AWS Integration - Delete account config returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "aws_account_v2" in the system
	AwsAccountV2DataAttributesAwsAccountID := os.Getenv("AWS_ACCOUNT_V2_DATA_ATTRIBUTES_AWS_ACCOUNT_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteAWSAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAWSIntegrationApi(apiClient)
	r, err := api.DeleteAWSAccount(ctx, AwsAccountV2DataAttributesAwsAccountID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.DeleteAWSAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
