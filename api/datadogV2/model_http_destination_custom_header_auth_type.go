// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"
)

// HttpDestinationCustomHeaderAuthType The HTTP destination custom header auth type.
type HttpDestinationCustomHeaderAuthType string

// List of HttpDestinationCustomHeaderAuthType.
const (
	HTTPDESTINATIONCUSTOMHEADERAUTHTYPE_CUSTOM_HEADER HttpDestinationCustomHeaderAuthType = "custom_header"
)

var allowedHttpDestinationCustomHeaderAuthTypeEnumValues = []HttpDestinationCustomHeaderAuthType{
	HTTPDESTINATIONCUSTOMHEADERAUTHTYPE_CUSTOM_HEADER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HttpDestinationCustomHeaderAuthType) GetAllowedValues() []HttpDestinationCustomHeaderAuthType {
	return allowedHttpDestinationCustomHeaderAuthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HttpDestinationCustomHeaderAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HttpDestinationCustomHeaderAuthType(value)
	return nil
}

// NewHttpDestinationCustomHeaderAuthTypeFromValue returns a pointer to a valid HttpDestinationCustomHeaderAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHttpDestinationCustomHeaderAuthTypeFromValue(v string) (*HttpDestinationCustomHeaderAuthType, error) {
	ev := HttpDestinationCustomHeaderAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HttpDestinationCustomHeaderAuthType: valid values are %v", v, allowedHttpDestinationCustomHeaderAuthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HttpDestinationCustomHeaderAuthType) IsValid() bool {
	for _, existing := range allowedHttpDestinationCustomHeaderAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HttpDestinationCustomHeaderAuthType value.
func (v HttpDestinationCustomHeaderAuthType) Ptr() *HttpDestinationCustomHeaderAuthType {
	return &v
}
