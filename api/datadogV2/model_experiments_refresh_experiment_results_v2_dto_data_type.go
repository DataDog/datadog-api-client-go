// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsRefreshExperimentResultsV2DTODataType Experiment results refresh resource type.
type ExperimentsRefreshExperimentResultsV2DTODataType string

// List of ExperimentsRefreshExperimentResultsV2DTODataType.
const (
	EXPERIMENTSREFRESHEXPERIMENTRESULTSV2DTODATATYPE_EXPERIMENT_RESULTS_REFRESH ExperimentsRefreshExperimentResultsV2DTODataType = "experiment-results-refresh"
)

var allowedExperimentsRefreshExperimentResultsV2DTODataTypeEnumValues = []ExperimentsRefreshExperimentResultsV2DTODataType{
	EXPERIMENTSREFRESHEXPERIMENTRESULTSV2DTODATATYPE_EXPERIMENT_RESULTS_REFRESH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsRefreshExperimentResultsV2DTODataType) GetAllowedValues() []ExperimentsRefreshExperimentResultsV2DTODataType {
	return allowedExperimentsRefreshExperimentResultsV2DTODataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsRefreshExperimentResultsV2DTODataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsRefreshExperimentResultsV2DTODataType(value)
	return nil
}

// NewExperimentsRefreshExperimentResultsV2DTODataTypeFromValue returns a pointer to a valid ExperimentsRefreshExperimentResultsV2DTODataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsRefreshExperimentResultsV2DTODataTypeFromValue(v string) (*ExperimentsRefreshExperimentResultsV2DTODataType, error) {
	ev := ExperimentsRefreshExperimentResultsV2DTODataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsRefreshExperimentResultsV2DTODataType: valid values are %v", v, allowedExperimentsRefreshExperimentResultsV2DTODataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsRefreshExperimentResultsV2DTODataType) IsValid() bool {
	for _, existing := range allowedExperimentsRefreshExperimentResultsV2DTODataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsRefreshExperimentResultsV2DTODataType value.
func (v ExperimentsRefreshExperimentResultsV2DTODataType) Ptr() *ExperimentsRefreshExperimentResultsV2DTODataType {
	return &v
}
