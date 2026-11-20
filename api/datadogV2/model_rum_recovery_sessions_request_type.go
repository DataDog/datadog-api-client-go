// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMRecoverySessionsRequestType Targeted recovery. Omitted request_type defaults to sessions.
type RUMRecoverySessionsRequestType string

// List of RUMRecoverySessionsRequestType.
const (
	RUMRECOVERYSESSIONSREQUESTTYPE_SESSIONS RUMRecoverySessionsRequestType = "sessions"
)

var allowedRUMRecoverySessionsRequestTypeEnumValues = []RUMRecoverySessionsRequestType{
	RUMRECOVERYSESSIONSREQUESTTYPE_SESSIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMRecoverySessionsRequestType) GetAllowedValues() []RUMRecoverySessionsRequestType {
	return allowedRUMRecoverySessionsRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMRecoverySessionsRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMRecoverySessionsRequestType(value)
	return nil
}

// NewRUMRecoverySessionsRequestTypeFromValue returns a pointer to a valid RUMRecoverySessionsRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMRecoverySessionsRequestTypeFromValue(v string) (*RUMRecoverySessionsRequestType, error) {
	ev := RUMRecoverySessionsRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMRecoverySessionsRequestType: valid values are %v", v, allowedRUMRecoverySessionsRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMRecoverySessionsRequestType) IsValid() bool {
	for _, existing := range allowedRUMRecoverySessionsRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMRecoverySessionsRequestType value.
func (v RUMRecoverySessionsRequestType) Ptr() *RUMRecoverySessionsRequestType {
	return &v
}
