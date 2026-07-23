// Update degradation template returns "OK" response

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
	body := datadogV2.PatchDegradationTemplateRequest{
		Data: &datadogV2.PatchDegradationTemplateRequestData{
			Attributes: &datadogV2.PatchDegradationTemplateRequestDataAttributes{
				ComponentsAffected: []datadogV2.PatchDegradationTemplateRequestDataAttributesComponentsAffectedItems{
					{
						Id:     "",
						Status: datadogV2.PATCHDEGRADATIONTEMPLATEREQUESTDATAATTRIBUTESCOMPONENTSAFFECTEDITEMSSTATUS_OPERATIONAL,
					},
				},
				Updates: []datadogV2.PatchDegradationTemplateRequestDataAttributesUpdatesItems{
					{
						Status: datadogV2.CREATEDEGRADATIONREQUESTDATAATTRIBUTESSTATUS_INVESTIGATING,
					},
				},
			},
			Id:   "",
			Type: datadogV2.PATCHDEGRADATIONTEMPLATEREQUESTDATATYPE_DEGRADATION_TEMPLATES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStatusPagesApi(apiClient)
	resp, r, err := api.UpdateDegradationTemplate(ctx, uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), body, *datadogV2.NewUpdateDegradationTemplateOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusPagesApi.UpdateDegradationTemplate`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `StatusPagesApi.UpdateDegradationTemplate`:\n%s\n", responseContent)
}
