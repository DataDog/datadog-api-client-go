// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumArchiveSessionType JSON:API resource type.
type RumArchiveSessionType string

// List of RumArchiveSessionType.
const (
	RUMARCHIVESESSIONTYPE_RUM_ARCHIVE_SESSION RumArchiveSessionType = "rum_archive_session"
)

var allowedRumArchiveSessionTypeEnumValues = []RumArchiveSessionType{
	RUMARCHIVESESSIONTYPE_RUM_ARCHIVE_SESSION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumArchiveSessionType) GetAllowedValues() []RumArchiveSessionType {
	return allowedRumArchiveSessionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumArchiveSessionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumArchiveSessionType(value)
	return nil
}

// NewRumArchiveSessionTypeFromValue returns a pointer to a valid RumArchiveSessionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumArchiveSessionTypeFromValue(v string) (*RumArchiveSessionType, error) {
	ev := RumArchiveSessionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumArchiveSessionType: valid values are %v", v, allowedRumArchiveSessionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumArchiveSessionType) IsValid() bool {
	for _, existing := range allowedRumArchiveSessionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumArchiveSessionType value.
func (v RumArchiveSessionType) Ptr() *RumArchiveSessionType {
	return &v
}
