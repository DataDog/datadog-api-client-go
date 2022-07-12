// Create an index returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.LogsIndex{
		DailyLimit: datadog.PtrInt64(300000000),
		ExclusionFilters: []datadog.LogsExclusion{
			{
				Filter: &datadog.LogsExclusionFilter{
					Query:      datadog.PtrString("*"),
					SampleRate: 1.0,
				},
				Name: "payment",
			},
		},
		Filter: datadog.LogsFilter{
			Query: datadog.PtrString("source:python"),
		},
		Name:             "main",
		NumRetentionDays: datadog.PtrInt64(15),
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewLogsIndexesApi(apiClient)
	resp, r, err := api.CreateLogsIndex(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsIndexesApi.CreateLogsIndex`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsIndexesApi.CreateLogsIndex`:\n%s\n", responseContent)
}
