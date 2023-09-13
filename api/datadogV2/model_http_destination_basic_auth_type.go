// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"
)

// HttpDestinationBasicAuthType The HTTP destination basic auth type.
type HttpDestinationBasicAuthType string

// List of HttpDestinationBasicAuthType.
const (
	HTTPDESTINATIONBASICAUTHTYPE_BASIC HttpDestinationBasicAuthType = "basic"
)

var allowedHttpDestinationBasicAuthTypeEnumValues = []HttpDestinationBasicAuthType{
	HTTPDESTINATIONBASICAUTHTYPE_BASIC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HttpDestinationBasicAuthType) GetAllowedValues() []HttpDestinationBasicAuthType {
	return allowedHttpDestinationBasicAuthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HttpDestinationBasicAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HttpDestinationBasicAuthType(value)
	return nil
}

// NewHttpDestinationBasicAuthTypeFromValue returns a pointer to a valid HttpDestinationBasicAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHttpDestinationBasicAuthTypeFromValue(v string) (*HttpDestinationBasicAuthType, error) {
	ev := HttpDestinationBasicAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HttpDestinationBasicAuthType: valid values are %v", v, allowedHttpDestinationBasicAuthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HttpDestinationBasicAuthType) IsValid() bool {
	for _, existing := range allowedHttpDestinationBasicAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HttpDestinationBasicAuthType value.
func (v HttpDestinationBasicAuthType) Ptr() *HttpDestinationBasicAuthType {
	return &v
}
