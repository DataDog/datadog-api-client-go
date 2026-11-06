// Create a recovery request returns "Accepted" response

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
	body := datadogV2.RumRecoveryCreateRequest{
		Data: datadogV2.RumRecoveryCreateData{
			Attributes: datadogV2.RumRecoveryCreateAttributes{
				RumRecoverySessionsCreateAttributes: &datadogV2.RumRecoverySessionsCreateAttributes{
					RequestType: datadogV2.RUMRECOVERYSESSIONSREQUESTTYPE_SESSIONS.Ptr(),
					Sessions: []datadogV2.RumRecoverySessionInput{
						{
							ApplicationId: uuid.MustParse("a33671aa-24fd-4dcd-ba4b-5bbdbafe7690"),
							SessionId:     uuid.MustParse("00000000-0000-4000-8000-000000000001"),
						},
					},
				}},
			Type: datadog.PtrString("rum_recovery_request"),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMRecoveryApi(apiClient)
	resp, r, err := api.CreateRUMRecoveryRequest(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMRecoveryApi.CreateRUMRecoveryRequest`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMRecoveryApi.CreateRUMRecoveryRequest`:\n%s\n", responseContent)
}
