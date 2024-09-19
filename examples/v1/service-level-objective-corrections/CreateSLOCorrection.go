// Create an SLO correction returns "OK" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "slo" in the system
	SloData0ID := os.Getenv("SLO_DATA_0_ID")


	body := datadogV1.SLOCorrectionCreateRequest{
Data: &datadogV1.SLOCorrectionCreateData{
Attributes: &datadogV1.SLOCorrectionCreateRequestAttributes{
Category: datadogV1.SLOCORRECTIONCATEGORY_SCHEDULED_MAINTENANCE,
Description: datadog.PtrString("Example-Service-Level-Objective-Correction"),
End: datadog.PtrInt64(time.Now().Add(time.Hour*1).Unix()),
SloId: SloData0ID,
Start: time.Now().Unix(),
Timezone: datadog.PtrString("UTC"),
},
Type: datadogV1.SLOCORRECTIONTYPE_CORRECTION,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewServiceLevelObjectiveCorrectionsApi(apiClient)
	resp, r, err := api.CreateSLOCorrection(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectiveCorrectionsApi.CreateSLOCorrection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectiveCorrectionsApi.CreateSLOCorrection`:\n%s\n", responseContent)
}