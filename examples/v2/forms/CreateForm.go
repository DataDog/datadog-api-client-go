// Create a form returns "OK" response

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
	body := datadogV2.CreateFormRequest{
		Data: datadogV2.CreateFormData{
			Attributes: datadogV2.CreateFormDataAttributes{
				Anonymous:      datadog.PtrBool(false),
				DataDefinition: datadogV2.FormDataDefinition{},
				Description:    datadog.PtrString("A form to collect user feedback."),
				IdpSurvey:      datadog.PtrBool(false),
				Name:           "User Feedback Form",
				SingleResponse: datadog.PtrBool(false),
				UiDefinition:   datadogV2.FormUiDefinition{},
			},
			Type: datadogV2.FORMTYPE_FORMS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateForm", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFormsApi(apiClient)
	resp, r, err := api.CreateForm(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormsApi.CreateForm`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FormsApi.CreateForm`:\n%s\n", responseContent)
}
