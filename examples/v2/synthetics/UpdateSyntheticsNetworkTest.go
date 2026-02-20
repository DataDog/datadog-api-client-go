// Edit a Network Path test returns "OK" response

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
	body := datadogV2.SyntheticsNetworkTestEditRequest{
		Data: datadogV2.SyntheticsNetworkTestEdit{
			Attributes: datadogV2.SyntheticsNetworkTest{
				Config: datadogV2.SyntheticsNetworkTestConfig{
					Assertions: []datadogV2.SyntheticsNetworkAssertion{
						datadogV2.SyntheticsNetworkAssertion{
							SyntheticsNetworkAssertionLatency: &datadogV2.SyntheticsNetworkAssertionLatency{
								Operator: datadogV2.SYNTHETICSNETWORKASSERTIONOPERATOR_LESS_THAN,
								Property: datadogV2.SYNTHETICSNETWORKASSERTIONPROPERTY_AVG,
								Target:   500,
								Type:     datadogV2.SYNTHETICSNETWORKASSERTIONLATENCYTYPE_LATENCY,
							}},
					},
					Request: &datadogV2.SyntheticsNetworkTestRequest{
						E2eQueries:        50,
						Host:              "",
						MaxTtl:            30,
						Port:              datadog.PtrInt64(443),
						TcpMethod:         datadogV2.SYNTHETICSNETWORKTESTREQUESTTCPMETHOD_PREFER_SACK.Ptr(),
						TracerouteQueries: 3,
					},
				},
				Locations: []string{
					"aws:us-east-1",
					"agent:my-agent-name",
				},
				Message: "Network Path test notification",
				Name:    "Example Network Path test",
				Options: datadogV2.SyntheticsTestOptions{
					MonitorOptions: &datadogV2.SyntheticsTestOptionsMonitorOptions{
						NotificationPresetName: datadogV2.SYNTHETICSTESTOPTIONSMONITOROPTIONSNOTIFICATIONPRESETNAME_SHOW_ALL.Ptr(),
					},
					RestrictedRoles: []string{
						"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
					},
					Retry: &datadogV2.SyntheticsTestOptionsRetry{},
					Scheduling: &datadogV2.SyntheticsTestOptionsScheduling{
						Timeframes: []datadogV2.SyntheticsTestOptionsSchedulingTimeframe{
							{
								Day:  1,
								From: "07:00",
								To:   "16:00",
							},
							{
								Day:  3,
								From: "07:00",
								To:   "16:00",
							},
						},
						Timezone: "America/New_York",
					},
				},
				Status:  datadogV2.SYNTHETICSTESTPAUSESTATUS_LIVE.Ptr(),
				Subtype: datadogV2.SYNTHETICSNETWORKTESTSUBTYPE_TCP.Ptr(),
				Tags: []string{
					"env:production",
				},
				Type: datadogV2.SYNTHETICSNETWORKTESTTYPE_NETWORK,
			},
			Type: datadogV2.SYNTHETICSNETWORKTESTTYPE_NETWORK,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSyntheticsApi(apiClient)
	resp, r, err := api.UpdateSyntheticsNetworkTest(ctx, "public_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.UpdateSyntheticsNetworkTest`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.UpdateSyntheticsNetworkTest`:\n%s\n", responseContent)
}
