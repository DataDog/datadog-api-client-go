// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AccessPolicyEffect The effect applied when an email address matches the policy. Deny policies take precedence over allow policies.
type AccessPolicyEffect string

// List of AccessPolicyEffect.
const (
	ACCESSPOLICYEFFECT_ALLOW AccessPolicyEffect = "allow"
	ACCESSPOLICYEFFECT_DENY  AccessPolicyEffect = "deny"
)

var allowedAccessPolicyEffectEnumValues = []AccessPolicyEffect{
	ACCESSPOLICYEFFECT_ALLOW,
	ACCESSPOLICYEFFECT_DENY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AccessPolicyEffect) GetAllowedValues() []AccessPolicyEffect {
	return allowedAccessPolicyEffectEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AccessPolicyEffect) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AccessPolicyEffect(value)
	return nil
}

// NewAccessPolicyEffectFromValue returns a pointer to a valid AccessPolicyEffect
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAccessPolicyEffectFromValue(v string) (*AccessPolicyEffect, error) {
	ev := AccessPolicyEffect(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AccessPolicyEffect: valid values are %v", v, allowedAccessPolicyEffectEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AccessPolicyEffect) IsValid() bool {
	for _, existing := range allowedAccessPolicyEffectEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccessPolicyEffect value.
func (v AccessPolicyEffect) Ptr() *AccessPolicyEffect {
	return &v
}
