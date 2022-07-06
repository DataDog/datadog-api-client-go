// Delete an SLO correction returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.ServiceLevelObjectiveCorrectionsApi(apiClient)
	r, err := api.DeleteSLOCorrection(ctx, "slo_correction_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectiveCorrectionsApi.DeleteSLOCorrection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
