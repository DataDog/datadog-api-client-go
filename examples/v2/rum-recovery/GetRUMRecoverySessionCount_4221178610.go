// Count recoverable sessions accepts exactly 180 days

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
	body := datadogV2.RumRecoveryCountRequest{
		Data: datadogV2.RumRecoveryCountRequestData{
			Type: datadog.PtrString("rum_recovery_count"),
			Attributes: datadogV2.RumRecoveryCountAttributes{
				Query: "service:rum18600-fixture",
				From:  1621077071000,
				To:    1636629071000,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMRecoveryApi(apiClient)
	resp, r, err := api.GetRUMRecoverySessionCount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMRecoveryApi.GetRUMRecoverySessionCount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMRecoveryApi.GetRUMRecoverySessionCount`:\n%s\n", responseContent)
}
