// Search LLM Observability spans returns "OK" response

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
	body := datadogV2.LLMObsSearchSpansRequest{
		Data: datadogV2.LLMObsSearchSpansRequestData{
			Attributes: datadogV2.LLMObsSearchSpansRequestAttributes{
				Filter: &datadogV2.LLMObsSpanFilter{
					From:     datadog.PtrString("now-900s"),
					MlApp:    datadog.PtrString("my-llm-app"),
					Query:    datadog.PtrString("@session_id:abc123def456"),
					SpanId:   datadog.PtrString("abc123def456"),
					SpanKind: datadog.PtrString("llm"),
					SpanName: datadog.PtrString("llm_call"),
					To:       datadog.PtrString("now"),
					TraceId:  datadog.PtrString("trace-9a8b7c6d5e4f"),
				},
				Options: &datadogV2.LLMObsSpanSearchOptions{
					IncludeAttachments: datadog.PtrBool(true),
					TimeOffset:         datadog.PtrInt64(0),
				},
				Page: &datadogV2.LLMObsSpanPageQuery{
					Cursor: datadog.PtrString("eyJzdGFydCI6MTAwfQ=="),
					Limit:  datadog.PtrInt64(10),
				},
				Sort: datadog.PtrString("-start_ns"),
			},
			Type: datadogV2.LLMOBSSEARCHSPANSREQUESTTYPE_SPANS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.SearchLLMObsSpans", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.SearchLLMObsSpans(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.SearchLLMObsSpans`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.SearchLLMObsSpans`:\n%s\n", responseContent)
}
