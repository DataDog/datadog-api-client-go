// Update ServiceNow template returns "OK" response

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
	body := datadogV2.ServiceNowTemplateUpdateRequest{
		Data: datadogV2.ServiceNowTemplateUpdateRequestData{
			Attributes: datadogV2.ServiceNowTemplateUpdateRequestAttributes{
				AssignmentGroupId: datadog.PtrUUID(uuid.MustParse("65b3341b-0680-47f9-a6d4-134db45c603e")),
				BusinessServiceId: datadog.PtrUUID(uuid.MustParse("65b3341b-0680-47f9-a6d4-134db45c603e")),
				FieldsMapping: map[string]string{
					"category": "hardware",
					"priority": "2",
				},
				HandleName:          "incident-template-updated",
				InstanceId:          uuid.MustParse("65b3341b-0680-47f9-a6d4-134db45c603e"),
				ServicenowTablename: "incident",
				UserId:              datadog.PtrUUID(uuid.MustParse("65b3341b-0680-47f9-a6d4-134db45c603e")),
			},
			Type: datadogV2.SERVICENOWTEMPLATETYPE_SERVICENOW_TEMPLATES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateServiceNowTemplate", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewServiceNowIntegrationApi(apiClient)
	resp, r, err := api.UpdateServiceNowTemplate(ctx, uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceNowIntegrationApi.UpdateServiceNowTemplate`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceNowIntegrationApi.UpdateServiceNowTemplate`:\n%s\n", responseContent)
}
