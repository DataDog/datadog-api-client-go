// Create a target audience returns "CREATED" response

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
	body := datadogV2.GoogleChatTargetAudienceCreateRequest{
		Data: datadogV2.GoogleChatTargetAudienceCreateRequestData{
			Attributes: datadogV2.GoogleChatTargetAudienceCreateRequestAttributes{
				AudienceId:   "fake-audience-id-1",
				AudienceName: "fake audience name 1",
			},
			Type: datadogV2.GOOGLECHATTARGETAUDIENCETYPE_GOOGLE_CHAT_TARGET_AUDIENCE_TYPE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewGoogleChatIntegrationApi(apiClient)
	resp, r, err := api.CreateGoogleChatTargetAudience(ctx, "organization_binding_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoogleChatIntegrationApi.CreateGoogleChatTargetAudience`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `GoogleChatIntegrationApi.CreateGoogleChatTargetAudience`:\n%s\n", responseContent)
}
