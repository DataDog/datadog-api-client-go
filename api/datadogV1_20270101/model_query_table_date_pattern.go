// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryTableDatePattern Test-only generated schema for the dashboard versioning proof of concept.
type QueryTableDatePattern string

// List of QueryTableDatePattern.
const (
	QUERYTABLEDATEPATTERN_MMM_DD_YYYY  QueryTableDatePattern = "MMM dd yyyy"
	QUERYTABLEDATEPATTERN_MMM_D        QueryTableDatePattern = "MMM d"
	QUERYTABLEDATEPATTERN_MMM_D__YYYY  QueryTableDatePattern = "MMM d, yyyy"
	QUERYTABLEDATEPATTERN_MM_DD_YYYY   QueryTableDatePattern = "MM/dd/yyyy"
	QUERYTABLEDATEPATTERN_DD_MM_YYYY   QueryTableDatePattern = "dd/MM/yyyy"
	QUERYTABLEDATEPATTERN_MMMM_D__YYYY QueryTableDatePattern = "MMMM d, yyyy"
)

var allowedQueryTableDatePatternEnumValues = []QueryTableDatePattern{
	QUERYTABLEDATEPATTERN_MMM_DD_YYYY,
	QUERYTABLEDATEPATTERN_MMM_D,
	QUERYTABLEDATEPATTERN_MMM_D__YYYY,
	QUERYTABLEDATEPATTERN_MM_DD_YYYY,
	QUERYTABLEDATEPATTERN_DD_MM_YYYY,
	QUERYTABLEDATEPATTERN_MMMM_D__YYYY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QueryTableDatePattern) GetAllowedValues() []QueryTableDatePattern {
	return allowedQueryTableDatePatternEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryTableDatePattern) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryTableDatePattern(value)
	return nil
}

// NewQueryTableDatePatternFromValue returns a pointer to a valid QueryTableDatePattern
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryTableDatePatternFromValue(v string) (*QueryTableDatePattern, error) {
	ev := QueryTableDatePattern(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryTableDatePattern: valid values are %v", v, allowedQueryTableDatePatternEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryTableDatePattern) IsValid() bool {
	for _, existing := range allowedQueryTableDatePatternEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryTableDatePattern value.
func (v QueryTableDatePattern) Ptr() *QueryTableDatePattern {
	return &v
}
