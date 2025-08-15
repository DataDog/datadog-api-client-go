// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateAppsDatastoreFromImportResponseDataType Datastores resource type.
type CreateAppsDatastoreFromImportResponseDataType string

// List of CreateAppsDatastoreFromImportResponseDataType.
const (
	CREATEAPPSDATASTOREFROMIMPORTRESPONSEDATATYPE_DATASTORES CreateAppsDatastoreFromImportResponseDataType = "datastores"
)

var allowedCreateAppsDatastoreFromImportResponseDataTypeEnumValues = []CreateAppsDatastoreFromImportResponseDataType{
	CREATEAPPSDATASTOREFROMIMPORTRESPONSEDATATYPE_DATASTORES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CreateAppsDatastoreFromImportResponseDataType) GetAllowedValues() []CreateAppsDatastoreFromImportResponseDataType {
	return allowedCreateAppsDatastoreFromImportResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CreateAppsDatastoreFromImportResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CreateAppsDatastoreFromImportResponseDataType(value)
	return nil
}

// NewCreateAppsDatastoreFromImportResponseDataTypeFromValue returns a pointer to a valid CreateAppsDatastoreFromImportResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCreateAppsDatastoreFromImportResponseDataTypeFromValue(v string) (*CreateAppsDatastoreFromImportResponseDataType, error) {
	ev := CreateAppsDatastoreFromImportResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CreateAppsDatastoreFromImportResponseDataType: valid values are %v", v, allowedCreateAppsDatastoreFromImportResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CreateAppsDatastoreFromImportResponseDataType) IsValid() bool {
	for _, existing := range allowedCreateAppsDatastoreFromImportResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateAppsDatastoreFromImportResponseDataType value.
func (v CreateAppsDatastoreFromImportResponseDataType) Ptr() *CreateAppsDatastoreFromImportResponseDataType {
	return &v
}
