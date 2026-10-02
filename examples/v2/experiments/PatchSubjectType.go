// Patch subject type returns "OK" response

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
	// there is a valid "experiment_subject_type" in the system
	ExperimentSubjectTypeDataID := uuid.MustParse(os.Getenv("EXPERIMENT_SUBJECT_TYPE_DATA_ID"))

	body := datadogV2.ExperimentsPatchSubjectTypeV2Request{
		Data: datadogV2.ExperimentsPatchSubjectTypeV2RequestData{
			Type: datadogV2.EXPERIMENTSSUBJECTTYPEV2DTODATATYPE_SUBJECT_TYPES,
			Attributes: datadogV2.ExperimentsPatchSubjectTypeV2RequestDataAttributes{
				Name: datadog.PtrString("ex-14bb9543f523edde updated"),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExperimentsApi(apiClient)
	resp, r, err := api.PatchSubjectType(ctx, ExperimentSubjectTypeDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsApi.PatchSubjectType`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsApi.PatchSubjectType`:\n%s\n", responseContent)
}
