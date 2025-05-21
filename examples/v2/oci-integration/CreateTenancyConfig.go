// Create tenancy config returns "Created" response

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
	body := datadogV2.CreateTenancyConfig{
		Data: &datadogV2.CreateTenancyConfigData{
			Attributes: &datadogV2.CreateTenancyConfigDataAttributes{
				AuthCredentials: datadogV2.AuthCredentials{
					Fingerprint: "a7:b5:54:f2:da:a2:d7:b0:ed:f4:79:47:93:64:12:b1",
					PrivateKey: `-----BEGIN PRIVATE KEY-----
o9kEwoumo8yHVn5Ztp4F2cxaD6+MzSJ/I6WesPyePUD7sPeorXByg1UNOXyzqDub
/aU4/sNo2f8epM9l7QGiCtY=
-----END PRIVATE KEY-----`,
				},
				ConfigVersion: datadog.PtrInt64(2),
				HomeRegion:    "us-ashburn-1",
				LogsConfig: &datadogV2.OCILogsConfig{
					CompartmentTagFilters: []string{
						"datadog:true",
						"env:prod",
					},
					Enabled: datadog.PtrBool(true),
					EnabledServices: []string{
						"oacnativeproduction",
					},
				},
				MetricsConfig: &datadogV2.OCIMetricsConfig{
					CompartmentTagFilters: []string{
						"datadog:true",
						"env:prod",
					},
					Enabled: datadog.PtrBool(true),
					ExcludedServices: []string{
						"oci_compute",
					},
				},
				ResourceCollectionEnabled: datadog.PtrBool(true),
				UserOcid:                  "ocid1.user.test",
			},
			Id:   "ocid1.tenancy.dummy_value",
			Type: datadogV2.CREATETENANCYCONFIGDATATYPE_OCI_TENANCY,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOCIIntegrationApi(apiClient)
	resp, r, err := api.CreateTenancyConfig(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OCIIntegrationApi.CreateTenancyConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OCIIntegrationApi.CreateTenancyConfig`:\n%s\n", responseContent)
}
