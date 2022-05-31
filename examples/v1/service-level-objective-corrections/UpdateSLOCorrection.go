// Update an SLO correction returns "OK" response

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
	// there is a valid "correction" for "slo"
	CorrectionDataID := os.Getenv("CORRECTION_DATA_ID")

	body := datadog.SLOCorrectionUpdateRequest{
		Data: &datadog.SLOCorrectionUpdateData{
			Attributes: &datadog.SLOCorrectionUpdateRequestAttributes{
				Category:    datadog.SLOCORRECTIONCATEGORY_DEPLOYMENT.Ptr(),
				Description: datadog.PtrString("Example-Update_an_SLO_correction_returns_OK_response"),
				End:         datadog.PtrInt64(time.Now().Add(time.Hour * 1).Unix()),
				Start:       datadog.PtrInt64(time.Now().Unix()),
				Timezone:    datadog.PtrString("UTC"),
			},
			Type: datadog.SLOCORRECTIONTYPE_CORRECTION.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceLevelObjectiveCorrectionsApi.UpdateSLOCorrection(ctx, CorrectionDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectiveCorrectionsApi.UpdateSLOCorrection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectiveCorrectionsApi.UpdateSLOCorrection`:\n%s\n", responseContent)
}
