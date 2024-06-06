// AWS Integration - Create account returns "AWS Account object" response

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
	body := datadogV2.AWSAccountCreateRequest{
		Data: datadogV2.AWSAccountCreate{
			Attributes: datadogV2.AWSAccountCreateAttributes{
				AccountTags:  []string{},
				AuthConfig:   &datadogV2.AWSAuthConfig{},
				AwsAccountId: "123456789012",
				AwsRegions: &datadogV2.AWSRegionsList{
					IncludeOnly: []string{
						"us-east-1",
					},
				},
				LogsConfig: &datadogV2.AWSLogs{
					LambdaForwarder: &datadogV2.AWSLambdaForwarder{
						Lambdas: []string{},
						Sources: []string{
							"s3",
						},
					},
				},
				MetricsConfig: &datadogV2.AWSMetrics{
					NamespaceFilters: &datadogV2.AWSNamespacesList{
						ExcludeOnly: []string{
							"AWS/EC2",
						},
						IncludeOnly: []string{
							"AWS/EC2",
						},
					},
					TagFilters: []datadogV2.AWSNamespaceTagFilter{
						{
							Namespace: datadog.PtrString("AWS/EC2"),
							Tags:      []string{},
						},
					},
				},
				ResourcesConfig: &datadogV2.AWSResources{},
				TracesConfig: &datadogV2.AWSTraces{
					XrayServices: &datadogV2.AWSXRayServicesList{
						IncludeOnly: []string{
							"AWS/AppSync",
						},
					},
				},
			},
			Type: datadogV2.AWSACCOUNTTYPE_AWS_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateAWSAccountv2", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAWSIntegrationApi(apiClient)
	resp, r, err := api.CreateAWSAccountv2(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.CreateAWSAccountv2`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSIntegrationApi.CreateAWSAccountv2`:\n%s\n", responseContent)
}
