// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"
)

// HttpDestinationType The HTTP destination type.
type HttpDestinationType string

// List of HttpDestinationType.
const (
	HTTPDESTINATIONTYPE_HTTP HttpDestinationType = "http"
)

var allowedHttpDestinationTypeEnumValues = []HttpDestinationType{
	HTTPDESTINATIONTYPE_HTTP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HttpDestinationType) GetAllowedValues() []HttpDestinationType {
	return allowedHttpDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HttpDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HttpDestinationType(value)
	return nil
}

// NewHttpDestinationTypeFromValue returns a pointer to a valid HttpDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHttpDestinationTypeFromValue(v string) (*HttpDestinationType, error) {
	ev := HttpDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HttpDestinationType: valid values are %v", v, allowedHttpDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HttpDestinationType) IsValid() bool {
	for _, existing := range allowedHttpDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HttpDestinationType value.
func (v HttpDestinationType) Ptr() *HttpDestinationType {
	return &v
}
