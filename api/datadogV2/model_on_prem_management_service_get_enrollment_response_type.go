// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnPremManagementServiceGetEnrollmentResponseType The type of the resource. The value should always be getEnrollmentResponse.
type OnPremManagementServiceGetEnrollmentResponseType string

// List of OnPremManagementServiceGetEnrollmentResponseType.
const (
	ONPREMMANAGEMENTSERVICEGETENROLLMENTRESPONSETYPE_GETENROLLMENTRESPONSE OnPremManagementServiceGetEnrollmentResponseType = "getEnrollmentResponse"
)

var allowedOnPremManagementServiceGetEnrollmentResponseTypeEnumValues = []OnPremManagementServiceGetEnrollmentResponseType{
	ONPREMMANAGEMENTSERVICEGETENROLLMENTRESPONSETYPE_GETENROLLMENTRESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OnPremManagementServiceGetEnrollmentResponseType) GetAllowedValues() []OnPremManagementServiceGetEnrollmentResponseType {
	return allowedOnPremManagementServiceGetEnrollmentResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OnPremManagementServiceGetEnrollmentResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OnPremManagementServiceGetEnrollmentResponseType(value)
	return nil
}

// NewOnPremManagementServiceGetEnrollmentResponseTypeFromValue returns a pointer to a valid OnPremManagementServiceGetEnrollmentResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOnPremManagementServiceGetEnrollmentResponseTypeFromValue(v string) (*OnPremManagementServiceGetEnrollmentResponseType, error) {
	ev := OnPremManagementServiceGetEnrollmentResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OnPremManagementServiceGetEnrollmentResponseType: valid values are %v", v, allowedOnPremManagementServiceGetEnrollmentResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OnPremManagementServiceGetEnrollmentResponseType) IsValid() bool {
	for _, existing := range allowedOnPremManagementServiceGetEnrollmentResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnPremManagementServiceGetEnrollmentResponseType value.
func (v OnPremManagementServiceGetEnrollmentResponseType) Ptr() *OnPremManagementServiceGetEnrollmentResponseType {
	return &v
}
