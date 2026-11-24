// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMSdkConfigTracingUrlPropagatorType A trace propagator type.
type RUMSdkConfigTracingUrlPropagatorType string

// List of RUMSdkConfigTracingUrlPropagatorType.
const (
	RUMSDKCONFIGTRACINGURLPROPAGATORTYPE_DATADOG      RUMSdkConfigTracingUrlPropagatorType = "datadog"
	RUMSDKCONFIGTRACINGURLPROPAGATORTYPE_B3           RUMSdkConfigTracingUrlPropagatorType = "b3"
	RUMSDKCONFIGTRACINGURLPROPAGATORTYPE_B3MULTI      RUMSdkConfigTracingUrlPropagatorType = "b3multi"
	RUMSDKCONFIGTRACINGURLPROPAGATORTYPE_TRACECONTEXT RUMSdkConfigTracingUrlPropagatorType = "tracecontext"
)

var allowedRUMSdkConfigTracingUrlPropagatorTypeEnumValues = []RUMSdkConfigTracingUrlPropagatorType{
	RUMSDKCONFIGTRACINGURLPROPAGATORTYPE_DATADOG,
	RUMSDKCONFIGTRACINGURLPROPAGATORTYPE_B3,
	RUMSDKCONFIGTRACINGURLPROPAGATORTYPE_B3MULTI,
	RUMSDKCONFIGTRACINGURLPROPAGATORTYPE_TRACECONTEXT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMSdkConfigTracingUrlPropagatorType) GetAllowedValues() []RUMSdkConfigTracingUrlPropagatorType {
	return allowedRUMSdkConfigTracingUrlPropagatorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMSdkConfigTracingUrlPropagatorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMSdkConfigTracingUrlPropagatorType(value)
	return nil
}

// NewRUMSdkConfigTracingUrlPropagatorTypeFromValue returns a pointer to a valid RUMSdkConfigTracingUrlPropagatorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMSdkConfigTracingUrlPropagatorTypeFromValue(v string) (*RUMSdkConfigTracingUrlPropagatorType, error) {
	ev := RUMSdkConfigTracingUrlPropagatorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMSdkConfigTracingUrlPropagatorType: valid values are %v", v, allowedRUMSdkConfigTracingUrlPropagatorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMSdkConfigTracingUrlPropagatorType) IsValid() bool {
	for _, existing := range allowedRUMSdkConfigTracingUrlPropagatorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMSdkConfigTracingUrlPropagatorType value.
func (v RUMSdkConfigTracingUrlPropagatorType) Ptr() *RUMSdkConfigTracingUrlPropagatorType {
	return &v
}
