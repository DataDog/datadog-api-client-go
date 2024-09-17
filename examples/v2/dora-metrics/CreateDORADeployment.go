// Send a deployment event for DORA Metrics returns "OK" response

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
	body := datadogV2.DORADeploymentRequest{
		Data: datadogV2.DORADeploymentRequestData{
			Attributes: datadogV2.DORADeploymentRequestAttributes{
				FinishedAt: 1693491984000000000,
				Git: &datadogV2.DORAGitInfo{
					CommitSha:     "66adc9350f2cc9b250b69abddab733dd55e1a588",
					RepositoryUrl: "https://github.com/organization/example-repository",
				},
				Service:   "shopist",
				StartedAt: 1693491974000000000,
				Version:   datadog.PtrString("v1.12.07"),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateDORADeployment", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDORAMetricsApi(apiClient)
	resp, r, err := api.CreateDORADeployment(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DORAMetricsApi.CreateDORADeployment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DORAMetricsApi.CreateDORADeployment`:\n%s\n", responseContent)
}
