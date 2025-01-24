// Update App returns "OK" response

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
	// there is a valid "app" in the system
	AppDataID := os.Getenv("APP_DATA_ID")

	body := datadogV2.UpdateAppRequest{
		Data: &datadogV2.UpdateAppRequestData{
			Attributes: &datadogV2.UpdateAppRequestDataAttributes{
				Name:             datadog.PtrString("Updated Name"),
				RootInstanceName: datadog.PtrString("grid0"),
			},
			Id:   datadog.PtrString(AppDataID),
			Type: datadogV2.UPDATEAPPREQUESTDATATYPE_APPDEFINITIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateApp", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAppsApi(apiClient)
	resp, r, err := api.UpdateApp(ctx, AppDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.UpdateApp`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AppsApi.UpdateApp`:\n%s\n", responseContent)
}
