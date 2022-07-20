// Create an SLO correction returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	// there is a valid "slo" in the system
	SloData0ID := os.Getenv("SLO_DATA_0_ID")

	body := datadog.SLOCorrectionCreateRequest{
		Data: &datadog.SLOCorrectionCreateData{
			Attributes: &datadog.SLOCorrectionCreateRequestAttributes{
				Category:    datadog.SLOCORRECTIONCATEGORY_SCHEDULED_MAINTENANCE,
				Description: common.PtrString("Example-Create_an_SLO_correction_returns_OK_response"),
				End:         common.PtrInt64(time.Now().Add(time.Hour * 1).Unix()),
				SloId:       SloData0ID,
				Start:       time.Now().Unix(),
				Timezone:    common.PtrString("UTC"),
			},
			Type: datadog.SLOCORRECTIONTYPE_CORRECTION,
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewServiceLevelObjectiveCorrectionsApi(apiClient)
	resp, r, err := api.CreateSLOCorrection(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectiveCorrectionsApi.CreateSLOCorrection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectiveCorrectionsApi.CreateSLOCorrection`:\n%s\n", responseContent)
}
