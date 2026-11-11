// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsExperimentHoldoutV2DTODataType Experiment holdout resource type.
type ExperimentsExperimentHoldoutV2DTODataType string

// List of ExperimentsExperimentHoldoutV2DTODataType.
const (
	EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATATYPE_EXPERIMENT_HOLDOUTS ExperimentsExperimentHoldoutV2DTODataType = "experiment-holdouts"
)

var allowedExperimentsExperimentHoldoutV2DTODataTypeEnumValues = []ExperimentsExperimentHoldoutV2DTODataType{
	EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATATYPE_EXPERIMENT_HOLDOUTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsExperimentHoldoutV2DTODataType) GetAllowedValues() []ExperimentsExperimentHoldoutV2DTODataType {
	return allowedExperimentsExperimentHoldoutV2DTODataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsExperimentHoldoutV2DTODataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsExperimentHoldoutV2DTODataType(value)
	return nil
}

// NewExperimentsExperimentHoldoutV2DTODataTypeFromValue returns a pointer to a valid ExperimentsExperimentHoldoutV2DTODataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsExperimentHoldoutV2DTODataTypeFromValue(v string) (*ExperimentsExperimentHoldoutV2DTODataType, error) {
	ev := ExperimentsExperimentHoldoutV2DTODataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsExperimentHoldoutV2DTODataType: valid values are %v", v, allowedExperimentsExperimentHoldoutV2DTODataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsExperimentHoldoutV2DTODataType) IsValid() bool {
	for _, existing := range allowedExperimentsExperimentHoldoutV2DTODataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsExperimentHoldoutV2DTODataType value.
func (v ExperimentsExperimentHoldoutV2DTODataType) Ptr() *ExperimentsExperimentHoldoutV2DTODataType {
	return &v
}
