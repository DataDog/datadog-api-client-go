// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateAppsDatastoreFromImportRequestDataType Datastores resource type.
type CreateAppsDatastoreFromImportRequestDataType string

// List of CreateAppsDatastoreFromImportRequestDataType.
const (
	CREATEAPPSDATASTOREFROMIMPORTREQUESTDATATYPE_DATASTORES CreateAppsDatastoreFromImportRequestDataType = "datastores"
)

var allowedCreateAppsDatastoreFromImportRequestDataTypeEnumValues = []CreateAppsDatastoreFromImportRequestDataType{
	CREATEAPPSDATASTOREFROMIMPORTREQUESTDATATYPE_DATASTORES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CreateAppsDatastoreFromImportRequestDataType) GetAllowedValues() []CreateAppsDatastoreFromImportRequestDataType {
	return allowedCreateAppsDatastoreFromImportRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CreateAppsDatastoreFromImportRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CreateAppsDatastoreFromImportRequestDataType(value)
	return nil
}

// NewCreateAppsDatastoreFromImportRequestDataTypeFromValue returns a pointer to a valid CreateAppsDatastoreFromImportRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCreateAppsDatastoreFromImportRequestDataTypeFromValue(v string) (*CreateAppsDatastoreFromImportRequestDataType, error) {
	ev := CreateAppsDatastoreFromImportRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CreateAppsDatastoreFromImportRequestDataType: valid values are %v", v, allowedCreateAppsDatastoreFromImportRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CreateAppsDatastoreFromImportRequestDataType) IsValid() bool {
	for _, existing := range allowedCreateAppsDatastoreFromImportRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateAppsDatastoreFromImportRequestDataType value.
func (v CreateAppsDatastoreFromImportRequestDataType) Ptr() *CreateAppsDatastoreFromImportRequestDataType {
	return &v
}
