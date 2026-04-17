// Export security monitoring resources to Terraform returns "OK" response

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "suppression" in the system
	SuppressionDataID := os.Getenv("SUPPRESSION_DATA_ID")

	body := datadogV2.SecurityMonitoringTerraformBulkExportRequest{
		Data: datadogV2.SecurityMonitoringTerraformBulkExportData{
			Attributes: datadogV2.SecurityMonitoringTerraformBulkExportAttributes{
				ResourceIds: []string{
					SuppressionDataID,
				},
			},
			Type: "bulk_export_resources",
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.BulkExportSecurityMonitoringTerraformResources", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.BulkExportSecurityMonitoringTerraformResources(ctx, datadogV2.SECURITYMONITORINGTERRAFORMRESOURCETYPE_SUPPRESSIONS, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.BulkExportSecurityMonitoringTerraformResources`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := ioutil.ReadAll(resp)
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.BulkExportSecurityMonitoringTerraformResources`:\n%s\n", responseContent)
}
