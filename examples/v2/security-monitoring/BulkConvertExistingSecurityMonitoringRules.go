// Bulk convert rules to Terraform returns "OK" response

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
	body := datadogV2.SecurityMonitoringRuleConvertBulkPayload{
		Data: datadogV2.SecurityMonitoringRuleConvertBulkData{
			Attributes: datadogV2.SecurityMonitoringRuleConvertBulkAttributes{
				RuleIds: []string{
					"def-000-u7q",
					"def-000-7dd",
				},
			},
			Id:   datadog.PtrString("convert_bulk"),
			Type: datadogV2.SECURITYMONITORINGRULECONVERTBULKDATATYPE_SECURITY_MONITORING_RULES_CONVERT_BULK,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.BulkConvertExistingSecurityMonitoringRules(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.BulkConvertExistingSecurityMonitoringRules`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := ioutil.ReadAll(resp)
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.BulkConvertExistingSecurityMonitoringRules`:\n%s\n", responseContent)
}
