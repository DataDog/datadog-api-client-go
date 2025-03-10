// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatadogAgentSourceType The type of source.
type DatadogAgentSourceType string

// List of DatadogAgentSourceType.
const (
	DATADOGAGENTSOURCETYPE_DATADOG_AGENT DatadogAgentSourceType = "datadog_agent"
)

var allowedDatadogAgentSourceTypeEnumValues = []DatadogAgentSourceType{
	DATADOGAGENTSOURCETYPE_DATADOG_AGENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DatadogAgentSourceType) GetAllowedValues() []DatadogAgentSourceType {
	return allowedDatadogAgentSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DatadogAgentSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DatadogAgentSourceType(value)
	return nil
}

// NewDatadogAgentSourceTypeFromValue returns a pointer to a valid DatadogAgentSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDatadogAgentSourceTypeFromValue(v string) (*DatadogAgentSourceType, error) {
	ev := DatadogAgentSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DatadogAgentSourceType: valid values are %v", v, allowedDatadogAgentSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DatadogAgentSourceType) IsValid() bool {
	for _, existing := range allowedDatadogAgentSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatadogAgentSourceType value.
func (v DatadogAgentSourceType) Ptr() *DatadogAgentSourceType {
	return &v
}
