// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMRetentionFilterEditionMode The types of changes the retention filter accepts. `complete` filters can be updated and deleted.
// `limited` filters are default retention filters that can't be deleted.
type RUMRetentionFilterEditionMode string

// List of RUMRetentionFilterEditionMode.
const (
	RUMRETENTIONFILTEREDITIONMODE_UNKNOWN  RUMRetentionFilterEditionMode = "unknown"
	RUMRETENTIONFILTEREDITIONMODE_COMPLETE RUMRetentionFilterEditionMode = "complete"
	RUMRETENTIONFILTEREDITIONMODE_LIMITED  RUMRetentionFilterEditionMode = "limited"
)

var allowedRUMRetentionFilterEditionModeEnumValues = []RUMRetentionFilterEditionMode{
	RUMRETENTIONFILTEREDITIONMODE_UNKNOWN,
	RUMRETENTIONFILTEREDITIONMODE_COMPLETE,
	RUMRETENTIONFILTEREDITIONMODE_LIMITED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMRetentionFilterEditionMode) GetAllowedValues() []RUMRetentionFilterEditionMode {
	return allowedRUMRetentionFilterEditionModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMRetentionFilterEditionMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMRetentionFilterEditionMode(value)
	return nil
}

// NewRUMRetentionFilterEditionModeFromValue returns a pointer to a valid RUMRetentionFilterEditionMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMRetentionFilterEditionModeFromValue(v string) (*RUMRetentionFilterEditionMode, error) {
	ev := RUMRetentionFilterEditionMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMRetentionFilterEditionMode: valid values are %v", v, allowedRUMRetentionFilterEditionModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMRetentionFilterEditionMode) IsValid() bool {
	for _, existing := range allowedRUMRetentionFilterEditionModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMRetentionFilterEditionMode value.
func (v RUMRetentionFilterEditionMode) Ptr() *RUMRetentionFilterEditionMode {
	return &v
}
