// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreemapBackendDefinitionGroupBy Test-only generated schema for the dashboard versioning proof of concept.
type TreemapBackendDefinitionGroupBy string

// List of TreemapBackendDefinitionGroupBy.
const (
	TREEMAPBACKENDDEFINITIONGROUPBY_PROCESS TreemapBackendDefinitionGroupBy = "process"
	TREEMAPBACKENDDEFINITIONGROUPBY_USER    TreemapBackendDefinitionGroupBy = "user"
	TREEMAPBACKENDDEFINITIONGROUPBY_FAMILY  TreemapBackendDefinitionGroupBy = "family"
)

var allowedTreemapBackendDefinitionGroupByEnumValues = []TreemapBackendDefinitionGroupBy{
	TREEMAPBACKENDDEFINITIONGROUPBY_PROCESS,
	TREEMAPBACKENDDEFINITIONGROUPBY_USER,
	TREEMAPBACKENDDEFINITIONGROUPBY_FAMILY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TreemapBackendDefinitionGroupBy) GetAllowedValues() []TreemapBackendDefinitionGroupBy {
	return allowedTreemapBackendDefinitionGroupByEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TreemapBackendDefinitionGroupBy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TreemapBackendDefinitionGroupBy(value)
	return nil
}

// NewTreemapBackendDefinitionGroupByFromValue returns a pointer to a valid TreemapBackendDefinitionGroupBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTreemapBackendDefinitionGroupByFromValue(v string) (*TreemapBackendDefinitionGroupBy, error) {
	ev := TreemapBackendDefinitionGroupBy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TreemapBackendDefinitionGroupBy: valid values are %v", v, allowedTreemapBackendDefinitionGroupByEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TreemapBackendDefinitionGroupBy) IsValid() bool {
	for _, existing := range allowedTreemapBackendDefinitionGroupByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TreemapBackendDefinitionGroupBy value.
func (v TreemapBackendDefinitionGroupBy) Ptr() *TreemapBackendDefinitionGroupBy {
	return &v
}
