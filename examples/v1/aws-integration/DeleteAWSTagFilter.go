// Delete a tag filtering entry returns "OK" response

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
	body := datadog.AWSTagFilterDeleteRequest{
		AccountId: datadog.PtrString("FAKEAC0FAKEAC2FAKEAC"),
		Namespace: datadog.AWSNAMESPACE_ELB.Ptr(),
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewAWSIntegrationApi(apiClient)
	resp, r, err := api.DeleteAWSTagFilter(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.DeleteAWSTagFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSIntegrationApi.DeleteAWSTagFilter`:\n%s\n", responseContent)
}
