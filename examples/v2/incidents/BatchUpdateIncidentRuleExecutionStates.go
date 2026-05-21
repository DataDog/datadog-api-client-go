// Batch update incident rule execution states returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.IncidentBatchUpdateRuleExecutionStatesRequest{
		Data: datadogV2.IncidentBatchUpdateRuleExecutionStatesData{
			Attributes: datadogV2.IncidentBatchUpdateRuleExecutionStatesDataAttributes{
				Rules: []datadogV2.IncidentRuleExecutionStateRule{
					{
						LastExecutedAt: *datadog.NewNullableTime(datadog.PtrTime(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC))),
						RuleUuid:       uuid.MustParse("00000000-0000-0000-0000-000000000000"),
					},
				},
			},
			Type: datadogV2.INCIDENTRULEEXECUTIONSTATETYPE_INCIDENT_RULE_EXECUTION_STATES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.BatchUpdateIncidentRuleExecutionStates", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.BatchUpdateIncidentRuleExecutionStates(ctx, "incident_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.BatchUpdateIncidentRuleExecutionStates`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.BatchUpdateIncidentRuleExecutionStates`:\n%s\n", responseContent)
}
