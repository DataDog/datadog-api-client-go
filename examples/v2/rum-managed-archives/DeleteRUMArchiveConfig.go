// Delete an archive configuration returns "No Content" response

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMManagedArchivesApi(apiClient)
	r, err := api.DeleteRUMArchiveConfig(ctx, uuid.MustParse("a33671aa-24fd-4dcd-ba4b-5bbdbafe7690"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMManagedArchivesApi.DeleteRUMArchiveConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
