// Get a list of RUM events returns "OK" response

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
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewRUMApi(apiClient)
	resp, r, err := api.ListRUMEvents(ctx, *datadog.NewListRUMEventsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMApi.ListRUMEvents`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMApi.ListRUMEvents`:\n%s\n", responseContent)
}
