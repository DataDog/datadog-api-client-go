// Create an HCP Terraform persona mapping returns "Created" response

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
	// there is a valid "service_account_user_with_role" in the system
	ServiceAccountUserWithRoleDataAttributesEmail := os.Getenv("SERVICE_ACCOUNT_USER_WITH_ROLE_DATA_ATTRIBUTES_EMAIL")

	body := datadogV2.TerraformCloudAuthPersonaMappingCreateRequest{
		Data: datadogV2.TerraformCloudAuthPersonaMappingCreateData{
			Type: datadogV2.TERRAFORMCLOUDAUTHPERSONAMAPPINGTYPE_TERRAFORM_CLOUD_AUTH_CONFIG,
			Attributes: datadogV2.TerraformCloudAuthPersonaMappingCreateAttributes{
				IssuerUrl:         datadogV2.TERRAFORMCLOUDAUTHISSUERURL_US,
				SubPattern:        "organization:datadog-api-spec:project:Default Project:workspace:sdk-examplecloudauthentication:run_phase:*",
				AccountIdentifier: ServiceAccountUserWithRoleDataAttributesEmail,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateTerraformCloudAuthPersonaMapping", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudAuthenticationApi(apiClient)
	resp, r, err := api.CreateTerraformCloudAuthPersonaMapping(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAuthenticationApi.CreateTerraformCloudAuthPersonaMapping`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudAuthenticationApi.CreateTerraformCloudAuthPersonaMapping`:\n%s\n", responseContent)
}
