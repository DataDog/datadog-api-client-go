// Delete an app key owned by this service account returns "No Content" response

package main

import (
	"context"
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
	r, err := api.DeleteServiceAccountApplicationKey(ctx, ServiceAccountUserDataID, ServiceAccountApplicationKeyDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.DeleteServiceAccountApplicationKey`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
