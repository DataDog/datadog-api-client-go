// Delete Multiple Apps returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "app" in the system
	AppDataID := uuid.MustParse(os.Getenv("APP_DATA_ID"))

	body := datadogV2.DeleteAppsRequest{
		Data: []datadogV2.DeleteAppsRequestDataItems{
			{
				Id:   AppDataID,
				Type: datadogV2.APPDEFINITIONTYPE_APPDEFINITIONS,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteApps", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAppBuilderApi(apiClient)
	resp, r, err := api.DeleteApps(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppBuilderApi.DeleteApps`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AppBuilderApi.DeleteApps`:\n%s\n", responseContent)
}
