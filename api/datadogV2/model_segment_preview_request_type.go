// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SegmentPreviewRequestType Segment preview request resource type.
type SegmentPreviewRequestType string

// List of SegmentPreviewRequestType.
const (
	SEGMENTPREVIEWREQUESTTYPE_SEGMENT_PREVIEW_REQUEST SegmentPreviewRequestType = "segment_preview_request"
)

var allowedSegmentPreviewRequestTypeEnumValues = []SegmentPreviewRequestType{
	SEGMENTPREVIEWREQUESTTYPE_SEGMENT_PREVIEW_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SegmentPreviewRequestType) GetAllowedValues() []SegmentPreviewRequestType {
	return allowedSegmentPreviewRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SegmentPreviewRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SegmentPreviewRequestType(value)
	return nil
}

// NewSegmentPreviewRequestTypeFromValue returns a pointer to a valid SegmentPreviewRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSegmentPreviewRequestTypeFromValue(v string) (*SegmentPreviewRequestType, error) {
	ev := SegmentPreviewRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SegmentPreviewRequestType: valid values are %v", v, allowedSegmentPreviewRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SegmentPreviewRequestType) IsValid() bool {
	for _, existing := range allowedSegmentPreviewRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SegmentPreviewRequestType value.
func (v SegmentPreviewRequestType) Ptr() *SegmentPreviewRequestType {
	return &v
}
