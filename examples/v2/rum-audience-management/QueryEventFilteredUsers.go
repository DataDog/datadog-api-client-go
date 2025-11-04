// Query event filtered users returns "Successful response with filtered user data" response

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
	body := datadogV2.QueryEventFilteredUsersRequest{
		Data: &datadogV2.QueryEventFilteredUsersRequestData{
			Attributes: &datadogV2.QueryEventFilteredUsersRequestDataAttributes{
				EventQuery: &datadogV2.QueryEventFilteredUsersRequestDataAttributesEventQuery{
					Query: datadog.PtrString("@type:view AND @view.loading_time:>3000 AND @application.name:ecommerce-platform"),
					TimeFrame: &datadogV2.QueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame{
						End:   datadog.PtrInt64(1761309676),
						Start: datadog.PtrInt64(1760100076),
					},
				},
				IncludeRowCount: datadog.PtrBool(true),
				Limit:           datadog.PtrInt64(25),
				Query:           datadog.PtrString("user_org_id:5001 AND first_country_code:US AND first_browser_name:Chrome"),
				SelectColumns: []string{
					"user_id",
					"user_email",
					"first_country_code",
					"first_browser_name",
					"events_count",
					"session_count",
					"error_count",
					"avg_loading_time",
				},
			},
			Id:   datadog.PtrString("query_event_filtered_users_request"),
			Type: datadogV2.QUERYEVENTFILTEREDUSERSREQUESTDATATYPE_QUERY_EVENT_FILTERED_USERS_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.QueryEventFilteredUsers", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumAudienceManagementApi(apiClient)
	resp, r, err := api.QueryEventFilteredUsers(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumAudienceManagementApi.QueryEventFilteredUsers`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumAudienceManagementApi.QueryEventFilteredUsers`:\n%s\n", responseContent)
}
