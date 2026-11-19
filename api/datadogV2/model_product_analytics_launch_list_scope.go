// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsLaunchListScope Scope of a launch list. `liftoffs` leaves out launches in the `plan` phase.
type ProductAnalyticsLaunchListScope string

// List of ProductAnalyticsLaunchListScope.
const (
	PRODUCTANALYTICSLAUNCHLISTSCOPE_LIFTOFFS ProductAnalyticsLaunchListScope = "liftoffs"
)

var allowedProductAnalyticsLaunchListScopeEnumValues = []ProductAnalyticsLaunchListScope{
	PRODUCTANALYTICSLAUNCHLISTSCOPE_LIFTOFFS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsLaunchListScope) GetAllowedValues() []ProductAnalyticsLaunchListScope {
	return allowedProductAnalyticsLaunchListScopeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsLaunchListScope) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsLaunchListScope(value)
	return nil
}

// NewProductAnalyticsLaunchListScopeFromValue returns a pointer to a valid ProductAnalyticsLaunchListScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsLaunchListScopeFromValue(v string) (*ProductAnalyticsLaunchListScope, error) {
	ev := ProductAnalyticsLaunchListScope(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsLaunchListScope: valid values are %v", v, allowedProductAnalyticsLaunchListScopeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsLaunchListScope) IsValid() bool {
	for _, existing := range allowedProductAnalyticsLaunchListScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsLaunchListScope value.
func (v ProductAnalyticsLaunchListScope) Ptr() *ProductAnalyticsLaunchListScope {
	return &v
}
