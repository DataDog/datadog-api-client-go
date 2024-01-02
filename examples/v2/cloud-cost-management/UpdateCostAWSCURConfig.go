// Update Cloud Cost Management AWS CUR config returns "OK" response

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
	body := datadogV2.AwsCURConfigPatchRequest{
		Data: datadogV2.AwsCURConfigPatchData{
			Attributes: datadogV2.AwsCURConfigPatchRequestAttributes{
				IsEnabled: true,
			},
			Type: datadogV2.AWSCURCONFIGPATCHREQUESTTYPE_AWS_CUR_CONFIG_PATCH_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	resp, r, err := api.UpdateCostAWSCURConfig(ctx, "100", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.UpdateCostAWSCURConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudCostManagementApi.UpdateCostAWSCURConfig`:\n%s\n", responseContent)
}
