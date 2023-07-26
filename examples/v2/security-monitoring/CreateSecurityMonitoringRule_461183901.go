// Create a detection rule with type 'impossible_travel' returns "OK" response

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
	body := datadogV2.SecurityMonitoringRuleCreatePayload{
		SecurityMonitoringStandardRuleCreatePayload: &datadogV2.SecurityMonitoringStandardRuleCreatePayload{
			Queries: []datadogV2.SecurityMonitoringStandardRuleQuery{
				{
					Aggregation: datadogV2.SECURITYMONITORINGRULEQUERYAGGREGATION_GEO_DATA.Ptr(),
					GroupByFields: []string{
						"@usr.id",
					},
					DistinctFields: []string{},
					Metric:         datadog.PtrString("@network.client.geoip"),
					Query:          datadog.PtrString("*"),
				},
			},
			Cases: []datadogV2.SecurityMonitoringRuleCaseCreate{
				{
					Name:          datadog.PtrString(""),
					Status:        datadogV2.SECURITYMONITORINGRULESEVERITY_INFO,
					Notifications: []string{},
				},
			},
			HasExtendedTitle: datadog.PtrBool(true),
			Message:          "test",
			IsEnabled:        true,
			Options: datadogV2.SecurityMonitoringRuleOptions{
				MaxSignalDuration: datadogV2.SECURITYMONITORINGRULEMAXSIGNALDURATION_ONE_DAY.Ptr(),
				EvaluationWindow:  datadogV2.SECURITYMONITORINGRULEEVALUATIONWINDOW_FIFTEEN_MINUTES.Ptr(),
				KeepAlive:         datadogV2.SECURITYMONITORINGRULEKEEPALIVE_ONE_HOUR.Ptr(),
				DetectionMethod:   datadogV2.SECURITYMONITORINGRULEDETECTIONMETHOD_IMPOSSIBLE_TRAVEL.Ptr(),
				ImpossibleTravelOptions: &datadogV2.SecurityMonitoringRuleImpossibleTravelOptions{
					BaselineUserLocations: datadog.PtrBool(false),
				},
			},
			Name:    "Example-Security-Monitoring",
			Type:    datadogV2.SECURITYMONITORINGRULETYPECREATE_LOG_DETECTION.Ptr(),
			Tags:    []string{},
			Filters: []datadogV2.SecurityMonitoringFilter{},
		}}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateSecurityMonitoringRule(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateSecurityMonitoringRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateSecurityMonitoringRule`:\n%s\n", responseContent)
}
