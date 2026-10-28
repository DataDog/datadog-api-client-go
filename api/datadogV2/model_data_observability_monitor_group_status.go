// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DataObservabilityMonitorGroupStatus The state of a monitor group, from the alerting engine.
type DataObservabilityMonitorGroupStatus string

// List of DataObservabilityMonitorGroupStatus.
const (
	DATAOBSERVABILITYMONITORGROUPSTATUS_OK      DataObservabilityMonitorGroupStatus = "ok"
	DATAOBSERVABILITYMONITORGROUPSTATUS_WARN    DataObservabilityMonitorGroupStatus = "warn"
	DATAOBSERVABILITYMONITORGROUPSTATUS_ALERT   DataObservabilityMonitorGroupStatus = "alert"
	DATAOBSERVABILITYMONITORGROUPSTATUS_NO_DATA DataObservabilityMonitorGroupStatus = "no_data"
	DATAOBSERVABILITYMONITORGROUPSTATUS_UNKNOWN DataObservabilityMonitorGroupStatus = "unknown"
)

var allowedDataObservabilityMonitorGroupStatusEnumValues = []DataObservabilityMonitorGroupStatus{
	DATAOBSERVABILITYMONITORGROUPSTATUS_OK,
	DATAOBSERVABILITYMONITORGROUPSTATUS_WARN,
	DATAOBSERVABILITYMONITORGROUPSTATUS_ALERT,
	DATAOBSERVABILITYMONITORGROUPSTATUS_NO_DATA,
	DATAOBSERVABILITYMONITORGROUPSTATUS_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DataObservabilityMonitorGroupStatus) GetAllowedValues() []DataObservabilityMonitorGroupStatus {
	return allowedDataObservabilityMonitorGroupStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DataObservabilityMonitorGroupStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DataObservabilityMonitorGroupStatus(value)
	return nil
}

// NewDataObservabilityMonitorGroupStatusFromValue returns a pointer to a valid DataObservabilityMonitorGroupStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDataObservabilityMonitorGroupStatusFromValue(v string) (*DataObservabilityMonitorGroupStatus, error) {
	ev := DataObservabilityMonitorGroupStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DataObservabilityMonitorGroupStatus: valid values are %v", v, allowedDataObservabilityMonitorGroupStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DataObservabilityMonitorGroupStatus) IsValid() bool {
	for _, existing := range allowedDataObservabilityMonitorGroupStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataObservabilityMonitorGroupStatus value.
func (v DataObservabilityMonitorGroupStatus) Ptr() *DataObservabilityMonitorGroupStatus {
	return &v
}
