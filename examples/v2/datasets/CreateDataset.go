// Create a dataset returns "OK" response

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
	body := datadogV2.DatasetCreateRequest{
Data: datadogV2.DatasetRequest{
Attributes: datadogV2.DatasetAttributesRequest{
Name: "Security Audit Dataset",
Principals: []string{
"role:94172442-be03-11e9-a77a-3b7612558ac1",
},
ProductFilters: []datadogV2.FiltersPerProduct{
{
Filters: []string{
"@application.id:ABCD",
},
Product: "metrics",
},
},
},
Type: datadogV2.DATASETTYPE_DATASET,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateDataset", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDatasetsApi(apiClient)
	resp, r, err := api.CreateDataset(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasetsApi.CreateDataset`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DatasetsApi.CreateDataset`:\n%s\n", responseContent)
}