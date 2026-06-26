// Mitigate governance detections returns "Accepted" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.GovernanceMitigationRequest{
		Data: datadogV2.GovernanceMitigationRequestData{
			Attributes: &datadogV2.GovernanceMitigationRequestAttributes{
				DetectionIds: []string{
					"3f9b2c1a-8d4e-4a6b-9c2f-1e7d5a0b3c4d",
				},
				DetectionType:  datadog.PtrString("unused_api_keys"),
				MitigationType: datadog.PtrString("revoke_api_key"),
			},
			Type: datadogV2.GOVERNANCECONTROLDETECTIONRESOURCETYPE_GOVERNANCE_CONTROL_DETECTION,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateGovernanceMitigation", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewGovernanceControlsApi(apiClient)
	r, err := api.CreateGovernanceMitigation(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceControlsApi.CreateGovernanceMitigation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
