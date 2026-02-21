// Create a timestamp override for an incident returns "Created" response

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
	body := datadogV2.IncidentTimestampOverrideCreateRequest{
		Data: datadogV2.IncidentTimestampOverrideCreateData{
			Attributes: datadogV2.IncidentTimestampOverrideCreateAttributes{
				TimestampType:  datadogV2.TIMESTAMPTYPE_CREATED,
				TimestampValue: time.Date(2024, 12, 29, 10, 0, 0, 0, time.UTC),
			},
			Type: datadogV2.INCIDENTSTIMESTAMPOVERRIDESTYPE_INCIDENTS_TIMESTAMP_OVERRIDES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentTimestampOverride", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.CreateIncidentTimestampOverride(ctx, uuid.MustParse("9cecfde8-e35d-4387-8985-9b30dcb9cb1c"), body, *datadogV2.NewCreateIncidentTimestampOverrideOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncidentTimestampOverride`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncidentTimestampOverride`:\n%s\n", responseContent)
}
