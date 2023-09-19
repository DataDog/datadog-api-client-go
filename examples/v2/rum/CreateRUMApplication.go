// Create a new RUM application returns "OK" response

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
	body := datadogV2.RUMApplicationCreateRequest{
		Data: datadogV2.RUMApplicationCreate{
			Attributes: datadogV2.RUMApplicationCreateAttributes{
				Name: "test-rum-5c67ebb32077e1d9",
				Type: datadog.PtrString("ios"),
			},
			Type: datadogV2.RUMAPPLICATIONCREATETYPE_RUM_APPLICATION_CREATE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMApi(apiClient)
	resp, r, err := api.CreateRUMApplication(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMApi.CreateRUMApplication`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMApi.CreateRUMApplication`:\n%s\n", responseContent)
}
