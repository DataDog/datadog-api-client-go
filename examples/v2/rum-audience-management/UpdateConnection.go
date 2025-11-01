// Update connection returns "Connection updated successfully" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.UpdateConnectionRequest{
		Data: &datadogV2.UpdateConnectionRequestData{
			Attributes: &datadogV2.UpdateConnectionRequestDataAttributes{
				FieldsToAdd: []datadogV2.CreateConnectionRequestDataAttributesFieldsItems{
					{
						Description: datadog.PtrString("Net Promoter Score from customer surveys"),
						DisplayName: datadog.PtrString("NPS Score"),
						Groups: []string{
							"Satisfaction",
							"Metrics",
						},
						Id:         "nps_score",
						SourceName: "net_promoter_score",
						Type:       "number",
					},
				},
				FieldsToDelete: []string{
					"old_revenue_field",
				},
				FieldsToUpdate: []datadogV2.UpdateConnectionRequestDataAttributesFieldsToUpdateItems{
					{
						FieldId:            "lifetime_value",
						UpdatedDisplayName: datadog.PtrString(`Customer Lifetime Value (` + "`" + `USD` + "`" + `)`),
						UpdatedGroups: []string{
							"Financial",
							"Metrics",
						},
					},
				},
			},
			Id:   "crm-integration",
			Type: datadogV2.UPDATECONNECTIONREQUESTDATATYPE_CONNECTION_ID,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateConnection", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumAudienceManagementApi(apiClient)
	r, err := api.UpdateConnection(ctx, "users", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumAudienceManagementApi.UpdateConnection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
