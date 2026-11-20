// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreemapBackendDefinitionSizeBy Test-only generated schema for the dashboard versioning proof of concept.
type TreemapBackendDefinitionSizeBy string

// List of TreemapBackendDefinitionSizeBy.
const (
	TREEMAPBACKENDDEFINITIONSIZEBY_PCT_CPU TreemapBackendDefinitionSizeBy = "pct_cpu"
	TREEMAPBACKENDDEFINITIONSIZEBY_PCT_MEM TreemapBackendDefinitionSizeBy = "pct_mem"
)

var allowedTreemapBackendDefinitionSizeByEnumValues = []TreemapBackendDefinitionSizeBy{
	TREEMAPBACKENDDEFINITIONSIZEBY_PCT_CPU,
	TREEMAPBACKENDDEFINITIONSIZEBY_PCT_MEM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TreemapBackendDefinitionSizeBy) GetAllowedValues() []TreemapBackendDefinitionSizeBy {
	return allowedTreemapBackendDefinitionSizeByEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TreemapBackendDefinitionSizeBy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TreemapBackendDefinitionSizeBy(value)
	return nil
}

// NewTreemapBackendDefinitionSizeByFromValue returns a pointer to a valid TreemapBackendDefinitionSizeBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTreemapBackendDefinitionSizeByFromValue(v string) (*TreemapBackendDefinitionSizeBy, error) {
	ev := TreemapBackendDefinitionSizeBy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TreemapBackendDefinitionSizeBy: valid values are %v", v, allowedTreemapBackendDefinitionSizeByEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TreemapBackendDefinitionSizeBy) IsValid() bool {
	for _, existing := range allowedTreemapBackendDefinitionSizeByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TreemapBackendDefinitionSizeBy value.
func (v TreemapBackendDefinitionSizeBy) Ptr() *TreemapBackendDefinitionSizeBy {
	return &v
}
