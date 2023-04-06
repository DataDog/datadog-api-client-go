// Create Scanning Group returns "OK" response

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
	// a valid "configuration" in the system
	ConfigurationDataID := os.Getenv("CONFIGURATION_DATA_ID")

	body := datadogV2.SensitiveDataScannerGroupCreateRequest{
		Meta: &datadogV2.SensitiveDataScannerMetaVersionOnly{},
		Data: &datadogV2.SensitiveDataScannerGroupCreate{
			Type: datadogV2.SENSITIVEDATASCANNERGROUPTYPE_SENSITIVE_DATA_SCANNER_GROUP,
			Attributes: datadogV2.SensitiveDataScannerGroupAttributes{
				Name:      datadog.PtrString("Example-Sensitive-Data-Scanner"),
				IsEnabled: datadog.PtrBool(false),
				ProductList: []datadogV2.SensitiveDataScannerProduct{
					datadogV2.SENSITIVEDATASCANNERPRODUCT_LOGS,
				},
				Filter: &datadogV2.SensitiveDataScannerFilter{
					Query: datadog.PtrString("*"),
				},
			},
			Relationships: &datadogV2.SensitiveDataScannerGroupRelationships{
				Configuration: &datadogV2.SensitiveDataScannerConfigurationData{
					Data: &datadogV2.SensitiveDataScannerConfiguration{
						Type: datadogV2.SENSITIVEDATASCANNERCONFIGURATIONTYPE_SENSITIVE_DATA_SCANNER_CONFIGURATIONS.Ptr(),
						Id:   datadog.PtrString(ConfigurationDataID),
					},
				},
				Rules: &datadogV2.SensitiveDataScannerRuleData{
					Data: []datadogV2.SensitiveDataScannerRule{},
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSensitiveDataScannerApi(apiClient)
	resp, r, err := api.CreateScanningGroup(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SensitiveDataScannerApi.CreateScanningGroup`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SensitiveDataScannerApi.CreateScanningGroup`:\n%s\n", responseContent)
}
