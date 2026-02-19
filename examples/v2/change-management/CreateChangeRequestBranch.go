// Create a change request branch returns "OK" response

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
	body := datadogV2.ChangeRequestBranchCreateRequest{
		Data: datadogV2.ChangeRequestBranchCreateData{
			Attributes: datadogV2.ChangeRequestBranchCreateAttributes{
				BranchName: "chm/CHM-1234",
				RepoId:     "DataDog/dd-source",
			},
			Type: datadogV2.CHANGEREQUESTBRANCHRESOURCETYPE_CHANGE_REQUEST_BRANCH,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateChangeRequestBranch", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewChangeManagementApi(apiClient)
	resp, r, err := api.CreateChangeRequestBranch(ctx, "change_request_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeManagementApi.CreateChangeRequestBranch`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ChangeManagementApi.CreateChangeRequestBranch`:\n%s\n", responseContent)
}
