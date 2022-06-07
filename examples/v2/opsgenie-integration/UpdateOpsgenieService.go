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
	body := datadog.OpsgenieServiceUpdateRequest{
		Data: datadog.OpsgenieServiceUpdateData{
			Attributes: datadog.OpsgenieServiceUpdateAttributes{
				CustomUrl:      *datadog.NewNullableString(datadog.PtrString("https://example.com")),
				Name:           datadog.PtrString("fake-opsgenie-service-name"),
				OpsgenieApiKey: datadog.PtrString("00000000-0000-0000-0000-000000000000"),
				Region:         datadog.OPSGENIESERVICEREGIONTYPE_US.Ptr(),
			},
			Id:   "596da4af-0563-4097-90ff-07230c3f9db3",
			Type: datadog.OPSGENIESERVICETYPE_OPSGENIE_SERVICE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsgenieIntegrationApi.UpdateOpsgenieService(ctx, "integration_service_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsgenieIntegrationApi.UpdateOpsgenieService`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OpsgenieIntegrationApi.UpdateOpsgenieService`:\n%s\n", responseContent)
}
