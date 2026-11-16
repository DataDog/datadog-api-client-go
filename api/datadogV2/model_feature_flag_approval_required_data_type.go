// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagApprovalRequiredDataType The resource type for an approval-required response.
type FeatureFlagApprovalRequiredDataType string

// List of FeatureFlagApprovalRequiredDataType.
const (
	FEATUREFLAGAPPROVALREQUIREDDATATYPE_APPROVAL_REQUIRED FeatureFlagApprovalRequiredDataType = "approval-required"
)

var allowedFeatureFlagApprovalRequiredDataTypeEnumValues = []FeatureFlagApprovalRequiredDataType{
	FEATUREFLAGAPPROVALREQUIREDDATATYPE_APPROVAL_REQUIRED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FeatureFlagApprovalRequiredDataType) GetAllowedValues() []FeatureFlagApprovalRequiredDataType {
	return allowedFeatureFlagApprovalRequiredDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FeatureFlagApprovalRequiredDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FeatureFlagApprovalRequiredDataType(value)
	return nil
}

// NewFeatureFlagApprovalRequiredDataTypeFromValue returns a pointer to a valid FeatureFlagApprovalRequiredDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFeatureFlagApprovalRequiredDataTypeFromValue(v string) (*FeatureFlagApprovalRequiredDataType, error) {
	ev := FeatureFlagApprovalRequiredDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FeatureFlagApprovalRequiredDataType: valid values are %v", v, allowedFeatureFlagApprovalRequiredDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FeatureFlagApprovalRequiredDataType) IsValid() bool {
	for _, existing := range allowedFeatureFlagApprovalRequiredDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeatureFlagApprovalRequiredDataType value.
func (v FeatureFlagApprovalRequiredDataType) Ptr() *FeatureFlagApprovalRequiredDataType {
	return &v
}
