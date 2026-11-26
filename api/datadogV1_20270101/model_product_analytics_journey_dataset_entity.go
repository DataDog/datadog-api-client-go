// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyDatasetEntity Test-only generated schema for the dashboard versioning proof of concept.
type ProductAnalyticsJourneyDatasetEntity string

// List of ProductAnalyticsJourneyDatasetEntity.
const (
	PRODUCTANALYTICSJOURNEYDATASETENTITY_SESSION ProductAnalyticsJourneyDatasetEntity = "session"
	PRODUCTANALYTICSJOURNEYDATASETENTITY_USER    ProductAnalyticsJourneyDatasetEntity = "user"
	PRODUCTANALYTICSJOURNEYDATASETENTITY_ACCOUNT ProductAnalyticsJourneyDatasetEntity = "account"
)

var allowedProductAnalyticsJourneyDatasetEntityEnumValues = []ProductAnalyticsJourneyDatasetEntity{
	PRODUCTANALYTICSJOURNEYDATASETENTITY_SESSION,
	PRODUCTANALYTICSJOURNEYDATASETENTITY_USER,
	PRODUCTANALYTICSJOURNEYDATASETENTITY_ACCOUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsJourneyDatasetEntity) GetAllowedValues() []ProductAnalyticsJourneyDatasetEntity {
	return allowedProductAnalyticsJourneyDatasetEntityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsJourneyDatasetEntity) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsJourneyDatasetEntity(value)
	return nil
}

// NewProductAnalyticsJourneyDatasetEntityFromValue returns a pointer to a valid ProductAnalyticsJourneyDatasetEntity
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsJourneyDatasetEntityFromValue(v string) (*ProductAnalyticsJourneyDatasetEntity, error) {
	ev := ProductAnalyticsJourneyDatasetEntity(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsJourneyDatasetEntity: valid values are %v", v, allowedProductAnalyticsJourneyDatasetEntityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsJourneyDatasetEntity) IsValid() bool {
	for _, existing := range allowedProductAnalyticsJourneyDatasetEntityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsJourneyDatasetEntity value.
func (v ProductAnalyticsJourneyDatasetEntity) Ptr() *ProductAnalyticsJourneyDatasetEntity {
	return &v
}
