// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumStaticSegmentRequestType Type of the static segment creation request resource.
type RumStaticSegmentRequestType string

// List of RumStaticSegmentRequestType.
const (
	RUMSTATICSEGMENTREQUESTTYPE_CREATE_STATIC_SEGMENT_REQUEST RumStaticSegmentRequestType = "create_static_segment_request"
)

var allowedRumStaticSegmentRequestTypeEnumValues = []RumStaticSegmentRequestType{
	RUMSTATICSEGMENTREQUESTTYPE_CREATE_STATIC_SEGMENT_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumStaticSegmentRequestType) GetAllowedValues() []RumStaticSegmentRequestType {
	return allowedRumStaticSegmentRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumStaticSegmentRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumStaticSegmentRequestType(value)
	return nil
}

// NewRumStaticSegmentRequestTypeFromValue returns a pointer to a valid RumStaticSegmentRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumStaticSegmentRequestTypeFromValue(v string) (*RumStaticSegmentRequestType, error) {
	ev := RumStaticSegmentRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumStaticSegmentRequestType: valid values are %v", v, allowedRumStaticSegmentRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumStaticSegmentRequestType) IsValid() bool {
	for _, existing := range allowedRumStaticSegmentRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumStaticSegmentRequestType value.
func (v RumStaticSegmentRequestType) Ptr() *RumStaticSegmentRequestType {
	return &v
}
