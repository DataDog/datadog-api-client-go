// Create Cloud Cost Management AWS CUR config returns "OK" response

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
	body := datadogV2.AwsCURConfigPostRequest{
		Data: datadogV2.AwsCURConfigPostData{
			Attributes: datadogV2.AwsCURConfigPostRequestAttributes{
				AccountId:    "123456789123",
				BucketName:   "dd-cost-bucket",
				BucketRegion: datadog.PtrString("us-east-1"),
				ReportName:   "dd-report-name",
				ReportPrefix: "dd-report-prefix",
			},
			Type: datadogV2.AWSCURCONFIGPOSTREQUESTTYPE_AWS_CUR_CONFIG_POST_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	resp, r, err := api.CreateCostAWSCURConfig(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.CreateCostAWSCURConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudCostManagementApi.CreateCostAWSCURConfig`:\n%s\n", responseContent)
}
