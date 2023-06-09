// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// DowntimeAttributeNotifyEndStateActions Action that will trigger a monitor notification if the downtime is in the `notify_end_types` state.
type DowntimeAttributeNotifyEndStateActions string

// List of DowntimeAttributeNotifyEndStateActions.
const (
	DOWNTIMEATTRIBUTENOTIFYENDSTATEACTIONS_CANCELED DowntimeAttributeNotifyEndStateActions = "canceled"
	DOWNTIMEATTRIBUTENOTIFYENDSTATEACTIONS_EXPIRED  DowntimeAttributeNotifyEndStateActions = "expired"
)

var allowedDowntimeAttributeNotifyEndStateActionsEnumValues = []DowntimeAttributeNotifyEndStateActions{
	DOWNTIMEATTRIBUTENOTIFYENDSTATEACTIONS_CANCELED,
	DOWNTIMEATTRIBUTENOTIFYENDSTATEACTIONS_EXPIRED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DowntimeAttributeNotifyEndStateActions) GetAllowedValues() []DowntimeAttributeNotifyEndStateActions {
	return allowedDowntimeAttributeNotifyEndStateActionsEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DowntimeAttributeNotifyEndStateActions) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DowntimeAttributeNotifyEndStateActions(value)
	return nil
}

// NewDowntimeAttributeNotifyEndStateActionsFromValue returns a pointer to a valid DowntimeAttributeNotifyEndStateActions
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDowntimeAttributeNotifyEndStateActionsFromValue(v string) (*DowntimeAttributeNotifyEndStateActions, error) {
	ev := DowntimeAttributeNotifyEndStateActions(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DowntimeAttributeNotifyEndStateActions: valid values are %v", v, allowedDowntimeAttributeNotifyEndStateActionsEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DowntimeAttributeNotifyEndStateActions) IsValid() bool {
	for _, existing := range allowedDowntimeAttributeNotifyEndStateActionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DowntimeAttributeNotifyEndStateActions value.
func (v DowntimeAttributeNotifyEndStateActions) Ptr() *DowntimeAttributeNotifyEndStateActions {
	return &v
}

// NullableDowntimeAttributeNotifyEndStateActions handles when a null is used for DowntimeAttributeNotifyEndStateActions.
type NullableDowntimeAttributeNotifyEndStateActions struct {
	value *DowntimeAttributeNotifyEndStateActions
	isSet bool
}

// Get returns the associated value.
func (v NullableDowntimeAttributeNotifyEndStateActions) Get() *DowntimeAttributeNotifyEndStateActions {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDowntimeAttributeNotifyEndStateActions) Set(val *DowntimeAttributeNotifyEndStateActions) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDowntimeAttributeNotifyEndStateActions) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableDowntimeAttributeNotifyEndStateActions) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDowntimeAttributeNotifyEndStateActions initializes the struct as if Set has been called.
func NewNullableDowntimeAttributeNotifyEndStateActions(val *DowntimeAttributeNotifyEndStateActions) *NullableDowntimeAttributeNotifyEndStateActions {
	return &NullableDowntimeAttributeNotifyEndStateActions{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDowntimeAttributeNotifyEndStateActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDowntimeAttributeNotifyEndStateActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
