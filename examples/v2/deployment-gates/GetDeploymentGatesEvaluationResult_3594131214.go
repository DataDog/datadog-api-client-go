// Get a deployment gates evaluation result returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "deployment_gates_evaluation" in the system
	DeploymentGatesEvaluationDataID := uuid.MustParse(os.Getenv("DEPLOYMENT_GATES_EVALUATION_DATA_ID"))

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetDeploymentGatesEvaluationResult", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDeploymentGatesApi(apiClient)
	resp, r, err := api.GetDeploymentGatesEvaluationResult(ctx, DeploymentGatesEvaluationDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentGatesApi.GetDeploymentGatesEvaluationResult`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DeploymentGatesApi.GetDeploymentGatesEvaluationResult`:\n%s\n", responseContent)
}
