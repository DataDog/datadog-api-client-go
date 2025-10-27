// Ruleset get multiple returns "OK" response

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
	body := datadogV2.GetMultipleRulesetsRequest{
		Data: &datadogV2.GetMultipleRulesetsRequestData{
			Attributes: &datadogV2.GetMultipleRulesetsRequestDataAttributes{
				Rulesets: []string{},
			},
			Type: datadogV2.GETMULTIPLERULESETSREQUESTDATATYPE_GET_MULTIPLE_RULESETS_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ListMultipleRulesets", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.ListMultipleRulesets(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.ListMultipleRulesets`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.ListMultipleRulesets`:\n%s\n", responseContent)
}
