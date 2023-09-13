// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"
)

// SplunkHecDestinationType The Splunk destination type.
type SplunkHecDestinationType string

// List of SplunkHecDestinationType.
const (
	SPLUNKHECDESTINATIONTYPE_SPLUNK_HEC SplunkHecDestinationType = "splunk_hec"
)

var allowedSplunkHecDestinationTypeEnumValues = []SplunkHecDestinationType{
	SPLUNKHECDESTINATIONTYPE_SPLUNK_HEC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SplunkHecDestinationType) GetAllowedValues() []SplunkHecDestinationType {
	return allowedSplunkHecDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SplunkHecDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SplunkHecDestinationType(value)
	return nil
}

// NewSplunkHecDestinationTypeFromValue returns a pointer to a valid SplunkHecDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSplunkHecDestinationTypeFromValue(v string) (*SplunkHecDestinationType, error) {
	ev := SplunkHecDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SplunkHecDestinationType: valid values are %v", v, allowedSplunkHecDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SplunkHecDestinationType) IsValid() bool {
	for _, existing := range allowedSplunkHecDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SplunkHecDestinationType value.
func (v SplunkHecDestinationType) Ptr() *SplunkHecDestinationType {
	return &v
}
