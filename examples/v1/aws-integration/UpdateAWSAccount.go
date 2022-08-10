// Update an AWS integration returns "OK" response

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
		AccountId: datadog.PtrString("123456789012"),
		AccountSpecificNamespaceRules: map[string]bool{
			"auto_scaling": false,
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
		RoleName:                  datadog.PtrString("datadog-role"),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewAWSIntegrationApi(apiClient)
	resp, r, err := api.UpdateAWSAccount(ctx, body, *datadogV1.NewUpdateAWSAccountOptionalParameters().WithAccountId("123456789012").WithRoleName("datadog-role"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.UpdateAWSAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSIntegrationApi.UpdateAWSAccount`:\n%s\n", responseContent)
}
