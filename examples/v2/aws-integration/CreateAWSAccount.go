// AWS Integration - Create account config returns "AWS Account object" response

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
		Data: &datadogV2.AWSAccountCreateRequestData{
			Attributes: datadogV2.AWSAccountCreateRequestAttributes{
				AccountTags: *datadog.NewNullableList(&[]string{}),
				AuthConfig: datadogV2.AWSAuthConfig{
					AWSAuthConfigRole: &datadogV2.AWSAuthConfigRole{
						RoleName: "test",
					}},
				AwsAccountId: "123456789012",
				AwsPartition: datadogV2.AWSACCOUNTPARTITION_AWS,
				AwsRegions: &datadogV2.AWSRegions{
					AWSRegionsIncludeOnly: &datadogV2.AWSRegionsIncludeOnly{
						IncludeOnly: []string{
							"us-east-1",
						},
					}},
				LogsConfig: &datadogV2.AWSLogsConfig{
					LambdaForwarder: &datadogV2.AWSLambdaForwarderConfig{
						Lambdas: []string{},
						Sources: []string{
							"s3",
						},
					},
				},
				MetricsConfig: &datadogV2.AWSMetricsConfig{
					NamespaceFilters: &datadogV2.AWSNamespaceFilters{
						AWSNamespaceFiltersIncludeOnly: &datadogV2.AWSNamespaceFiltersIncludeOnly{
							IncludeOnly: []string{
								"AWS/EC2",
							},
						}},
					TagFilters: []datadogV2.AWSNamespaceTagFilter{
						{
							Namespace: datadog.PtrString("AWS/EC2"),
							Tags:      *datadog.NewNullableList(&[]string{}),
						},
					},
				},
				ResourcesConfig: &datadogV2.AWSResourcesConfig{},
				TracesConfig: &datadogV2.AWSTracesConfig{
					XrayServices: &datadogV2.XRayServicesList{
						XRayServicesIncludeOnly: &datadogV2.XRayServicesIncludeOnly{
							IncludeOnly: []string{
								"AWS/AppSync",
							},
						}},
				},
			},
			Id:   datadog.PtrString("123456789012"),
			Type: datadog.PtrString("account"),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateAWSAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAWSIntegrationApi(apiClient)
	resp, r, err := api.CreateAWSAccount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.CreateAWSAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSIntegrationApi.CreateAWSAccount`:\n%s\n", responseContent)
}
