// Change the related incidents of a security signal returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
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
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.EditSecurityMonitoringSignalIncidents(ctx, "AQAAAYG1bl5K4HuUewAAAABBWUcxYmw1S0FBQmt2RmhRN0V4ZUVnQUE", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.EditSecurityMonitoringSignalIncidents`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.EditSecurityMonitoringSignalIncidents`:\n%s\n", responseContent)
}
