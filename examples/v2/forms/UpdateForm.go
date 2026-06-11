// Update a form returns "OK" response

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

	body := datadogV2.UpdateFormRequest{
		Data: datadogV2.UpdateFormData{
			Attributes: datadogV2.UpdateFormDataAttributes{
				FormUpdate: datadogV2.FormUpdateAttributes{
					DatastoreConfig: &datadogV2.FormDatastoreConfigAttributes{
						DatastoreId:                  uuid.MustParse("5108ea24-dd83-4696-9caa-f069f73d0fad"),
						PrimaryColumnName:            "id",
						PrimaryKeyGenerationStrategy: "none",
					},
					Description: datadog.PtrString("An updated description."),
					Name:        datadog.PtrString("Updated Form Name"),
				},
			},
			Id:   datadog.PtrUUID(FormDataID),
			Type: datadogV2.FORMTYPE_FORMS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateForm", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFormsApi(apiClient)
	resp, r, err := api.UpdateForm(ctx, FormDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormsApi.UpdateForm`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FormsApi.UpdateForm`:\n%s\n", responseContent)
}
