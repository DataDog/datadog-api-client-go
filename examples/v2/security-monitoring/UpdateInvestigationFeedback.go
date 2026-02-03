// Update investigation feedback returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.SecurityMonitoringSignalInvestigationFeedbackRequest{
		Data: datadogV2.SecurityMonitoringSignalInvestigationFeedbackRequestData{
			Attributes: datadogV2.SecurityMonitoringSignalInvestigationFeedbackRequestAttributes{
				Feedback: "positive",
				FeedbackContent: []datadogV2.SecurityMonitoringSignalInvestigationFeedbackSection{
					{
						Id: "section-1",
						Metrics: []datadogV2.SecurityMonitoringSignalInvestigationFeedbackMetric{
							{
								Placeholder: datadog.PtrString("Enter your feedback here"),
								Prompt:      "How helpful was this investigation?",
								Response:    "Very helpful",
								Type:        "sentiment",
							},
						},
						Title: "Investigation Quality",
					},
				},
				Incomplete: datadog.PtrBool(false),
				Rating:     datadog.PtrString("positive"),
				SignalId:   "AAAAAWgN8Xwgr1vKDQAAAABBV2dOOFh3ZzZobm1mWXJFYTR0OA",
				Type:       datadog.PtrString("metrics"),
			},
			Type: datadogV2.SECURITYMONITORINGSIGNALINVESTIGATIONFEEDBACKTYPE_INVESTIGATION_FEEDBACK,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateInvestigationFeedback", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	r, err := api.UpdateInvestigationFeedback(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.UpdateInvestigationFeedback`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
