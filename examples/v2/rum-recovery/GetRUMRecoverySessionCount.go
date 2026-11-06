// Estimate recoverable sessions returns "OK" response

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
			Attributes: datadogV2.RumRecoveryCountAttributes{
				From:  1789905600000,
				Query: "service:web",
				To:    1789992000000,
			},
			Id:   datadog.PtrString("11111111-1111-4111-8111-111111111111"),
			Type: datadog.PtrString("rum_recovery_count"),
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
