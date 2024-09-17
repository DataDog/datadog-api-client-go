// Create an application key for this service account returns "Created" response

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

	body := datadogV2.ApplicationKeyCreateRequest{
		Data: datadogV2.ApplicationKeyCreateData{
			Attributes: datadogV2.ApplicationKeyCreateAttributes{
				Name: "Example-Service-Account",
			},
			Type: datadogV2.APPLICATIONKEYSTYPE_APPLICATION_KEYS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewServiceAccountsApi(apiClient)
	resp, r, err := api.CreateServiceAccountApplicationKey(ctx, ServiceAccountUserDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.CreateServiceAccountApplicationKey`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.CreateServiceAccountApplicationKey`:\n%s\n", responseContent)
}
