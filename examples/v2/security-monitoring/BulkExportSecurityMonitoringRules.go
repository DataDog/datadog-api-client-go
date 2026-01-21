// Bulk export security monitoring rules returns "OK" response

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
	// there is a valid "security_rule" in the system
	SecurityRuleID := os.Getenv("SECURITY_RULE_ID")

	body := datadogV2.SecurityMonitoringRuleBulkExportPayload{
		Data: datadogV2.SecurityMonitoringRuleBulkExportData{
			Attributes: datadogV2.SecurityMonitoringRuleBulkExportAttributes{
				RuleIds: []string{
					SecurityRuleID,
				},
			},
			Type: datadogV2.SECURITYMONITORINGRULEBULKEXPORTDATATYPE_SECURITY_MONITORING_RULES_BULK_EXPORT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.BulkExportSecurityMonitoringRules(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.BulkExportSecurityMonitoringRules`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := ioutil.ReadAll(resp)
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.BulkExportSecurityMonitoringRules`:\n%s\n", responseContent)
}
