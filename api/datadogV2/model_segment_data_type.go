// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SegmentDataType
type SegmentDataType string

// List of SegmentDataType.
const (
	SEGMENTDATATYPE_SEGMENT SegmentDataType = "segment"
)

var allowedSegmentDataTypeEnumValues = []SegmentDataType{
	SEGMENTDATATYPE_SEGMENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SegmentDataType) GetAllowedValues() []SegmentDataType {
	return allowedSegmentDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SegmentDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SegmentDataType(value)
	return nil
}

// NewSegmentDataTypeFromValue returns a pointer to a valid SegmentDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSegmentDataTypeFromValue(v string) (*SegmentDataType, error) {
	ev := SegmentDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SegmentDataType: valid values are %v", v, allowedSegmentDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SegmentDataType) IsValid() bool {
	for _, existing := range allowedSegmentDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SegmentDataType value.
func (v SegmentDataType) Ptr() *SegmentDataType {
	return &v
}
