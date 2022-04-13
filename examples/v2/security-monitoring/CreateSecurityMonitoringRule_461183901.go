// Create a detection rule with type 'impossible_travel' returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	body := datadog.SecurityMonitoringRuleCreatePayload{
		Queries: []datadog.SecurityMonitoringRuleQueryCreate{
			{
				Aggregation: datadog.SECURITYMONITORINGRULEQUERYAGGREGATION_GEO_DATA.Ptr(),
				GroupByFields: []string{
					"@usr.id",
				},
				DistinctFields: []string{},
				Metric:         datadog.PtrString("@network.client.geoip"),
				Query:          "*",
			},
		},
		Cases: []datadog.SecurityMonitoringRuleCaseCreate{
			{
				Name:          datadog.PtrString(""),
				Status:        datadog.SECURITYMONITORINGRULESEVERITY_INFO,
				Notifications: []string{},
			},
		},
		HasExtendedTitle: datadog.PtrBool(true),
		Message:          "test",
		IsEnabled:        true,
		Options: datadog.SecurityMonitoringRuleOptions{
			MaxSignalDuration: datadog.SECURITYMONITORINGRULEMAXSIGNALDURATION_ONE_DAY.Ptr(),
			EvaluationWindow:  datadog.SECURITYMONITORINGRULEEVALUATIONWINDOW_FIFTEEN_MINUTES.Ptr(),
			KeepAlive:         datadog.SECURITYMONITORINGRULEKEEPALIVE_ONE_HOUR.Ptr(),
			DetectionMethod:   datadog.SECURITYMONITORINGRULEDETECTIONMETHOD_IMPOSSIBLE_TRAVEL.Ptr(),
			ImpossibleTravelOptions: &datadog.SecurityMonitoringRuleImpossibleTravelOptions{
				BaselineUserLocations: datadog.PtrBool(false),
			},
		},
		Name:    "Example-Create_a_detection_rule_with_type_impossible_travel_returns_OK_response",
		Type:    datadog.SECURITYMONITORINGRULETYPECREATE_LOG_DETECTION.Ptr(),
		Tags:    []string{},
		Filters: []datadog.SecurityMonitoringFilter{},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityMonitoringApi.CreateSecurityMonitoringRule(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateSecurityMonitoringRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateSecurityMonitoringRule`:\n%s\n", responseContent)
}
