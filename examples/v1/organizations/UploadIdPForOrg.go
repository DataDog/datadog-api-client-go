// Upload IdP metadata returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewOrganizationsApi(apiClient)
	resp, r, err := api.UploadIdPForOrg(ctx, "abc123", func() io.Reader { fp, _ := os.Open("./idp_metadata.xml"); return fp }())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UploadIdPForOrg`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UploadIdPForOrg`:\n%s\n", responseContent)
}
