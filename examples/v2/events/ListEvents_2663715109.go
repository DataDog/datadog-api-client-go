// Get a quick list of events returns "OK" response

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
	resp, r, err := api.ListEvents(ctx, *datadog.NewListEventsOptionalParameters().WithFilterQuery("datadog-agent").WithFilterFrom("2020-09-17T11:48:36+01:00").WithFilterTo("2020-09-17T12:48:36+01:00").WithPageLimit(5))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ListEvents`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `EventsApi.ListEvents`:\n%s\n", responseContent)
}
