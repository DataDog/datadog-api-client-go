// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageServiceNowIncidentType ServiceNow incident resource type.
type OnCallPageServiceNowIncidentType string

// List of OnCallPageServiceNowIncidentType.
const (
	ONCALLPAGESERVICENOWINCIDENTTYPE_SERVICE_NOW_INCIDENTS OnCallPageServiceNowIncidentType = "service_now_incidents"
)

var allowedOnCallPageServiceNowIncidentTypeEnumValues = []OnCallPageServiceNowIncidentType{
	ONCALLPAGESERVICENOWINCIDENTTYPE_SERVICE_NOW_INCIDENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OnCallPageServiceNowIncidentType) GetAllowedValues() []OnCallPageServiceNowIncidentType {
	return allowedOnCallPageServiceNowIncidentTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OnCallPageServiceNowIncidentType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OnCallPageServiceNowIncidentType(value)
	return nil
}

// NewOnCallPageServiceNowIncidentTypeFromValue returns a pointer to a valid OnCallPageServiceNowIncidentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOnCallPageServiceNowIncidentTypeFromValue(v string) (*OnCallPageServiceNowIncidentType, error) {
	ev := OnCallPageServiceNowIncidentType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OnCallPageServiceNowIncidentType: valid values are %v", v, allowedOnCallPageServiceNowIncidentTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OnCallPageServiceNowIncidentType) IsValid() bool {
	for _, existing := range allowedOnCallPageServiceNowIncidentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnCallPageServiceNowIncidentType value.
func (v OnCallPageServiceNowIncidentType) Ptr() *OnCallPageServiceNowIncidentType {
	return &v
}
