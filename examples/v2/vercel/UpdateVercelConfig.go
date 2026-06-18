// Update Vercel configuration returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.VercelConfigAttributes{
		ApiKey: datadogV2.VercelApiKey{
			Id:    "00000000-0000-0000-0000-000000000001",
			Value: "",
		},
		LogsConfig: datadogV2.VercelLogsConfig{
			Enabled: true,
			Environments: []datadogV2.VercelEnvironment{
				datadogV2.VERCELENVIRONMENT_PRODUCTION,
			},
			LogSources: []datadogV2.VercelLogSource{
				datadogV2.VERCELLOGSOURCE_LAMBDA,
			},
			SamplingPercentage: 100,
		},
		TraceConfig: datadogV2.VercelTraceConfig{
			Enabled:            true,
			IsDeprecatedOtel:   false,
			SamplingPercentage: 100,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewVercelApi(apiClient)
	r, err := api.UpdateVercelConfig(ctx, "configuration_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VercelApi.UpdateVercelConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
