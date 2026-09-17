// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyComputedColumnName Name of a computed column to add to each row.
type ProductAnalyticsJourneyComputedColumnName string

// List of ProductAnalyticsJourneyComputedColumnName.
const (
	PRODUCTANALYTICSJOURNEYCOMPUTEDCOLUMNNAME_FIRST_CONVERSION_TIMESTAMPS ProductAnalyticsJourneyComputedColumnName = "first_conversion_timestamps"
)

var allowedProductAnalyticsJourneyComputedColumnNameEnumValues = []ProductAnalyticsJourneyComputedColumnName{
	PRODUCTANALYTICSJOURNEYCOMPUTEDCOLUMNNAME_FIRST_CONVERSION_TIMESTAMPS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsJourneyComputedColumnName) GetAllowedValues() []ProductAnalyticsJourneyComputedColumnName {
	return allowedProductAnalyticsJourneyComputedColumnNameEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsJourneyComputedColumnName) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsJourneyComputedColumnName(value)
	return nil
}

// NewProductAnalyticsJourneyComputedColumnNameFromValue returns a pointer to a valid ProductAnalyticsJourneyComputedColumnName
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsJourneyComputedColumnNameFromValue(v string) (*ProductAnalyticsJourneyComputedColumnName, error) {
	ev := ProductAnalyticsJourneyComputedColumnName(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsJourneyComputedColumnName: valid values are %v", v, allowedProductAnalyticsJourneyComputedColumnNameEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsJourneyComputedColumnName) IsValid() bool {
	for _, existing := range allowedProductAnalyticsJourneyComputedColumnNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsJourneyComputedColumnName value.
func (v ProductAnalyticsJourneyComputedColumnName) Ptr() *ProductAnalyticsJourneyComputedColumnName {
	return &v
}
