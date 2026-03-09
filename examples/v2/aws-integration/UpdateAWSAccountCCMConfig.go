// Update AWS CCM config returns "AWS CCM Config object" response

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
	body := datadogV2.AWSCcmConfigRequest{
		Data: datadogV2.AWSCcmConfigRequestData{
			Attributes: datadogV2.AWSCcmConfigRequestAttributes{
				CcmConfig: datadogV2.AWSCcmConfig{
					DataExportConfigs: []datadogV2.DataExportConfig{
						{
							BucketName:   "billing",
							BucketRegion: "us-east-1",
							ReportName:   "cost-and-usage-report",
							ReportPrefix: "reports",
							ReportType:   "CUR2.0",
						},
					},
				},
			},
			Type: datadogV2.AWSCCMCONFIGTYPE_CCM_CONFIG,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateAWSAccountCCMConfig", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAWSIntegrationApi(apiClient)
	resp, r, err := api.UpdateAWSAccountCCMConfig(ctx, "aws_account_config_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.UpdateAWSAccountCCMConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSIntegrationApi.UpdateAWSAccountCCMConfig`:\n%s\n", responseContent)
}
