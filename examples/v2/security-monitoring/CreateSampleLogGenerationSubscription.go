// Subscribe to sample log generation returns "OK" response

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
	body := datadogV2.SampleLogGenerationSubscriptionCreateRequest{
		Data: datadogV2.SampleLogGenerationSubscriptionCreateData{
			Attributes: datadogV2.SampleLogGenerationSubscriptionCreateAttributes{
				ContentPackId: "aws-cloudtrail",
				Duration:      datadogV2.SAMPLELOGGENERATIONDURATION_THREE_DAYS.Ptr(),
			},
			Type: datadogV2.SAMPLELOGGENERATIONSUBSCRIPTIONREQUESTTYPE_SUBSCRIPTION_REQUESTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateSampleLogGenerationSubscription", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateSampleLogGenerationSubscription(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateSampleLogGenerationSubscription`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateSampleLogGenerationSubscription`:\n%s\n", responseContent)
}
