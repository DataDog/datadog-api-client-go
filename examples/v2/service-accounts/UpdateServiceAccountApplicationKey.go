// Edit an application key for this service account returns "OK" response

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
	// there is a valid "service_account_user" in the system
	ServiceAccountUserDataID := os.Getenv("SERVICE_ACCOUNT_USER_DATA_ID")

	// there is a valid "service_account_application_key" for "service_account_user"
	ServiceAccountApplicationKeyDataID := os.Getenv("SERVICE_ACCOUNT_APPLICATION_KEY_DATA_ID")

	body := datadogV2.ApplicationKeyUpdateRequest{
		Data: datadogV2.ApplicationKeyUpdateData{
			Id:   ServiceAccountApplicationKeyDataID,
			Type: datadogV2.APPLICATIONKEYSTYPE_APPLICATION_KEYS,
			Attributes: datadogV2.ApplicationKeyUpdateAttributes{
				Name: datadog.PtrString("Application Key for managing dashboards-updated"),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewServiceAccountsApi(apiClient)
	resp, r, err := api.UpdateServiceAccountApplicationKey(ctx, ServiceAccountUserDataID, ServiceAccountApplicationKeyDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.UpdateServiceAccountApplicationKey`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.UpdateServiceAccountApplicationKey`:\n%s\n", responseContent)
}
