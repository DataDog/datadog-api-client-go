// Create a custom destination returns "OK" response

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
				Enabled:     datadog.PtrBool(true),
				ForwardTags: datadog.PtrBool(true),
				ForwardTagsRestrictionList: []string{
					"datacenter",
					"host",
				},
				ForwardTagsRestrictionListType: datadogV2.CUSTOMDESTINATIONATTRIBUTETAGSRESTRICTIONLISTTYPE_ALLOW_LIST.Ptr(),
				ForwarderDestination: datadogV2.CustomDestinationForwardDestination{
					CustomDestinationForwardDestinationHttp: &datadogV2.CustomDestinationForwardDestinationHttp{
						Auth: datadogV2.CustomDestinationHttpDestinationAuth{
							CustomDestinationHttpDestinationAuthBasic: &datadogV2.CustomDestinationHttpDestinationAuthBasic{
								Password: "datadog-custom-destination-password",
								Type:     datadogV2.CUSTOMDESTINATIONHTTPDESTINATIONAUTHBASICTYPE_BASIC,
								Username: "datadog-custom-destination-username",
							}},
						Endpoint: "https://example.com",
						Type:     datadogV2.CUSTOMDESTINATIONFORWARDDESTINATIONHTTPTYPE_HTTP,
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
