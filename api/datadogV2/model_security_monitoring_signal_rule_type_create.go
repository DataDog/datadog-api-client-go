// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// SecurityMonitoringSignalRuleTypeCreate The rule type.
type SecurityMonitoringSignalRuleTypeCreate string

// List of SecurityMonitoringSignalRuleTypeCreate.
const (
	SECURITYMONITORINGSIGNALRULETYPECREATE_SIGNAL_CORRELATION SecurityMonitoringSignalRuleTypeCreate = "signal_correlation"
)

var allowedSecurityMonitoringSignalRuleTypeCreateEnumValues = []SecurityMonitoringSignalRuleTypeCreate{
	SECURITYMONITORINGSIGNALRULETYPECREATE_SIGNAL_CORRELATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringSignalRuleTypeCreate) GetAllowedValues() []SecurityMonitoringSignalRuleTypeCreate {
	return allowedSecurityMonitoringSignalRuleTypeCreateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringSignalRuleTypeCreate) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringSignalRuleTypeCreate(value)
	return nil
}

// NewSecurityMonitoringSignalRuleTypeCreateFromValue returns a pointer to a valid SecurityMonitoringSignalRuleTypeCreate
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringSignalRuleTypeCreateFromValue(v string) (*SecurityMonitoringSignalRuleTypeCreate, error) {
	ev := SecurityMonitoringSignalRuleTypeCreate(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringSignalRuleTypeCreate: valid values are %v", v, allowedSecurityMonitoringSignalRuleTypeCreateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringSignalRuleTypeCreate) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringSignalRuleTypeCreateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringSignalRuleTypeCreate value.
func (v SecurityMonitoringSignalRuleTypeCreate) Ptr() *SecurityMonitoringSignalRuleTypeCreate {
	return &v
}

// NullableSecurityMonitoringSignalRuleTypeCreate handles when a null is used for SecurityMonitoringSignalRuleTypeCreate.
type NullableSecurityMonitoringSignalRuleTypeCreate struct {
	value *SecurityMonitoringSignalRuleTypeCreate
	isSet bool
}

// Get returns the associated value.
func (v NullableSecurityMonitoringSignalRuleTypeCreate) Get() *SecurityMonitoringSignalRuleTypeCreate {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableSecurityMonitoringSignalRuleTypeCreate) Set(val *SecurityMonitoringSignalRuleTypeCreate) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableSecurityMonitoringSignalRuleTypeCreate) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableSecurityMonitoringSignalRuleTypeCreate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSecurityMonitoringSignalRuleTypeCreate initializes the struct as if Set has been called.
func NewNullableSecurityMonitoringSignalRuleTypeCreate(val *SecurityMonitoringSignalRuleTypeCreate) *NullableSecurityMonitoringSignalRuleTypeCreate {
	return &NullableSecurityMonitoringSignalRuleTypeCreate{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableSecurityMonitoringSignalRuleTypeCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableSecurityMonitoringSignalRuleTypeCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
