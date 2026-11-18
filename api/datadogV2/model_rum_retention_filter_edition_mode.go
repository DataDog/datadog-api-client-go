// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRetentionFilterEditionMode The types of changes the retention filter accepts. `complete` filters can be updated and deleted.
// `limited` filters are default retention filters that can't be deleted.
type RumRetentionFilterEditionMode string

// List of RumRetentionFilterEditionMode.
const (
	RUMRETENTIONFILTEREDITIONMODE_UNKNOWN  RumRetentionFilterEditionMode = "unknown"
	RUMRETENTIONFILTEREDITIONMODE_COMPLETE RumRetentionFilterEditionMode = "complete"
	RUMRETENTIONFILTEREDITIONMODE_LIMITED  RumRetentionFilterEditionMode = "limited"
)

var allowedRumRetentionFilterEditionModeEnumValues = []RumRetentionFilterEditionMode{
	RUMRETENTIONFILTEREDITIONMODE_UNKNOWN,
	RUMRETENTIONFILTEREDITIONMODE_COMPLETE,
	RUMRETENTIONFILTEREDITIONMODE_LIMITED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumRetentionFilterEditionMode) GetAllowedValues() []RumRetentionFilterEditionMode {
	return allowedRumRetentionFilterEditionModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumRetentionFilterEditionMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumRetentionFilterEditionMode(value)
	return nil
}

// NewRumRetentionFilterEditionModeFromValue returns a pointer to a valid RumRetentionFilterEditionMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumRetentionFilterEditionModeFromValue(v string) (*RumRetentionFilterEditionMode, error) {
	ev := RumRetentionFilterEditionMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumRetentionFilterEditionMode: valid values are %v", v, allowedRumRetentionFilterEditionModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumRetentionFilterEditionMode) IsValid() bool {
	for _, existing := range allowedRumRetentionFilterEditionModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumRetentionFilterEditionMode value.
func (v RumRetentionFilterEditionMode) Ptr() *RumRetentionFilterEditionMode {
	return &v
}
