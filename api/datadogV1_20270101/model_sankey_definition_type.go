// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SankeyDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type SankeyDefinitionType string

// List of SankeyDefinitionType.
const (
	SANKEYDEFINITIONTYPE_SANKEY SankeyDefinitionType = "sankey"
)

var allowedSankeyDefinitionTypeEnumValues = []SankeyDefinitionType{
	SANKEYDEFINITIONTYPE_SANKEY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SankeyDefinitionType) GetAllowedValues() []SankeyDefinitionType {
	return allowedSankeyDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SankeyDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SankeyDefinitionType(value)
	return nil
}

// NewSankeyDefinitionTypeFromValue returns a pointer to a valid SankeyDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSankeyDefinitionTypeFromValue(v string) (*SankeyDefinitionType, error) {
	ev := SankeyDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SankeyDefinitionType: valid values are %v", v, allowedSankeyDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SankeyDefinitionType) IsValid() bool {
	for _, existing := range allowedSankeyDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SankeyDefinitionType value.
func (v SankeyDefinitionType) Ptr() *SankeyDefinitionType {
	return &v
}
