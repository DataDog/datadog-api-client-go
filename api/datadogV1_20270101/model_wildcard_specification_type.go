// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WildcardSpecificationType Test-only generated schema for the dashboard versioning proof of concept.
type WildcardSpecificationType string

// List of WildcardSpecificationType.
const (
	WILDCARDSPECIFICATIONTYPE_VEGA      WildcardSpecificationType = "vega"
	WILDCARDSPECIFICATIONTYPE_VEGA_LITE WildcardSpecificationType = "vega-lite"
)

var allowedWildcardSpecificationTypeEnumValues = []WildcardSpecificationType{
	WILDCARDSPECIFICATIONTYPE_VEGA,
	WILDCARDSPECIFICATIONTYPE_VEGA_LITE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WildcardSpecificationType) GetAllowedValues() []WildcardSpecificationType {
	return allowedWildcardSpecificationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WildcardSpecificationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WildcardSpecificationType(value)
	return nil
}

// NewWildcardSpecificationTypeFromValue returns a pointer to a valid WildcardSpecificationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWildcardSpecificationTypeFromValue(v string) (*WildcardSpecificationType, error) {
	ev := WildcardSpecificationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WildcardSpecificationType: valid values are %v", v, allowedWildcardSpecificationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WildcardSpecificationType) IsValid() bool {
	for _, existing := range allowedWildcardSpecificationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WildcardSpecificationType value.
func (v WildcardSpecificationType) Ptr() *WildcardSpecificationType {
	return &v
}
