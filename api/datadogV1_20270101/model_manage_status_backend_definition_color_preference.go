// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ManageStatusBackendDefinitionColorPreference Test-only generated schema for the dashboard versioning proof of concept.
type ManageStatusBackendDefinitionColorPreference string

// List of ManageStatusBackendDefinitionColorPreference.
const (
	MANAGESTATUSBACKENDDEFINITIONCOLORPREFERENCE_TEXT       ManageStatusBackendDefinitionColorPreference = "text"
	MANAGESTATUSBACKENDDEFINITIONCOLORPREFERENCE_BACKGROUND ManageStatusBackendDefinitionColorPreference = "background"
)

var allowedManageStatusBackendDefinitionColorPreferenceEnumValues = []ManageStatusBackendDefinitionColorPreference{
	MANAGESTATUSBACKENDDEFINITIONCOLORPREFERENCE_TEXT,
	MANAGESTATUSBACKENDDEFINITIONCOLORPREFERENCE_BACKGROUND,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ManageStatusBackendDefinitionColorPreference) GetAllowedValues() []ManageStatusBackendDefinitionColorPreference {
	return allowedManageStatusBackendDefinitionColorPreferenceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ManageStatusBackendDefinitionColorPreference) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ManageStatusBackendDefinitionColorPreference(value)
	return nil
}

// NewManageStatusBackendDefinitionColorPreferenceFromValue returns a pointer to a valid ManageStatusBackendDefinitionColorPreference
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewManageStatusBackendDefinitionColorPreferenceFromValue(v string) (*ManageStatusBackendDefinitionColorPreference, error) {
	ev := ManageStatusBackendDefinitionColorPreference(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ManageStatusBackendDefinitionColorPreference: valid values are %v", v, allowedManageStatusBackendDefinitionColorPreferenceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ManageStatusBackendDefinitionColorPreference) IsValid() bool {
	for _, existing := range allowedManageStatusBackendDefinitionColorPreferenceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ManageStatusBackendDefinitionColorPreference value.
func (v ManageStatusBackendDefinitionColorPreference) Ptr() *ManageStatusBackendDefinitionColorPreference {
	return &v
}
