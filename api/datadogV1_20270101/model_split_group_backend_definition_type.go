// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SplitGroupBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type SplitGroupBackendDefinitionType string

// List of SplitGroupBackendDefinitionType.
const (
	SPLITGROUPBACKENDDEFINITIONTYPE_SPLIT_GROUP SplitGroupBackendDefinitionType = "split_group"
)

var allowedSplitGroupBackendDefinitionTypeEnumValues = []SplitGroupBackendDefinitionType{
	SPLITGROUPBACKENDDEFINITIONTYPE_SPLIT_GROUP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SplitGroupBackendDefinitionType) GetAllowedValues() []SplitGroupBackendDefinitionType {
	return allowedSplitGroupBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SplitGroupBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SplitGroupBackendDefinitionType(value)
	return nil
}

// NewSplitGroupBackendDefinitionTypeFromValue returns a pointer to a valid SplitGroupBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSplitGroupBackendDefinitionTypeFromValue(v string) (*SplitGroupBackendDefinitionType, error) {
	ev := SplitGroupBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SplitGroupBackendDefinitionType: valid values are %v", v, allowedSplitGroupBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SplitGroupBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedSplitGroupBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SplitGroupBackendDefinitionType value.
func (v SplitGroupBackendDefinitionType) Ptr() *SplitGroupBackendDefinitionType {
	return &v
}
