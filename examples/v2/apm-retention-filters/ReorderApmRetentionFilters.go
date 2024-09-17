// Re-order retention filters returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.ReorderRetentionFiltersRequest{
		Data: []datadogV2.RetentionFilterWithoutAttributes{
			{
				Id:   "jdZrilSJQLqzb6Cu7aub9Q",
				Type: datadogV2.APMRETENTIONFILTERTYPE_apm_retention_filter,
			},
			{
				Id:   "7RBOb7dLSYWI01yc3pIH8w",
				Type: datadogV2.APMRETENTIONFILTERTYPE_apm_retention_filter,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAPMRetentionFiltersApi(apiClient)
	r, err := api.ReorderApmRetentionFilters(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APMRetentionFiltersApi.ReorderApmRetentionFilters`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
