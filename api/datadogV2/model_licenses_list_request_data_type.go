// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LicensesListRequestDataType
type LicensesListRequestDataType string

// List of LicensesListRequestDataType.
const (
	LICENSESLISTREQUESTDATATYPE_LICENSEREQUEST LicensesListRequestDataType = "licenserequest"
)

var allowedLicensesListRequestDataTypeEnumValues = []LicensesListRequestDataType{
	LICENSESLISTREQUESTDATATYPE_LICENSEREQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LicensesListRequestDataType) GetAllowedValues() []LicensesListRequestDataType {
	return allowedLicensesListRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LicensesListRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LicensesListRequestDataType(value)
	return nil
}

// NewLicensesListRequestDataTypeFromValue returns a pointer to a valid LicensesListRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLicensesListRequestDataTypeFromValue(v string) (*LicensesListRequestDataType, error) {
	ev := LicensesListRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LicensesListRequestDataType: valid values are %v", v, allowedLicensesListRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LicensesListRequestDataType) IsValid() bool {
	for _, existing := range allowedLicensesListRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LicensesListRequestDataType value.
func (v LicensesListRequestDataType) Ptr() *LicensesListRequestDataType {
	return &v
}
