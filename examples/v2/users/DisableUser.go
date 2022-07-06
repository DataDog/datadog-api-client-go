// Disable a user returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "user" in the system
	UserDataID := os.Getenv("USER_DATA_ID")

	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.UsersApi(apiClient)
	r, err := api.DisableUser(ctx, UserDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.DisableUser`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
