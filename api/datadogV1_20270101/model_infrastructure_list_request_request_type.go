// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// InfrastructureListRequestRequestType Test-only generated schema for the dashboard versioning proof of concept.
type InfrastructureListRequestRequestType string

// List of InfrastructureListRequestRequestType.
const (
	INFRASTRUCTURELISTREQUESTREQUESTTYPE_INFRASTRUCTURE_LIST InfrastructureListRequestRequestType = "infrastructure_list"
)

var allowedInfrastructureListRequestRequestTypeEnumValues = []InfrastructureListRequestRequestType{
	INFRASTRUCTURELISTREQUESTREQUESTTYPE_INFRASTRUCTURE_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *InfrastructureListRequestRequestType) GetAllowedValues() []InfrastructureListRequestRequestType {
	return allowedInfrastructureListRequestRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *InfrastructureListRequestRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = InfrastructureListRequestRequestType(value)
	return nil
}

// NewInfrastructureListRequestRequestTypeFromValue returns a pointer to a valid InfrastructureListRequestRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewInfrastructureListRequestRequestTypeFromValue(v string) (*InfrastructureListRequestRequestType, error) {
	ev := InfrastructureListRequestRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for InfrastructureListRequestRequestType: valid values are %v", v, allowedInfrastructureListRequestRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v InfrastructureListRequestRequestType) IsValid() bool {
	for _, existing := range allowedInfrastructureListRequestRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InfrastructureListRequestRequestType value.
func (v InfrastructureListRequestRequestType) Ptr() *InfrastructureListRequestRequestType {
	return &v
}
