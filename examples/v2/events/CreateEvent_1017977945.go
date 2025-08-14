// Post an event with metric_configuration resource type returns "OK" response

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
	body := datadogV2.EventCreateRequestPayload{
		Data: datadogV2.EventCreateRequest{
			Attributes: datadogV2.EventPayload{
				AggregationKey: datadog.PtrString("aggregation_key_123"),
				Attributes: datadogV2.EventPayloadAttributes{
					ChangeEventCustomAttributes: &datadogV2.ChangeEventCustomAttributes{
						Author: &datadogV2.ChangeEventCustomAttributesAuthor{
							Name: "example@datadog.com",
							Type: datadogV2.CHANGEEVENTCUSTOMATTRIBUTESAUTHORTYPE_USER,
						},
						ChangeMetadata: map[string]interface{}{
							"dd":            "{'team': 'datadog_team', 'user_email': 'datadog@datadog.com', 'user_id': 'datadog_user_id', 'user_name': 'datadog_username'}",
							"resource_link": "datadog.com/metric/config_test",
						},
						ChangedResource: datadogV2.ChangeEventCustomAttributesChangedResource{
							Name: "config_test",
							Type: datadogV2.CHANGEEVENTCUSTOMATTRIBUTESCHANGEDRESOURCETYPE_METRIC_CONFIGURATION,
						},
						ImpactedResources: []datadogV2.ChangeEventCustomAttributesImpactedResourcesItems{
							{
								Name: "system.cpu.usage",
								Type: datadogV2.CHANGEEVENTCUSTOMATTRIBUTESIMPACTEDRESOURCESITEMSTYPE_SERVICE,
							},
						},
						NewValue: map[string]interface{}{
							"aggregation": "avg",
							"tags":        "['env:prod', 'service:web']",
							"unit":        "percent",
						},
						PrevValue: map[string]interface{}{
							"aggregation": "sum",
							"tags":        "['env:prod']",
							"unit":        "percent",
						},
					}},
				Category:      datadogV2.EVENTCATEGORY_CHANGE,
				IntegrationId: datadogV2.EVENTPAYLOADINTEGRATIONID_CUSTOM_EVENTS.Ptr(),
				Message:       datadog.PtrString("metric configuration has been updated"),
				Tags: []string{
					"env:api_client_test",
				},
				Title: "metric configuration updated",
			},
			Type: datadogV2.EVENTCREATEREQUESTTYPE_EVENT,
		},
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
