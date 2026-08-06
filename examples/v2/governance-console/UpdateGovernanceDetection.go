// Update a detection returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.GovernanceControlDetectionUpdateRequest{
		Data: datadogV2.GovernanceControlDetectionUpdateData{
			Attributes: &datadogV2.GovernanceControlDetectionUpdateAttributes{
				AssignedTeam:  datadog.PtrString("platform-security"),
				AssignedTo:    datadog.PtrString("11111111-2222-3333-4444-555555555555"),
				MitigateAfter: datadog.PtrTime(time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC)),
				State:         datadogV2.GOVERNANCECONTROLDETECTIONUPDATESTATE_EXCEPTION.Ptr(),
			},
			Type: datadogV2.GOVERNANCECONTROLDETECTIONRESOURCETYPE_GOVERNANCE_CONTROL_DETECTION,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateGovernanceDetection", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewGovernanceConsoleApi(apiClient)
	resp, r, err := api.UpdateGovernanceDetection(ctx, "detection_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceConsoleApi.UpdateGovernanceDetection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `GovernanceConsoleApi.UpdateGovernanceDetection`:\n%s\n", responseContent)
}
