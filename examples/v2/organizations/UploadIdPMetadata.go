// Upload IdP metadata returns "OK" response

package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrganizationsApi(apiClient)
	r, err := api.UploadIdPMetadata(ctx, *datadogV2.NewUploadIdPMetadataOptionalParameters().WithIdpFile(func() io.Reader {
		fp, _ := os.Open("fixtures/organizations/saml_configurations/valid_idp_metadata.xml")
		return fp
	}()))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UploadIdPMetadata`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
