// Create an AI Gateway data capture override returns "Created" response

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
	body := datadogV2.AiGatewayCaptureCreateRequest{
		Data: datadogV2.AiGatewayCaptureCreateRequestData{
			Attributes: datadogV2.AiGatewayCaptureCreateRequestDataAttributes{
				ConfigurationRevision: "capture-revision-1",
				MetadataCapture: []datadogV2.AiGatewayMetadataCaptureItem{
					datadogV2.AIGATEWAYMETADATACAPTUREITEM_TOOL_USAGE,
					datadogV2.AIGATEWAYMETADATACAPTUREITEM_TOOL_OUTPUT_SIZE,
					datadogV2.AIGATEWAYMETADATACAPTUREITEM_TOOL_NAMES,
				},
				Mode:      datadogV2.AIGATEWAYCAPTURECREATEREQUESTDATAATTRIBUTESMODE_METADATA_ONLY,
				Name:      "Research",
				Operation: datadogV2.AIGATEWAYCAPTURECREATEREQUESTDATAATTRIBUTESOPERATION_CREATE.Ptr(),
				Revision:  datadog.PtrInt64(1),
				Scope: datadogV2.AiGatewayCaptureScope{
					Tags: map[string]string{
						"env": "prod",
					},
				},
			},
			Id:   uuid.MustParse("11111111-1111-4111-8111-111111111111"),
			Type: datadogV2.AIGATEWAYCAPTURECREATEREQUESTDATATYPE_DATA_CAPTURE_OVERRIDES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateAiGatewayDataCaptureOverride", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAIGatewayApi(apiClient)
	resp, r, err := api.CreateAiGatewayDataCaptureOverride(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayApi.CreateAiGatewayDataCaptureOverride`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AIGatewayApi.CreateAiGatewayDataCaptureOverride`:\n%s\n", responseContent)
}
