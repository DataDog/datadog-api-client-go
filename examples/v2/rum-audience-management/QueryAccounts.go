// Query accounts returns "Successful response with account data" response

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
	body := datadogV2.QueryAccountRequest{
		Data: &datadogV2.QueryAccountRequestData{
			Attributes: &datadogV2.QueryAccountRequestDataAttributes{
				Limit: datadog.PtrInt64(20),
				Query: datadog.PtrString("plan_type:enterprise AND user_count:>100 AND subscription_status:active"),
				SelectColumns: []string{
					"account_id",
					"account_name",
					"user_count",
					"plan_type",
					"subscription_status",
					"created_at",
					"mrr",
					"industry",
				},
				Sort: &datadogV2.QueryAccountRequestDataAttributesSort{
					Field: datadog.PtrString("user_count"),
					Order: datadog.PtrString("DESC"),
				},
				WildcardSearchTerm: datadog.PtrString("tech"),
			},
			Id:   datadog.PtrString("query_account_request"),
			Type: datadogV2.QUERYACCOUNTREQUESTDATATYPE_QUERY_ACCOUNT_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.QueryAccounts", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumAudienceManagementApi(apiClient)
	resp, r, err := api.QueryAccounts(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumAudienceManagementApi.QueryAccounts`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumAudienceManagementApi.QueryAccounts`:\n%s\n", responseContent)
}
