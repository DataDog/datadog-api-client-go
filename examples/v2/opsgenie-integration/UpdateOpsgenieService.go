// Update a single service object returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "opsgenie_service" in the system
	OpsgenieServiceDataID := os.Getenv("OPSGENIE_SERVICE_DATA_ID")

	body := datadog.OpsgenieServiceUpdateRequest{
		Data: datadog.OpsgenieServiceUpdateData{
			Attributes: datadog.OpsgenieServiceUpdateAttributes{
				Name:           datadog.PtrString("fake-opsgenie-service-name--updated"),
				OpsgenieApiKey: datadog.PtrString("00000000-0000-0000-0000-000000000000"),
				Region:         datadog.OPSGENIESERVICEREGIONTYPE_EU.Ptr(),
			},
			Id:   OpsgenieServiceDataID,
			Type: datadog.OPSGENIESERVICETYPE_OPSGENIE_SERVICE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsgenieIntegrationApi.UpdateOpsgenieService(ctx, OpsgenieServiceDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsgenieIntegrationApi.UpdateOpsgenieService`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OpsgenieIntegrationApi.UpdateOpsgenieService`:\n%s\n", responseContent)
}
