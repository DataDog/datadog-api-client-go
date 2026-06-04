// Validate AWS CCM config returns "AWS CCM Config validation result" response

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
	body := datadogV2.AWSCcmConfigValidationRequest{
		Data: datadogV2.AWSCcmConfigValidationRequestData{
			Attributes: datadogV2.AWSCcmConfigValidationRequestAttributes{
				AccountId:    "123456789012",
				BucketName:   "billing",
				BucketRegion: "us-east-1",
				ReportName:   "cost-and-usage-report",
				ReportPrefix: datadog.PtrString("reports"),
			},
			Type: datadogV2.AWSCCMCONFIGVALIDATIONTYPE_CCM_CONFIG_VALIDATION,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ValidateAWSCCMConfig", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAWSIntegrationApi(apiClient)
	resp, r, err := api.ValidateAWSCCMConfig(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.ValidateAWSCCMConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSIntegrationApi.ValidateAWSCCMConfig`:\n%s\n", responseContent)
}
