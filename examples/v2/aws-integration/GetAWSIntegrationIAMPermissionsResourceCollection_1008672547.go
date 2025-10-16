// Get resource collection IAM permissions returns "AWS IAM Permissions object" response

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
	api := datadogV2.NewAWSIntegrationApi(apiClient)
	resp, r, err := api.GetAWSIntegrationIAMPermissionsResourceCollection(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.GetAWSIntegrationIAMPermissionsResourceCollection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSIntegrationApi.GetAWSIntegrationIAMPermissionsResourceCollection`:\n%s\n", responseContent)
}