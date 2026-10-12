// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SourcemapSearchBy The search mode for listing JavaScript source maps.
type SourcemapSearchBy string

// List of SourcemapSearchBy.
const (
	SOURCEMAPSEARCHBY_DEBUG_ID SourcemapSearchBy = "debug_id"
)

var allowedSourcemapSearchByEnumValues = []SourcemapSearchBy{
	SOURCEMAPSEARCHBY_DEBUG_ID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SourcemapSearchBy) GetAllowedValues() []SourcemapSearchBy {
	return allowedSourcemapSearchByEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SourcemapSearchBy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SourcemapSearchBy(value)
	return nil
}

// NewSourcemapSearchByFromValue returns a pointer to a valid SourcemapSearchBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSourcemapSearchByFromValue(v string) (*SourcemapSearchBy, error) {
	ev := SourcemapSearchBy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SourcemapSearchBy: valid values are %v", v, allowedSourcemapSearchByEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SourcemapSearchBy) IsValid() bool {
	for _, existing := range allowedSourcemapSearchByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SourcemapSearchBy value.
func (v SourcemapSearchBy) Ptr() *SourcemapSearchBy {
	return &v
}
