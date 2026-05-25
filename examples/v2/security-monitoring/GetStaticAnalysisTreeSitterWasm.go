// Get tree-sitter WASM file returns "BLOB with the content of the WASM file" response

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetStaticAnalysisTreeSitterWasm", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.GetStaticAnalysisTreeSitterWasm(ctx, "tree-sitter-python.wasm")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.GetStaticAnalysisTreeSitterWasm`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := ioutil.ReadAll(resp)
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.GetStaticAnalysisTreeSitterWasm`:\n%s\n", responseContent)
}
