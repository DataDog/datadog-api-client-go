// Update a custom destination returns "OK" response

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
	// there is a valid "custom_destination" in the system
	CustomDestinationDataID := os.Getenv("CUSTOM_DESTINATION_DATA_ID")

	body := datadogV2.CustomDestinationUpdateRequest{
		Data: &datadogV2.CustomDestinationUpdateRequestDefinition{
			Attributes: &datadogV2.CustomDestinationUpdateRequestAttributes{
				Name:                           datadog.PtrString("Nginx logs (Updated)"),
				Query:                          datadog.PtrString("source:nginx"),
				Enabled:                        datadog.PtrBool(false),
				ForwardTags:                    datadog.PtrBool(false),
				ForwardTagsRestrictionListType: datadogV2.CUSTOMDESTINATIONATTRIBUTETAGSRESTRICTIONLISTTYPE_BLOCK_LIST.Ptr(),
			},
			Type: datadogV2.CUSTOMDESTINATIONTYPE_CUSTOM_DESTINATION,
			Id:   CustomDestinationDataID,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsCustomDestinationsApi(apiClient)
	resp, r, err := api.UpdateLogsCustomDestination(ctx, CustomDestinationDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsCustomDestinationsApi.UpdateLogsCustomDestination`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsCustomDestinationsApi.UpdateLogsCustomDestination`:\n%s\n", responseContent)
}
