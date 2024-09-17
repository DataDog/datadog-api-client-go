// Modify the triage assignee of a security signal returns "OK" response

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
	body := datadogV2.SecurityMonitoringSignalAssigneeUpdateRequest{
		Data: datadogV2.SecurityMonitoringSignalAssigneeUpdateData{
			Attributes: datadogV2.SecurityMonitoringSignalAssigneeUpdateAttributes{
				Assignee: datadogV2.SecurityMonitoringTriageUser{
					Uuid: "",
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.EditSecurityMonitoringSignalAssignee(ctx, "AQAAAYG1bl5K4HuUewAAAABBWUcxYmw1S0FBQmt2RmhRN0V4ZUVnQUE", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.EditSecurityMonitoringSignalAssignee`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.EditSecurityMonitoringSignalAssignee`:\n%s\n", responseContent)
}
