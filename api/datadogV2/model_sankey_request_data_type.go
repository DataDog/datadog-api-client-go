// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SankeyRequestDataType
type SankeyRequestDataType string

// List of SankeyRequestDataType.
const (
	SANKEYREQUESTDATATYPE_SANKEY_REQUEST SankeyRequestDataType = "sankey_request"
)

var allowedSankeyRequestDataTypeEnumValues = []SankeyRequestDataType{
	SANKEYREQUESTDATATYPE_SANKEY_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SankeyRequestDataType) GetAllowedValues() []SankeyRequestDataType {
	return allowedSankeyRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SankeyRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SankeyRequestDataType(value)
	return nil
}

// NewSankeyRequestDataTypeFromValue returns a pointer to a valid SankeyRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSankeyRequestDataTypeFromValue(v string) (*SankeyRequestDataType, error) {
	ev := SankeyRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SankeyRequestDataType: valid values are %v", v, allowedSankeyRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SankeyRequestDataType) IsValid() bool {
	for _, existing := range allowedSankeyRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SankeyRequestDataType value.
func (v SankeyRequestDataType) Ptr() *SankeyRequestDataType {
	return &v
}
