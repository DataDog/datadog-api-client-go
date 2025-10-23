// Post dependencies for analysis returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.ScaRequest{
		Data: &datadogV2.ScaRequestData{
			Attributes: &datadogV2.ScaRequestDataAttributes{
				Commit: &datadogV2.ScaRequestDataAttributesCommit{},
				Dependencies: []datadogV2.ScaRequestDataAttributesDependenciesItems{
					{
						Exclusions: []string{},
						Locations: []datadogV2.ScaRequestDataAttributesDependenciesItemsLocationsItems{
							{
								Block: &datadogV2.ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition{
									End:   &datadogV2.ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition{},
									Start: &datadogV2.ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition{},
								},
								Name: &datadogV2.ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition{
									End:   &datadogV2.ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition{},
									Start: &datadogV2.ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition{},
								},
								Namespace: &datadogV2.ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition{
									End:   &datadogV2.ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition{},
									Start: &datadogV2.ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition{},
								},
								Version: &datadogV2.ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition{
									End:   &datadogV2.ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition{},
									Start: &datadogV2.ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition{},
								},
							},
						},
						ReachableSymbolProperties: []datadogV2.ScaRequestDataAttributesDependenciesItemsReachableSymbolPropertiesItems{
							{},
						},
					},
				},
				Files: []datadogV2.ScaRequestDataAttributesFilesItems{
					{},
				},
				Relations: []datadogV2.ScaRequestDataAttributesRelationsItems{
					{
						DependsOn: []string{},
					},
				},
				Repository: &datadogV2.ScaRequestDataAttributesRepository{},
				Vulnerabilities: []datadogV2.ScaRequestDataAttributesVulnerabilitiesItems{
					{
						Affects: []datadogV2.ScaRequestDataAttributesVulnerabilitiesItemsAffectsItems{
							{},
						},
					},
				},
			},
			Type: datadogV2.SCAREQUESTDATATYPE_SCAREQUESTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateSCAResult", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStaticAnalysisApi(apiClient)
	r, err := api.CreateSCAResult(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticAnalysisApi.CreateSCAResult`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
