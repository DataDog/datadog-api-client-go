// Convert security monitoring resource to Terraform returns "OK" response

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
	body := datadogV2.SecurityMonitoringTerraformConvertRequest{
		Data: datadogV2.SecurityMonitoringTerraformConvertData{
			Type: "convert_resource",
			Id:   "abc-123",
			Attributes: datadogV2.SecurityMonitoringTerraformConvertAttributes{
				ResourceJson: map[string]interface{}{
					"enabled":           true,
					"name":              "Example-Security-Monitoring",
					"rule_query":        "source:cloudtrail",
					"suppression_query": "env:test",
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ConvertSecurityMonitoringTerraformResource", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.ConvertSecurityMonitoringTerraformResource(ctx, datadogV2.SECURITYMONITORINGTERRAFORMRESOURCETYPE_SUPPRESSIONS, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.ConvertSecurityMonitoringTerraformResource`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.ConvertSecurityMonitoringTerraformResource`:\n%s\n", responseContent)
}
