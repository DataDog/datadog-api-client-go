// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SetupRulesRequestDataType
type SetupRulesRequestDataType string

// List of SetupRulesRequestDataType.
const (
	SETUPRULESREQUESTDATATYPE_SETUP SetupRulesRequestDataType = "setup"
)

var allowedSetupRulesRequestDataTypeEnumValues = []SetupRulesRequestDataType{
	SETUPRULESREQUESTDATATYPE_SETUP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SetupRulesRequestDataType) GetAllowedValues() []SetupRulesRequestDataType {
	return allowedSetupRulesRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SetupRulesRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SetupRulesRequestDataType(value)
	return nil
}

// NewSetupRulesRequestDataTypeFromValue returns a pointer to a valid SetupRulesRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSetupRulesRequestDataTypeFromValue(v string) (*SetupRulesRequestDataType, error) {
	ev := SetupRulesRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SetupRulesRequestDataType: valid values are %v", v, allowedSetupRulesRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SetupRulesRequestDataType) IsValid() bool {
	for _, existing := range allowedSetupRulesRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SetupRulesRequestDataType value.
func (v SetupRulesRequestDataType) Ptr() *SetupRulesRequestDataType {
	return &v
}
