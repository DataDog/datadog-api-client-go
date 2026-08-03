// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetAgentV2AttributesInstrumentationStatus The single-step instrumentation status of the Agent.
type FleetAgentV2AttributesInstrumentationStatus string

// List of FleetAgentV2AttributesInstrumentationStatus.
const (
	FLEETAGENTV2ATTRIBUTESINSTRUMENTATIONSTATUS_SUCCESS FleetAgentV2AttributesInstrumentationStatus = "success"
	FLEETAGENTV2ATTRIBUTESINSTRUMENTATIONSTATUS_FAILURE FleetAgentV2AttributesInstrumentationStatus = "failure"
)

var allowedFleetAgentV2AttributesInstrumentationStatusEnumValues = []FleetAgentV2AttributesInstrumentationStatus{
	FLEETAGENTV2ATTRIBUTESINSTRUMENTATIONSTATUS_SUCCESS,
	FLEETAGENTV2ATTRIBUTESINSTRUMENTATIONSTATUS_FAILURE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FleetAgentV2AttributesInstrumentationStatus) GetAllowedValues() []FleetAgentV2AttributesInstrumentationStatus {
	return allowedFleetAgentV2AttributesInstrumentationStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FleetAgentV2AttributesInstrumentationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FleetAgentV2AttributesInstrumentationStatus(value)
	return nil
}

// NewFleetAgentV2AttributesInstrumentationStatusFromValue returns a pointer to a valid FleetAgentV2AttributesInstrumentationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFleetAgentV2AttributesInstrumentationStatusFromValue(v string) (*FleetAgentV2AttributesInstrumentationStatus, error) {
	ev := FleetAgentV2AttributesInstrumentationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FleetAgentV2AttributesInstrumentationStatus: valid values are %v", v, allowedFleetAgentV2AttributesInstrumentationStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FleetAgentV2AttributesInstrumentationStatus) IsValid() bool {
	for _, existing := range allowedFleetAgentV2AttributesInstrumentationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FleetAgentV2AttributesInstrumentationStatus value.
func (v FleetAgentV2AttributesInstrumentationStatus) Ptr() *FleetAgentV2AttributesInstrumentationStatus {
	return &v
}
