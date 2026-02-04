// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnPremManagementServiceCreateEnrollmentResponseType The type of the resource. The value should always be createEnrollmentResponse.
type OnPremManagementServiceCreateEnrollmentResponseType string

// List of OnPremManagementServiceCreateEnrollmentResponseType.
const (
	ONPREMMANAGEMENTSERVICECREATEENROLLMENTRESPONSETYPE_CREATEENROLLMENTRESPONSE OnPremManagementServiceCreateEnrollmentResponseType = "createEnrollmentResponse"
)

var allowedOnPremManagementServiceCreateEnrollmentResponseTypeEnumValues = []OnPremManagementServiceCreateEnrollmentResponseType{
	ONPREMMANAGEMENTSERVICECREATEENROLLMENTRESPONSETYPE_CREATEENROLLMENTRESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OnPremManagementServiceCreateEnrollmentResponseType) GetAllowedValues() []OnPremManagementServiceCreateEnrollmentResponseType {
	return allowedOnPremManagementServiceCreateEnrollmentResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OnPremManagementServiceCreateEnrollmentResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OnPremManagementServiceCreateEnrollmentResponseType(value)
	return nil
}

// NewOnPremManagementServiceCreateEnrollmentResponseTypeFromValue returns a pointer to a valid OnPremManagementServiceCreateEnrollmentResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOnPremManagementServiceCreateEnrollmentResponseTypeFromValue(v string) (*OnPremManagementServiceCreateEnrollmentResponseType, error) {
	ev := OnPremManagementServiceCreateEnrollmentResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OnPremManagementServiceCreateEnrollmentResponseType: valid values are %v", v, allowedOnPremManagementServiceCreateEnrollmentResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OnPremManagementServiceCreateEnrollmentResponseType) IsValid() bool {
	for _, existing := range allowedOnPremManagementServiceCreateEnrollmentResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnPremManagementServiceCreateEnrollmentResponseType value.
func (v OnPremManagementServiceCreateEnrollmentResponseType) Ptr() *OnPremManagementServiceCreateEnrollmentResponseType {
	return &v
}
