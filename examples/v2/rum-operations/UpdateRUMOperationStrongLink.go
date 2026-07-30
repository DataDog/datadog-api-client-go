// Update a RUM operation strong link returns "OK" response

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
	body := datadogV2.RUMOperationStrongLinkUpdateRequest{
		Data: datadogV2.RUMOperationStrongLinkUpdateRequestData{
			Attributes: datadogV2.RUMOperationStrongLinkUpdateRequestAttributes{
				Status: datadogV2.RUMOPERATIONSTRONGLINKUPDATESTATUS_CONFIRMED,
			},
			Type: datadogV2.RUMOPERATIONSTRONGLINKTYPE_STRONG_LINKS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateRUMOperationStrongLink", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMOperationsApi(apiClient)
	resp, r, err := api.UpdateRUMOperationStrongLink(ctx, "rum_operation_id", "feature_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMOperationsApi.UpdateRUMOperationStrongLink`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMOperationsApi.UpdateRUMOperationStrongLink`:\n%s\n", responseContent)
}
