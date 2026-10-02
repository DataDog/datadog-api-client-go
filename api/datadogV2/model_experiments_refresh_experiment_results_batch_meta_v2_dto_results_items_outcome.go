// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome Outcome of attempting to refresh one experiment.
type ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome string

// List of ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome.
const (
	EXPERIMENTSREFRESHEXPERIMENTRESULTSBATCHMETAV2DTORESULTSITEMSOUTCOME_TRIGGERED               ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome = "TRIGGERED"
	EXPERIMENTSREFRESHEXPERIMENTRESULTSBATCHMETAV2DTORESULTSITEMSOUTCOME_SKIPPED_ALREADY_RUNNING ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome = "SKIPPED_ALREADY_RUNNING"
	EXPERIMENTSREFRESHEXPERIMENTRESULTSBATCHMETAV2DTORESULTSITEMSOUTCOME_SKIPPED_NOT_EDITABLE    ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome = "SKIPPED_NOT_EDITABLE"
	EXPERIMENTSREFRESHEXPERIMENTRESULTSBATCHMETAV2DTORESULTSITEMSOUTCOME_SKIPPED_ORG_AT_CAPACITY ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome = "SKIPPED_ORG_AT_CAPACITY"
	EXPERIMENTSREFRESHEXPERIMENTRESULTSBATCHMETAV2DTORESULTSITEMSOUTCOME_FAILED                  ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome = "FAILED"
)

var allowedExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcomeEnumValues = []ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome{
	EXPERIMENTSREFRESHEXPERIMENTRESULTSBATCHMETAV2DTORESULTSITEMSOUTCOME_TRIGGERED,
	EXPERIMENTSREFRESHEXPERIMENTRESULTSBATCHMETAV2DTORESULTSITEMSOUTCOME_SKIPPED_ALREADY_RUNNING,
	EXPERIMENTSREFRESHEXPERIMENTRESULTSBATCHMETAV2DTORESULTSITEMSOUTCOME_SKIPPED_NOT_EDITABLE,
	EXPERIMENTSREFRESHEXPERIMENTRESULTSBATCHMETAV2DTORESULTSITEMSOUTCOME_SKIPPED_ORG_AT_CAPACITY,
	EXPERIMENTSREFRESHEXPERIMENTRESULTSBATCHMETAV2DTORESULTSITEMSOUTCOME_FAILED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome) GetAllowedValues() []ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome {
	return allowedExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcomeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome(value)
	return nil
}

// NewExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcomeFromValue returns a pointer to a valid ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcomeFromValue(v string) (*ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome, error) {
	ev := ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome: valid values are %v", v, allowedExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcomeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome) IsValid() bool {
	for _, existing := range allowedExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcomeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome value.
func (v ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome) Ptr() *ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome {
	return &v
}
