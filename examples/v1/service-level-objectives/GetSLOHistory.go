// Get an SLO's history returns "OK" response

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

	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v1.GetSLOHistory", true)
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewServiceLevelObjectivesApi(apiClient)
	resp, r, err := api.GetSLOHistory(ctx, SloData0ID, time.Now().AddDate(0, 0, -1).Unix(), time.Now().Unix(), *datadog.NewGetSLOHistoryOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.GetSLOHistory`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectivesApi.GetSLOHistory`:\n%s\n", responseContent)
}
