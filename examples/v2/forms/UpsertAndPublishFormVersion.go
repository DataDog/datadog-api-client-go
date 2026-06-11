// Upsert and publish a form version returns "OK" response

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

	body := datadogV2.UpsertAndPublishFormVersionRequest{
		Data: datadogV2.UpsertAndPublishFormVersionData{
			Attributes: datadogV2.UpsertAndPublishFormVersionDataAttributes{
				DataDefinition: datadogV2.FormDataDefinition{
					Description: datadog.PtrString("Welcome to the Engineering Experience Survey."),
					Required:    []string{},
					Title:       datadog.PtrString("Developer Experience Survey"),
					Type:        datadogV2.FORMDATADEFINITIONTYPE_OBJECT.Ptr(),
				},
				UiDefinition: datadogV2.FormUiDefinition{
					UiOrder: []string{},
					UiTheme: &datadogV2.FormUiDefinitionUiTheme{
						PrimaryColor: datadogV2.FORMUIDEFINITIONUITHEMEPRIMARYCOLOR_GRAY.Ptr(),
					},
				},
				UpsertParams: datadogV2.UpsertAndPublishFormVersionUpsertParams{
					Etag: "b51f08b698d88d8027a935d9db649774949f5fb41a0c559bfee6a9a13225c72d",
				},
			},
			Type: datadogV2.FORMVERSIONTYPE_FORM_VERSIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpsertAndPublishFormVersion", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFormsApi(apiClient)
	resp, r, err := api.UpsertAndPublishFormVersion(ctx, FormDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormsApi.UpsertAndPublishFormVersion`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FormsApi.UpsertAndPublishFormVersion`:\n%s\n", responseContent)
}
