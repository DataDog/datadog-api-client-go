// Update tenancy config returns "OK" response

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
	// there is a valid "oci_tenancy" resource in the system
	OciTenancyDataID := os.Getenv("OCI_TENANCY_DATA_ID")

	body := datadogV2.UpdateTenancyConfig{
		Data: &datadogV2.UpdateTenancyConfigData{
			Attributes: &datadogV2.UpdateTenancyConfigDataAttributes{
				HomeRegion: datadog.PtrString("us-sanjose-1"),
				MetricsConfig: &datadogV2.OCIMetricsConfig{
					CompartmentTagFilters: []string{
						"datadog:true",
						"env:prod",
					},
					Enabled:          datadog.PtrBool(false),
					ExcludedServices: []string{},
				},
				ResourceCollectionEnabled: datadog.PtrBool(false),
				UserOcid:                  datadog.PtrString("ocid1.user.test_updated"),
			},
			Id:   OciTenancyDataID,
			Type: datadogV2.UPDATETENANCYCONFIGDATATYPE_OCI_TENANCY,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOCIIntegrationApi(apiClient)
	resp, r, err := api.UpdateTenancyConfig(ctx, OciTenancyDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OCIIntegrationApi.UpdateTenancyConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OCIIntegrationApi.UpdateTenancyConfig`:\n%s\n", responseContent)
}
