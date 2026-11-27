// Update a RUM operation link returns "OK" response

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
	// there is a valid "rum_operation_strong_link" in the system
	RumOperationStrongLinkDataAttributesFeatureID := os.Getenv("RUM_OPERATION_STRONG_LINK_DATA_ATTRIBUTES_FEATURE_ID")
	RumOperationStrongLinkDataAttributesOperationID := os.Getenv("RUM_OPERATION_STRONG_LINK_DATA_ATTRIBUTES_OPERATION_ID")

	body := datadogV2.RUMOperationStrongLinkUpdateRequest{
		Data: datadogV2.RUMOperationStrongLinkUpdateRequestData{
			Type: datadogV2.RUMOPERATIONSTRONGLINKTYPE_STRONG_LINKS,
			Attributes: datadogV2.RUMOperationStrongLinkUpdateRequestAttributes{
				Status: datadogV2.RUMOPERATIONSTRONGLINKUPDATESTATUS_REJECTED,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMOperationsApi(apiClient)
	resp, r, err := api.UpdateRUMOperationStrongLink(ctx, RumOperationStrongLinkDataAttributesOperationID, RumOperationStrongLinkDataAttributesFeatureID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMOperationsApi.UpdateRUMOperationStrongLink`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMOperationsApi.UpdateRUMOperationStrongLink`:\n%s\n", responseContent)
}
