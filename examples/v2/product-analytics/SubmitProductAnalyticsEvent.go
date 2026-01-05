// Send server-side events returns "Request accepted for processing (always 202 empty JSON)." response

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
	body := datadogV2.ProductAnalyticsServerSideEventItem{
		Account: &datadogV2.ProductAnalyticsServerSideEventItemAccount{
			Id: "account-67890",
		},
		Application: datadogV2.ProductAnalyticsServerSideEventItemApplication{
			Id: "123abcde-123a-123b-1234-123456789abc",
		},
		Event: datadogV2.ProductAnalyticsServerSideEventItemEvent{
			Name: "payment.processed",
		},
		Session: &datadogV2.ProductAnalyticsServerSideEventItemSession{
			Id: "session-abcdef",
		},
		Type: datadogV2.PRODUCTANALYTICSSERVERSIDEEVENTITEMTYPE_SERVER,
		Usr: &datadogV2.ProductAnalyticsServerSideEventItemUsr{
			Id: "user-12345",
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewProductAnalyticsApi(apiClient)
	resp, r, err := api.SubmitProductAnalyticsEvent(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAnalyticsApi.SubmitProductAnalyticsEvent`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ProductAnalyticsApi.SubmitProductAnalyticsEvent`:\n%s\n", responseContent)
}
