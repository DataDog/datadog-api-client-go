// Replace AI Gateway governance controls returns "OK" response

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
	body := datadogV2.AiGatewayGovernanceReplaceRequest{
		Data: datadogV2.AiGatewayGovernanceReplaceData{
			Attributes: datadogV2.AiGatewayGovernanceReplaceAttributes{
				DataCaptureMode: datadogV2.AIGATEWAYDATACAPTUREMODE_METADATA_ONLY,
				MetadataCapture: []datadogV2.AiGatewayMetadataCaptureItem{
					datadogV2.AIGATEWAYMETADATACAPTUREITEM_TOOL_USAGE,
					datadogV2.AIGATEWAYMETADATACAPTUREITEM_TOOL_OUTPUT_SIZE,
					datadogV2.AIGATEWAYMETADATACAPTUREITEM_TOOL_NAMES,
				},
				ZdrMode: datadogV2.AIGATEWAYZDRMODE_OFF.Ptr(),
			},
			Id:   datadog.PtrString("organization"),
			Type: datadogV2.AIGATEWAYGOVERNANCETYPE_GOVERNANCE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateAiGatewayGovernance", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAIGatewayApi(apiClient)
	resp, r, err := api.UpdateAiGatewayGovernance(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayApi.UpdateAiGatewayGovernance`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AIGatewayApi.UpdateAiGatewayGovernance`:\n%s\n", responseContent)
}
