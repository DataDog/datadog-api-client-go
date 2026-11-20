// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterplotBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type ScatterplotBackendDefinitionType string

// List of ScatterplotBackendDefinitionType.
const (
	SCATTERPLOTBACKENDDEFINITIONTYPE_SCATTERPLOT ScatterplotBackendDefinitionType = "scatterplot"
)

var allowedScatterplotBackendDefinitionTypeEnumValues = []ScatterplotBackendDefinitionType{
	SCATTERPLOTBACKENDDEFINITIONTYPE_SCATTERPLOT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScatterplotBackendDefinitionType) GetAllowedValues() []ScatterplotBackendDefinitionType {
	return allowedScatterplotBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScatterplotBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScatterplotBackendDefinitionType(value)
	return nil
}

// NewScatterplotBackendDefinitionTypeFromValue returns a pointer to a valid ScatterplotBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScatterplotBackendDefinitionTypeFromValue(v string) (*ScatterplotBackendDefinitionType, error) {
	ev := ScatterplotBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScatterplotBackendDefinitionType: valid values are %v", v, allowedScatterplotBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScatterplotBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedScatterplotBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScatterplotBackendDefinitionType value.
func (v ScatterplotBackendDefinitionType) Ptr() *ScatterplotBackendDefinitionType {
	return &v
}
