// Create a GitHub cloud auth persona mapping returns "Created" response

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
	body := datadogV2.GitHubCloudAuthPersonaMappingCreateRequest{
		Data: datadogV2.GitHubCloudAuthPersonaMappingCreateData{
			Attributes: datadogV2.GitHubCloudAuthPersonaMappingCreateAttributes{
				AccountIdentifier: "test@example.com",
				ClaimMatchers: datadogV2.GitHubOIDCClaimPatterns{
					Actor:                datadog.PtrString("octocat"),
					ActorId:              datadog.PtrString("1234567"),
					Enterprise:           datadog.PtrString("test_enterprise"),
					EnterpriseId:         datadog.PtrString("42"),
					Environment:          datadog.PtrString("production"),
					EventName:            datadog.PtrString("push"),
					JobWorkflowRef:       datadog.PtrString("test_owner/test_repo/.github/workflows/jobs.yml@refs/heads/main"),
					Ref:                  datadog.PtrString("refs/heads/main"),
					RefType:              datadog.PtrString("branch"),
					Repository:           datadog.PtrString("test_owner/test_repo"),
					RepositoryId:         datadog.PtrString("123456789"),
					RepositoryOwner:      datadog.PtrString("test_owner"),
					RepositoryOwnerId:    datadog.PtrString("987654321"),
					RepositoryVisibility: datadog.PtrString("public"),
					RunnerEnvironment:    datadog.PtrString("github-hosted"),
					Sub:                  "repo:test_owner/test_repo:(ref:refs/heads/main|pull_request)",
					Workflow:             datadog.PtrString("CI"),
					WorkflowRef:          datadog.PtrString("test_owner/test_repo/.github/workflows/ci.yml@refs/heads/main"),
				},
			},
			Type: datadogV2.GITHUBCLOUDAUTHPERSONAMAPPINGTYPE_GITHUB_OIDC_AUTH_CONFIG,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateGitHubCloudAuthPersonaMapping", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudAuthenticationApi(apiClient)
	resp, r, err := api.CreateGitHubCloudAuthPersonaMapping(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAuthenticationApi.CreateGitHubCloudAuthPersonaMapping`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudAuthenticationApi.CreateGitHubCloudAuthPersonaMapping`:\n%s\n", responseContent)
}
