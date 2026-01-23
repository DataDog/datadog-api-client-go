// Create Custom Rule Revision returns "Successfully created" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.CustomRuleRevisionRequest{
		Data: &datadogV2.CustomRuleRevisionRequestData{
			Attributes: &datadogV2.CustomRuleRevisionInputAttributes{
				Arguments: []datadogV2.Argument{
					{
						Description: "YXJndW1lbnQgZGVzY3JpcHRpb24=",
						Name:        "YXJndW1lbnRfbmFtZQ==",
					},
				},
				Category:         datadogV2.CUSTOMRULEREVISIONATTRIBUTESCATEGORY_SECURITY,
				Code:             "Y29uZHVjdG9yOgogICAgLSBkZXBsb3lfb25seTogdHJ1ZQ==",
				CreationMessage:  "Initial revision",
				Cve:              *datadog.NewNullableString(datadog.PtrString("CVE-2024-1234")),
				Cwe:              *datadog.NewNullableString(datadog.PtrString("CWE-79")),
				Description:      "bG9uZyBkZXNjcmlwdGlvbg==",
				DocumentationUrl: *datadog.NewNullableString(datadog.PtrString("https://docs.example.com/rules/my-rule")),
				IsPublished:      false,
				IsTesting:        false,
				Language:         datadogV2.LANGUAGE_PYTHON,
				Severity:         datadogV2.CUSTOMRULEREVISIONATTRIBUTESSEVERITY_ERROR,
				ShortDescription: "c2hvcnQgZGVzY3JpcHRpb24=",
				ShouldUseAiFix:   false,
				Tags: []string{
					"security",
					"custom",
				},
				Tests: []datadogV2.CustomRuleRevisionTest{
					{
						AnnotationCount: 1,
						Code:            "Y29uZHVjdG9yOgogICAgLSBkZXBsb3lfb25seTogdHJ1ZQ==",
						Filename:        "test.yaml",
					},
				},
				TreeSitterQuery: "Y29uZHVjdG9yOgogICAgLSBkZXBsb3lfb25seTogdHJ1ZQ==",
			},
			Type: datadogV2.CUSTOMRULEREVISIONDATATYPE_CUSTOM_RULE_REVISION.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateCustomRuleRevision", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStaticAnalysisApi(apiClient)
	r, err := api.CreateCustomRuleRevision(ctx, "ruleset_name", "rule_name", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticAnalysisApi.CreateCustomRuleRevision`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
