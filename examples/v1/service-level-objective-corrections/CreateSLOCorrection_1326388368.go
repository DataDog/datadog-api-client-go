// Create an SLO correction with rrule returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	// there is a valid "slo" in the system
	SloData0ID := os.Getenv("SLO_DATA_0_ID")

	body := datadog.SLOCorrectionCreateRequest{
		Data: &datadog.SLOCorrectionCreateData{
			Attributes: &datadog.SLOCorrectionCreateRequestAttributes{
				Category:    datadog.SLOCORRECTIONCATEGORY_SCHEDULED_MAINTENANCE,
				Description: datadog.PtrString("Example-Create_an_SLO_correction_with_rrule_returns_OK_response"),
				SloId:       SloData0ID,
				Start:       time.Now().Unix(),
				Duration:    datadog.PtrInt64(3600),
				Rrule:       datadog.PtrString("FREQ=DAILY;INTERVAL=10;COUNT=5"),
				Timezone:    datadog.PtrString("UTC"),
			},
			Type: datadog.SLOCORRECTIONTYPE_CORRECTION,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceLevelObjectiveCorrectionsApi.CreateSLOCorrection(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectiveCorrectionsApi.CreateSLOCorrection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectiveCorrectionsApi.CreateSLOCorrection`:\n%s\n", responseContent)
}
