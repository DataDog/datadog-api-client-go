// Create a new Opsgenie account returns "CREATED" response

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
	body := datadogV2.OpsgenieAccountCreateRequest{
		Data: datadogV2.OpsgenieAccountCreateData{
			Attributes: datadogV2.OpsgenieAccountCreateAttributes{
				ApiKey: "00000000-0000-0000-0000-000000000000",
				Region: datadogV2.OPSGENIESERVICEREGIONTYPE_US,
			},
			Type: datadogV2.OPSGENIEACCOUNTTYPE_OPSGENIE_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOpsgenieIntegrationApi(apiClient)
	resp, r, err := api.CreateOpsgenieAccount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsgenieIntegrationApi.CreateOpsgenieAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OpsgenieIntegrationApi.CreateOpsgenieAccount`:\n%s\n", responseContent)
}
