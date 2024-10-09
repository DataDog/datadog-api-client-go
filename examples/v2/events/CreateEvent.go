// Post a change event returns "OK" response

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
	body := datadogV2.ChangeEventCreateRequest{
		Attributes: &datadogV2.ChangeEvent{
			Attributes: datadogV2.ChangeEventCustomAttributes{
				Author: &datadogV2.ChangeEventCustomAttributesAuthor{
					Name: "",
					Type: datadogV2.CHANGEEVENTCUSTOMATTRIBUTESAUTHORTYPE_USER,
				},
				ChangeMetadata: new(interface{}),
				ChangedResource: datadogV2.ChangeEventCustomAttributesChangedResource{
					Name: "fallback_payments_test",
					Type: datadogV2.CHANGEEVENTCUSTOMATTRIBUTESCHANGEDRESOURCETYPE_FEATURE_FLAG,
				},
				ImpactedResources: []datadogV2.ChangeEventCustomAttributesImpactedResourcesItems{
					{
						Name: "payments_api",
						Type: datadogV2.CHANGEEVENTCUSTOMATTRIBUTESIMPACTEDRESOURCESITEMSTYPE_SERVICE,
					},
				},
				NewValue:  new(interface{}),
				PrevValue: new(interface{}),
			},
			Category: datadogV2.CHANGEEVENTCATEGORY_CHANGE,
			Message:  datadog.PtrString("payment_processed feature flag has been enabled"),
			Tags: []string{
				"environment:test",
			},
			Title: "payment_processed feature flag updated",
		},
		Type: datadogV2.CHANGEEVENTCREATEREQUESTTYPE_EVENT.Ptr(),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewEventsApi(apiClient)
	resp, r, err := api.CreateEvent(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.CreateEvent`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `EventsApi.CreateEvent`:\n%s\n", responseContent)
}
