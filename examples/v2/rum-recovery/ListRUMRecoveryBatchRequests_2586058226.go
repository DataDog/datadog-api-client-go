// List batch history filters request creation time independently of session time

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMRecoveryApi(apiClient)
	resp, r, err := api.ListRUMRecoveryBatchRequests(ctx, datadogV2.RUMRECOVERYBATCHREQUESTTYPE_BATCH, *datadogV2.NewListRUMRecoveryBatchRequestsOptionalParameters().WithFrom(1).WithTo(2).WithQuery("service:rum18600-fixture AND @usr.id:(123 OR 456)"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMRecoveryApi.ListRUMRecoveryBatchRequests`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMRecoveryApi.ListRUMRecoveryBatchRequests`:\n%s\n", responseContent)
}
