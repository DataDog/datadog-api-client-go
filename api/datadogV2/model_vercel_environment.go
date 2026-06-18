// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// VercelEnvironment Vercel deployment environment.
type VercelEnvironment string

// List of VercelEnvironment.
const (
	VERCELENVIRONMENT_PRODUCTION VercelEnvironment = "production"
	VERCELENVIRONMENT_PREVIEW    VercelEnvironment = "preview"
)

var allowedVercelEnvironmentEnumValues = []VercelEnvironment{
	VERCELENVIRONMENT_PRODUCTION,
	VERCELENVIRONMENT_PREVIEW,
}

// GetAllowedValues reeturns the list of possible values.
func (v *VercelEnvironment) GetAllowedValues() []VercelEnvironment {
	return allowedVercelEnvironmentEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *VercelEnvironment) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = VercelEnvironment(value)
	return nil
}

// NewVercelEnvironmentFromValue returns a pointer to a valid VercelEnvironment
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewVercelEnvironmentFromValue(v string) (*VercelEnvironment, error) {
	ev := VercelEnvironment(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for VercelEnvironment: valid values are %v", v, allowedVercelEnvironmentEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v VercelEnvironment) IsValid() bool {
	for _, existing := range allowedVercelEnvironmentEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VercelEnvironment value.
func (v VercelEnvironment) Ptr() *VercelEnvironment {
	return &v
}
