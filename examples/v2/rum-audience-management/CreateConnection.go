// Create connection returns "Connection created successfully" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.CreateConnectionRequest{
		Data: &datadogV2.CreateConnectionRequestData{
			Attributes: &datadogV2.CreateConnectionRequestDataAttributes{
				Fields: []datadogV2.CreateConnectionRequestDataAttributesFieldsItems{
					{
						Description: datadog.PtrString(`Customer subscription tier from ` + "`" + `CRM` + "`"),
						DisplayName: datadog.PtrString("Customer Tier"),
						Id:          "customer_tier",
						SourceName:  "subscription_tier",
						Type:        "string",
					},
					{
						Description: datadog.PtrString(`Customer lifetime value in ` + "`" + `USD` + "`"),
						DisplayName: datadog.PtrString("Lifetime Value"),
						Id:          "lifetime_value",
						SourceName:  "ltv",
						Type:        "number",
					},
				},
				JoinAttribute: "user_email",
				JoinType:      "email",
				Type:          "ref_table",
			},
			Id:   datadog.PtrString("crm-integration"),
			Type: datadogV2.UPDATECONNECTIONREQUESTDATATYPE_CONNECTION_ID,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateConnection", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumAudienceManagementApi(apiClient)
	r, err := api.CreateConnection(ctx, "users", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumAudienceManagementApi.CreateConnection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
