// List SKUs returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ListProductCatalogSKUs", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewProductCatalogApi(apiClient)
	resp, r, err := api.ListProductCatalogSKUs(ctx, datadogV2.PRODUCTCATALOGSKUSAPIVERSION_V1, *datadogV2.NewListProductCatalogSKUsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogApi.ListProductCatalogSKUs`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogApi.ListProductCatalogSKUs`:\n%s\n", responseContent)
}
