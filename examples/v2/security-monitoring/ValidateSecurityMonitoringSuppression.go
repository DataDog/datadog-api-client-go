// Validate a suppression rule returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.SecurityMonitoringSuppressionCreateRequest{
		Data: datadogV2.SecurityMonitoringSuppressionCreateData{
			Attributes: datadogV2.SecurityMonitoringSuppressionCreateAttributes{
				DataExclusionQuery: datadog.PtrString("source:cloudtrail account_id:12345"),
				Description:        datadog.PtrString("This rule suppresses low-severity signals in staging environments."),
				Enabled:            true,
				Name:               "Custom suppression",
				RuleQuery:          "type:log_detection source:cloudtrail",
			},
			Type: datadogV2.SECURITYMONITORINGSUPPRESSIONTYPE_SUPPRESSIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	r, err := api.ValidateSecurityMonitoringSuppression(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.ValidateSecurityMonitoringSuppression`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
