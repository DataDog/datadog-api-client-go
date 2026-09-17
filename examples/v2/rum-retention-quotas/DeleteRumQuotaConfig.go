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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMRetentionQuotasApi(apiClient)
	r, err := api.DeleteRumQuotaConfig(ctx, datadogV2.RUMRETENTIONQUOTASCOPETYPE_APPLICATION, "cd73a516-a481-4af5-8352-9b577465c77b")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMRetentionQuotasApi.DeleteRumQuotaConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
