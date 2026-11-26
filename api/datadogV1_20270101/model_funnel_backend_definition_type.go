// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FunnelBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type FunnelBackendDefinitionType string

// List of FunnelBackendDefinitionType.
const (
	FUNNELBACKENDDEFINITIONTYPE_FUNNEL FunnelBackendDefinitionType = "funnel"
)

var allowedFunnelBackendDefinitionTypeEnumValues = []FunnelBackendDefinitionType{
	FUNNELBACKENDDEFINITIONTYPE_FUNNEL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FunnelBackendDefinitionType) GetAllowedValues() []FunnelBackendDefinitionType {
	return allowedFunnelBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FunnelBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FunnelBackendDefinitionType(value)
	return nil
}

// NewFunnelBackendDefinitionTypeFromValue returns a pointer to a valid FunnelBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFunnelBackendDefinitionTypeFromValue(v string) (*FunnelBackendDefinitionType, error) {
	ev := FunnelBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FunnelBackendDefinitionType: valid values are %v", v, allowedFunnelBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FunnelBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedFunnelBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FunnelBackendDefinitionType value.
func (v FunnelBackendDefinitionType) Ptr() *FunnelBackendDefinitionType {
	return &v
}
