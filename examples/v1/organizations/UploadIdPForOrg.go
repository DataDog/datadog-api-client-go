// Upload IdP metadata returns "OK" response

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
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.OrganizationsApi(apiClient)
	resp, r, err := api.UploadIdPForOrg(ctx, "abc123", func() *os.File { fp, _ := os.Open("./idp_metadata.xml"); return fp }())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UploadIdPForOrg`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UploadIdPForOrg`:\n%s\n", responseContent)
}
