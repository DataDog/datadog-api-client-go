// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSegmentTemplateResourceType Type of the segment template resource.
type RumSegmentTemplateResourceType string

// List of RumSegmentTemplateResourceType.
const (
	RUMSEGMENTTEMPLATERESOURCETYPE_TEMPLATE_METADATA RumSegmentTemplateResourceType = "template_metadata"
)

var allowedRumSegmentTemplateResourceTypeEnumValues = []RumSegmentTemplateResourceType{
	RUMSEGMENTTEMPLATERESOURCETYPE_TEMPLATE_METADATA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumSegmentTemplateResourceType) GetAllowedValues() []RumSegmentTemplateResourceType {
	return allowedRumSegmentTemplateResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumSegmentTemplateResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumSegmentTemplateResourceType(value)
	return nil
}

// NewRumSegmentTemplateResourceTypeFromValue returns a pointer to a valid RumSegmentTemplateResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumSegmentTemplateResourceTypeFromValue(v string) (*RumSegmentTemplateResourceType, error) {
	ev := RumSegmentTemplateResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumSegmentTemplateResourceType: valid values are %v", v, allowedRumSegmentTemplateResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumSegmentTemplateResourceType) IsValid() bool {
	for _, existing := range allowedRumSegmentTemplateResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumSegmentTemplateResourceType value.
func (v RumSegmentTemplateResourceType) Ptr() *RumSegmentTemplateResourceType {
	return &v
}
