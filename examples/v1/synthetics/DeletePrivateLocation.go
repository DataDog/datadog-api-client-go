// Delete a private location returns "OK" response

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
	api := datadog.NewSyntheticsApi(apiClient)
	r, err := api.DeletePrivateLocation(ctx, "location_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.DeletePrivateLocation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
