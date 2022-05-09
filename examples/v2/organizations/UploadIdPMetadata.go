// Upload IdP metadata returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsApi.UploadIdPMetadata(ctx, *datadog.NewUploadIdPMetadataOptionalParameters().WithIdpFile(func() *os.File {
		fp, _ := os.Open("fixtures/organizations/saml_configurations/valid_idp_metadata.xml")
		return fp
	}()))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UploadIdPMetadata`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
