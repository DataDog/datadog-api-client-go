// Edit an application key owned by current user returns "OK" response

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
	// there is a valid "application_key" in the system
	ApplicationKeyDataID := os.Getenv("APPLICATION_KEY_DATA_ID")

	body := datadogV2.ApplicationKeyUpdateRequest{
		Data: datadogV2.ApplicationKeyUpdateData{
			Id:   ApplicationKeyDataID,
			Type: datadogV2.APPLICATIONKEYSTYPE_APPLICATION_KEYS,
			Attributes: datadogV2.ApplicationKeyUpdateAttributes{
				Name: datadog.PtrString("Application Key for managing dashboards-updated"),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewKeyManagementApi(apiClient)
	resp, r, err := api.UpdateCurrentUserApplicationKey(ctx, ApplicationKeyDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.UpdateCurrentUserApplicationKey`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `KeyManagementApi.UpdateCurrentUserApplicationKey`:\n%s\n", responseContent)
}
