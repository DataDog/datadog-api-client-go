// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IDPConfigType Resource type.
type IDPConfigType string

// List of IDPConfigType.
const (
	IDPCONFIGTYPE_IDP_CONFIG IDPConfigType = "idp_config"
)

var allowedIDPConfigTypeEnumValues = []IDPConfigType{
	IDPCONFIGTYPE_IDP_CONFIG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IDPConfigType) GetAllowedValues() []IDPConfigType {
	return allowedIDPConfigTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IDPConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IDPConfigType(value)
	return nil
}

// NewIDPConfigTypeFromValue returns a pointer to a valid IDPConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIDPConfigTypeFromValue(v string) (*IDPConfigType, error) {
	ev := IDPConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IDPConfigType: valid values are %v", v, allowedIDPConfigTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IDPConfigType) IsValid() bool {
	for _, existing := range allowedIDPConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IDPConfigType value.
func (v IDPConfigType) Ptr() *IDPConfigType {
	return &v
}
