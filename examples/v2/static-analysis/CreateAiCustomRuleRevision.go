// Create an AI custom rule revision returns "Successfully created" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.AiCustomRuleRevisionRequest{
		Data: &datadogV2.AiCustomRuleRevisionRequestData{
			Attributes: &datadogV2.AiCustomRuleRevisionRequestAttributes{
				Category:      datadogV2.CUSTOMRULEREVISIONATTRIBUTESCATEGORY_SECURITY,
				Content:       "Content",
				Cwe:           *datadog.NewNullableString(datadog.PtrString("79")),
				Description:   "Ruleset description",
				Directories:   []string{},
				ExecutionMode: datadogV2.AICUSTOMRULEREVISIONEXECUTIONMODE_AUTO,
				Globs: []string{
					"**/*.py",
				},
				IsPublished:      false,
				IsTesting:        false,
				Severity:         datadogV2.CUSTOMRULEREVISIONATTRIBUTESSEVERITY_ERROR,
				ShortDescription: "Ruleset short description",
				VersionId:        datadog.PtrInt64(1),
			},
			Id:   datadog.PtrString("revision-abc-123"),
			Type: datadogV2.AICUSTOMRULEREVISIONDATATYPE_AI_RULE_REVISION.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateAiCustomRuleRevision", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStaticAnalysisApi(apiClient)
	r, err := api.CreateAiCustomRuleRevision(ctx, "my-ai-ruleset", "my-ai-rule", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticAnalysisApi.CreateAiCustomRuleRevision`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
