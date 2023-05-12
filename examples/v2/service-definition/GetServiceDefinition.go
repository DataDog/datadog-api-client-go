// Get a single service definition returns "OK" response

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
	api := datadogV2.NewServiceDefinitionApi(apiClient)
	resp, r, err := api.GetServiceDefinition(ctx, "service-definition-test", *datadogV2.NewGetServiceDefinitionOptionalParameters().WithSchemaVersion(datadogV2.SERVICEDEFINITIONSCHEMAVERSIONS_V2_1))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceDefinitionApi.GetServiceDefinition`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceDefinitionApi.GetServiceDefinition`:\n%s\n", responseContent)
}
