// Mute or unmute a finding returns "OK" response

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
	body := datadogV2.MuteFindingRequest{
		Data: datadogV2.MuteFindingRequestData{
			Attributes: datadogV2.MuteFindingRequestAttributes{
				Mute: datadogV2.MuteFindingRequestProperties{
					Description:    datadog.PtrString("To be resolved later"),
					ExpirationDate: datadog.PtrInt64(1778721573794),
					Muted:          true,
					Reason:         datadogV2.FINDINGMUTEREASON_ACCEPTED_RISK,
				},
			},
			Id:   "AQAAAYbb1i0gijTHEQAAAABBWWJiMWkwZ0FBQ2FuNzBJVGZXZ3B3QUE",
			Type: datadogV2.FINDINGTYPE_FINDING,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateFinding", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.UpdateFinding(ctx, "AQAAAYbb1i0gijTHEQAAAABBWWJiMWkwZ0FBQ2FuNzBJVGZXZ3B3QUE", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.UpdateFinding`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.UpdateFinding`:\n%s\n", responseContent)
}
