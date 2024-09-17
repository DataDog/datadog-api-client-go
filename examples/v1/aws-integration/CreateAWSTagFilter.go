// Set an AWS tag filter returns "OK" response

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
	body := datadogV1.AWSTagFilterCreateRequest{
		AccountId:    datadog.PtrString("123456789012"),
		Namespace:    datadogV1.AWSNAMESPACE_ELB.Ptr(),
		TagFilterStr: datadog.PtrString("prod*"),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewAWSIntegrationApi(apiClient)
	resp, r, err := api.CreateAWSTagFilter(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.CreateAWSTagFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSIntegrationApi.CreateAWSTagFilter`:\n%s\n", responseContent)
}
