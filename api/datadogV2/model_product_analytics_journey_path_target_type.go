// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyPathTargetType The discriminator identifying a target that references a range of steps.
type ProductAnalyticsJourneyPathTargetType string

// List of ProductAnalyticsJourneyPathTargetType.
const (
	PRODUCTANALYTICSJOURNEYPATHTARGETTYPE_PATH ProductAnalyticsJourneyPathTargetType = "path"
)

var allowedProductAnalyticsJourneyPathTargetTypeEnumValues = []ProductAnalyticsJourneyPathTargetType{
	PRODUCTANALYTICSJOURNEYPATHTARGETTYPE_PATH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsJourneyPathTargetType) GetAllowedValues() []ProductAnalyticsJourneyPathTargetType {
	return allowedProductAnalyticsJourneyPathTargetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsJourneyPathTargetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsJourneyPathTargetType(value)
	return nil
}

// NewProductAnalyticsJourneyPathTargetTypeFromValue returns a pointer to a valid ProductAnalyticsJourneyPathTargetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsJourneyPathTargetTypeFromValue(v string) (*ProductAnalyticsJourneyPathTargetType, error) {
	ev := ProductAnalyticsJourneyPathTargetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsJourneyPathTargetType: valid values are %v", v, allowedProductAnalyticsJourneyPathTargetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsJourneyPathTargetType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsJourneyPathTargetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsJourneyPathTargetType value.
func (v ProductAnalyticsJourneyPathTargetType) Ptr() *ProductAnalyticsJourneyPathTargetType {
	return &v
}
