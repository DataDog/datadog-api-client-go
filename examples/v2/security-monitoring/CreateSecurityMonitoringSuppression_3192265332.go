// Create a suppression rule with an exclusion query returns "OK" response

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
	body := datadogV2.SecurityMonitoringSuppressionCreateRequest{
		Data: datadogV2.SecurityMonitoringSuppressionCreateData{
			Attributes: datadogV2.SecurityMonitoringSuppressionCreateAttributes{
				Description:        datadog.PtrString("This rule suppresses low-severity signals in staging environments."),
				Enabled:            true,
				ExpirationDate:     datadog.PtrInt64(1638443471000),
				Name:               "Example-Security-Monitoring",
				RuleQuery:          "type:log_detection source:cloudtrail",
				DataExclusionQuery: datadog.PtrString("account_id:12345"),
			},
			Type: datadogV2.SECURITYMONITORINGSUPPRESSIONTYPE_SUPPRESSIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateSecurityMonitoringSuppression(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateSecurityMonitoringSuppression`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateSecurityMonitoringSuppression`:\n%s\n", responseContent)
}
