// Get rum funnel returns "Successful response with funnel analysis data" response

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
	body := datadogV2.FunnelRequest{
		Data: &datadogV2.FunnelRequestData{
			Attributes: &datadogV2.FunnelRequestDataAttributes{
				DataSource:            datadog.PtrString("rum"),
				EnforcedExecutionType: datadog.PtrString(""),
				RequestId:             datadog.PtrString(""),
				Search: &datadogV2.FunnelRequestDataAttributesSearch{
					CrossSessionFilter: datadog.PtrString(""),
					QueryString:        datadog.PtrString("@type:view"),
					Steps: []datadogV2.FunnelRequestDataAttributesSearchStepsItems{
						{
							Facet:      datadog.PtrString("@view.name"),
							StepFilter: datadog.PtrString(""),
							Value:      datadog.PtrString("/apm/home"),
						},
						{
							Facet:      datadog.PtrString("@view.name"),
							StepFilter: datadog.PtrString(""),
							Value:      datadog.PtrString("/apm/traces"),
						},
					},
					SubqueryId: datadog.PtrString(""),
				},
				Time: &datadogV2.FunnelRequestDataAttributesTime{
					From: datadog.PtrInt64(1756425600000),
					To:   datadog.PtrInt64(1756857600000),
				},
			},
			Id:   datadog.PtrString("funnel_request"),
			Type: datadogV2.FUNNELREQUESTDATATYPE_FUNNEL_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetRumFunnel", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFunnelApi(apiClient)
	resp, r, err := api.GetRumFunnel(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunnelApi.GetRumFunnel`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FunnelApi.GetRumFunnel`:\n%s\n", responseContent)
}
