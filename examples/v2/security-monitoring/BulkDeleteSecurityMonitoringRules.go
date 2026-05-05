// Bulk delete security monitoring rules returns "OK" response

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
	body := datadogV2.SecurityMonitoringRuleBulkDeletePayload{
		Data: datadogV2.SecurityMonitoringRuleBulkDeleteData{
			Attributes: datadogV2.SecurityMonitoringRuleBulkDeleteAttributes{
				RuleIds: []string{
					"abc-000-u7q",
					"abc-000-7dd",
				},
			},
			Type: datadogV2.SECURITYMONITORINGRULEBULKDELETEREQUESTDATATYPE_BULK_DELETE_RULES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.BulkDeleteSecurityMonitoringRules(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.BulkDeleteSecurityMonitoringRules`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.BulkDeleteSecurityMonitoringRules`:\n%s\n", responseContent)
}
