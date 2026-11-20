// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceMapDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type ServiceMapDefinitionType string

// List of ServiceMapDefinitionType.
const (
	SERVICEMAPDEFINITIONTYPE_SERVICEMAP ServiceMapDefinitionType = "servicemap"
)

var allowedServiceMapDefinitionTypeEnumValues = []ServiceMapDefinitionType{
	SERVICEMAPDEFINITIONTYPE_SERVICEMAP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceMapDefinitionType) GetAllowedValues() []ServiceMapDefinitionType {
	return allowedServiceMapDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceMapDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceMapDefinitionType(value)
	return nil
}

// NewServiceMapDefinitionTypeFromValue returns a pointer to a valid ServiceMapDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceMapDefinitionTypeFromValue(v string) (*ServiceMapDefinitionType, error) {
	ev := ServiceMapDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceMapDefinitionType: valid values are %v", v, allowedServiceMapDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceMapDefinitionType) IsValid() bool {
	for _, existing := range allowedServiceMapDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceMapDefinitionType value.
func (v ServiceMapDefinitionType) Ptr() *ServiceMapDefinitionType {
	return &v
}
