// Create a widget returns "OK" response

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
	body := datadogV2.CreateOrUpdateWidgetRequest{
		Data: datadogV2.CreateOrUpdateWidgetRequestData{
			Attributes: datadogV2.CreateOrUpdateWidgetRequestAttributes{
				Definition: datadogV2.WidgetDefinition{
					Title: "My Widget",
					Type:  datadogV2.WIDGETTYPE_BAR_CHART,
				},
				Tags: *datadog.NewNullableList(&[]string{}),
			},
			Type: "widgets",
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewWidgetsApi(apiClient)
	resp, r, err := api.CreateWidget(ctx, datadogV2.WIDGETEXPERIENCETYPE_CCM_REPORTS, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WidgetsApi.CreateWidget`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `WidgetsApi.CreateWidget`:\n%s\n", responseContent)
}
