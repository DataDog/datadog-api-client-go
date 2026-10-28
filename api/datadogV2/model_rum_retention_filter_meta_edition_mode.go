// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRetentionFilterMetaEditionMode Which attributes of the retention filter can be edited. `complete`: all attributes can be edited. `limited`: only the sample rates can be edited; the name, event type, and query are fixed and the filter cannot be reordered (for example, default filters). `unknown`: the edition mode is not set.
type RumRetentionFilterMetaEditionMode string

// List of RumRetentionFilterMetaEditionMode.
const (
	RUMRETENTIONFILTERMETAEDITIONMODE_UNKNOWN  RumRetentionFilterMetaEditionMode = "unknown"
	RUMRETENTIONFILTERMETAEDITIONMODE_COMPLETE RumRetentionFilterMetaEditionMode = "complete"
	RUMRETENTIONFILTERMETAEDITIONMODE_LIMITED  RumRetentionFilterMetaEditionMode = "limited"
)

var allowedRumRetentionFilterMetaEditionModeEnumValues = []RumRetentionFilterMetaEditionMode{
	RUMRETENTIONFILTERMETAEDITIONMODE_UNKNOWN,
	RUMRETENTIONFILTERMETAEDITIONMODE_COMPLETE,
	RUMRETENTIONFILTERMETAEDITIONMODE_LIMITED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumRetentionFilterMetaEditionMode) GetAllowedValues() []RumRetentionFilterMetaEditionMode {
	return allowedRumRetentionFilterMetaEditionModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumRetentionFilterMetaEditionMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumRetentionFilterMetaEditionMode(value)
	return nil
}

// NewRumRetentionFilterMetaEditionModeFromValue returns a pointer to a valid RumRetentionFilterMetaEditionMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumRetentionFilterMetaEditionModeFromValue(v string) (*RumRetentionFilterMetaEditionMode, error) {
	ev := RumRetentionFilterMetaEditionMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumRetentionFilterMetaEditionMode: valid values are %v", v, allowedRumRetentionFilterMetaEditionModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumRetentionFilterMetaEditionMode) IsValid() bool {
	for _, existing := range allowedRumRetentionFilterMetaEditionModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumRetentionFilterMetaEditionMode value.
func (v RumRetentionFilterMetaEditionMode) Ptr() *RumRetentionFilterMetaEditionMode {
	return &v
}
