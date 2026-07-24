// Create a RUM operation strong link returns "Created" response

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
	body := datadogV2.RUMOperationStrongLinkCreateRequest{
		Data: datadogV2.RUMOperationStrongLinkCreateRequestData{
			Attributes: datadogV2.RUMOperationStrongLinkCreateRequestAttributes{
				Description: *datadog.NewNullableString(nil),
				FeatureId:   "feature-123",
				OperationId: datadog.PtrString("abc12345-1234-5678-abcd-ef1234567890"),
				Status:      datadogV2.RUMOPERATIONSTRONGLINKSTATUS_CONFIRMED.Ptr(),
				Tags:        []string{},
			},
			Type: datadogV2.RUMOPERATIONSTRONGLINKTYPE_STRONG_LINKS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateRUMOperationStrongLink", true)
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
