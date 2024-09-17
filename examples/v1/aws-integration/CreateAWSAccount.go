// Create an AWS integration returns "OK" response

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
	body := datadogV1.AWSAccount{
		AccountId: datadog.PtrString("163662907100"),
		AccountSpecificNamespaceRules: map[string]bool{
			"auto_scaling": false,
		},
		CspmResourceCollectionEnabled: datadog.PtrBool(true),
		ExcludedRegions: []string{
			"us-east-1",
			"us-west-2",
		},
		ExtendedResourceCollectionEnabled: datadog.PtrBool(true),
		FilterTags: []string{
			"$KEY:$VALUE",
		},
		HostTags: []string{
			"$KEY:$VALUE",
		},
		MetricsCollectionEnabled: datadog.PtrBool(false),
		RoleName:                 datadog.PtrString("DatadogAWSIntegrationRole"),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewAWSIntegrationApi(apiClient)
	resp, r, err := api.CreateAWSAccount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.CreateAWSAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSIntegrationApi.CreateAWSAccount`:\n%s\n", responseContent)
}
