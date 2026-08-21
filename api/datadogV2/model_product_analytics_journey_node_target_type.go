// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyNodeTargetType The discriminator identifying a target that references a single step.
type ProductAnalyticsJourneyNodeTargetType string

// List of ProductAnalyticsJourneyNodeTargetType.
const (
	PRODUCTANALYTICSJOURNEYNODETARGETTYPE_NODE ProductAnalyticsJourneyNodeTargetType = "node"
)

var allowedProductAnalyticsJourneyNodeTargetTypeEnumValues = []ProductAnalyticsJourneyNodeTargetType{
	PRODUCTANALYTICSJOURNEYNODETARGETTYPE_NODE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsJourneyNodeTargetType) GetAllowedValues() []ProductAnalyticsJourneyNodeTargetType {
	return allowedProductAnalyticsJourneyNodeTargetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsJourneyNodeTargetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsJourneyNodeTargetType(value)
	return nil
}

// NewProductAnalyticsJourneyNodeTargetTypeFromValue returns a pointer to a valid ProductAnalyticsJourneyNodeTargetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsJourneyNodeTargetTypeFromValue(v string) (*ProductAnalyticsJourneyNodeTargetType, error) {
	ev := ProductAnalyticsJourneyNodeTargetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsJourneyNodeTargetType: valid values are %v", v, allowedProductAnalyticsJourneyNodeTargetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsJourneyNodeTargetType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsJourneyNodeTargetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsJourneyNodeTargetType value.
func (v ProductAnalyticsJourneyNodeTargetType) Ptr() *ProductAnalyticsJourneyNodeTargetType {
	return &v
}
