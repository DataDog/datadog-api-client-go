// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagApprovalsDataType The resource type for feature flag approval settings.
type FeatureFlagApprovalsDataType string

// List of FeatureFlagApprovalsDataType.
const (
	FEATUREFLAGAPPROVALSDATATYPE_FEATURE_FLAG_APPROVALS FeatureFlagApprovalsDataType = "feature-flag-approvals"
)

var allowedFeatureFlagApprovalsDataTypeEnumValues = []FeatureFlagApprovalsDataType{
	FEATUREFLAGAPPROVALSDATATYPE_FEATURE_FLAG_APPROVALS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FeatureFlagApprovalsDataType) GetAllowedValues() []FeatureFlagApprovalsDataType {
	return allowedFeatureFlagApprovalsDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FeatureFlagApprovalsDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FeatureFlagApprovalsDataType(value)
	return nil
}

// NewFeatureFlagApprovalsDataTypeFromValue returns a pointer to a valid FeatureFlagApprovalsDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFeatureFlagApprovalsDataTypeFromValue(v string) (*FeatureFlagApprovalsDataType, error) {
	ev := FeatureFlagApprovalsDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FeatureFlagApprovalsDataType: valid values are %v", v, allowedFeatureFlagApprovalsDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FeatureFlagApprovalsDataType) IsValid() bool {
	for _, existing := range allowedFeatureFlagApprovalsDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeatureFlagApprovalsDataType value.
func (v FeatureFlagApprovalsDataType) Ptr() *FeatureFlagApprovalsDataType {
	return &v
}
