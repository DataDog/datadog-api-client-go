// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateAppsDatastoreRequestDataType Datastores resource type.
type UpdateAppsDatastoreRequestDataType string

// List of UpdateAppsDatastoreRequestDataType.
const (
	UPDATEAPPSDATASTOREREQUESTDATATYPE_DATASTORES UpdateAppsDatastoreRequestDataType = "datastores"
)

var allowedUpdateAppsDatastoreRequestDataTypeEnumValues = []UpdateAppsDatastoreRequestDataType{
	UPDATEAPPSDATASTOREREQUESTDATATYPE_DATASTORES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UpdateAppsDatastoreRequestDataType) GetAllowedValues() []UpdateAppsDatastoreRequestDataType {
	return allowedUpdateAppsDatastoreRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UpdateAppsDatastoreRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UpdateAppsDatastoreRequestDataType(value)
	return nil
}

// NewUpdateAppsDatastoreRequestDataTypeFromValue returns a pointer to a valid UpdateAppsDatastoreRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUpdateAppsDatastoreRequestDataTypeFromValue(v string) (*UpdateAppsDatastoreRequestDataType, error) {
	ev := UpdateAppsDatastoreRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UpdateAppsDatastoreRequestDataType: valid values are %v", v, allowedUpdateAppsDatastoreRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UpdateAppsDatastoreRequestDataType) IsValid() bool {
	for _, existing := range allowedUpdateAppsDatastoreRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UpdateAppsDatastoreRequestDataType value.
func (v UpdateAppsDatastoreRequestDataType) Ptr() *UpdateAppsDatastoreRequestDataType {
	return &v
}
