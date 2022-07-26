// Get a list of events returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ListEvents", true)
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewEventsApi(apiClient)
	resp, r, err := api.ListEvents(ctx, *datadog.NewListEventsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ListEvents`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `EventsApi.ListEvents`:\n%s\n", responseContent)
}
