// Create and publish a form returns "OK" response

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
	body := datadogV2.FormCreateRequest{
		Data: datadogV2.FormDataRequest{
			Attributes: datadogV2.FormDataAttributesRequest{
				DataDefinition: new(interface{}),
				Description:    "test description",
				Name:           "test form happy path",
				UiDefinition:   new(interface{}),
			},
			Id:   datadog.PtrUUID(uuid.MustParse("00000000-0000-0000-0000-000000000000")),
			Type: datadogV2.FORMTYPE_FORMS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateAndPublishForm", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFormsApi(apiClient)
	resp, r, err := api.CreateAndPublishForm(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormsApi.CreateAndPublishForm`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FormsApi.CreateAndPublishForm`:\n%s\n", responseContent)
}
