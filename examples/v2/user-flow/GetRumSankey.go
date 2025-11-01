// Get rum sankey returns "Successful response with Sankey diagram data" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.SankeyRequest{
		Data: &datadogV2.SankeyRequestData{
			Attributes: &datadogV2.SankeyRequestDataAttributes{
				DataSource: datadog.PtrString(""),
				Definition: &datadogV2.SankeyRequestDataAttributesDefinition{
					EntriesPerStep: datadog.PtrInt64(10),
					NumberOfSteps:  datadog.PtrInt64(5),
					Source:         datadog.PtrString("@view.name"),
					Target:         datadog.PtrString("@view.name"),
				},
				EnforcedExecutionType: datadog.PtrString(""),
				RequestId:             datadog.PtrString(""),
				Sampling: &datadogV2.SankeyRequestDataAttributesSampling{
					Enabled: datadog.PtrBool(true),
				},
				Search: &datadogV2.SankeyRequestDataAttributesSearch{
					AudienceFilters: &datadogV2.SankeyRequestDataAttributesSearchAudienceFilters{},
					Query:           datadog.PtrString("@type:view @application.id:*"),
					SubqueryId:      datadog.PtrString(""),
				},
				Time: &datadogV2.SankeyRequestDataAttributesTime{
					From: datadog.PtrInt64(1756425600000),
					To:   datadog.PtrInt64(1756857600000),
				},
			},
			Id:   datadog.PtrString("sankey_request"),
			Type: datadogV2.SANKEYREQUESTDATATYPE_SANKEY_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetRumSankey", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewUserFlowApi(apiClient)
	resp, r, err := api.GetRumSankey(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserFlowApi.GetRumSankey`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UserFlowApi.GetRumSankey`:\n%s\n", responseContent)
}
