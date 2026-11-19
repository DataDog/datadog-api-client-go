// Stop bandit returns "The bandit is stopped and serves the chosen variant." response

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
	body := datadogV2.ExperimentsStopBanditV2Request{
		Data: datadogV2.ExperimentsStopBanditV2RequestData{
			Attributes: datadogV2.ExperimentsStopBanditV2RequestDataAttributes{
				ServedVariantKey: "green",
			},
			Type: datadogV2.EXPERIMENTSSTOPBANDITV2REQUESTDATATYPE_STOP_BANDIT_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.StopBandit", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExperimentsApi(apiClient)
	r, err := api.StopBandit(ctx, uuid.MustParse("4b1f3c2e-8d5a-4f6b-9c7e-1a2b3c4d5e6f"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsApi.StopBandit`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
