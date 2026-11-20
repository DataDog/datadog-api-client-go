// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExcludeNullTableColumnFunctionName Test-only generated schema for the dashboard versioning proof of concept.
type ExcludeNullTableColumnFunctionName string

// List of ExcludeNullTableColumnFunctionName.
const (
	EXCLUDENULLTABLECOLUMNFUNCTIONNAME_EXCLUDE_NULL ExcludeNullTableColumnFunctionName = "exclude_null"
)

var allowedExcludeNullTableColumnFunctionNameEnumValues = []ExcludeNullTableColumnFunctionName{
	EXCLUDENULLTABLECOLUMNFUNCTIONNAME_EXCLUDE_NULL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExcludeNullTableColumnFunctionName) GetAllowedValues() []ExcludeNullTableColumnFunctionName {
	return allowedExcludeNullTableColumnFunctionNameEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExcludeNullTableColumnFunctionName) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExcludeNullTableColumnFunctionName(value)
	return nil
}

// NewExcludeNullTableColumnFunctionNameFromValue returns a pointer to a valid ExcludeNullTableColumnFunctionName
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExcludeNullTableColumnFunctionNameFromValue(v string) (*ExcludeNullTableColumnFunctionName, error) {
	ev := ExcludeNullTableColumnFunctionName(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExcludeNullTableColumnFunctionName: valid values are %v", v, allowedExcludeNullTableColumnFunctionNameEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExcludeNullTableColumnFunctionName) IsValid() bool {
	for _, existing := range allowedExcludeNullTableColumnFunctionNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExcludeNullTableColumnFunctionName value.
func (v ExcludeNullTableColumnFunctionName) Ptr() *ExcludeNullTableColumnFunctionName {
	return &v
}
