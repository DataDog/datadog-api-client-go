// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PutAppsDatastoreItemRequestDataType Items resource type.
type PutAppsDatastoreItemRequestDataType string

// List of PutAppsDatastoreItemRequestDataType.
const (
	PUTAPPSDATASTOREITEMREQUESTDATATYPE_ITEMS PutAppsDatastoreItemRequestDataType = "items"
)

var allowedPutAppsDatastoreItemRequestDataTypeEnumValues = []PutAppsDatastoreItemRequestDataType{
	PUTAPPSDATASTOREITEMREQUESTDATATYPE_ITEMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PutAppsDatastoreItemRequestDataType) GetAllowedValues() []PutAppsDatastoreItemRequestDataType {
	return allowedPutAppsDatastoreItemRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PutAppsDatastoreItemRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PutAppsDatastoreItemRequestDataType(value)
	return nil
}

// NewPutAppsDatastoreItemRequestDataTypeFromValue returns a pointer to a valid PutAppsDatastoreItemRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPutAppsDatastoreItemRequestDataTypeFromValue(v string) (*PutAppsDatastoreItemRequestDataType, error) {
	ev := PutAppsDatastoreItemRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PutAppsDatastoreItemRequestDataType: valid values are %v", v, allowedPutAppsDatastoreItemRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PutAppsDatastoreItemRequestDataType) IsValid() bool {
	for _, existing := range allowedPutAppsDatastoreItemRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PutAppsDatastoreItemRequestDataType value.
func (v PutAppsDatastoreItemRequestDataType) Ptr() *PutAppsDatastoreItemRequestDataType {
	return &v
}
