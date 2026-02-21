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
	body := datadogV2.FormVersionRequest{
		Data: datadogV2.FormVersionDataRequest{
			Attributes: datadogV2.FormVersionAttributes{
				DataDefinition: new(interface{}),
				State:          datadogV2.FORMVERSIONSTATE_DRAFT.Ptr(),
				UiDefinition:   new(interface{}),
				UpsertParams: datadogV2.FormVersionUpsertParams{
					Etag:        datadog.PtrString("b51f08b698d88d8027a935d9db649774949f5fb41a0c559bfee6a9a13225c72d"),
					MatchPolicy: datadog.PtrString("none"),
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
	resp, r, err := api.UpsertAndPublishFormVersion(ctx, uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormsApi.UpsertAndPublishFormVersion`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FormsApi.UpsertAndPublishFormVersion`:\n%s\n", responseContent)
}
