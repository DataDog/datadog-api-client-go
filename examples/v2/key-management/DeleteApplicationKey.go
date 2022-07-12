// Delete an application key returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewKeyManagementApi(apiClient)
	r, err := api.DeleteApplicationKey(ctx, "app_key_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.DeleteApplicationKey`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
