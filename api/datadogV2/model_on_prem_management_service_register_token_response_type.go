// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnPremManagementServiceRegisterTokenResponseType The type of the resource. The value should always be tokenDefinitions.
type OnPremManagementServiceRegisterTokenResponseType string

// List of OnPremManagementServiceRegisterTokenResponseType.
const (
	ONPREMMANAGEMENTSERVICEREGISTERTOKENRESPONSETYPE_TOKENDEFINITIONS OnPremManagementServiceRegisterTokenResponseType = "tokenDefinitions"
)

var allowedOnPremManagementServiceRegisterTokenResponseTypeEnumValues = []OnPremManagementServiceRegisterTokenResponseType{
	ONPREMMANAGEMENTSERVICEREGISTERTOKENRESPONSETYPE_TOKENDEFINITIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OnPremManagementServiceRegisterTokenResponseType) GetAllowedValues() []OnPremManagementServiceRegisterTokenResponseType {
	return allowedOnPremManagementServiceRegisterTokenResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OnPremManagementServiceRegisterTokenResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OnPremManagementServiceRegisterTokenResponseType(value)
	return nil
}

// NewOnPremManagementServiceRegisterTokenResponseTypeFromValue returns a pointer to a valid OnPremManagementServiceRegisterTokenResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOnPremManagementServiceRegisterTokenResponseTypeFromValue(v string) (*OnPremManagementServiceRegisterTokenResponseType, error) {
	ev := OnPremManagementServiceRegisterTokenResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OnPremManagementServiceRegisterTokenResponseType: valid values are %v", v, allowedOnPremManagementServiceRegisterTokenResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OnPremManagementServiceRegisterTokenResponseType) IsValid() bool {
	for _, existing := range allowedOnPremManagementServiceRegisterTokenResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnPremManagementServiceRegisterTokenResponseType value.
func (v OnPremManagementServiceRegisterTokenResponseType) Ptr() *OnPremManagementServiceRegisterTokenResponseType {
	return &v
}
