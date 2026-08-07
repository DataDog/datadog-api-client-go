// Update GitHub CI Visibility status returns "OK" response

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
	body := datadogV2.CIAppGitHubAccountUpdateRequest{
		Data: datadogV2.CIAppGitHubAccountUpdateRequestData{
			Attributes: datadogV2.CIAppGitHubAccountUpdateRequestAttributes{
				Account: "datadog",
				Enabled: datadog.PtrBool(true),
				Host:    datadog.PtrString("github.com"),
				Repository: &datadogV2.CIAppGitHubAccountUpdateRequestRepository{
					Enabled: true,
					Name:    "shopist",
				},
			},
			Type: datadogV2.CIAPPGITHUBACCOUNTTYPE_CI_GITHUB_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCIVisibilityGitHubAccountsApi(apiClient)
	resp, r, err := api.UpdateCIAppGitHubAccount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CIVisibilityGitHubAccountsApi.UpdateCIAppGitHubAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CIVisibilityGitHubAccountsApi.UpdateCIAppGitHubAccount`:\n%s\n", responseContent)
}
