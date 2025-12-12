// Update an AWS integration returns "AWS Account object" response

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
	AwsAccountV2DataID := os.Getenv("AWS_ACCOUNT_V2_DATA_ID")

	body := datadogV2.AWSAccountUpdateRequest{
		Data: datadogV2.AWSAccountUpdateRequestData{
			Attributes: datadogV2.AWSAccountUpdateRequestAttributes{
				AccountTags: *datadog.NewNullableList(&[]string{
					"key:value",
				}),
				AuthConfig: &datadogV2.AWSAuthConfig{
					AWSAuthConfigRole: &datadogV2.AWSAuthConfigRole{
						RoleName: "DatadogIntegrationRole",
					}},
				AwsAccountId: "123456789012",
				AwsPartition: datadogV2.AWSACCOUNTPARTITION_AWS.Ptr(),
				CcmConfig: &datadogV2.AWSCCMConfig{
					DataExportConfigs: []datadogV2.DataExportConfig{
						{
							BucketName:   datadog.PtrString("updated-bucket"),
							BucketRegion: datadog.PtrString("us-west-2"),
							ReportName:   datadog.PtrString("updated-report"),
							ReportPrefix: datadog.PtrString("cost-reports"),
							ReportType:   datadog.PtrString("CUR2.0"),
						},
					},
				},
				LogsConfig: &datadogV2.AWSLogsConfig{
					LambdaForwarder: &datadogV2.AWSLambdaForwarderConfig{
						Lambdas: []string{
							"arn:aws:lambda:us-east-1:123456789012:function:DatadogLambdaLogForwarder",
						},
						LogSourceConfig: &datadogV2.AWSLambdaForwarderConfigLogSourceConfig{
							TagFilters: []datadogV2.AWSLogSourceTagFilter{
								{
									Source: datadog.PtrString("s3"),
									Tags: *datadog.NewNullableList(&[]string{
										"test:test",
									}),
								},
							},
						},
						Sources: []string{
							"s3",
						},
					},
				},
				MetricsConfig: &datadogV2.AWSMetricsConfig{
					AutomuteEnabled:         datadog.PtrBool(true),
					CollectCloudwatchAlarms: datadog.PtrBool(true),
					CollectCustomMetrics:    datadog.PtrBool(true),
					Enabled:                 datadog.PtrBool(true),
					TagFilters: []datadogV2.AWSNamespaceTagFilter{
						{
							Namespace: datadog.PtrString("AWS/EC2"),
							Tags: *datadog.NewNullableList(&[]string{
								"key:value",
							}),
						},
					},
				},
				ResourcesConfig: &datadogV2.AWSResourcesConfig{
					CloudSecurityPostureManagementCollection: datadog.PtrBool(false),
					ExtendedCollection:                       datadog.PtrBool(false),
				},
				TracesConfig: &datadogV2.AWSTracesConfig{},
			},
			Type: datadogV2.AWSACCOUNTTYPE_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAWSIntegrationApi(apiClient)
	resp, r, err := api.UpdateAWSAccount(ctx, AwsAccountV2DataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.UpdateAWSAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSIntegrationApi.UpdateAWSAccount`:\n%s\n", responseContent)
}
