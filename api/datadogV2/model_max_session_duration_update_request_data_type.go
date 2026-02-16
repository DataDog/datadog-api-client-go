// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MaxSessionDurationUpdateRequestDataType Type of the resource.
type MaxSessionDurationUpdateRequestDataType string

// List of MaxSessionDurationUpdateRequestDataType.
const (
	MAXSESSIONDURATIONUPDATEREQUESTDATATYPE_MAX_SESSION_DURATION MaxSessionDurationUpdateRequestDataType = "max_session_duration"
)

var allowedMaxSessionDurationUpdateRequestDataTypeEnumValues = []MaxSessionDurationUpdateRequestDataType{
	MAXSESSIONDURATIONUPDATEREQUESTDATATYPE_MAX_SESSION_DURATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MaxSessionDurationUpdateRequestDataType) GetAllowedValues() []MaxSessionDurationUpdateRequestDataType {
	return allowedMaxSessionDurationUpdateRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MaxSessionDurationUpdateRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MaxSessionDurationUpdateRequestDataType(value)
	return nil
}

// NewMaxSessionDurationUpdateRequestDataTypeFromValue returns a pointer to a valid MaxSessionDurationUpdateRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMaxSessionDurationUpdateRequestDataTypeFromValue(v string) (*MaxSessionDurationUpdateRequestDataType, error) {
	ev := MaxSessionDurationUpdateRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MaxSessionDurationUpdateRequestDataType: valid values are %v", v, allowedMaxSessionDurationUpdateRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MaxSessionDurationUpdateRequestDataType) IsValid() bool {
	for _, existing := range allowedMaxSessionDurationUpdateRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MaxSessionDurationUpdateRequestDataType value.
func (v MaxSessionDurationUpdateRequestDataType) Ptr() *MaxSessionDurationUpdateRequestDataType {
	return &v
}
