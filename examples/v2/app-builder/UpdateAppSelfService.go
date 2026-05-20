// Update App Self-Service Status returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "app" in the system
	AppDataID := uuid.MustParse(os.Getenv("APP_DATA_ID"))

	body := datadogV2.UpdateAppSelfServiceRequest{
		Data: &datadogV2.UpdateAppSelfServiceRequestData{
			Attributes: &datadogV2.UpdateAppSelfServiceRequestDataAttributes{
				SelfService: true,
			},
			Type: datadogV2.APPSELFSERVICETYPE_SELFSERVICE.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAppBuilderApi(apiClient)
	r, err := api.UpdateAppSelfService(ctx, AppDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppBuilderApi.UpdateAppSelfService`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
