// Patch a deployment event by version returns "Accepted" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.DORADeploymentPatchByVersionRequest{
		Data: datadogV2.DORADeploymentPatchByVersionRequestData{
			Attributes: datadogV2.DORADeploymentPatchByVersionRequestAttributes{
				ChangeFailure: true,
				Env:           "production",
				Remediation: &datadogV2.DORADeploymentPatchByVersionRemediation{
					DORADeploymentPatchByVersionRemediationByVersion: &datadogV2.DORADeploymentPatchByVersionRemediationByVersion{
						Type:    datadogV2.DORADEPLOYMENTPATCHREMEDIATIONTYPE_ROLLBACK,
						Version: "v1.2.2",
					}},
				Service: "my-service",
				Version: "v1.2.3",
			},
			Type: datadogV2.DORADEPLOYMENTPATCHREQUESTDATATYPE_DORA_DEPLOYMENT_PATCH_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.PatchDORADeploymentByVersion", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDORAMetricsApi(apiClient)
	r, err := api.PatchDORADeploymentByVersion(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DORAMetricsApi.PatchDORADeploymentByVersion`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
