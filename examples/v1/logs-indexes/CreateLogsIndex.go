// Create an index returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := datadogV1.LogsIndex{
		DailyLimit: datadog.PtrInt64(300000000),
		DailyLimitReset: &datadogV1.LogsDailyLimitReset{
			ResetTime:      datadog.PtrString("14:00"),
			ResetUtcOffset: datadog.PtrString("+02:00"),
		},
		DailyLimitWarningThresholdPercentage: datadog.PtrFloat64(70),
		ExclusionFilters: []datadogV1.LogsExclusion{
			{
				Filter: &datadogV1.LogsExclusionFilter{
					Query:      datadog.PtrString("*"),
					SampleRate: 1.0,
				},
				Name: "payment",
			},
		},
		Filter: datadogV1.LogsFilter{
			Query: datadog.PtrString("source:python"),
		},
		Name:             "main",
		NumRetentionDays: datadog.PtrInt64(15),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewLogsIndexesApi(apiClient)
	resp, r, err := api.CreateLogsIndex(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsIndexesApi.CreateLogsIndex`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsIndexesApi.CreateLogsIndex`:\n%s\n", responseContent)
}
