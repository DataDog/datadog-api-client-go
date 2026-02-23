// Patch a deployment event returns "Accepted" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.DORADeploymentPatchRequest{
		Data: datadogV2.DORADeploymentPatchRequestData{
			Attributes: datadogV2.DORADeploymentPatchRequestAttributes{
				ChangeFailure: datadog.PtrBool(true),
				Remediation: &datadogV2.DORADeploymentPatchRemediation{
					Id:   datadog.PtrString("eG42zNIkVjM"),
					Type: datadogV2.DORADEPLOYMENTPATCHREMEDIATIONTYPE_ROLLBACK.Ptr(),
				},
			},
			Id:   "z_RwVLi7v4Y",
			Type: datadogV2.DORADEPLOYMENTPATCHREQUESTDATATYPE_DORA_DEPLOYMENT_PATCH_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDORAMetricsApi(apiClient)
	r, err := api.PatchDORADeployment(ctx, "deployment_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DORAMetricsApi.PatchDORADeployment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
