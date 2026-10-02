// Conclude experiment returns "The experiment was concluded and the winning variant was rolled out to its linked feature
// flag allocation." response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.ExperimentsConcludeExperimentV2Request{
		Data: datadogV2.ExperimentsConcludeExperimentV2RequestData{
			Attributes: datadogV2.ExperimentsConcludeExperimentV2RequestDataAttributes{
				DecisionVariantKey: "treatment",
			},
			Type: datadogV2.EXPERIMENTSCONCLUDEEXPERIMENTV2REQUESTDATATYPE_CONCLUDE_EXPERIMENT_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExperimentsApi(apiClient)
	r, err := api.ConcludeExperiment(ctx, uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsApi.ConcludeExperiment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
