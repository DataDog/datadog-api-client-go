// Set up rules for organization returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.SetupRulesRequest{
		Data: datadogV2.SetupRulesRequestData{
			Attributes: datadogV2.SetupRulesRequestAttributes{
				DisabledDefaultRules: []string{
					"q8MQxk8TCqrHnWkx",
					"r9NRyl9UDrsIoXly",
				},
			},
			Type: datadogV2.SETUPRULESREQUESTDATATYPE_SETUP,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.SetupScorecardRules", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewServiceScorecardsApi(apiClient)
	r, err := api.SetupScorecardRules(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceScorecardsApi.SetupScorecardRules`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
