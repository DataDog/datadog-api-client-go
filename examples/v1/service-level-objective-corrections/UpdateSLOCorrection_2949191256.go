// Update an SLO correction with slo_query returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	// there is a valid "correction_with_query" in the system
	CorrectionWithQueryDataID := os.Getenv("CORRECTION_WITH_QUERY_DATA_ID")

	body := datadogV1.SLOCorrectionUpdateRequest{
		Data: &datadogV1.SLOCorrectionUpdateData{
			Attributes: &datadogV1.SLOCorrectionUpdateRequestAttributes{
				Category:    datadogV1.SLOCORRECTIONCATEGORY_SCHEDULED_MAINTENANCE.Ptr(),
				Description: datadog.PtrString("Example-Service-Level-Objective-Correction"),
				End:         datadog.PtrInt64(time.Now().Add(time.Hour * 1).Unix()),
				SloQuery:    datadog.PtrString("env:staging service:checkout"),
				Start:       datadog.PtrInt64(time.Now().Unix()),
				Timezone:    datadog.PtrString("UTC"),
			},
			Type: datadogV1.SLOCORRECTIONTYPE_CORRECTION.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewServiceLevelObjectiveCorrectionsApi(apiClient)
	resp, r, err := api.UpdateSLOCorrection(ctx, CorrectionWithQueryDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectiveCorrectionsApi.UpdateSLOCorrection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectiveCorrectionsApi.UpdateSLOCorrection`:\n%s\n", responseContent)
}
