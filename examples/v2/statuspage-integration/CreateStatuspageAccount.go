// Create the Statuspage account returns "CREATED" response

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
	body := datadogV2.StatuspageAccountCreateRequest{
		Data: datadogV2.StatuspageAccountCreateData{
			Attributes: datadogV2.StatuspageAccountCreateAttributes{
				ApiKey: "00000000-0000-0000-0000-000000000000",
			},
			Type: datadogV2.STATUSPAGEACCOUNTTYPE_STATUSPAGE_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStatuspageIntegrationApi(apiClient)
	resp, r, err := api.CreateStatuspageAccount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatuspageIntegrationApi.CreateStatuspageAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `StatuspageIntegrationApi.CreateStatuspageAccount`:\n%s\n", responseContent)
}
