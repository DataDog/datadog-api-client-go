// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SegmentPreviewResponseType Segment preview response resource type.
type SegmentPreviewResponseType string

// List of SegmentPreviewResponseType.
const (
	SEGMENTPREVIEWRESPONSETYPE_UNIFIED_SEGMENT_PREVIEW_RESPONSE SegmentPreviewResponseType = "unified_segment_preview_response"
)

var allowedSegmentPreviewResponseTypeEnumValues = []SegmentPreviewResponseType{
	SEGMENTPREVIEWRESPONSETYPE_UNIFIED_SEGMENT_PREVIEW_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SegmentPreviewResponseType) GetAllowedValues() []SegmentPreviewResponseType {
	return allowedSegmentPreviewResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SegmentPreviewResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SegmentPreviewResponseType(value)
	return nil
}

// NewSegmentPreviewResponseTypeFromValue returns a pointer to a valid SegmentPreviewResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSegmentPreviewResponseTypeFromValue(v string) (*SegmentPreviewResponseType, error) {
	ev := SegmentPreviewResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SegmentPreviewResponseType: valid values are %v", v, allowedSegmentPreviewResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SegmentPreviewResponseType) IsValid() bool {
	for _, existing := range allowedSegmentPreviewResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SegmentPreviewResponseType value.
func (v SegmentPreviewResponseType) Ptr() *SegmentPreviewResponseType {
	return &v
}
