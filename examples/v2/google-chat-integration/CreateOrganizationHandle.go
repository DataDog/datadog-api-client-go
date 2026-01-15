// Create organization handle returns "CREATED" response

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
	body := datadogV2.GoogleChatCreateOrganizationHandleRequest{
		Data: datadogV2.GoogleChatCreateOrganizationHandleRequestData{
			Attributes: datadogV2.GoogleChatCreateOrganizationHandleRequestAttributes{
				Name:              "Example-Google-Chat-Integration",
				SpaceResourceName: "spaces/AAQA-zFIks8",
			},
		},
		Type: datadogV2.GOOGLECHATORGANIZATIONHANDLETYPE_GOOGLE_CHAT_ORGANIZATION_HANDLE_TYPE,
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewGoogleChatIntegrationApi(apiClient)
	resp, r, err := api.CreateOrganizationHandle(ctx, "e54cb570-c674-529c-769d-84b312288ed7", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoogleChatIntegrationApi.CreateOrganizationHandle`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `GoogleChatIntegrationApi.CreateOrganizationHandle`:\n%s\n", responseContent)
}
