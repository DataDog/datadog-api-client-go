// Update a pipeline returns "OK" response

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
	body := datadogV1.LogsPipeline{
		Filter: &datadogV1.LogsFilter{
			Query: datadog.PtrString("source:python"),
		},
		Name: "",
		Processors: []datadogV1.LogsProcessor{
			datadogV1.LogsProcessor{
				LogsGrokParser: &datadogV1.LogsGrokParser{
					Grok: datadogV1.LogsGrokParserRules{
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
					Type:      datadogV1.LOGSGROKPARSERTYPE_GROK_PARSER,
				}},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewLogsPipelinesApi(apiClient)
	resp, r, err := api.UpdateLogsPipeline(ctx, "pipeline_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsPipelinesApi.UpdateLogsPipeline`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsPipelinesApi.UpdateLogsPipeline`:\n%s\n", responseContent)
}
