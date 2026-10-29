// List source maps returns "OK" response

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
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ListSourcemaps", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMApi(apiClient)
	resp, r, err := api.ListSourcemaps(ctx, *datadogV2.NewListSourcemapsOptionalParameters().WithMapkind(datadogV2.SOURCEMAPMAPKIND_WASM).WithSymbolType(datadogV2.SOURCEMAPSYMBOLTYPE_SOURCEMAP).WithFilterService([]string{
		"my-wasm-app",
	}).WithFilterVersion([]string{
		"1.0.0",
	}))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMApi.ListSourcemaps`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMApi.ListSourcemaps`:\n%s\n", responseContent)
}
