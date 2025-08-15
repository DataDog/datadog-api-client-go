// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ItemApiPayloadDataType Items resource type.
type ItemApiPayloadDataType string

// List of ItemApiPayloadDataType.
const (
	ITEMAPIPAYLOADDATATYPE_ITEMS ItemApiPayloadDataType = "items"
)

var allowedItemApiPayloadDataTypeEnumValues = []ItemApiPayloadDataType{
	ITEMAPIPAYLOADDATATYPE_ITEMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ItemApiPayloadDataType) GetAllowedValues() []ItemApiPayloadDataType {
	return allowedItemApiPayloadDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ItemApiPayloadDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ItemApiPayloadDataType(value)
	return nil
}

// NewItemApiPayloadDataTypeFromValue returns a pointer to a valid ItemApiPayloadDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewItemApiPayloadDataTypeFromValue(v string) (*ItemApiPayloadDataType, error) {
	ev := ItemApiPayloadDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ItemApiPayloadDataType: valid values are %v", v, allowedItemApiPayloadDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ItemApiPayloadDataType) IsValid() bool {
	for _, existing := range allowedItemApiPayloadDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ItemApiPayloadDataType value.
func (v ItemApiPayloadDataType) Ptr() *ItemApiPayloadDataType {
	return &v
}
