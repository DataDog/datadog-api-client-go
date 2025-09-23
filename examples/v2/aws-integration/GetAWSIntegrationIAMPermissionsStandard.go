// Get AWS integration standard IAM permissions returns "AWS integration standard IAM permissions." response

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAWSIntegrationApi(apiClient)
	resp, r, err := api.GetAWSIntegrationIAMPermissionsStandard(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.GetAWSIntegrationIAMPermissionsStandard`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSIntegrationApi.GetAWSIntegrationIAMPermissionsStandard`:\n%s\n", responseContent)
}
