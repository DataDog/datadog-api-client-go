// Query the event stream returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewEventsApi(apiClient)
	resp, r, err := api.ListEvents(ctx, 9223372036854775807, 9223372036854775807, *datadog.NewListEventsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ListEvents`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `EventsApi.ListEvents`:\n%s\n", responseContent)
}
