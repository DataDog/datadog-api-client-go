// Delete Scanning Rule returns "OK" response

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
	// the "scanning_group" has a "scanning_rule"
	RuleDataID := os.Getenv("RULE_DATA_ID")

	body := datadogV2.SensitiveDataScannerRuleDeleteRequest{
		Meta: datadogV2.SensitiveDataScannerMetaVersionOnly{},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSensitiveDataScannerApi(apiClient)
	resp, r, err := api.DeleteScanningRule(ctx, RuleDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SensitiveDataScannerApi.DeleteScanningRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SensitiveDataScannerApi.DeleteScanningRule`:\n%s\n", responseContent)
}
