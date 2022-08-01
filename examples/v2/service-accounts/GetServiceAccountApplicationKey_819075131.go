// Get one app key owned by this service account returns "OK" response

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
	// there is a valid "service_account_user" in the system
	ServiceAccountUserDataID := os.Getenv("SERVICE_ACCOUNT_USER_DATA_ID")

	// there is a valid "service_account_application_key" for "service_account_user"
	ServiceAccountApplicationKeyDataID := os.Getenv("SERVICE_ACCOUNT_APPLICATION_KEY_DATA_ID")

	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewServiceAccountsApi(apiClient)
	resp, r, err := api.GetServiceAccountApplicationKey(ctx, ServiceAccountUserDataID, ServiceAccountApplicationKeyDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.GetServiceAccountApplicationKey`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.GetServiceAccountApplicationKey`:\n%s\n", responseContent)
}
