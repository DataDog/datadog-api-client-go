// Remove host tags returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v1/datadog"
)

func main() {
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewTagsApi(apiClient)
	r, err := api.DeleteHostTags(ctx, "host_name", *datadog.NewDeleteHostTagsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.DeleteHostTags`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
