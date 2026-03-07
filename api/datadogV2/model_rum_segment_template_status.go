// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSegmentTemplateStatus The status of a segment template.
type RumSegmentTemplateStatus string

// List of RumSegmentTemplateStatus.
const (
	RUMSEGMENTTEMPLATESTATUS_ACTIVE     RumSegmentTemplateStatus = "active"
	RUMSEGMENTTEMPLATESTATUS_DEPRECATED RumSegmentTemplateStatus = "deprecated"
	RUMSEGMENTTEMPLATESTATUS_ARCHIVED   RumSegmentTemplateStatus = "archived"
)

var allowedRumSegmentTemplateStatusEnumValues = []RumSegmentTemplateStatus{
	RUMSEGMENTTEMPLATESTATUS_ACTIVE,
	RUMSEGMENTTEMPLATESTATUS_DEPRECATED,
	RUMSEGMENTTEMPLATESTATUS_ARCHIVED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumSegmentTemplateStatus) GetAllowedValues() []RumSegmentTemplateStatus {
	return allowedRumSegmentTemplateStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumSegmentTemplateStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumSegmentTemplateStatus(value)
	return nil
}

// NewRumSegmentTemplateStatusFromValue returns a pointer to a valid RumSegmentTemplateStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumSegmentTemplateStatusFromValue(v string) (*RumSegmentTemplateStatus, error) {
	ev := RumSegmentTemplateStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumSegmentTemplateStatus: valid values are %v", v, allowedRumSegmentTemplateStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumSegmentTemplateStatus) IsValid() bool {
	for _, existing := range allowedRumSegmentTemplateStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumSegmentTemplateStatus value.
func (v RumSegmentTemplateStatus) Ptr() *RumSegmentTemplateStatus {
	return &v
}
