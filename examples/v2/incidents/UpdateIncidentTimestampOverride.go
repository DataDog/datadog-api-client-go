// Update a timestamp override for an incident returns "OK" response

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
	body := datadogV2.IncidentTimestampOverridePatchRequest{
		Data: datadogV2.IncidentTimestampOverridePatchData{
			Attributes: datadogV2.IncidentTimestampOverridePatchAttributes{
				TimestampValue: time.Date(2024, 12, 29, 11, 0, 0, 0, time.UTC),
			},
			Type: datadogV2.INCIDENTSTIMESTAMPOVERRIDESTYPE_INCIDENTS_TIMESTAMP_OVERRIDES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateIncidentTimestampOverride", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.UpdateIncidentTimestampOverride(ctx, uuid.MustParse("9cecfde8-e35d-4387-8985-9b30dcb9cb1c"), uuid.MustParse("6f48a86f-9a39-4bcf-a76b-9a1b20188995"), body, *datadogV2.NewUpdateIncidentTimestampOverrideOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateIncidentTimestampOverride`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateIncidentTimestampOverride`:\n%s\n", responseContent)
}
