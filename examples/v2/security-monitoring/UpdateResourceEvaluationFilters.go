// Update resource filters returns "OK" response

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
	body := datadogV2.UpdateResourceEvaluationFiltersRequest{
Data: datadogV2.UpdateResourceEvaluationFiltersRequestData{
Attributes: datadogV2.ResourceFilterAttributes{
CloudProvider: map[string]map[string][]string{
"aws": map[string][]string{
"aws_account_id": []string{
"tag1:v1",
},
},
},
},
Id: datadog.PtrString("csm_resource_filter"),
Type: datadogV2.RESOURCEFILTERREQUESTTYPE_CSM_RESOURCE_FILTER,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.UpdateResourceEvaluationFilters(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.UpdateResourceEvaluationFilters`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.UpdateResourceEvaluationFilters`:\n%s\n", responseContent)
}