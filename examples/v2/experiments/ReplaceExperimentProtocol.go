// Replace experiment protocol returns "OK" response

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
	// there is a valid "published_protocol" in the system
	PublishedProtocolDataAttributesUpdatedAt := os.Getenv("PUBLISHED_PROTOCOL_DATA_ATTRIBUTES_UPDATED_AT")
	PublishedProtocolDataID := uuid.MustParse(os.Getenv("PUBLISHED_PROTOCOL_DATA_ID"))

	// there is a valid "experiment_subject_type" in the system
	ExperimentSubjectTypeDataID := uuid.MustParse(os.Getenv("EXPERIMENT_SUBJECT_TYPE_DATA_ID"))

	body := datadogV2.ExperimentsReplaceExperimentProtocolRequest{
		Data: datadogV2.ExperimentsReplaceExperimentProtocolRequestData{
			Type: datadogV2.EXPERIMENTSPUBLICPROTOCOLRESPONSEDATATYPE_PROTOCOLS,
			Id:   datadog.PtrUUID(PublishedProtocolDataID),
			Attributes: datadogV2.ExperimentsReplaceExperimentProtocolRequestDataAttributes{
				SubjectTypeId:             ExperimentSubjectTypeDataID,
				IsEqualSplitEnforced:      false,
				IsMinimumDurationEnabled:  false,
				IsDurationRequiredToStart: false,
				IsMetricLimitEnabled:      false,
				MetricGroups:              []datadogV2.ExperimentsProtocolMetricGroupInput{},
				DefaultDurationDays:       *datadog.NewNullableInt64(datadog.PtrInt64(14)),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExperimentsApi(apiClient)
	resp, r, err := api.ReplaceExperimentProtocol(ctx, PublishedProtocolDataID, body, *datadogV2.NewReplaceExperimentProtocolOptionalParameters().WithExpectedUpdatedAt(PublishedProtocolDataAttributesUpdatedAt))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsApi.ReplaceExperimentProtocol`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsApi.ReplaceExperimentProtocol`:\n%s\n", responseContent)
}
