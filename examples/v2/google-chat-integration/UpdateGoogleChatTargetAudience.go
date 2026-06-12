// Update a target audience returns "OK" response

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
	body := datadogV2.GoogleChatTargetAudienceUpdateRequest{
		Data: datadogV2.GoogleChatTargetAudienceUpdateRequestData{
			Attributes: datadogV2.GoogleChatTargetAudienceUpdateRequestAttributes{
				AudienceId:   datadog.PtrString("fake-audience-id-1"),
				AudienceName: datadog.PtrString("fake audience name 1"),
			},
			Type: datadogV2.GOOGLECHATTARGETAUDIENCETYPE_GOOGLE_CHAT_TARGET_AUDIENCE_TYPE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewGoogleChatIntegrationApi(apiClient)
	resp, r, err := api.UpdateGoogleChatTargetAudience(ctx, "organization_binding_id", "target_audience_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoogleChatIntegrationApi.UpdateGoogleChatTargetAudience`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `GoogleChatIntegrationApi.UpdateGoogleChatTargetAudience`:\n%s\n", responseContent)
}
