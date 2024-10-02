// Update api handle returns "OK" response

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
	// there is a valid "api_handle" in the system
	APIHandleDataID := os.Getenv("API_HANDLE_DATA_ID")

	body := datadogV2.MicrosoftTeamsUpdateApiHandleRequest{
		Data: datadogV2.MicrosoftTeamsUpdateApiHandleRequestData{
			Attributes: datadogV2.MicrosoftTeamsApiHandleAttributes{
				Name: datadog.PtrString("fake-handle-name--updated"),
			},
			Type: datadogV2.MICROSOFTTEAMSAPIHANDLETYPE_HANDLE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMicrosoftTeamsIntegrationApi(apiClient)
	resp, r, err := api.UpdateApiHandle(ctx, APIHandleDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MicrosoftTeamsIntegrationApi.UpdateApiHandle`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MicrosoftTeamsIntegrationApi.UpdateApiHandle`:\n%s\n", responseContent)
}
