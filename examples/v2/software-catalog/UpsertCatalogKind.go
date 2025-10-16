// Create or update kinds returns "ACCEPTED" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.UpsertCatalogKindRequest{
KindObj: &datadogV2.KindObj{
Kind: "my-job",
}}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSoftwareCatalogApi(apiClient)
	resp, r, err := api.UpsertCatalogKind(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SoftwareCatalogApi.UpsertCatalogKind`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SoftwareCatalogApi.UpsertCatalogKind`:\n%s\n", responseContent)
}