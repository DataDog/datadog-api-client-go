// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnPremManagementServiceRegisterTokenType The type of the resource. The value should always be registerTokenRequest.
type OnPremManagementServiceRegisterTokenType string

// List of OnPremManagementServiceRegisterTokenType.
const (
	ONPREMMANAGEMENTSERVICEREGISTERTOKENTYPE_REGISTERTOKENREQUEST OnPremManagementServiceRegisterTokenType = "registerTokenRequest"
)

var allowedOnPremManagementServiceRegisterTokenTypeEnumValues = []OnPremManagementServiceRegisterTokenType{
	ONPREMMANAGEMENTSERVICEREGISTERTOKENTYPE_REGISTERTOKENREQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OnPremManagementServiceRegisterTokenType) GetAllowedValues() []OnPremManagementServiceRegisterTokenType {
	return allowedOnPremManagementServiceRegisterTokenTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OnPremManagementServiceRegisterTokenType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OnPremManagementServiceRegisterTokenType(value)
	return nil
}

// NewOnPremManagementServiceRegisterTokenTypeFromValue returns a pointer to a valid OnPremManagementServiceRegisterTokenType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOnPremManagementServiceRegisterTokenTypeFromValue(v string) (*OnPremManagementServiceRegisterTokenType, error) {
	ev := OnPremManagementServiceRegisterTokenType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OnPremManagementServiceRegisterTokenType: valid values are %v", v, allowedOnPremManagementServiceRegisterTokenTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OnPremManagementServiceRegisterTokenType) IsValid() bool {
	for _, existing := range allowedOnPremManagementServiceRegisterTokenTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnPremManagementServiceRegisterTokenType value.
func (v OnPremManagementServiceRegisterTokenType) Ptr() *OnPremManagementServiceRegisterTokenType {
	return &v
}
