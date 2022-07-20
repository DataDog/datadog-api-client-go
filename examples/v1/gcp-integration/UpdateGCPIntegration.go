// Update a GCP integration returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.GCPAccount{
		AuthProviderX509CertUrl: common.PtrString("https://www.googleapis.com/oauth2/v1/certs"),
		AuthUri:                 common.PtrString("https://accounts.google.com/o/oauth2/auth"),
		ClientEmail:             common.PtrString("api-dev@datadog-sandbox.iam.gserviceaccount.com"),
		ClientId:                common.PtrString("123456712345671234567"),
		ClientX509CertUrl:       common.PtrString("https://www.googleapis.com/robot/v1/metadata/x509/$CLIENT_EMAIL"),
		Errors: []string{
			"*",
		},
		HostFilters:  common.PtrString("key:value,filter:example"),
		PrivateKey:   common.PtrString("private_key"),
		PrivateKeyId: common.PtrString("123456789abcdefghi123456789abcdefghijklm"),
		ProjectId:    common.PtrString("datadog-apitest"),
		TokenUri:     common.PtrString("https://accounts.google.com/o/oauth2/token"),
		Type:         common.PtrString("service_account"),
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewGCPIntegrationApi(apiClient)
	resp, r, err := api.UpdateGCPIntegration(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GCPIntegrationApi.UpdateGCPIntegration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `GCPIntegrationApi.UpdateGCPIntegration`:\n%s\n", responseContent)
}
