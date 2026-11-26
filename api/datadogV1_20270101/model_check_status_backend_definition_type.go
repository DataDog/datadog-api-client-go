// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CheckStatusBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type CheckStatusBackendDefinitionType string

// List of CheckStatusBackendDefinitionType.
const (
	CHECKSTATUSBACKENDDEFINITIONTYPE_CHECK_STATUS CheckStatusBackendDefinitionType = "check_status"
)

var allowedCheckStatusBackendDefinitionTypeEnumValues = []CheckStatusBackendDefinitionType{
	CHECKSTATUSBACKENDDEFINITIONTYPE_CHECK_STATUS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CheckStatusBackendDefinitionType) GetAllowedValues() []CheckStatusBackendDefinitionType {
	return allowedCheckStatusBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CheckStatusBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CheckStatusBackendDefinitionType(value)
	return nil
}

// NewCheckStatusBackendDefinitionTypeFromValue returns a pointer to a valid CheckStatusBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCheckStatusBackendDefinitionTypeFromValue(v string) (*CheckStatusBackendDefinitionType, error) {
	ev := CheckStatusBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CheckStatusBackendDefinitionType: valid values are %v", v, allowedCheckStatusBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CheckStatusBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedCheckStatusBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CheckStatusBackendDefinitionType value.
func (v CheckStatusBackendDefinitionType) Ptr() *CheckStatusBackendDefinitionType {
	return &v
}
