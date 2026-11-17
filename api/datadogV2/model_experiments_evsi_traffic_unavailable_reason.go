// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsEVSITrafficUnavailableReason Recorded traffic forecast unavailability reason. Per-analysis reasons are reported separately.
type ExperimentsEVSITrafficUnavailableReason string

// List of ExperimentsEVSITrafficUnavailableReason.
const (
	EXPERIMENTSEVSITRAFFICUNAVAILABLEREASON_TRAFFIC_UNAVAILABLE        ExperimentsEVSITrafficUnavailableReason = "traffic_unavailable"
	EXPERIMENTSEVSITRAFFICUNAVAILABLEREASON_INSUFFICIENT_COMPLETE_DAYS ExperimentsEVSITrafficUnavailableReason = "insufficient_complete_days"
	EXPERIMENTSEVSITRAFFICUNAVAILABLEREASON_NO_TOTAL_EXPOSURES         ExperimentsEVSITrafficUnavailableReason = "no_total_exposures"
	EXPERIMENTSEVSITRAFFICUNAVAILABLEREASON_NO_LOOKBACK_EXPOSURES      ExperimentsEVSITrafficUnavailableReason = "no_lookback_exposures"
	EXPERIMENTSEVSITRAFFICUNAVAILABLEREASON_INVALID_TRAFFIC            ExperimentsEVSITrafficUnavailableReason = "invalid_traffic"
)

var allowedExperimentsEVSITrafficUnavailableReasonEnumValues = []ExperimentsEVSITrafficUnavailableReason{
	EXPERIMENTSEVSITRAFFICUNAVAILABLEREASON_TRAFFIC_UNAVAILABLE,
	EXPERIMENTSEVSITRAFFICUNAVAILABLEREASON_INSUFFICIENT_COMPLETE_DAYS,
	EXPERIMENTSEVSITRAFFICUNAVAILABLEREASON_NO_TOTAL_EXPOSURES,
	EXPERIMENTSEVSITRAFFICUNAVAILABLEREASON_NO_LOOKBACK_EXPOSURES,
	EXPERIMENTSEVSITRAFFICUNAVAILABLEREASON_INVALID_TRAFFIC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsEVSITrafficUnavailableReason) GetAllowedValues() []ExperimentsEVSITrafficUnavailableReason {
	return allowedExperimentsEVSITrafficUnavailableReasonEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsEVSITrafficUnavailableReason) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsEVSITrafficUnavailableReason(value)
	return nil
}

// NewExperimentsEVSITrafficUnavailableReasonFromValue returns a pointer to a valid ExperimentsEVSITrafficUnavailableReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsEVSITrafficUnavailableReasonFromValue(v string) (*ExperimentsEVSITrafficUnavailableReason, error) {
	ev := ExperimentsEVSITrafficUnavailableReason(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsEVSITrafficUnavailableReason: valid values are %v", v, allowedExperimentsEVSITrafficUnavailableReasonEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsEVSITrafficUnavailableReason) IsValid() bool {
	for _, existing := range allowedExperimentsEVSITrafficUnavailableReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsEVSITrafficUnavailableReason value.
func (v ExperimentsEVSITrafficUnavailableReason) Ptr() *ExperimentsEVSITrafficUnavailableReason {
	return &v
}

// NullableExperimentsEVSITrafficUnavailableReason handles when a null is used for ExperimentsEVSITrafficUnavailableReason.
type NullableExperimentsEVSITrafficUnavailableReason struct {
	value *ExperimentsEVSITrafficUnavailableReason
	isSet bool
}

// Get returns the associated value.
func (v NullableExperimentsEVSITrafficUnavailableReason) Get() *ExperimentsEVSITrafficUnavailableReason {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableExperimentsEVSITrafficUnavailableReason) Set(val *ExperimentsEVSITrafficUnavailableReason) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableExperimentsEVSITrafficUnavailableReason) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableExperimentsEVSITrafficUnavailableReason) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableExperimentsEVSITrafficUnavailableReason initializes the struct as if Set has been called.
func NewNullableExperimentsEVSITrafficUnavailableReason(val *ExperimentsEVSITrafficUnavailableReason) *NullableExperimentsEVSITrafficUnavailableReason {
	return &NullableExperimentsEVSITrafficUnavailableReason{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableExperimentsEVSITrafficUnavailableReason) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableExperimentsEVSITrafficUnavailableReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return datadog.Unmarshal(src, &v.value)
}
