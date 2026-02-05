// Get seats users for a product code returns "OK" response

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
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSeatsApi(apiClient)
	resp, r, err := api.GetSeatsUsersV2(ctx, "product_code", *datadogV2.NewGetSeatsUsersV2OptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeatsApi.GetSeatsUsersV2`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SeatsApi.GetSeatsUsersV2`:\n%s\n", responseContent)
}
