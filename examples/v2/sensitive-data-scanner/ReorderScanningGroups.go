// Reorder Groups returns "OK" response

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
	// there is a valid "scanning_group" in the system
	GroupDataID := os.Getenv("GROUP_DATA_ID")

	// a valid "configuration" in the system
	ConfigurationDataID := os.Getenv("CONFIGURATION_DATA_ID")

	body := datadogV2.SensitiveDataScannerConfigRequest{
		Data: datadogV2.SensitiveDataScannerReorderConfig{
			Relationships: &datadogV2.SensitiveDataScannerConfigurationRelationships{
				Groups: &datadogV2.SensitiveDataScannerGroupList{
					Data: []datadogV2.SensitiveDataScannerGroupItem{
						{
							Type: datadogV2.SENSITIVEDATASCANNERGROUPTYPE_SENSITIVE_DATA_SCANNER_GROUP.Ptr(),
							Id:   datadog.PtrString(GroupDataID),
						},
					},
				},
			},
			Type: datadogV2.SENSITIVEDATASCANNERCONFIGURATIONTYPE_SENSITIVE_DATA_SCANNER_CONFIGURATIONS.Ptr(),
			Id:   datadog.PtrString(ConfigurationDataID),
		},
		Meta: datadogV2.SensitiveDataScannerMetaVersionOnly{},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSensitiveDataScannerApi(apiClient)
	resp, r, err := api.ReorderScanningGroups(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SensitiveDataScannerApi.ReorderScanningGroups`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SensitiveDataScannerApi.ReorderScanningGroups`:\n%s\n", responseContent)
}
