// Update Scorecard outcomes asynchronously returns "Accepted" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.UpdateOutcomesAsyncRequest{
		Data: &datadogV2.UpdateOutcomesAsyncRequestData{
			Attributes: &datadogV2.UpdateOutcomesAsyncAttributes{
				Results: []datadogV2.UpdateOutcomesAsyncRequestItem{
					{
						EntityReference: "service:my-service",
						Remarks:         datadog.PtrString(`See: <a href="https://app.datadoghq.com/services">Services</a>`),
						RuleId:          "q8MQxk8TCqrHnWkx",
						State:           datadogV2.STATE_PASS,
					},
				},
			},
			Type: datadogV2.UPDATEOUTCOMESASYNCTYPE_BATCHED_OUTCOME.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateScorecardOutcomesAsync", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewServiceScorecardsApi(apiClient)
	r, err := api.UpdateScorecardOutcomesAsync(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceScorecardsApi.UpdateScorecardOutcomesAsync`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
