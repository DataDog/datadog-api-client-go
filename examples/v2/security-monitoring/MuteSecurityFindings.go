// Mute or unmute security findings returns "Accepted" response

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
	body := datadogV2.MuteFindingsRequest{
		Data: datadogV2.MuteFindingsRequestData{
			Attributes: datadogV2.MuteFindingsRequestDataAttributes{
				Mute: datadogV2.MuteFindingsMuteAttributes{
					Description: datadog.PtrString("To be resolved later."),
					ExpireAt:    datadog.PtrInt64(1778721573794),
					IsMuted:     true,
					Reason:      datadogV2.MUTEFINDINGSREASON_PENDING_FIX,
				},
			},
			Id: datadog.PtrString("93bfeb70-af47-424d-908a-948d3f08e37f"),
			Relationships: datadogV2.MuteFindingsRequestDataRelationships{
				Findings: datadogV2.Findings{
					Data: []datadogV2.FindingData{
						{
							Id:   "ZGVmLTAwcC1pZXJ-aS0wZjhjNjMyZDNmMzRlZTgzNw==",
							Type: datadogV2.FINDINGDATATYPE_FINDINGS,
						},
					},
				},
			},
			Type: datadogV2.MUTEDATATYPE_MUTE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.MuteSecurityFindings", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.MuteSecurityFindings(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.MuteSecurityFindings`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.MuteSecurityFindings`:\n%s\n", responseContent)
}
