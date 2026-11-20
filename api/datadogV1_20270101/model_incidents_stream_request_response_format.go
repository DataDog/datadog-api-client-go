// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentsStreamRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type IncidentsStreamRequestResponseFormat string

// List of IncidentsStreamRequestResponseFormat.
const (
	INCIDENTSSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST IncidentsStreamRequestResponseFormat = "event_list"
)

var allowedIncidentsStreamRequestResponseFormatEnumValues = []IncidentsStreamRequestResponseFormat{
	INCIDENTSSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentsStreamRequestResponseFormat) GetAllowedValues() []IncidentsStreamRequestResponseFormat {
	return allowedIncidentsStreamRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentsStreamRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentsStreamRequestResponseFormat(value)
	return nil
}

// NewIncidentsStreamRequestResponseFormatFromValue returns a pointer to a valid IncidentsStreamRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentsStreamRequestResponseFormatFromValue(v string) (*IncidentsStreamRequestResponseFormat, error) {
	ev := IncidentsStreamRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentsStreamRequestResponseFormat: valid values are %v", v, allowedIncidentsStreamRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentsStreamRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedIncidentsStreamRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentsStreamRequestResponseFormat value.
func (v IncidentsStreamRequestResponseFormat) Ptr() *IncidentsStreamRequestResponseFormat {
	return &v
}
