// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeletedSuiteType
type DeletedSuiteType string

// List of DeletedSuiteType.
const (
	DELETEDSUITETYPE_SUITES DeletedSuiteType = "suites"
)

var allowedDeletedSuiteTypeEnumValues = []DeletedSuiteType{
	DELETEDSUITETYPE_SUITES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DeletedSuiteType) GetAllowedValues() []DeletedSuiteType {
	return allowedDeletedSuiteTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DeletedSuiteType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DeletedSuiteType(value)
	return nil
}

// NewDeletedSuiteTypeFromValue returns a pointer to a valid DeletedSuiteType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDeletedSuiteTypeFromValue(v string) (*DeletedSuiteType, error) {
	ev := DeletedSuiteType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DeletedSuiteType: valid values are %v", v, allowedDeletedSuiteTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DeletedSuiteType) IsValid() bool {
	for _, existing := range allowedDeletedSuiteTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeletedSuiteType value.
func (v DeletedSuiteType) Ptr() *DeletedSuiteType {
	return &v
}
