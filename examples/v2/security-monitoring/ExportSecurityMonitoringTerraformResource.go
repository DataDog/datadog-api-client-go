// Export security monitoring resource to Terraform returns "OK" response

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
	// there is a valid "suppression" in the system
	SuppressionDataID := os.Getenv("SUPPRESSION_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ExportSecurityMonitoringTerraformResource", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.ExportSecurityMonitoringTerraformResource(ctx, datadogV2.SECURITYMONITORINGTERRAFORMRESOURCETYPE_SUPPRESSIONS, SuppressionDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.ExportSecurityMonitoringTerraformResource`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.ExportSecurityMonitoringTerraformResource`:\n%s\n", responseContent)
}
