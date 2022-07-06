// Edit an application key returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "application_key" in the system
	ApplicationKeyDataID := os.Getenv("APPLICATION_KEY_DATA_ID")

	body := datadog.ApplicationKeyUpdateRequest{
		Data: datadog.ApplicationKeyUpdateData{
			Id:   ApplicationKeyDataID,
			Type: datadog.APPLICATIONKEYSTYPE_APPLICATION_KEYS,
			Attributes: datadog.ApplicationKeyUpdateAttributes{
				Name: datadog.PtrString("Application Key for managing dashboards-updated"),
			},
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.KeyManagementApi(apiClient)
	resp, r, err := api.UpdateApplicationKey(ctx, ApplicationKeyDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.UpdateApplicationKey`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `KeyManagementApi.UpdateApplicationKey`:\n%s\n", responseContent)
}
