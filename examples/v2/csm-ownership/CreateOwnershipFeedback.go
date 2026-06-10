// Submit feedback on an ownership inference returns "Created" response

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
	body := datadogV2.OwnershipFeedbackRequest{
		Data: datadogV2.OwnershipFeedbackRequestData{
			Attributes: datadogV2.OwnershipFeedbackRequestAttributes{
				Action:               datadogV2.OWNERSHIPFEEDBACKACTION_CONFIRM,
				ActorHandle:          "user@example.com",
				ActorType:            "user",
				CorrectedOwnerHandle: *datadog.NewNullableString(datadog.PtrString("team-b")),
				CorrectedOwnerType:   *datadog.NewNullableString(datadog.PtrString("team")),
				InferenceChecksum:    "abc123",
				Reason:               *datadog.NewNullableString(datadog.PtrString("Confirmed by team lead.")),
			},
			Type: datadogV2.OWNERSHIPFEEDBACKTYPE_OWNERSHIP_FEEDBACK,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateOwnershipFeedback", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCSMOwnershipApi(apiClient)
	resp, r, err := api.CreateOwnershipFeedback(ctx, "res-1", datadogV2.OWNERSHIPOWNERTYPE_TEAM, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CSMOwnershipApi.CreateOwnershipFeedback`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CSMOwnershipApi.CreateOwnershipFeedback`:\n%s\n", responseContent)
}
