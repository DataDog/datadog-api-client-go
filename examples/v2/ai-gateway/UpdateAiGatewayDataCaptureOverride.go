// Replace an AI Gateway data capture override returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.AiGatewayCaptureUpdateRequest{
		Data: datadogV2.AiGatewayCaptureUpdateRequestData{
			Attributes: datadogV2.AiGatewayCaptureUpdateRequestDataAttributes{
				ConfigurationRevision: "capture-revision-1",
				MetadataCapture: []datadogV2.AiGatewayMetadataCaptureItem{
					datadogV2.AIGATEWAYMETADATACAPTUREITEM_TOOL_USAGE,
					datadogV2.AIGATEWAYMETADATACAPTUREITEM_TOOL_OUTPUT_SIZE,
					datadogV2.AIGATEWAYMETADATACAPTUREITEM_TOOL_NAMES,
				},
				Mode:      datadogV2.AIGATEWAYCAPTUREUPDATEREQUESTDATAATTRIBUTESMODE_METADATA_ONLY,
				Name:      "Research",
				Operation: datadogV2.AIGATEWAYCAPTUREUPDATEREQUESTDATAATTRIBUTESOPERATION_CREATE.Ptr(),
				Revision:  1,
				Scope: &datadogV2.AiGatewayCaptureScope{
					Tags: map[string]string{
						"env": "prod",
					},
				},
			},
			Id:   uuid.MustParse("11111111-1111-4111-8111-111111111111"),
			Type: datadogV2.AIGATEWAYCAPTUREUPDATEREQUESTDATATYPE_DATA_CAPTURE_OVERRIDES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateAiGatewayDataCaptureOverride", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAIGatewayApi(apiClient)
	resp, r, err := api.UpdateAiGatewayDataCaptureOverride(ctx, uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayApi.UpdateAiGatewayDataCaptureOverride`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AIGatewayApi.UpdateAiGatewayDataCaptureOverride`:\n%s\n", responseContent)
}
