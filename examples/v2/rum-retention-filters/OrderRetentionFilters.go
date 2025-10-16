// Order RUM retention filters returns "Ordered" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.RumRetentionFiltersOrderRequest{
Data: []datadogV2.RumRetentionFiltersOrderData{
{
Type: datadogV2.RUMRETENTIONFILTERTYPE_RETENTION_FILTERS,
Id: "325631eb-94c9-49c0-93f9-ab7e4fd24529",
},
{
Type: datadogV2.RUMRETENTIONFILTERTYPE_RETENTION_FILTERS,
Id: "42d89430-5b80-426e-a44b-ba3b417ece25",
},
{
Type: datadogV2.RUMRETENTIONFILTERTYPE_RETENTION_FILTERS,
Id: "bff0bc34-99e9-4c16-adce-f47e71948c23",
},
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumRetentionFiltersApi(apiClient)
	resp, r, err := api.OrderRetentionFilters(ctx, "1d4b9c34-7ac4-423a-91cf-9902d926e9b3", body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumRetentionFiltersApi.OrderRetentionFilters`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumRetentionFiltersApi.OrderRetentionFilters`:\n%s\n", responseContent)
}