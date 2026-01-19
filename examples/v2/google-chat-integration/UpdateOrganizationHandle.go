// Update organization handle returns "OK" response

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
	// there is a valid "organization_handle" in the system
	OrganizationHandleDataID := os.Getenv("ORGANIZATION_HANDLE_DATA_ID")

	body := datadogV2.GoogleChatUpdateOrganizationHandleRequest{
		Data: datadogV2.GoogleChatUpdateOrganizationHandleRequestData{
			Attributes: datadogV2.GoogleChatUpdateOrganizationHandleRequestAttributes{
				Name: datadog.PtrString("fake-handle-name--updated"),
			},
		},
		Type: datadogV2.GOOGLECHATORGANIZATIONHANDLETYPE_GOOGLE_CHAT_ORGANIZATION_HANDLE_TYPE,
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewGoogleChatIntegrationApi(apiClient)
	resp, r, err := api.UpdateOrganizationHandle(ctx, "e54cb570-c674-529c-769d-84b312288ed7", OrganizationHandleDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoogleChatIntegrationApi.UpdateOrganizationHandle`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `GoogleChatIntegrationApi.UpdateOrganizationHandle`:\n%s\n", responseContent)
}
