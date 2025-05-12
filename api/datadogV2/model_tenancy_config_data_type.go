// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TenancyConfigDataType OCI tenancy resource type.
type TenancyConfigDataType string

// List of TenancyConfigDataType.
const (
	TENANCYCONFIGDATATYPE_OCI_TENANCY TenancyConfigDataType = "oci_tenancy"
)

var allowedTenancyConfigDataTypeEnumValues = []TenancyConfigDataType{
	TENANCYCONFIGDATATYPE_OCI_TENANCY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TenancyConfigDataType) GetAllowedValues() []TenancyConfigDataType {
	return allowedTenancyConfigDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TenancyConfigDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TenancyConfigDataType(value)
	return nil
}

// NewTenancyConfigDataTypeFromValue returns a pointer to a valid TenancyConfigDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTenancyConfigDataTypeFromValue(v string) (*TenancyConfigDataType, error) {
	ev := TenancyConfigDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TenancyConfigDataType: valid values are %v", v, allowedTenancyConfigDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TenancyConfigDataType) IsValid() bool {
	for _, existing := range allowedTenancyConfigDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TenancyConfigDataType value.
func (v TenancyConfigDataType) Ptr() *TenancyConfigDataType {
	return &v
}
