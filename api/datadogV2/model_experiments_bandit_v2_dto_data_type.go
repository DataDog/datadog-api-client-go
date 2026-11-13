// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditV2DTODataType Bandit resource type.
type ExperimentsBanditV2DTODataType string

// List of ExperimentsBanditV2DTODataType.
const (
	EXPERIMENTSBANDITV2DTODATATYPE_BANDITS ExperimentsBanditV2DTODataType = "bandits"
)

var allowedExperimentsBanditV2DTODataTypeEnumValues = []ExperimentsBanditV2DTODataType{
	EXPERIMENTSBANDITV2DTODATATYPE_BANDITS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsBanditV2DTODataType) GetAllowedValues() []ExperimentsBanditV2DTODataType {
	return allowedExperimentsBanditV2DTODataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsBanditV2DTODataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsBanditV2DTODataType(value)
	return nil
}

// NewExperimentsBanditV2DTODataTypeFromValue returns a pointer to a valid ExperimentsBanditV2DTODataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsBanditV2DTODataTypeFromValue(v string) (*ExperimentsBanditV2DTODataType, error) {
	ev := ExperimentsBanditV2DTODataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsBanditV2DTODataType: valid values are %v", v, allowedExperimentsBanditV2DTODataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsBanditV2DTODataType) IsValid() bool {
	for _, existing := range allowedExperimentsBanditV2DTODataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsBanditV2DTODataType value.
func (v ExperimentsBanditV2DTODataType) Ptr() *ExperimentsBanditV2DTODataType {
	return &v
}
