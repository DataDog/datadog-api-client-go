// Get a list of events returns "OK" response with pagination

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
	configuration.SetUnstableOperationEnabled("v2.ListEvents", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewEventsApi(apiClient)
	resp, _, err := api.ListEventsWithPagination(ctx, *datadogV2.NewListEventsOptionalParameters().WithFilterFrom("now-15m").WithFilterTo("now").WithPageLimit(2))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ListEvents`: %v\n", err)
	}

	for item := range resp {
		responseContent, _ := json.MarshalIndent(item, "", "  ")
		fmt.Fprintf(os.Stdout, "%s\n", responseContent)
	}
}
