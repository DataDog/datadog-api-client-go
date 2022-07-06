// Delete a global variable returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.SyntheticsApi(apiClient)
	r, err := api.DeleteGlobalVariable(ctx, "variable_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.DeleteGlobalVariable`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
