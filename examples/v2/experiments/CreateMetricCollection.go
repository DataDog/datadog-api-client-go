// Create metric collection returns "Created" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.ExperimentsCreateMetricCollectionV2Request{
		Data: datadogV2.ExperimentsCreateMetricCollectionV2RequestData{
			Type: datadogV2.EXPERIMENTSPATCHMETRICCOLLECTIONV2REQUESTDATATYPE_METRIC_COLLECTIONS,
			Attributes: datadogV2.ExperimentsCreateMetricCollectionV2RequestDataAttributes{
				Name: "ex-14bb9543f523edde",
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExperimentsApi(apiClient)
	resp, r, err := api.CreateMetricCollection(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsApi.CreateMetricCollection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsApi.CreateMetricCollection`:\n%s\n", responseContent)
}
