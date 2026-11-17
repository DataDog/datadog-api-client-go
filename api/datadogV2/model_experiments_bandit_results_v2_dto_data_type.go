// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditResultsV2DTODataType Bandit results resource type.
type ExperimentsBanditResultsV2DTODataType string

// List of ExperimentsBanditResultsV2DTODataType.
const (
	EXPERIMENTSBANDITRESULTSV2DTODATATYPE_BANDIT_RESULTS ExperimentsBanditResultsV2DTODataType = "bandit-results"
)

var allowedExperimentsBanditResultsV2DTODataTypeEnumValues = []ExperimentsBanditResultsV2DTODataType{
	EXPERIMENTSBANDITRESULTSV2DTODATATYPE_BANDIT_RESULTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsBanditResultsV2DTODataType) GetAllowedValues() []ExperimentsBanditResultsV2DTODataType {
	return allowedExperimentsBanditResultsV2DTODataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsBanditResultsV2DTODataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsBanditResultsV2DTODataType(value)
	return nil
}

// NewExperimentsBanditResultsV2DTODataTypeFromValue returns a pointer to a valid ExperimentsBanditResultsV2DTODataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsBanditResultsV2DTODataTypeFromValue(v string) (*ExperimentsBanditResultsV2DTODataType, error) {
	ev := ExperimentsBanditResultsV2DTODataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsBanditResultsV2DTODataType: valid values are %v", v, allowedExperimentsBanditResultsV2DTODataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsBanditResultsV2DTODataType) IsValid() bool {
	for _, existing := range allowedExperimentsBanditResultsV2DTODataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsBanditResultsV2DTODataType value.
func (v ExperimentsBanditResultsV2DTODataType) Ptr() *ExperimentsBanditResultsV2DTODataType {
	return &v
}
