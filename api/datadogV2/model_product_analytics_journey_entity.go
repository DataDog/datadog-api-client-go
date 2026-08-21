// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyEntity The kind of entity returned by a journey list query.
type ProductAnalyticsJourneyEntity string

// List of ProductAnalyticsJourneyEntity.
const (
	PRODUCTANALYTICSJOURNEYENTITY_SESSION ProductAnalyticsJourneyEntity = "session"
	PRODUCTANALYTICSJOURNEYENTITY_USER    ProductAnalyticsJourneyEntity = "user"
	PRODUCTANALYTICSJOURNEYENTITY_ACCOUNT ProductAnalyticsJourneyEntity = "account"
)

var allowedProductAnalyticsJourneyEntityEnumValues = []ProductAnalyticsJourneyEntity{
	PRODUCTANALYTICSJOURNEYENTITY_SESSION,
	PRODUCTANALYTICSJOURNEYENTITY_USER,
	PRODUCTANALYTICSJOURNEYENTITY_ACCOUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsJourneyEntity) GetAllowedValues() []ProductAnalyticsJourneyEntity {
	return allowedProductAnalyticsJourneyEntityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsJourneyEntity) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsJourneyEntity(value)
	return nil
}

// NewProductAnalyticsJourneyEntityFromValue returns a pointer to a valid ProductAnalyticsJourneyEntity
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsJourneyEntityFromValue(v string) (*ProductAnalyticsJourneyEntity, error) {
	ev := ProductAnalyticsJourneyEntity(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsJourneyEntity: valid values are %v", v, allowedProductAnalyticsJourneyEntityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsJourneyEntity) IsValid() bool {
	for _, existing := range allowedProductAnalyticsJourneyEntityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsJourneyEntity value.
func (v ProductAnalyticsJourneyEntity) Ptr() *ProductAnalyticsJourneyEntity {
	return &v
}
