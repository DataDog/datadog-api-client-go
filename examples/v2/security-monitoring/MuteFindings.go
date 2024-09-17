// Mute or unmute a batch of findings returns "OK" response

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
	body := datadogV2.BulkMuteFindingsRequest{
		Data: datadogV2.BulkMuteFindingsRequestData{
			Attributes: datadogV2.BulkMuteFindingsRequestAttributes{
				Mute: datadogV2.BulkMuteFindingsRequestProperties{
					ExpirationDate: datadog.PtrInt64(1778721573794),
					Muted:          true,
					Reason:         datadogV2.FINDINGMUTEREASON_ACCEPTED_RISK,
				},
			},
			Id: "dbe5f567-192b-4404-b908-29b70e1c9f76",
			Meta: datadogV2.BulkMuteFindingsRequestMeta{
				Findings: []datadogV2.BulkMuteFindingsRequestMetaFindings{
					{
						FindingId: datadog.PtrString("ZGVmLTAwcC1pZXJ-aS0wZjhjNjMyZDNmMzRlZTgzNw=="),
					},
				},
			},
			Type: datadogV2.FINDINGTYPE_FINDING,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.MuteFindings", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.MuteFindings(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.MuteFindings`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.MuteFindings`:\n%s\n", responseContent)
}
