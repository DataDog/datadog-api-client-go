// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CompositeStreamRequestRequestType Test-only generated schema for the dashboard versioning proof of concept.
type CompositeStreamRequestRequestType string

// List of CompositeStreamRequestRequestType.
const (
	COMPOSITESTREAMREQUESTREQUESTTYPE_COMPOSITE_STREAM CompositeStreamRequestRequestType = "composite_stream"
)

var allowedCompositeStreamRequestRequestTypeEnumValues = []CompositeStreamRequestRequestType{
	COMPOSITESTREAMREQUESTREQUESTTYPE_COMPOSITE_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CompositeStreamRequestRequestType) GetAllowedValues() []CompositeStreamRequestRequestType {
	return allowedCompositeStreamRequestRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CompositeStreamRequestRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CompositeStreamRequestRequestType(value)
	return nil
}

// NewCompositeStreamRequestRequestTypeFromValue returns a pointer to a valid CompositeStreamRequestRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCompositeStreamRequestRequestTypeFromValue(v string) (*CompositeStreamRequestRequestType, error) {
	ev := CompositeStreamRequestRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CompositeStreamRequestRequestType: valid values are %v", v, allowedCompositeStreamRequestRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CompositeStreamRequestRequestType) IsValid() bool {
	for _, existing := range allowedCompositeStreamRequestRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CompositeStreamRequestRequestType value.
func (v CompositeStreamRequestRequestType) Ptr() *CompositeStreamRequestRequestType {
	return &v
}
