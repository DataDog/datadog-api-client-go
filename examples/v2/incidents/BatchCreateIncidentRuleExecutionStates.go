// Batch create incident rule execution states returns "Created" response

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
	body := datadogV2.IncidentBatchCreateRuleExecutionStatesRequest{
		Data: datadogV2.IncidentBatchCreateRuleExecutionStatesData{
			Attributes: datadogV2.IncidentBatchCreateRuleExecutionStatesDataAttributes{
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
	configuration.SetUnstableOperationEnabled("v2.BatchCreateIncidentRuleExecutionStates", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.BatchCreateIncidentRuleExecutionStates(ctx, "incident_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.BatchCreateIncidentRuleExecutionStates`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.BatchCreateIncidentRuleExecutionStates`:\n%s\n", responseContent)
}
