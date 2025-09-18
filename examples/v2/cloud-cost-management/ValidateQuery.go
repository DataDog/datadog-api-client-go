// Validate query returns "OK" response

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
	body := datadogV2.RulesValidateQueryRequest{
		Data: &datadogV2.RulesValidateQueryRequestData{
			Attributes: &datadogV2.RulesValidateQueryRequestDataAttributes{
				Query: "example:query AND test:true",
			},
			Type: datadogV2.RULESVALIDATEQUERYREQUESTDATATYPE_VALIDATE_QUERY,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	resp, r, err := api.ValidateQuery(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.ValidateQuery`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudCostManagementApi.ValidateQuery`:\n%s\n", responseContent)
}
