// Update Confluent account returns "OK" response

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
	ConfluentAccountDataAttributesAPIKey := os.Getenv("CONFLUENT_ACCOUNT_DATA_ATTRIBUTES_API_KEY")
	ConfluentAccountDataID := os.Getenv("CONFLUENT_ACCOUNT_DATA_ID")

	body := datadogV2.ConfluentAccountUpdateRequest{
		Data: datadogV2.ConfluentAccountUpdateRequestData{
			Attributes: datadogV2.ConfluentAccountUpdateRequestAttributes{
				ApiKey:    ConfluentAccountDataAttributesAPIKey,
				ApiSecret: "update-secret",
				Tags: []string{
					"updated_tag:val",
				},
			},
			Type: datadogV2.CONFLUENTACCOUNTTYPE_CONFLUENT_CLOUD_ACCOUNTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewConfluentCloudApi(apiClient)
	resp, r, err := api.UpdateConfluentAccount(ctx, ConfluentAccountDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfluentCloudApi.UpdateConfluentAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ConfluentCloudApi.UpdateConfluentAccount`:\n%s\n", responseContent)
}
