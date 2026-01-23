// Update Custom Ruleset returns "Successfully updated" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.CustomRulesetRequest{
		Data: &datadogV2.CustomRulesetRequestData{
			Attributes: &datadogV2.CustomRulesetRequestDataAttributes{
				Rules: []datadogV2.CustomRule{
					{
						CreatedAt: time.Date(2026, 1, 9, 13, 0, 57, 473141, time.UTC),
						CreatedBy: "foobarbaz",
						LastRevision: datadogV2.CustomRuleRevision{
							Attributes: datadogV2.CustomRuleRevisionAttributes{
								Arguments: []datadogV2.Argument{
									{
										Description: "YXJndW1lbnQgZGVzY3JpcHRpb24=",
										Name:        "YXJndW1lbnRfbmFtZQ==",
									},
								},
								Category:         datadogV2.CUSTOMRULEREVISIONATTRIBUTESCATEGORY_SECURITY,
								Checksum:         "8a66c4e4e631099ad71be3c1ea3ea8fc2d57193e56db2c296e2dd8a508b26b99",
								Code:             "Y29uZHVjdG9yOgogICAgLSBkZXBsb3lfb25seTogdHJ1ZQ==",
								CreatedAt:        time.Date(2026, 1, 9, 13, 0, 57, 473141, time.UTC),
								CreatedBy:        "foobarbaz",
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
							Id:   "revision-123",
							Type: datadogV2.CUSTOMRULEREVISIONDATATYPE_CUSTOM_RULE_REVISION,
						},
						Name: "my-rule",
					},
				},
			},
			Type: datadogV2.CUSTOMRULESETDATATYPE_CUSTOM_RULESET.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateCustomRuleset", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStaticAnalysisApi(apiClient)
	resp, r, err := api.UpdateCustomRuleset(ctx, "ruleset_name", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticAnalysisApi.UpdateCustomRuleset`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `StaticAnalysisApi.UpdateCustomRuleset`:\n%s\n", responseContent)
}
