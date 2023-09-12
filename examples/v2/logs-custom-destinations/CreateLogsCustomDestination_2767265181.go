// Create a custom destination with Elasticsearch forwarding returns "OK" response

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
	body := datadogV2.CustomDestinationCreatePayload{
		Data: datadogV2.CustomDestinationWithoutId{
			Attributes: datadogV2.CustomDestinationAttributes{
				Name:    "my-destination",
				Version: 0,
				ForwarderDestination: datadogV2.CustomDestinationForwarderDestination{
					ElasticsearchDestination: &datadogV2.ElasticsearchDestination{
						Type:          datadogV2.ELASTICSEARCHDESTINATIONTYPE_ELASTICSEARCH,
						Endpoint:      "https://example.org/my-intake",
						IndexName:     "name",
						IndexRotation: "yyyy-MM-dd",
						Auth: &datadogV2.HttpDestinationBasicAuth{
							Type:     datadogV2.HTTPDESTINATIONBASICAUTHTYPE_BASIC,
							Username: "username",
							Password: "password",
						},
					}},
			},
			Type: "custom_destination",
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
