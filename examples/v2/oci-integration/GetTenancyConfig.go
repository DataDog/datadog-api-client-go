// Get tenancy config returns "OK" response

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

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOCIIntegrationApi(apiClient)
	resp, r, err := api.GetTenancyConfig(ctx, OciTenancyDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OCIIntegrationApi.GetTenancyConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OCIIntegrationApi.GetTenancyConfig`:\n%s\n", responseContent)
}
