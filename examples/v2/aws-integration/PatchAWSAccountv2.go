// AWS Integration - Patch account returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.AWSAccountPatchRequest{
		Data: &datadogV2.AWSAccountPatch{
			Attributes: &datadogV2.AWSAccountPatchAttributes{
				AccountTags:  []string{},
				AuthConfig:   &datadogV2.AWSAuthConfig{},
				AwsAccountId: datadog.PtrString("123456789012"),
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
			Type: datadogV2.AWSACCOUNTTYPE_AWS_ACCOUNT.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.PatchAWSAccountv2", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAWSIntegrationApi(apiClient)
	r, err := api.PatchAWSAccountv2(ctx, "aws_account_config_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.PatchAWSAccountv2`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
