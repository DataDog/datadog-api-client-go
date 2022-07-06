// Delete an AWS integration returns "OK" response

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
	body := datadog.AWSAccountDeleteRequest{
		AccountId: datadog.PtrString("1234567"),
		RoleName:  datadog.PtrString("DatadogAWSIntegrationRole"),
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.AWSIntegrationApi(apiClient)
	resp, r, err := api.DeleteAWSAccount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.DeleteAWSAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSIntegrationApi.DeleteAWSAccount`:\n%s\n", responseContent)
}
