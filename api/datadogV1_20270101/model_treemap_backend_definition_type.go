// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreemapBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type TreemapBackendDefinitionType string

// List of TreemapBackendDefinitionType.
const (
	TREEMAPBACKENDDEFINITIONTYPE_TREEMAP TreemapBackendDefinitionType = "treemap"
)

var allowedTreemapBackendDefinitionTypeEnumValues = []TreemapBackendDefinitionType{
	TREEMAPBACKENDDEFINITIONTYPE_TREEMAP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TreemapBackendDefinitionType) GetAllowedValues() []TreemapBackendDefinitionType {
	return allowedTreemapBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TreemapBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TreemapBackendDefinitionType(value)
	return nil
}

// NewTreemapBackendDefinitionTypeFromValue returns a pointer to a valid TreemapBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTreemapBackendDefinitionTypeFromValue(v string) (*TreemapBackendDefinitionType, error) {
	ev := TreemapBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TreemapBackendDefinitionType: valid values are %v", v, allowedTreemapBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TreemapBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedTreemapBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TreemapBackendDefinitionType value.
func (v TreemapBackendDefinitionType) Ptr() *TreemapBackendDefinitionType {
	return &v
}
