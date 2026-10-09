// Apply a severity override to security findings returns "Accepted" response

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
	body := datadogV2.SeverityOverrideRequest{
		Data: datadogV2.SeverityOverrideRequestData{
			Attributes: datadogV2.SeverityOverrideRequestDataAttributes{
				Severity: datadogV2.SeverityOverrideAttributes{
					SeverityOverrideSet: &datadogV2.SeverityOverrideSet{
						Action:      datadogV2.SEVERITYOVERRIDESETACTIONTYPE_SET,
						Description: datadog.PtrString("Database contains sensitive data."),
						Value:       datadogV2.SEVERITYOVERRIDEVALUE_HIGH,
					}},
			},
			Relationships: datadogV2.SeverityOverrideRequestDataRelationships{
				Findings: datadogV2.Findings{
					Data: []datadogV2.FindingData{
						{
							Id:   "ZGVmLTAwMC0wYmd-MDE4NjcyMDJkMzE4MDE5ODY5MGE4ZmQ2MmFlMjg0Y2M=",
							Type: datadogV2.FINDINGDATATYPE_FINDINGS,
						},
					},
				},
			},
			Type: datadogV2.SEVERITYOVERRIDEDATATYPE_SEVERITY_OVERRIDE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateFindingsSeverity", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.UpdateFindingsSeverity(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.UpdateFindingsSeverity`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.UpdateFindingsSeverity`:\n%s\n", responseContent)
}
