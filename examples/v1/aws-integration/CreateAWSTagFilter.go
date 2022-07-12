// Set an AWS tag filter returns "OK" response

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
	body := datadog.AWSTagFilterCreateRequest{
		AccountId:    datadog.PtrString("1234567"),
		Namespace:    datadog.AWSNAMESPACE_ELB.Ptr(),
		TagFilterStr: datadog.PtrString("prod*"),
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewAWSIntegrationApi(apiClient)
	resp, r, err := api.CreateAWSTagFilter(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.CreateAWSTagFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSIntegrationApi.CreateAWSTagFilter`:\n%s\n", responseContent)
}
