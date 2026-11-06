// Cancel a batch recovery request returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.RumRecoveryCancelRequest{
		Data: datadogV2.RumRecoveryCancelData{
			Attributes: datadogV2.RumRecoveryCancelAttributes{
				Status: datadogV2.RUMRECOVERYCANCELSTATUS_CANCELLED,
			},
			Id:   datadog.PtrString("11111111-1111-4111-8111-111111111111"),
			Type: datadog.PtrString("rum_recovery_request"),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMRecoveryApi(apiClient)
	resp, r, err := api.UpdateRUMRecoveryBatchRequest(ctx, uuid.MustParse("11111111-1111-4111-8111-111111111111"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMRecoveryApi.UpdateRUMRecoveryBatchRequest`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMRecoveryApi.UpdateRUMRecoveryBatchRequest`:\n%s\n", responseContent)
}
