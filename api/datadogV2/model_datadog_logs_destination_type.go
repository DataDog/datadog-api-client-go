// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatadogLogsDestinationType The type of destination.
type DatadogLogsDestinationType string

// List of DatadogLogsDestinationType.
const (
	DATADOGLOGSDESTINATIONTYPE_DATADOG_LOGS DatadogLogsDestinationType = "datadog_logs"
)

var allowedDatadogLogsDestinationTypeEnumValues = []DatadogLogsDestinationType{
	DATADOGLOGSDESTINATIONTYPE_DATADOG_LOGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DatadogLogsDestinationType) GetAllowedValues() []DatadogLogsDestinationType {
	return allowedDatadogLogsDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DatadogLogsDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DatadogLogsDestinationType(value)
	return nil
}

// NewDatadogLogsDestinationTypeFromValue returns a pointer to a valid DatadogLogsDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDatadogLogsDestinationTypeFromValue(v string) (*DatadogLogsDestinationType, error) {
	ev := DatadogLogsDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DatadogLogsDestinationType: valid values are %v", v, allowedDatadogLogsDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DatadogLogsDestinationType) IsValid() bool {
	for _, existing := range allowedDatadogLogsDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatadogLogsDestinationType value.
func (v DatadogLogsDestinationType) Ptr() *DatadogLogsDestinationType {
	return &v
}
