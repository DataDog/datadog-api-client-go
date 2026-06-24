// Update a governance control returns "OK" response

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
	body := datadogV2.GovernanceControlUpdateRequest{
		Data: datadogV2.GovernanceControlUpdateData{
			Attributes: &datadogV2.GovernanceControlUpdateAttributes{
				DetectionFrequency:    datadog.PtrString("daily"),
				MitigationType:        datadog.PtrString("revoke_api_key"),
				Name:                  datadog.PtrString("Unused API Keys"),
				NotificationFrequency: datadog.PtrString("daily"),
				NotificationType:      datadog.PtrString("slack"),
			},
			Id:   datadog.PtrString("0d4e6f8a-1b2c-3d4e-5f6a-7b8c9d0e1f2a"),
			Type: datadogV2.GOVERNANCECONTROLRESOURCETYPE_GOVERNANCE_CONTROL,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateGovernanceControl", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewGovernanceControlsApi(apiClient)
	resp, r, err := api.UpdateGovernanceControl(ctx, "detection_type", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceControlsApi.UpdateGovernanceControl`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `GovernanceControlsApi.UpdateGovernanceControl`:\n%s\n", responseContent)
}
