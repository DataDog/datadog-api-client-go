// Get a list of tests events returns "OK" response with pagination

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCIVisibilityTestsApi(apiClient)
	resp, _ := api.ListCIAppTestEventsWithPagination(ctx, *datadogV2.NewListCIAppTestEventsOptionalParameters().WithFilterFrom(time.Now().Add(time.Second * -30)).WithFilterTo(time.Now()).WithPageLimit(2))

	for paginationResult := range resp {
		if paginationResult.Error != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `CIVisibilityTestsApi.ListCIAppTestEvents`: %v\n", paginationResult.Error)
		}
		responseContent, _ := json.MarshalIndent(paginationResult.Item, "", "  ")
		fmt.Fprintf(os.Stdout, "%s\n", responseContent)
	}
}
