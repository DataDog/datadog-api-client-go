// Analyze code returns "OK" response

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
	body := datadogV2.AnalysisRequest{
		Data: datadogV2.AnalysisRequestData{
			Attributes: datadogV2.AnalysisRequestDataAttributes{
				Code:         "aW1wb3J0IHN5cw==",
				FileEncoding: "utf-8",
				Filename:     "test.py",
				Language:     "python",
				Rules: []datadogV2.AnalysisRequestRule{
					{
						Category:        "BEST_PRACTICES",
						Checksum:        "abc123def456",
						Code:            "ZnVuY3Rpb24gdmlzaXQobm9kZSkge30=",
						EntityChecked:   *datadog.NewNullableString(nil),
						Id:              "python-best-practices/no-exit",
						Language:        "python",
						Regex:           *datadog.NewNullableString(nil),
						Severity:        "WARNING",
						TreeSitterQuery: "KGNhbGwgbmFtZTogKGF0dHJpYnV0ZSkpQHZhbA==",
						Type:            "TREE_SITTER_QUERY",
					},
				},
			},
			Type: datadogV2.ANALYSISREQUESTDATATYPE_ANALYSIS_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateStaticAnalysisServerAnalysis", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateStaticAnalysisServerAnalysis(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateStaticAnalysisServerAnalysis`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateStaticAnalysisServerAnalysis`:\n%s\n", responseContent)
}
