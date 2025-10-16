// Get a RUM retention filter returns "OK" response

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumRetentionFiltersApi(apiClient)
	resp, r, err := api.GetRetentionFilter(ctx, "a33671aa-24fd-4dcd-ba4b-5bbdbafe7690", "4b95d361-f65d-4515-9824-c9aaeba5ac2a", )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumRetentionFiltersApi.GetRetentionFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumRetentionFiltersApi.GetRetentionFilter`:\n%s\n", responseContent)
}