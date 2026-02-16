// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ConnectionGroupType The definition of `ConnectionGroupType` object.
type ConnectionGroupType string

// List of ConnectionGroupType.
const (
	CONNECTIONGROUPTYPE_CONNECTION_GROUP ConnectionGroupType = "connection_group"
)

var allowedConnectionGroupTypeEnumValues = []ConnectionGroupType{
	CONNECTIONGROUPTYPE_CONNECTION_GROUP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ConnectionGroupType) GetAllowedValues() []ConnectionGroupType {
	return allowedConnectionGroupTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ConnectionGroupType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ConnectionGroupType(value)
	return nil
}

// NewConnectionGroupTypeFromValue returns a pointer to a valid ConnectionGroupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewConnectionGroupTypeFromValue(v string) (*ConnectionGroupType, error) {
	ev := ConnectionGroupType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ConnectionGroupType: valid values are %v", v, allowedConnectionGroupTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ConnectionGroupType) IsValid() bool {
	for _, existing := range allowedConnectionGroupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConnectionGroupType value.
func (v ConnectionGroupType) Ptr() *ConnectionGroupType {
	return &v
}
