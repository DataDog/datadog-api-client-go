// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StringReferenceTableColumnType Test-only generated schema for the dashboard versioning proof of concept.
type StringReferenceTableColumnType string

// List of StringReferenceTableColumnType.
const (
	STRINGREFERENCETABLECOLUMNTYPE_REFERENCE StringReferenceTableColumnType = "reference"
)

var allowedStringReferenceTableColumnTypeEnumValues = []StringReferenceTableColumnType{
	STRINGREFERENCETABLECOLUMNTYPE_REFERENCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *StringReferenceTableColumnType) GetAllowedValues() []StringReferenceTableColumnType {
	return allowedStringReferenceTableColumnTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StringReferenceTableColumnType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StringReferenceTableColumnType(value)
	return nil
}

// NewStringReferenceTableColumnTypeFromValue returns a pointer to a valid StringReferenceTableColumnType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStringReferenceTableColumnTypeFromValue(v string) (*StringReferenceTableColumnType, error) {
	ev := StringReferenceTableColumnType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StringReferenceTableColumnType: valid values are %v", v, allowedStringReferenceTableColumnTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StringReferenceTableColumnType) IsValid() bool {
	for _, existing := range allowedStringReferenceTableColumnTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StringReferenceTableColumnType value.
func (v StringReferenceTableColumnType) Ptr() *StringReferenceTableColumnType {
	return &v
}
