// AWS Integration - Patch account config returns "AWS Account object" response

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
	// there is a valid "aws_account_v2" in the system
	AwsAccountV2DataAttributesAwsAccountID := os.Getenv("AWS_ACCOUNT_V2_DATA_ATTRIBUTES_AWS_ACCOUNT_ID")

	body := datadogV2.AWSAccountPatchRequest{
		Data: &datadogV2.AWSAccountPatchRequestData{
			Attributes: datadogV2.AWSAccountPatchRequestAttributes{
				AccountTags: *datadog.NewNullableList(&[]string{}),
				AuthConfig: &datadogV2.AWSAuthConfig{
					AWSAuthConfigRole: &datadogV2.AWSAuthConfigRole{
						RoleName: "test",
					}},
				AwsAccountId: AwsAccountV2DataAttributesAwsAccountID,
				AwsPartition: datadogV2.AWSACCOUNTPARTITION_AWS.Ptr(),
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
			Id:   datadog.PtrString(AwsAccountV2DataAttributesAwsAccountID),
			Type: datadog.PtrString("account"),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateAWSAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAWSIntegrationApi(apiClient)
	resp, r, err := api.UpdateAWSAccount(ctx, AwsAccountV2DataAttributesAwsAccountID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.UpdateAWSAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSIntegrationApi.UpdateAWSAccount`:\n%s\n", responseContent)
}
