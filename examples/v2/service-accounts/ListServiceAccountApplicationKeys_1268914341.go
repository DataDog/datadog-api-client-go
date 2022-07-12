// Get all app keys owned by this service account returns "OK" response

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
	// there is a valid "service_account_user" in the system
	ServiceAccountUserDataID := os.Getenv("SERVICE_ACCOUNT_USER_DATA_ID")

	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewServiceAccountsApi(apiClient)
	resp, r, err := api.ListServiceAccountApplicationKeys(ctx, ServiceAccountUserDataID, *datadog.NewListServiceAccountApplicationKeysOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.ListServiceAccountApplicationKeys`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.ListServiceAccountApplicationKeys`:\n%s\n", responseContent)
}
