// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateAppsDatastoreResponseDataType Datastores resource type.
type CreateAppsDatastoreResponseDataType string

// List of CreateAppsDatastoreResponseDataType.
const (
	CREATEAPPSDATASTORERESPONSEDATATYPE_DATASTORES CreateAppsDatastoreResponseDataType = "datastores"
)

var allowedCreateAppsDatastoreResponseDataTypeEnumValues = []CreateAppsDatastoreResponseDataType{
	CREATEAPPSDATASTORERESPONSEDATATYPE_DATASTORES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CreateAppsDatastoreResponseDataType) GetAllowedValues() []CreateAppsDatastoreResponseDataType {
	return allowedCreateAppsDatastoreResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CreateAppsDatastoreResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CreateAppsDatastoreResponseDataType(value)
	return nil
}

// NewCreateAppsDatastoreResponseDataTypeFromValue returns a pointer to a valid CreateAppsDatastoreResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCreateAppsDatastoreResponseDataTypeFromValue(v string) (*CreateAppsDatastoreResponseDataType, error) {
	ev := CreateAppsDatastoreResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CreateAppsDatastoreResponseDataType: valid values are %v", v, allowedCreateAppsDatastoreResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CreateAppsDatastoreResponseDataType) IsValid() bool {
	for _, existing := range allowedCreateAppsDatastoreResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateAppsDatastoreResponseDataType value.
func (v CreateAppsDatastoreResponseDataType) Ptr() *CreateAppsDatastoreResponseDataType {
	return &v
}
