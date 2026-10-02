// Update metric collection returns "OK" response

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
	// there is a valid "experiment_metric_collection" in the system
	ExperimentMetricCollectionDataID := uuid.MustParse(os.Getenv("EXPERIMENT_METRIC_COLLECTION_DATA_ID"))

	body := datadogV2.ExperimentsPatchMetricCollectionV2Request{
		Data: datadogV2.ExperimentsPatchMetricCollectionV2RequestData{
			Type: datadogV2.EXPERIMENTSPATCHMETRICCOLLECTIONV2REQUESTDATATYPE_METRIC_COLLECTIONS,
			Attributes: &datadogV2.ExperimentsPatchMetricCollectionV2RequestDataAttributes{
				Name: datadog.PtrString("ex-14bb9543f523edde updated"),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExperimentsApi(apiClient)
	resp, r, err := api.UpdateMetricCollection(ctx, ExperimentMetricCollectionDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsApi.UpdateMetricCollection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsApi.UpdateMetricCollection`:\n%s\n", responseContent)
}
