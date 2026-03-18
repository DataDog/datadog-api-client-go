// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableRequestRequestType Identifies this as a guided table request.
type GuidedTableRequestRequestType string

// List of GuidedTableRequestRequestType.
const (
	GUIDEDTABLEREQUESTREQUESTTYPE_TABLE GuidedTableRequestRequestType = "table"
)

var allowedGuidedTableRequestRequestTypeEnumValues = []GuidedTableRequestRequestType{
	GUIDEDTABLEREQUESTREQUESTTYPE_TABLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GuidedTableRequestRequestType) GetAllowedValues() []GuidedTableRequestRequestType {
	return allowedGuidedTableRequestRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GuidedTableRequestRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GuidedTableRequestRequestType(value)
	return nil
}

// NewGuidedTableRequestRequestTypeFromValue returns a pointer to a valid GuidedTableRequestRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGuidedTableRequestRequestTypeFromValue(v string) (*GuidedTableRequestRequestType, error) {
	ev := GuidedTableRequestRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GuidedTableRequestRequestType: valid values are %v", v, allowedGuidedTableRequestRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GuidedTableRequestRequestType) IsValid() bool {
	for _, existing := range allowedGuidedTableRequestRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GuidedTableRequestRequestType value.
func (v GuidedTableRequestRequestType) Ptr() *GuidedTableRequestRequestType {
	return &v
}
