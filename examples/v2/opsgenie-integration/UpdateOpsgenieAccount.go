// Update an Opsgenie account returns "OK" response

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
	body := datadogV2.OpsgenieAccountUpdateRequest{
		Data: datadogV2.OpsgenieAccountUpdateData{
			Attributes: datadogV2.OpsgenieAccountUpdateAttributes{
				ApiKey: datadog.PtrString("00000000-0000-0000-0000-000000000000"),
				Region: datadogV2.OPSGENIESERVICEREGIONTYPE_US.Ptr(),
			},
			Id:   "596da4af-0563-4097-90ff-07230c3f9db3",
			Type: datadogV2.OPSGENIEACCOUNTTYPE_OPSGENIE_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOpsgenieIntegrationApi(apiClient)
	resp, r, err := api.UpdateOpsgenieAccount(ctx, "account_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsgenieIntegrationApi.UpdateOpsgenieAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OpsgenieIntegrationApi.UpdateOpsgenieAccount`:\n%s\n", responseContent)
}
