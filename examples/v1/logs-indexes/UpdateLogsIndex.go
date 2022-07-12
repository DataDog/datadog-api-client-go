// Update an index returns "OK" response

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
	body := datadog.LogsIndexUpdateRequest{
		ExclusionFilters: []datadog.LogsExclusion{
			{
				Filter: &datadog.LogsExclusionFilter{
					Query:      common.PtrString("*"),
					SampleRate: 1.0,
				},
				Name: "payment",
			},
		},
		Filter: datadog.LogsFilter{
			Query: common.PtrString("source:python"),
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewLogsIndexesApi(apiClient)
	resp, r, err := api.UpdateLogsIndex(ctx, "name", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsIndexesApi.UpdateLogsIndex`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsIndexesApi.UpdateLogsIndex`:\n%s\n", responseContent)
}
