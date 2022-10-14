// Delete Confluent account returns "OK" response

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
	api := datadogV2.NewConfluentCloudApi(apiClient)
	r, err := api.DeleteConfluentAccount(ctx, "account_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfluentCloudApi.DeleteConfluentAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
