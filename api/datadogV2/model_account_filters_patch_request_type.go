// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AccountFiltersPatchRequestType Type of account filters patch request.
type AccountFiltersPatchRequestType string

// List of AccountFiltersPatchRequestType.
const (
	ACCOUNTFILTERSPATCHREQUESTTYPE_ACCOUNT_FILTERS_PATCH_REQUEST AccountFiltersPatchRequestType = "account_filters_patch_request"
)

var allowedAccountFiltersPatchRequestTypeEnumValues = []AccountFiltersPatchRequestType{
	ACCOUNTFILTERSPATCHREQUESTTYPE_ACCOUNT_FILTERS_PATCH_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AccountFiltersPatchRequestType) GetAllowedValues() []AccountFiltersPatchRequestType {
	return allowedAccountFiltersPatchRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AccountFiltersPatchRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AccountFiltersPatchRequestType(value)
	return nil
}

// NewAccountFiltersPatchRequestTypeFromValue returns a pointer to a valid AccountFiltersPatchRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAccountFiltersPatchRequestTypeFromValue(v string) (*AccountFiltersPatchRequestType, error) {
	ev := AccountFiltersPatchRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AccountFiltersPatchRequestType: valid values are %v", v, allowedAccountFiltersPatchRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AccountFiltersPatchRequestType) IsValid() bool {
	for _, existing := range allowedAccountFiltersPatchRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountFiltersPatchRequestType value.
func (v AccountFiltersPatchRequestType) Ptr() *AccountFiltersPatchRequestType {
	return &v
}
