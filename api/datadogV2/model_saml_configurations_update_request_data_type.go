// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SamlConfigurationsUpdateRequestDataType Type of the resource.
type SamlConfigurationsUpdateRequestDataType string

// List of SamlConfigurationsUpdateRequestDataType.
const (
	SAMLCONFIGURATIONSUPDATEREQUESTDATATYPE_SAML_PREFERENCES SamlConfigurationsUpdateRequestDataType = "saml_preferences"
)

var allowedSamlConfigurationsUpdateRequestDataTypeEnumValues = []SamlConfigurationsUpdateRequestDataType{
	SAMLCONFIGURATIONSUPDATEREQUESTDATATYPE_SAML_PREFERENCES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SamlConfigurationsUpdateRequestDataType) GetAllowedValues() []SamlConfigurationsUpdateRequestDataType {
	return allowedSamlConfigurationsUpdateRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SamlConfigurationsUpdateRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SamlConfigurationsUpdateRequestDataType(value)
	return nil
}

// NewSamlConfigurationsUpdateRequestDataTypeFromValue returns a pointer to a valid SamlConfigurationsUpdateRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSamlConfigurationsUpdateRequestDataTypeFromValue(v string) (*SamlConfigurationsUpdateRequestDataType, error) {
	ev := SamlConfigurationsUpdateRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SamlConfigurationsUpdateRequestDataType: valid values are %v", v, allowedSamlConfigurationsUpdateRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SamlConfigurationsUpdateRequestDataType) IsValid() bool {
	for _, existing := range allowedSamlConfigurationsUpdateRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SamlConfigurationsUpdateRequestDataType value.
func (v SamlConfigurationsUpdateRequestDataType) Ptr() *SamlConfigurationsUpdateRequestDataType {
	return &v
}
