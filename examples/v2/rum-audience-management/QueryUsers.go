// Query users returns "Successful response with user data" response

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
	body := datadogV2.QueryUsersRequest{
		Data: &datadogV2.QueryUsersRequestData{
			Attributes: &datadogV2.QueryUsersRequestDataAttributes{
				Limit: datadog.PtrInt64(25),
				Query: datadog.PtrString("user_email:*@techcorp.com AND first_country_code:US AND first_browser_name:Chrome"),
				SelectColumns: []string{
					"user_id",
					"user_email",
					"user_name",
					"user_org_id",
					"first_country_code",
					"first_browser_name",
					"first_device_type",
					"last_seen",
				},
				Sort: &datadogV2.QueryUsersRequestDataAttributesSort{
					Field: datadog.PtrString("first_seen"),
					Order: datadog.PtrString("DESC"),
				},
				WildcardSearchTerm: datadog.PtrString("john"),
			},
			Id:   datadog.PtrString("query_users_request"),
			Type: datadogV2.QUERYUSERSREQUESTDATATYPE_QUERY_USERS_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.QueryUsers", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumAudienceManagementApi(apiClient)
	resp, r, err := api.QueryUsers(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumAudienceManagementApi.QueryUsers`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumAudienceManagementApi.QueryUsers`:\n%s\n", responseContent)
}
