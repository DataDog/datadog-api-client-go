// Clone a form returns "OK" response

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
	body := datadogV2.CloneFormRequest{
		Data: datadogV2.CloneFormData{
			Attributes: &datadogV2.CloneFormDataAttributes{
				Name: datadog.PtrString("Copy of My Form"),
			},
			Type: datadogV2.FORMTYPE_FORMS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CloneForm", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFormsApi(apiClient)
	resp, r, err := api.CloneForm(ctx, uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormsApi.CloneForm`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FormsApi.CloneForm`:\n%s\n", responseContent)
}
