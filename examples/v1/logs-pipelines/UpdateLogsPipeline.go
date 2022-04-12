// Update a pipeline returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.LogsPipeline{
		Filter: &datadog.LogsFilter{
			Query: datadog.PtrString("source:python"),
		},
		Name: "",
		Processors: []datadog.LogsProcessor{
			datadog.LogsProcessor{
				LogsGrokParser: &datadog.LogsGrokParser{
					Grok: datadog.LogsGrokParserRules{
						MatchRules: `rule_name_1 foo
rule_name_2 bar
`,
						SupportRules: datadog.PtrString(`rule_name_1 foo
rule_name_2 bar
`),
					},
					IsEnabled: datadog.PtrBool(false),
					Samples:   []string{},
					Source:    "message",
					Type:      datadog.LOGSGROKPARSERTYPE_GROK_PARSER,
				}},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsPipelinesApi.UpdateLogsPipeline(ctx, "pipeline_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsPipelinesApi.UpdateLogsPipeline`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsPipelinesApi.UpdateLogsPipeline`:\n%s\n", responseContent)
}
