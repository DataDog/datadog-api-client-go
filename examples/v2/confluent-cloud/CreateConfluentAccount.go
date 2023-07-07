// Add Confluent account returns "OK" response

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
	body := datadogV2.ConfluentAccountCreateRequest{
		Data: datadogV2.ConfluentAccountCreateRequestData{
			Attributes: datadogV2.ConfluentAccountCreateRequestAttributes{
				ApiKey:    "TESTAPIKEY123",
				ApiSecret: "test-api-secret-123",
				Resources: []datadogV2.ConfluentAccountResourceAttributes{
					{
						EnableCustomMetrics: datadog.PtrBool(false),
						Id:                  datadog.PtrString("resource-id-123"),
						ResourceType:        "kafka",
						Tags: []string{
							"myTag",
							"myTag2:myValue",
						},
					},
				},
				Tags: []string{
					"myTag",
					"myTag2:myValue",
				},
			},
			Type: datadogV2.CONFLUENTACCOUNTTYPE_CONFLUENT_CLOUD_ACCOUNTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewConfluentCloudApi(apiClient)
	resp, r, err := api.CreateConfluentAccount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfluentCloudApi.CreateConfluentAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ConfluentCloudApi.CreateConfluentAccount`:\n%s\n", responseContent)
}
