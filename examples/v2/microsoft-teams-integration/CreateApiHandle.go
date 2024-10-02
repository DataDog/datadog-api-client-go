// Create handle returns "CREATED" response

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
	body := datadogV2.MicrosoftTeamsCreateApiHandleRequest{
		Data: datadogV2.MicrosoftTeamsApiHandleRequestData{
			Attributes: datadogV2.MicrosoftTeamsApiHandleRequestAttributes{
				ChannelId: "fake-channel-id",
				Name:      "fake-handle-name",
				TeamId:    "00000000-0000-0000-0000-000000000000",
				TenantId:  "00000000-0000-0000-0000-000000000001",
			},
			Type: datadogV2.MICROSOFTTEAMSAPIHANDLETYPE_HANDLE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMicrosoftTeamsIntegrationApi(apiClient)
	resp, r, err := api.CreateApiHandle(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MicrosoftTeamsIntegrationApi.CreateApiHandle`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MicrosoftTeamsIntegrationApi.CreateApiHandle`:\n%s\n", responseContent)
}
