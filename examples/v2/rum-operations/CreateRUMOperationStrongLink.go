// Create a RUM operation link returns "Created" response

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
	// there is a valid "rum_operation" in the system
	RumOperationDataID := os.Getenv("RUM_OPERATION_DATA_ID")

	body := datadogV2.RUMOperationStrongLinkCreateRequest{
		Data: datadogV2.RUMOperationStrongLinkCreateRequestData{
			Type: datadogV2.RUMOPERATIONSTRONGLINKTYPE_STRONG_LINKS,
			Attributes: datadogV2.RUMOperationStrongLinkCreateRequestAttributes{
				OperationId: datadog.PtrString(RumOperationDataID),
				FeatureId:   "bdd-examplerumoperation",
				Status:      datadogV2.RUMOPERATIONSTRONGLINKSTATUS_DRAFT.Ptr(),
				Tags: []string{
					"team:rum-backend",
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMOperationsApi(apiClient)
	resp, r, err := api.CreateRUMOperationStrongLink(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMOperationsApi.CreateRUMOperationStrongLink`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMOperationsApi.CreateRUMOperationStrongLink`:\n%s\n", responseContent)
}
