// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AlertValueDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type AlertValueDefinitionType string

// List of AlertValueDefinitionType.
const (
	ALERTVALUEDEFINITIONTYPE_ALERT_VALUE AlertValueDefinitionType = "alert_value"
)

var allowedAlertValueDefinitionTypeEnumValues = []AlertValueDefinitionType{
	ALERTVALUEDEFINITIONTYPE_ALERT_VALUE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AlertValueDefinitionType) GetAllowedValues() []AlertValueDefinitionType {
	return allowedAlertValueDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AlertValueDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AlertValueDefinitionType(value)
	return nil
}

// NewAlertValueDefinitionTypeFromValue returns a pointer to a valid AlertValueDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAlertValueDefinitionTypeFromValue(v string) (*AlertValueDefinitionType, error) {
	ev := AlertValueDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AlertValueDefinitionType: valid values are %v", v, allowedAlertValueDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AlertValueDefinitionType) IsValid() bool {
	for _, existing := range allowedAlertValueDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlertValueDefinitionType value.
func (v AlertValueDefinitionType) Ptr() *AlertValueDefinitionType {
	return &v
}
