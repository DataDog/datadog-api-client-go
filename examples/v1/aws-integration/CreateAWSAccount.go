// Create an AWS integration returns "OK" response

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
	body := datadog.AWSAccount{
		AccountId: datadog.PtrString("1234567"),
		AccountSpecificNamespaceRules: map[string]bool{
			"auto_scaling": false,
			"opswork":      false,
		},
		CspmResourceCollectionEnabled: datadog.PtrBool(true),
		ExcludedRegions: []string{
			"us-east-1",
			"us-west-2",
		},
		FilterTags: []string{
			"$KEY:$VALUE",
		},
		HostTags: []string{
			"$KEY:$VALUE",
		},
		MetricsCollectionEnabled:  datadog.PtrBool(false),
		ResourceCollectionEnabled: datadog.PtrBool(true),
		RoleName:                  datadog.PtrString("DatadogAWSIntegrationRole"),
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.AWSIntegrationApi(apiClient)
	resp, r, err := api.CreateAWSAccount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.CreateAWSAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSIntegrationApi.CreateAWSAccount`:\n%s\n", responseContent)
}
