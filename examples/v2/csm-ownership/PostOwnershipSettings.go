// Update ownership settings for the org returns "OK" response

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
	body := datadogV2.OwnershipSettingsRequest{
		Data: datadogV2.OwnershipSettingsRequestData{
			Attributes: datadogV2.OwnershipSettingsRequestAttributes{
				AutoTag:         true,
				ConfidenceLevel: datadogV2.OWNERSHIPCONFIDENCELEVEL_HIGH,
			},
			Type: datadogV2.OWNERSHIPSETTINGSTYPE_OWNERSHIP_SETTINGS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.PostOwnershipSettings", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCSMOwnershipApi(apiClient)
	resp, r, err := api.PostOwnershipSettings(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CSMOwnershipApi.PostOwnershipSettings`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CSMOwnershipApi.PostOwnershipSettings`:\n%s\n", responseContent)
}
