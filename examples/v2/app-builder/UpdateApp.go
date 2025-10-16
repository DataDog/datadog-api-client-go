// Update App returns "OK" response

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


	body := datadogV2.UpdateAppRequest{
Data: &datadogV2.UpdateAppRequestData{
Attributes: &datadogV2.UpdateAppRequestDataAttributes{
Name: datadog.PtrString("Updated Name"),
RootInstanceName: datadog.PtrString("grid0"),
},
Id: &AppDataID,
Type: datadogV2.APPDEFINITIONTYPE_APPDEFINITIONS,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAppBuilderApi(apiClient)
	resp, r, err := api.UpdateApp(ctx, AppDataID, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppBuilderApi.UpdateApp`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AppBuilderApi.UpdateApp`:\n%s\n", responseContent)
}