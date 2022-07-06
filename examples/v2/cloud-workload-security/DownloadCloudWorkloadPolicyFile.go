// Get the latest Cloud Workload Security policy returns "OK" response

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.CloudWorkloadSecurityApi(apiClient)
	resp, r, err := api.DownloadCloudWorkloadPolicyFile(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudWorkloadSecurityApi.DownloadCloudWorkloadPolicyFile`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := ioutil.ReadAll(resp)
	fmt.Fprintf(os.Stdout, "Response from `CloudWorkloadSecurityApi.DownloadCloudWorkloadPolicyFile`:\n%s\n", responseContent)
}
