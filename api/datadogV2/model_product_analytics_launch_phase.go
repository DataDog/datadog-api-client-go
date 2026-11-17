// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsLaunchPhase Lifecycle phase of a launch. A new launch starts in `plan` or `setup`, and defaults to `plan`.
// Allowed changes: `plan` to `setup`, `setup` to `ready`, `ready` to `setup` or `launched`, and `launched` to
// `ready`.
type ProductAnalyticsLaunchPhase string

// List of ProductAnalyticsLaunchPhase.
const (
	PRODUCTANALYTICSLAUNCHPHASE_PLAN     ProductAnalyticsLaunchPhase = "plan"
	PRODUCTANALYTICSLAUNCHPHASE_SETUP    ProductAnalyticsLaunchPhase = "setup"
	PRODUCTANALYTICSLAUNCHPHASE_READY    ProductAnalyticsLaunchPhase = "ready"
	PRODUCTANALYTICSLAUNCHPHASE_LAUNCHED ProductAnalyticsLaunchPhase = "launched"
)

var allowedProductAnalyticsLaunchPhaseEnumValues = []ProductAnalyticsLaunchPhase{
	PRODUCTANALYTICSLAUNCHPHASE_PLAN,
	PRODUCTANALYTICSLAUNCHPHASE_SETUP,
	PRODUCTANALYTICSLAUNCHPHASE_READY,
	PRODUCTANALYTICSLAUNCHPHASE_LAUNCHED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsLaunchPhase) GetAllowedValues() []ProductAnalyticsLaunchPhase {
	return allowedProductAnalyticsLaunchPhaseEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsLaunchPhase) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsLaunchPhase(value)
	return nil
}

// NewProductAnalyticsLaunchPhaseFromValue returns a pointer to a valid ProductAnalyticsLaunchPhase
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsLaunchPhaseFromValue(v string) (*ProductAnalyticsLaunchPhase, error) {
	ev := ProductAnalyticsLaunchPhase(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsLaunchPhase: valid values are %v", v, allowedProductAnalyticsLaunchPhaseEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsLaunchPhase) IsValid() bool {
	for _, existing := range allowedProductAnalyticsLaunchPhaseEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsLaunchPhase value.
func (v ProductAnalyticsLaunchPhase) Ptr() *ProductAnalyticsLaunchPhase {
	return &v
}
