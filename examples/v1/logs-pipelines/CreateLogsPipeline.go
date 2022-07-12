// Create a pipeline returns "OK" response

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
	body := datadog.LogsPipeline{
		Filter: &datadog.LogsFilter{
			Query: common.PtrString("source:python"),
		},
		Name: "",
		Processors: []datadog.LogsProcessor{
			datadog.LogsProcessor{
				LogsGrokParser: &datadog.LogsGrokParser{
					Grok: datadog.LogsGrokParserRules{
						MatchRules: `rule_name_1 foo
rule_name_2 bar
`,
						SupportRules: common.PtrString(`rule_name_1 foo
rule_name_2 bar
`),
					},
					IsEnabled: common.PtrBool(false),
					Samples:   []string{},
					Source:    "message",
					Type:      datadog.LOGSGROKPARSERTYPE_GROK_PARSER,
				}},
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewLogsPipelinesApi(apiClient)
	resp, r, err := api.CreateLogsPipeline(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsPipelinesApi.CreateLogsPipeline`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsPipelinesApi.CreateLogsPipeline`:\n%s\n", responseContent)
}
