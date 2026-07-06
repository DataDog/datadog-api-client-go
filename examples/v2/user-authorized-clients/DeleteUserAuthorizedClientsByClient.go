// Delete all user authorized clients for a client returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewUserAuthorizedClientsApi(apiClient)
	r, err := api.DeleteUserAuthorizedClientsByClient(ctx, "00000000-0000-0000-0000-000000000010")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAuthorizedClientsApi.DeleteUserAuthorizedClientsByClient`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
