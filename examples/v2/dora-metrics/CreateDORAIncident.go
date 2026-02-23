// Send an incident event (legacy) returns "OK" response

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
	body := datadogV2.DORAFailureRequest{
		Data: datadogV2.DORAFailureRequestData{
			Attributes: datadogV2.DORAFailureRequestAttributes{
				CustomTags: *datadog.NewNullableList(&[]string{
					"language:java",
					"department:engineering",
				}),
				Env:        datadog.PtrString("staging"),
				FinishedAt: datadog.PtrInt64(1693491984000000000),
				Git: &datadogV2.DORAGitInfo{
					CommitSha:     "66adc9350f2cc9b250b69abddab733dd55e1a588",
					RepositoryUrl: "https://github.com/organization/example-repository",
				},
				Name: datadog.PtrString("Webserver is down failing all requests."),
				Services: []string{
					"shopist",
				},
				Severity:  datadog.PtrString("High"),
				StartedAt: 1693491974000000000,
				Team:      datadog.PtrString("backend"),
				Version:   datadog.PtrString("v1.12.07"),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDORAMetricsApi(apiClient)
	resp, r, err := api.CreateDORAIncident(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DORAMetricsApi.CreateDORAIncident`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DORAMetricsApi.CreateDORAIncident`:\n%s\n", responseContent)
}
