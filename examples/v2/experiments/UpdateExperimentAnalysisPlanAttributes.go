// Update experiment analysis plan attributes returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "experiment" in the system
	ExperimentDataID := uuid.MustParse(os.Getenv("EXPERIMENT_DATA_ID"))

	body := datadogV2.ExperimentsAnalysisPlanWriteV2Request{
		Data: datadogV2.ExperimentsAnalysisPlanWriteV2RequestData{
			Type: datadogV2.EXPERIMENTSANALYSISPLANWRITEV2REQUESTDATATYPE_ANALYSIS_PLANS,
			Id:   datadog.PtrUUID(ExperimentDataID),
			Attributes: &datadogV2.ExperimentsAnalysisPlanWriteV2RequestDataAttributes{
				ConfidenceLevel: datadog.PtrFloat64(0.9),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExperimentsApi(apiClient)
	resp, r, err := api.UpdateExperimentAnalysisPlanAttributes(ctx, ExperimentDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsApi.UpdateExperimentAnalysisPlanAttributes`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsApi.UpdateExperimentAnalysisPlanAttributes`:\n%s\n", responseContent)
}
