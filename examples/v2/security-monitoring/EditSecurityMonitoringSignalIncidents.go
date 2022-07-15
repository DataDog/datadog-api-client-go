// Change the related incidents of a security signal returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	body := datadog.SecurityMonitoringSignalIncidentsUpdateRequest{
		Data: datadog.SecurityMonitoringSignalIncidentsUpdateData{
			Attributes: datadog.SecurityMonitoringSignalIncidentsUpdateAttributes{
				IncidentIds: []int64{
					2066,
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityMonitoringApi.EditSecurityMonitoringSignalIncidents(ctx, "AQAAAYG1bl5K4HuUewAAAABBWUcxYmw1S0FBQmt2RmhRN0V4ZUVnQUE", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.EditSecurityMonitoringSignalIncidents`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.EditSecurityMonitoringSignalIncidents`:\n%s\n", responseContent)
}
