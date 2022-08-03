// Upload IdP metadata returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v2/datadog"
)

func main() {
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewOrganizationsApi(apiClient)
	r, err := api.UploadIdPMetadata(ctx, *datadog.NewUploadIdPMetadataOptionalParameters().WithIdpFile(func() *os.File {
		fp, _ := os.Open("fixtures/organizations/saml_configurations/valid_idp_metadata.xml")
		return fp
	}()))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UploadIdPMetadata`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
