// Update a single service object returns "OK" response

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
	// there is a valid "opsgenie_service" in the system
	OpsgenieServiceDataID := os.Getenv("OPSGENIE_SERVICE_DATA_ID")

	body := datadogV2.OpsgenieServiceUpdateRequest{
		Data: datadogV2.OpsgenieServiceUpdateData{
			Attributes: datadogV2.OpsgenieServiceUpdateAttributes{
				Name:           datadog.PtrString("fake-opsgenie-service-name--updated"),
				OpsgenieApiKey: datadog.PtrString("00000000-0000-0000-0000-000000000000"),
				Region:         datadogV2.OPSGENIESERVICEREGIONTYPE_EU.Ptr(),
			},
			Id:   OpsgenieServiceDataID,
			Type: datadogV2.OPSGENIESERVICETYPE_OPSGENIE_SERVICE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOpsgenieIntegrationApi(apiClient)
	resp, r, err := api.UpdateOpsgenieService(ctx, OpsgenieServiceDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsgenieIntegrationApi.UpdateOpsgenieService`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OpsgenieIntegrationApi.UpdateOpsgenieService`:\n%s\n", responseContent)
}
