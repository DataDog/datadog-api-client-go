// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMOperationStrongLinkType The JSON:API type for RUM operation strong link resources.
type RUMOperationStrongLinkType string

// List of RUMOperationStrongLinkType.
const (
	RUMOPERATIONSTRONGLINKTYPE_STRONG_LINKS RUMOperationStrongLinkType = "strong_links"
)

var allowedRUMOperationStrongLinkTypeEnumValues = []RUMOperationStrongLinkType{
	RUMOPERATIONSTRONGLINKTYPE_STRONG_LINKS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMOperationStrongLinkType) GetAllowedValues() []RUMOperationStrongLinkType {
	return allowedRUMOperationStrongLinkTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMOperationStrongLinkType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMOperationStrongLinkType(value)
	return nil
}

// NewRUMOperationStrongLinkTypeFromValue returns a pointer to a valid RUMOperationStrongLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMOperationStrongLinkTypeFromValue(v string) (*RUMOperationStrongLinkType, error) {
	ev := RUMOperationStrongLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMOperationStrongLinkType: valid values are %v", v, allowedRUMOperationStrongLinkTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMOperationStrongLinkType) IsValid() bool {
	for _, existing := range allowedRUMOperationStrongLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMOperationStrongLinkType value.
func (v RUMOperationStrongLinkType) Ptr() *RUMOperationStrongLinkType {
	return &v
}
