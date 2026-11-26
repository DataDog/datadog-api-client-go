// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreemapBackendDefinitionColorBy Test-only generated schema for the dashboard versioning proof of concept.
type TreemapBackendDefinitionColorBy string

// List of TreemapBackendDefinitionColorBy.
const (
	TREEMAPBACKENDDEFINITIONCOLORBY_USER TreemapBackendDefinitionColorBy = "user"
)

var allowedTreemapBackendDefinitionColorByEnumValues = []TreemapBackendDefinitionColorBy{
	TREEMAPBACKENDDEFINITIONCOLORBY_USER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TreemapBackendDefinitionColorBy) GetAllowedValues() []TreemapBackendDefinitionColorBy {
	return allowedTreemapBackendDefinitionColorByEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TreemapBackendDefinitionColorBy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TreemapBackendDefinitionColorBy(value)
	return nil
}

// NewTreemapBackendDefinitionColorByFromValue returns a pointer to a valid TreemapBackendDefinitionColorBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTreemapBackendDefinitionColorByFromValue(v string) (*TreemapBackendDefinitionColorBy, error) {
	ev := TreemapBackendDefinitionColorBy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TreemapBackendDefinitionColorBy: valid values are %v", v, allowedTreemapBackendDefinitionColorByEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TreemapBackendDefinitionColorBy) IsValid() bool {
	for _, existing := range allowedTreemapBackendDefinitionColorByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TreemapBackendDefinitionColorBy value.
func (v TreemapBackendDefinitionColorBy) Ptr() *TreemapBackendDefinitionColorBy {
	return &v
}
