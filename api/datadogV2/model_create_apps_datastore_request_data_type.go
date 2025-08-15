// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateAppsDatastoreRequestDataType Datastores resource type.
type CreateAppsDatastoreRequestDataType string

// List of CreateAppsDatastoreRequestDataType.
const (
	CREATEAPPSDATASTOREREQUESTDATATYPE_DATASTORES CreateAppsDatastoreRequestDataType = "datastores"
)

var allowedCreateAppsDatastoreRequestDataTypeEnumValues = []CreateAppsDatastoreRequestDataType{
	CREATEAPPSDATASTOREREQUESTDATATYPE_DATASTORES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CreateAppsDatastoreRequestDataType) GetAllowedValues() []CreateAppsDatastoreRequestDataType {
	return allowedCreateAppsDatastoreRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CreateAppsDatastoreRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CreateAppsDatastoreRequestDataType(value)
	return nil
}

// NewCreateAppsDatastoreRequestDataTypeFromValue returns a pointer to a valid CreateAppsDatastoreRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCreateAppsDatastoreRequestDataTypeFromValue(v string) (*CreateAppsDatastoreRequestDataType, error) {
	ev := CreateAppsDatastoreRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CreateAppsDatastoreRequestDataType: valid values are %v", v, allowedCreateAppsDatastoreRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CreateAppsDatastoreRequestDataType) IsValid() bool {
	for _, existing := range allowedCreateAppsDatastoreRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateAppsDatastoreRequestDataType value.
func (v CreateAppsDatastoreRequestDataType) Ptr() *CreateAppsDatastoreRequestDataType {
	return &v
}
