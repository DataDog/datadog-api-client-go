// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SegmentDeletedType Deleted segment resource type.
type SegmentDeletedType string

// List of SegmentDeletedType.
const (
	SEGMENTDELETEDTYPE_DELETED_SEGMENT SegmentDeletedType = "deleted_segment"
)

var allowedSegmentDeletedTypeEnumValues = []SegmentDeletedType{
	SEGMENTDELETEDTYPE_DELETED_SEGMENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SegmentDeletedType) GetAllowedValues() []SegmentDeletedType {
	return allowedSegmentDeletedTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SegmentDeletedType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SegmentDeletedType(value)
	return nil
}

// NewSegmentDeletedTypeFromValue returns a pointer to a valid SegmentDeletedType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSegmentDeletedTypeFromValue(v string) (*SegmentDeletedType, error) {
	ev := SegmentDeletedType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SegmentDeletedType: valid values are %v", v, allowedSegmentDeletedTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SegmentDeletedType) IsValid() bool {
	for _, existing := range allowedSegmentDeletedTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SegmentDeletedType value.
func (v SegmentDeletedType) Ptr() *SegmentDeletedType {
	return &v
}
