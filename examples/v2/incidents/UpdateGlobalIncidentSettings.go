// Update global incident settings returns "OK" response

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
	body := datadogV2.GlobalIncidentSettingsRequest{
		Data: datadogV2.GlobalIncidentSettingsDataRequest{
			Attributes: &datadogV2.GlobalIncidentSettingsAttributesRequest{
				AnalyticsDashboardId: datadog.PtrString("abc-123-def"),
			},
			Type: datadogV2.GLOBALINCIDENTSETTINGSTYPE_INCIDENTS_GLOBAL_SETTINGS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateGlobalIncidentSettings", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.UpdateGlobalIncidentSettings(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateGlobalIncidentSettings`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateGlobalIncidentSettings`:\n%s\n", responseContent)
}
