// Name App Version returns "No Content" response

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

	body := datadogV2.UpdateAppVersionNameRequest{
		Data: &datadogV2.UpdateAppVersionNameRequestData{
			Attributes: &datadogV2.UpdateAppVersionNameRequestDataAttributes{
				Name: "v1.2.0 - bug fix release",
			},
			Type: datadogV2.APPVERSIONNAMETYPE_VERSIONNAMES.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAppBuilderApi(apiClient)
	r, err := api.UpdateAppVersionName(ctx, AppDataID, "latest", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppBuilderApi.UpdateAppVersionName`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
