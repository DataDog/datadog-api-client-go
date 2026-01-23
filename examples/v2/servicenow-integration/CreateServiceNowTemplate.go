// Create ServiceNow template returns "Created" response

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
	body := datadogV2.ServiceNowTemplateCreateRequest{
		Data: datadogV2.ServiceNowTemplateCreateRequestData{
			Attributes: datadogV2.ServiceNowTemplateCreateRequestAttributes{
				AssignmentGroupId: datadog.PtrUUID(uuid.MustParse("65b3341b-0680-47f9-a6d4-134db45c603e")),
				BusinessServiceId: datadog.PtrUUID(uuid.MustParse("65b3341b-0680-47f9-a6d4-134db45c603e")),
				FieldsMapping: map[string]string{
					"category": "software",
					"priority": "1",
				},
				HandleName:          "incident-template",
				InstanceId:          uuid.MustParse("65b3341b-0680-47f9-a6d4-134db45c603e"),
				ServicenowTablename: "incident",
				UserId:              datadog.PtrUUID(uuid.MustParse("65b3341b-0680-47f9-a6d4-134db45c603e")),
			},
			Type: datadogV2.SERVICENOWTEMPLATETYPE_SERVICENOW_TEMPLATES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateServiceNowTemplate", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewServiceNowIntegrationApi(apiClient)
	resp, r, err := api.CreateServiceNowTemplate(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceNowIntegrationApi.CreateServiceNowTemplate`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceNowIntegrationApi.CreateServiceNowTemplate`:\n%s\n", responseContent)
}
