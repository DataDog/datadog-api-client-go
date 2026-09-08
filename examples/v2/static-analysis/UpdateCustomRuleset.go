// Update Custom Ruleset returns "Successfully updated" response

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
	body := datadogV2.CustomRulesetRequest{
		Data: &datadogV2.CustomRulesetRequestData{
			Attributes: datadogV2.CustomRulesetRequestDataAttributes{
				Description: datadog.PtrString("bG9uZyBkZXNjcmlwdGlvbg=="),
				Name:        "my-ruleset",
				Rules: []datadogV2.CustomRule{
					{
						Id: "my-rule",
						LastRevision: &datadogV2.CustomRuleRevisionInput{
							Arguments: []datadogV2.Argument{
								{
									Description: "YXJndW1lbnQgZGVzY3JpcHRpb24=",
									Name:        "YXJndW1lbnRfbmFtZQ==",
								},
							},
							Category:         datadogV2.CUSTOMRULEREVISIONATTRIBUTESCATEGORY_SECURITY.Ptr(),
							Code:             datadog.PtrString("Y29uZHVjdG9yOgogICAgLSBkZXBsb3lfb25seTogdHJ1ZQ=="),
							CreationMessage:  datadog.PtrString("Initial revision"),
							Cve:              *datadog.NewNullableString(datadog.PtrString("CVE-2024-1234")),
							Cwe:              *datadog.NewNullableString(datadog.PtrString("CWE-79")),
							Description:      datadog.PtrString("bG9uZyBkZXNjcmlwdGlvbg=="),
							DocumentationUrl: *datadog.NewNullableString(datadog.PtrString("https://docs.example.com/rules/my-rule")),
							IsPublished:      datadog.PtrBool(false),
							IsTesting:        datadog.PtrBool(false),
							Language:         datadogV2.LANGUAGE_PYTHON.Ptr(),
							Severity:         datadogV2.CUSTOMRULEREVISIONATTRIBUTESSEVERITY_ERROR.Ptr(),
							ShortDescription: datadog.PtrString("c2hvcnQgZGVzY3JpcHRpb24="),
							ShouldUseAiFix:   datadog.PtrBool(false),
							Tags: *datadog.NewNullableList(&[]string{
								"security",
								"custom",
							}),
							Tests: []datadogV2.CustomRuleRevisionTest{
								{
									AnnotationCount: 1,
									Code:            "Y29uZHVjdG9yOgogICAgLSBkZXBsb3lfb25seTogdHJ1ZQ==",
									Filename:        "test.yaml",
								},
							},
							TreeSitterQuery: datadog.PtrString("Y29uZHVjdG9yOgogICAgLSBkZXBsb3lfb25seTogdHJ1ZQ=="),
						},
						Name: "my-rule",
						Revisions: []datadogV2.CustomRuleRevisionInput{
							{
								Arguments: []datadogV2.Argument{
									{
										Description: "YXJndW1lbnQgZGVzY3JpcHRpb24=",
										Name:        "YXJndW1lbnRfbmFtZQ==",
									},
								},
								Category:         datadogV2.CUSTOMRULEREVISIONATTRIBUTESCATEGORY_SECURITY.Ptr(),
								Code:             datadog.PtrString("Y29uZHVjdG9yOgogICAgLSBkZXBsb3lfb25seTogdHJ1ZQ=="),
								CreationMessage:  datadog.PtrString("Initial revision"),
								Cve:              *datadog.NewNullableString(datadog.PtrString("CVE-2024-1234")),
								Cwe:              *datadog.NewNullableString(datadog.PtrString("CWE-79")),
								Description:      datadog.PtrString("bG9uZyBkZXNjcmlwdGlvbg=="),
								DocumentationUrl: *datadog.NewNullableString(datadog.PtrString("https://docs.example.com/rules/my-rule")),
								IsPublished:      datadog.PtrBool(false),
								IsTesting:        datadog.PtrBool(false),
								Language:         datadogV2.LANGUAGE_PYTHON.Ptr(),
								Severity:         datadogV2.CUSTOMRULEREVISIONATTRIBUTESSEVERITY_ERROR.Ptr(),
								ShortDescription: datadog.PtrString("c2hvcnQgZGVzY3JpcHRpb24="),
								ShouldUseAiFix:   datadog.PtrBool(false),
								Tags: *datadog.NewNullableList(&[]string{
									"security",
									"custom",
								}),
								Tests: []datadogV2.CustomRuleRevisionTest{
									{
										AnnotationCount: 1,
										Code:            "Y29uZHVjdG9yOgogICAgLSBkZXBsb3lfb25seTogdHJ1ZQ==",
										Filename:        "test.yaml",
									},
								},
								TreeSitterQuery: datadog.PtrString("Y29uZHVjdG9yOgogICAgLSBkZXBsb3lfb25seTogdHJ1ZQ=="),
							},
						},
					},
				},
				ShortDescription: datadog.PtrString("c2hvcnQgZGVzY3JpcHRpb24="),
			},
			Id:   "my-ruleset",
			Type: datadogV2.CUSTOMRULESETDATATYPE_CUSTOM_RULESET,
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
