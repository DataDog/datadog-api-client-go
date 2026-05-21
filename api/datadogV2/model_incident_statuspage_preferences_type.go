// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentStatuspagePreferencesType Subscription preferences resource type.
type IncidentStatuspagePreferencesType string

// List of IncidentStatuspagePreferencesType.
const (
	INCIDENTSTATUSPAGEPREFERENCESTYPE_STATUSPAGE_SUBSCRIPTION_PREFERENCES IncidentStatuspagePreferencesType = "statuspage_subscription_preferences"
)

var allowedIncidentStatuspagePreferencesTypeEnumValues = []IncidentStatuspagePreferencesType{
	INCIDENTSTATUSPAGEPREFERENCESTYPE_STATUSPAGE_SUBSCRIPTION_PREFERENCES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentStatuspagePreferencesType) GetAllowedValues() []IncidentStatuspagePreferencesType {
	return allowedIncidentStatuspagePreferencesTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentStatuspagePreferencesType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentStatuspagePreferencesType(value)
	return nil
}

// NewIncidentStatuspagePreferencesTypeFromValue returns a pointer to a valid IncidentStatuspagePreferencesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentStatuspagePreferencesTypeFromValue(v string) (*IncidentStatuspagePreferencesType, error) {
	ev := IncidentStatuspagePreferencesType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentStatuspagePreferencesType: valid values are %v", v, allowedIncidentStatuspagePreferencesTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentStatuspagePreferencesType) IsValid() bool {
	for _, existing := range allowedIncidentStatuspagePreferencesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentStatuspagePreferencesType value.
func (v IncidentStatuspagePreferencesType) Ptr() *IncidentStatuspagePreferencesType {
	return &v
}
