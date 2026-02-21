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
	body := datadogV2.FormUpdateRequest{
		Data: datadogV2.FormUpdateDataRequest{
			Attributes: datadogV2.FormUpdateAttributes{
				FormUpdate: &datadogV2.FormUpdateAttributesFormUpdate{
					Description: datadog.PtrString("Updated description"),
					Name:        datadog.PtrString("New Form Name"),
				},
			},
			Id:   datadog.PtrUUID(uuid.MustParse("00000000-0000-0000-0000-000000000000")),
			Type: datadogV2.FORMTYPE_FORMS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateForm", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFormsApi(apiClient)
	resp, r, err := api.UpdateForm(ctx, uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormsApi.UpdateForm`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FormsApi.UpdateForm`:\n%s\n", responseContent)
}
