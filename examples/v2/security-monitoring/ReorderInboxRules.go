// Reorder the list of inbox rules in the pipeline returns "The list of inbox rules" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.ReorderInboxRulesParameters{
		Data: []datadogV2.ReorderInboxRulesParametersData{
			{
				Id:   uuid.MustParse("123e4567-e89b-12d3-a456-426655440000"),
				Type: datadogV2.INBOXRULESTYPE_INBOX_RULES,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.ReorderInboxRules(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.ReorderInboxRules`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.ReorderInboxRules`:\n%s\n", responseContent)
}
