// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumArchiveConfigType JSON:API resource type.
type RumArchiveConfigType string

// List of RumArchiveConfigType.
const (
	RUMARCHIVECONFIGTYPE_RUM_ARCHIVE_CONFIG RumArchiveConfigType = "rum_archive_config"
)

var allowedRumArchiveConfigTypeEnumValues = []RumArchiveConfigType{
	RUMARCHIVECONFIGTYPE_RUM_ARCHIVE_CONFIG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumArchiveConfigType) GetAllowedValues() []RumArchiveConfigType {
	return allowedRumArchiveConfigTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumArchiveConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumArchiveConfigType(value)
	return nil
}

// NewRumArchiveConfigTypeFromValue returns a pointer to a valid RumArchiveConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumArchiveConfigTypeFromValue(v string) (*RumArchiveConfigType, error) {
	ev := RumArchiveConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumArchiveConfigType: valid values are %v", v, allowedRumArchiveConfigTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumArchiveConfigType) IsValid() bool {
	for _, existing := range allowedRumArchiveConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumArchiveConfigType value.
func (v RumArchiveConfigType) Ptr() *RumArchiveConfigType {
	return &v
}
