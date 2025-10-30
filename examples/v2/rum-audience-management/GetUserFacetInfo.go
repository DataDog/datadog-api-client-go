// Get user facet info returns "Successful response with facet information" response

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
	body := datadogV2.FacetInfoRequest{
		Data: &datadogV2.FacetInfoRequestData{
			Attributes: &datadogV2.FacetInfoRequestDataAttributes{
				FacetId: "first_browser_name",
				Limit:   10,
				Search: &datadogV2.FacetInfoRequestDataAttributesSearch{
					Query: datadog.PtrString("user_org_id:5001 AND first_country_code:US"),
				},
				TermSearch: &datadogV2.FacetInfoRequestDataAttributesTermSearch{
					Value: datadog.PtrString("Chrome"),
				},
			},
			Id:   datadog.PtrString("facet_info_request"),
			Type: datadogV2.FACETINFOREQUESTDATATYPE_USERS_FACET_INFO_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetUserFacetInfo", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumAudienceManagementApi(apiClient)
	resp, r, err := api.GetUserFacetInfo(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumAudienceManagementApi.GetUserFacetInfo`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumAudienceManagementApi.GetUserFacetInfo`:\n%s\n", responseContent)
}
