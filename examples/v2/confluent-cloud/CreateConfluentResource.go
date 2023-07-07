// Add resource to Confluent account returns "OK" response

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
	// there is a valid "confluent_account" in the system
	ConfluentAccountDataID := os.Getenv("CONFLUENT_ACCOUNT_DATA_ID")

	body := datadogV2.ConfluentResourceRequest{
		Data: datadogV2.ConfluentResourceRequestData{
			Attributes: datadogV2.ConfluentResourceRequestAttributes{
				ResourceType: "kafka",
				Tags: []string{
					"myTag",
					"myTag2:myValue",
				},
				EnableCustomMetrics: datadog.PtrBool(false),
			},
			Id:   "exampleconfluentcloud",
			Type: datadogV2.CONFLUENTRESOURCETYPE_CONFLUENT_CLOUD_RESOURCES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewConfluentCloudApi(apiClient)
	resp, r, err := api.CreateConfluentResource(ctx, ConfluentAccountDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfluentCloudApi.CreateConfluentResource`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ConfluentCloudApi.CreateConfluentResource`:\n%s\n", responseContent)
}
