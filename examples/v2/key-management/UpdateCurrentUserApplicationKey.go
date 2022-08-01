// Edit an application key owned by current user returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v2/datadog"
)

func main() {
	// there is a valid "application_key" in the system
	ApplicationKeyDataID := os.Getenv("APPLICATION_KEY_DATA_ID")

	body := datadog.ApplicationKeyUpdateRequest{
		Data: datadog.ApplicationKeyUpdateData{
			Id:   ApplicationKeyDataID,
			Type: datadog.APPLICATIONKEYSTYPE_APPLICATION_KEYS,
			Attributes: datadog.ApplicationKeyUpdateAttributes{
				Name: common.PtrString("Application Key for managing dashboards-updated"),
			},
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewKeyManagementApi(apiClient)
	resp, r, err := api.UpdateCurrentUserApplicationKey(ctx, ApplicationKeyDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.UpdateCurrentUserApplicationKey`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `KeyManagementApi.UpdateCurrentUserApplicationKey`:\n%s\n", responseContent)
}
