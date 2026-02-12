// Create a detection rule with detection method 'anomaly_detection' with enabled feature 'instantaneousBaseline' returns
// "OK" response

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
			Name:      "Example-Security-Monitoring",
			Type:      datadogV2.SECURITYMONITORINGRULETYPECREATE_LOG_DETECTION.Ptr(),
			IsEnabled: true,
			Queries: []datadogV2.SecurityMonitoringStandardRuleQuery{
				{
					Aggregation:    datadogV2.SECURITYMONITORINGRULEQUERYAGGREGATION_COUNT.Ptr(),
					DataSource:     datadogV2.SECURITYMONITORINGSTANDARDDATASOURCE_LOGS.Ptr(),
					DistinctFields: []string{},
					GroupByFields: []string{
						"@usr.email",
						"@network.client.ip",
					},
					HasOptionalGroupByFields: datadog.PtrBool(false),
					Name:                     datadog.PtrString(""),
					Query:                    datadog.PtrString("service:app status:error"),
				},
			},
			Cases: []datadogV2.SecurityMonitoringRuleCaseCreate{
				{
					Name:          datadog.PtrString(""),
					Status:        datadogV2.SECURITYMONITORINGRULESEVERITY_INFO,
					Notifications: []string{},
					Condition:     datadog.PtrString("a > 0.995"),
				},
			},
			Message: "An anomaly detection rule",
			Options: datadogV2.SecurityMonitoringRuleOptions{
				DetectionMethod:   datadogV2.SECURITYMONITORINGRULEDETECTIONMETHOD_ANOMALY_DETECTION.Ptr(),
				EvaluationWindow:  datadogV2.SECURITYMONITORINGRULEEVALUATIONWINDOW_FIFTEEN_MINUTES.Ptr(),
				KeepAlive:         datadogV2.SECURITYMONITORINGRULEKEEPALIVE_ONE_HOUR.Ptr(),
				MaxSignalDuration: datadogV2.SECURITYMONITORINGRULEMAXSIGNALDURATION_ONE_DAY.Ptr(),
				AnomalyDetectionOptions: &datadogV2.SecurityMonitoringRuleAnomalyDetectionOptions{
					BucketDuration:        datadogV2.SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSBUCKETDURATION_FIVE_MINUTES.Ptr(),
					LearningDuration:      datadogV2.SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSLEARNINGDURATION_ONE_DAY.Ptr(),
					DetectionTolerance:    datadogV2.SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSDETECTIONTOLERANCE_THREE.Ptr(),
					InstantaneousBaseline: datadog.PtrBool(true),
				},
			},
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
