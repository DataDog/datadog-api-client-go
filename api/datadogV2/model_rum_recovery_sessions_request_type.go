// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRecoverySessionsRequestType Targeted recovery. Omitted request_type defaults to sessions.
type RumRecoverySessionsRequestType string

// List of RumRecoverySessionsRequestType.
const (
	RUMRECOVERYSESSIONSREQUESTTYPE_SESSIONS RumRecoverySessionsRequestType = "sessions"
)

var allowedRumRecoverySessionsRequestTypeEnumValues = []RumRecoverySessionsRequestType{
	RUMRECOVERYSESSIONSREQUESTTYPE_SESSIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumRecoverySessionsRequestType) GetAllowedValues() []RumRecoverySessionsRequestType {
	return allowedRumRecoverySessionsRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumRecoverySessionsRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumRecoverySessionsRequestType(value)
	return nil
}

// NewRumRecoverySessionsRequestTypeFromValue returns a pointer to a valid RumRecoverySessionsRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumRecoverySessionsRequestTypeFromValue(v string) (*RumRecoverySessionsRequestType, error) {
	ev := RumRecoverySessionsRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumRecoverySessionsRequestType: valid values are %v", v, allowedRumRecoverySessionsRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumRecoverySessionsRequestType) IsValid() bool {
	for _, existing := range allowedRumRecoverySessionsRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumRecoverySessionsRequestType value.
func (v RumRecoverySessionsRequestType) Ptr() *RumRecoverySessionsRequestType {
	return &v
}
