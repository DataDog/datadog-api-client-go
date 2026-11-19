// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditRefreshV2DTODataType Bandit refresh resource type.
type ExperimentsBanditRefreshV2DTODataType string

// List of ExperimentsBanditRefreshV2DTODataType.
const (
	EXPERIMENTSBANDITREFRESHV2DTODATATYPE_BANDIT_REFRESH ExperimentsBanditRefreshV2DTODataType = "bandit-refresh"
)

var allowedExperimentsBanditRefreshV2DTODataTypeEnumValues = []ExperimentsBanditRefreshV2DTODataType{
	EXPERIMENTSBANDITREFRESHV2DTODATATYPE_BANDIT_REFRESH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsBanditRefreshV2DTODataType) GetAllowedValues() []ExperimentsBanditRefreshV2DTODataType {
	return allowedExperimentsBanditRefreshV2DTODataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsBanditRefreshV2DTODataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsBanditRefreshV2DTODataType(value)
	return nil
}

// NewExperimentsBanditRefreshV2DTODataTypeFromValue returns a pointer to a valid ExperimentsBanditRefreshV2DTODataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsBanditRefreshV2DTODataTypeFromValue(v string) (*ExperimentsBanditRefreshV2DTODataType, error) {
	ev := ExperimentsBanditRefreshV2DTODataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsBanditRefreshV2DTODataType: valid values are %v", v, allowedExperimentsBanditRefreshV2DTODataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsBanditRefreshV2DTODataType) IsValid() bool {
	for _, existing := range allowedExperimentsBanditRefreshV2DTODataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsBanditRefreshV2DTODataType value.
func (v ExperimentsBanditRefreshV2DTODataType) Ptr() *ExperimentsBanditRefreshV2DTODataType {
	return &v
}
