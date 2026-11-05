// Delete an AI Gateway data capture override returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.AiGatewayCaptureDeleteRequest{
		Data: datadogV2.AiGatewayCaptureDeleteRequestData{
			Attributes: datadogV2.AiGatewayCaptureDeleteRequestDataAttributes{
				ConfigurationRevision: "capture-revision-1",
				Mode:                  datadogV2.AIGATEWAYCAPTUREDELETEREQUESTDATAATTRIBUTESMODE_METADATA_ONLY.Ptr(),
				Name:                  datadog.PtrString("Research"),
				Operation:             datadogV2.AIGATEWAYCAPTUREDELETEREQUESTDATAATTRIBUTESOPERATION_CREATE.Ptr(),
				Revision:              1,
				Scope: &datadogV2.AiGatewayCaptureScope{
					Tags: map[string]string{
						"env": "prod",
					},
				},
			},
			Id:   uuid.MustParse("11111111-1111-4111-8111-111111111111"),
			Type: datadogV2.AIGATEWAYCAPTUREDELETEREQUESTDATATYPE_DATA_CAPTURE_OVERRIDES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteAiGatewayDataCaptureOverride", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAIGatewayApi(apiClient)
	r, err := api.DeleteAiGatewayDataCaptureOverride(ctx, uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayApi.DeleteAiGatewayDataCaptureOverride`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
