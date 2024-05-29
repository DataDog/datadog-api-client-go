// Test an existing rule returns "OK" response

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
	body := datadogV2.SecurityMonitoringRuleTestRequest{
		RuleQueryPayloads: []datadogV2.SecurityMonitoringRuleQueryPayload{
			{
				ExpectedResult: datadog.PtrBool(true),
				Index:          datadog.PtrInt64(0),
				Payload: &datadogV2.SecurityMonitoringRuleQueryPayloadData{
					Ddsource: datadog.PtrString("nginx"),
					Ddtags:   datadog.PtrString("env:staging,version:5.1"),
					Hostname: datadog.PtrString("i-012345678"),
					Message:  datadog.PtrString("2019-11-19T14:37:58,995 INFO [process.name][20081] Hello World"),
					Service:  datadog.PtrString("payment"),
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.TestExistingSecurityMonitoringRule(ctx, "rule_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.TestExistingSecurityMonitoringRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.TestExistingSecurityMonitoringRule`:\n%s\n", responseContent)
}
