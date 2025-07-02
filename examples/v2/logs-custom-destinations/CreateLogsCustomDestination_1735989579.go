// Create a Microsoft Sentinel custom destination returns "OK" response

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
	body := datadogV2.CustomDestinationCreateRequest{
		Data: &datadogV2.CustomDestinationCreateRequestDefinition{
			Attributes: datadogV2.CustomDestinationCreateRequestAttributes{
				Enabled:     datadog.PtrBool(false),
				ForwardTags: datadog.PtrBool(false),
				ForwardTagsRestrictionList: []string{
					"datacenter",
					"host",
				},
				ForwardTagsRestrictionListType: datadogV2.CUSTOMDESTINATIONATTRIBUTETAGSRESTRICTIONLISTTYPE_ALLOW_LIST.Ptr(),
				ForwarderDestination: datadogV2.CustomDestinationForwardDestination{
					CustomDestinationForwardDestinationMicrosoftSentinel: &datadogV2.CustomDestinationForwardDestinationMicrosoftSentinel{
						Type:                   datadogV2.CUSTOMDESTINATIONFORWARDDESTINATIONMICROSOFTSENTINELTYPE_MICROSOFT_SENTINEL,
						TenantId:               "f3c9a8a1-4c2e-4d2e-b911-9f3c28c3c8b2",
						ClientId:               "9a2f4d83-2b5e-429e-a35a-2b3c4182db71",
						DataCollectionEndpoint: "https://my-dce-5kyl.eastus-1.ingest.monitor.azure.com",
						DataCollectionRuleId:   "dcr-000a00a000a00000a000000aa000a0aa",
						StreamName:             "Custom-MyTable",
					}},
				Name:  "Nginx logs",
				Query: datadog.PtrString("source:nginx"),
			},
			Type: datadogV2.CUSTOMDESTINATIONTYPE_CUSTOM_DESTINATION,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsCustomDestinationsApi(apiClient)
	resp, r, err := api.CreateLogsCustomDestination(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsCustomDestinationsApi.CreateLogsCustomDestination`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsCustomDestinationsApi.CreateLogsCustomDestination`:\n%s\n", responseContent)
}
