// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditPolicyHistoryV2DTODataType Bandit policy resource type.
type ExperimentsBanditPolicyHistoryV2DTODataType string

// List of ExperimentsBanditPolicyHistoryV2DTODataType.
const (
	EXPERIMENTSBANDITPOLICYHISTORYV2DTODATATYPE_BANDIT_POLICIES ExperimentsBanditPolicyHistoryV2DTODataType = "bandit-policies"
)

var allowedExperimentsBanditPolicyHistoryV2DTODataTypeEnumValues = []ExperimentsBanditPolicyHistoryV2DTODataType{
	EXPERIMENTSBANDITPOLICYHISTORYV2DTODATATYPE_BANDIT_POLICIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsBanditPolicyHistoryV2DTODataType) GetAllowedValues() []ExperimentsBanditPolicyHistoryV2DTODataType {
	return allowedExperimentsBanditPolicyHistoryV2DTODataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsBanditPolicyHistoryV2DTODataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsBanditPolicyHistoryV2DTODataType(value)
	return nil
}

// NewExperimentsBanditPolicyHistoryV2DTODataTypeFromValue returns a pointer to a valid ExperimentsBanditPolicyHistoryV2DTODataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsBanditPolicyHistoryV2DTODataTypeFromValue(v string) (*ExperimentsBanditPolicyHistoryV2DTODataType, error) {
	ev := ExperimentsBanditPolicyHistoryV2DTODataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsBanditPolicyHistoryV2DTODataType: valid values are %v", v, allowedExperimentsBanditPolicyHistoryV2DTODataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsBanditPolicyHistoryV2DTODataType) IsValid() bool {
	for _, existing := range allowedExperimentsBanditPolicyHistoryV2DTODataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsBanditPolicyHistoryV2DTODataType value.
func (v ExperimentsBanditPolicyHistoryV2DTODataType) Ptr() *ExperimentsBanditPolicyHistoryV2DTODataType {
	return &v
}
