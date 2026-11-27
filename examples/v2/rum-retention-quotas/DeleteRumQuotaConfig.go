// Delete a RUM retention quota configuration returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "rum_application" in the system
	RumApplicationDataID := os.Getenv("RUM_APPLICATION_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMRetentionQuotasApi(apiClient)
	r, err := api.DeleteRumQuotaConfig(ctx, datadogV2.RUMRETENTIONQUOTASCOPETYPE_APPLICATION, RumApplicationDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMRetentionQuotasApi.DeleteRumQuotaConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
