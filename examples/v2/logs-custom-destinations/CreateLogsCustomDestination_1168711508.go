// Create a custom destination with HTTP request header forwarding returns "OK" response

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
	body := datadogV2.CustomDestinationCreateRequest{
		Data: datadogV2.CustomDestinationCreateData{
			Attributes: datadogV2.CustomDestinationAttributes{
				Name:    "my-destination",
				Version: 0,
				ForwarderDestination: datadogV2.CustomDestinationForwarderDestination{
					HttpDestination: &datadogV2.HttpDestination{
						Type:     datadogV2.HTTPDESTINATIONTYPE_HTTP,
						Endpoint: "https://example.org/my-intake",
						Auth: datadogV2.HttpDestinationAuth{
							HttpDestinationCustomHeaderAuth: &datadogV2.HttpDestinationCustomHeaderAuth{
								Type:        datadogV2.HTTPDESTINATIONCUSTOMHEADERAUTHTYPE_CUSTOM_HEADER,
								HeaderName:  "header",
								HeaderValue: "value",
							}},
					}},
			},
			Type: datadogV2.CUSTOMDESTINATIONTYPE_CUSTOM_DESTINATION,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateLogsCustomDestination", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsCustomDestinationsApi(apiClient)
	resp, r, err := api.CreateLogsCustomDestination(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsCustomDestinationsApi.CreateLogsCustomDestination`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsCustomDestinationsApi.CreateLogsCustomDestination`:\n%s\n", responseContent)
}
