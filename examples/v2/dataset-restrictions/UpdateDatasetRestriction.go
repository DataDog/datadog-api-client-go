// Update a dataset restriction returns "OK" response

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
	body := datadogV2.DatasetRestrictionUpdateRequest{
		Data: datadogV2.DatasetRestrictionUpdateRequestData{
			Attributes: datadogV2.DatasetRestrictionUpdateRequestAttributes{
				OwnershipMode:          datadogV2.DATASETRESTRICTIONOWNERSHIPMODE_TEAM_TAG_BASED.Ptr(),
				RestrictionMode:        datadogV2.DATASETRESTRICTIONRESTRICTIONMODE_DEFAULT_HIDE,
				UnrestrictedPrincipals: []string{},
			},
			Type: datadogV2.DATASETRESTRICTIONSTYPE_DATASET_RESTRICTIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateDatasetRestriction", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDatasetRestrictionsApi(apiClient)
	resp, r, err := api.UpdateDatasetRestriction(ctx, "product_type", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasetRestrictionsApi.UpdateDatasetRestriction`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DatasetRestrictionsApi.UpdateDatasetRestriction`:\n%s\n", responseContent)
}
