// Publish a form version returns "OK" response

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
	// there is a valid "form" in the system
	FormDataID := uuid.MustParse(os.Getenv("FORM_DATA_ID"))

	body := datadogV2.PublishFormRequest{
		Data: datadogV2.PublishFormData{
			Attributes: datadogV2.PublishFormDataAttributes{
				Version: 1,
			},
			Type: datadogV2.FORMPUBLICATIONTYPE_FORM_PUBLICATIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.PublishForm", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFormsApi(apiClient)
	resp, r, err := api.PublishForm(ctx, FormDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormsApi.PublishForm`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FormsApi.PublishForm`:\n%s\n", responseContent)
}
