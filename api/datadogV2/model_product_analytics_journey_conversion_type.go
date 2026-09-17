// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyConversionType Whether to return the entities that converted at the target step, or those that dropped off.
type ProductAnalyticsJourneyConversionType string

// List of ProductAnalyticsJourneyConversionType.
const (
	PRODUCTANALYTICSJOURNEYCONVERSIONTYPE_CONVERSION ProductAnalyticsJourneyConversionType = "conversion"
	PRODUCTANALYTICSJOURNEYCONVERSIONTYPE_DROP_OFF   ProductAnalyticsJourneyConversionType = "drop-off"
)

var allowedProductAnalyticsJourneyConversionTypeEnumValues = []ProductAnalyticsJourneyConversionType{
	PRODUCTANALYTICSJOURNEYCONVERSIONTYPE_CONVERSION,
	PRODUCTANALYTICSJOURNEYCONVERSIONTYPE_DROP_OFF,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsJourneyConversionType) GetAllowedValues() []ProductAnalyticsJourneyConversionType {
	return allowedProductAnalyticsJourneyConversionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsJourneyConversionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsJourneyConversionType(value)
	return nil
}

// NewProductAnalyticsJourneyConversionTypeFromValue returns a pointer to a valid ProductAnalyticsJourneyConversionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsJourneyConversionTypeFromValue(v string) (*ProductAnalyticsJourneyConversionType, error) {
	ev := ProductAnalyticsJourneyConversionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsJourneyConversionType: valid values are %v", v, allowedProductAnalyticsJourneyConversionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsJourneyConversionType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsJourneyConversionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsJourneyConversionType value.
func (v ProductAnalyticsJourneyConversionType) Ptr() *ProductAnalyticsJourneyConversionType {
	return &v
}
