// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMSdkConfigDynamicOptionStrategy The strategy used to extract the dynamic value.
type RUMSdkConfigDynamicOptionStrategy string

// List of RUMSdkConfigDynamicOptionStrategy.
const (
	RUMSDKCONFIGDYNAMICOPTIONSTRATEGY_JS            RUMSdkConfigDynamicOptionStrategy = "js"
	RUMSDKCONFIGDYNAMICOPTIONSTRATEGY_COOKIE        RUMSdkConfigDynamicOptionStrategy = "cookie"
	RUMSDKCONFIGDYNAMICOPTIONSTRATEGY_DOM           RUMSdkConfigDynamicOptionStrategy = "dom"
	RUMSDKCONFIGDYNAMICOPTIONSTRATEGY_LOCAL_STORAGE RUMSdkConfigDynamicOptionStrategy = "localStorage"
)

var allowedRUMSdkConfigDynamicOptionStrategyEnumValues = []RUMSdkConfigDynamicOptionStrategy{
	RUMSDKCONFIGDYNAMICOPTIONSTRATEGY_JS,
	RUMSDKCONFIGDYNAMICOPTIONSTRATEGY_COOKIE,
	RUMSDKCONFIGDYNAMICOPTIONSTRATEGY_DOM,
	RUMSDKCONFIGDYNAMICOPTIONSTRATEGY_LOCAL_STORAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMSdkConfigDynamicOptionStrategy) GetAllowedValues() []RUMSdkConfigDynamicOptionStrategy {
	return allowedRUMSdkConfigDynamicOptionStrategyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMSdkConfigDynamicOptionStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMSdkConfigDynamicOptionStrategy(value)
	return nil
}

// NewRUMSdkConfigDynamicOptionStrategyFromValue returns a pointer to a valid RUMSdkConfigDynamicOptionStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMSdkConfigDynamicOptionStrategyFromValue(v string) (*RUMSdkConfigDynamicOptionStrategy, error) {
	ev := RUMSdkConfigDynamicOptionStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMSdkConfigDynamicOptionStrategy: valid values are %v", v, allowedRUMSdkConfigDynamicOptionStrategyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMSdkConfigDynamicOptionStrategy) IsValid() bool {
	for _, existing := range allowedRUMSdkConfigDynamicOptionStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMSdkConfigDynamicOptionStrategy value.
func (v RUMSdkConfigDynamicOptionStrategy) Ptr() *RUMSdkConfigDynamicOptionStrategy {
	return &v
}
