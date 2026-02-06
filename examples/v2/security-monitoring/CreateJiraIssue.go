// Create Jira issues for security findings returns "Accepted" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.JiraIssueRequest{
		Data: datadogV2.JiraIssueDataRequest{
			Attributes: datadogV2.JiraIssueDataAttributesRequest{
				AccountId:   "f7ccdf99-0e22-4378-bdf9-03fde5379fea",
				Fields:      *datadogV2.NewNullableJiraIssueCustomFields(nil),
				IssueType:   "story",
				IssuetypeId: "1235",
				Mode:        datadogV2.JIRAISSUEDATAATTRIBUTESREQUESTMODE_SINGLE.Ptr(),
				ProjectId:   "1234",
				ProjectKey:  "SEC",
			},
			Id: "ID",
			Meta: datadogV2.JiraIssueDataMeta{
				Findings: []datadogV2.JiraIssueFinding{
					{
						Description: "Description",
						Ids: []datadogV2.JiraIssueFindingId{
							{
								Discovered: 123213123,
								Id:         "afa-afa-hze",
								Resource:   "Resource",
								Tags:       "akjasd:asdsad",
							},
						},
						Impacted:    datadog.PtrInt64(1),
						References:  "",
						Remediation: "Remediation",
						Severity:    datadogV2.FINDINGSTATUS_CRITICAL,
						Title:       "Title",
						Type:        "ciem",
					},
				},
				Vulnerabilities: []datadogV2.JiraIssueFinding{
					{
						Description: "Description",
						Ids: []datadogV2.JiraIssueFindingId{
							{
								Discovered: 123213123,
								Id:         "afa-afa-hze",
								Resource:   "Resource",
								Tags:       "akjasd:asdsad",
							},
						},
						Impacted:    datadog.PtrInt64(1),
						References:  "",
						Remediation: "Remediation",
						Severity:    datadogV2.FINDINGSTATUS_CRITICAL,
						Title:       "Title",
						Type:        "ciem",
					},
				},
			},
			Type: datadogV2.JIRAISSUETYPE_JIRA_ISSUE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateJiraIssue", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	r, err := api.CreateJiraIssue(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateJiraIssue`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
