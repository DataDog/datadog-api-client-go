// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyListResponseAttributes Attributes of a journey list response.
type ProductAnalyticsJourneyListResponseAttributes struct {
	// The kind of entity returned by a journey list query.
	Entity ProductAnalyticsJourneyEntity `json:"entity"`
	// The returned rows.
	Records []map[string]interface{} `json:"records"`
	// Total number of rows matching the query, ignoring `limit`.
	TotalCount int64 `json:"total_count"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsJourneyListResponseAttributes instantiates a new ProductAnalyticsJourneyListResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsJourneyListResponseAttributes(entity ProductAnalyticsJourneyEntity, records []map[string]interface{}, totalCount int64) *ProductAnalyticsJourneyListResponseAttributes {
	this := ProductAnalyticsJourneyListResponseAttributes{}
	this.Entity = entity
	this.Records = records
	this.TotalCount = totalCount
	return &this
}

// NewProductAnalyticsJourneyListResponseAttributesWithDefaults instantiates a new ProductAnalyticsJourneyListResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsJourneyListResponseAttributesWithDefaults() *ProductAnalyticsJourneyListResponseAttributes {
	this := ProductAnalyticsJourneyListResponseAttributes{}
	return &this
}

// GetEntity returns the Entity field value.
func (o *ProductAnalyticsJourneyListResponseAttributes) GetEntity() ProductAnalyticsJourneyEntity {
	if o == nil {
		var ret ProductAnalyticsJourneyEntity
		return ret
	}
	return o.Entity
}

// GetEntityOk returns a tuple with the Entity field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyListResponseAttributes) GetEntityOk() (*ProductAnalyticsJourneyEntity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entity, true
}

// SetEntity sets field value.
func (o *ProductAnalyticsJourneyListResponseAttributes) SetEntity(v ProductAnalyticsJourneyEntity) {
	o.Entity = v
}

// GetRecords returns the Records field value.
func (o *ProductAnalyticsJourneyListResponseAttributes) GetRecords() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyListResponseAttributes) GetRecordsOk() (*[]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Records, true
}

// SetRecords sets field value.
func (o *ProductAnalyticsJourneyListResponseAttributes) SetRecords(v []map[string]interface{}) {
	o.Records = v
}

// GetTotalCount returns the TotalCount field value.
func (o *ProductAnalyticsJourneyListResponseAttributes) GetTotalCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyListResponseAttributes) GetTotalCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value.
func (o *ProductAnalyticsJourneyListResponseAttributes) SetTotalCount(v int64) {
	o.TotalCount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsJourneyListResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["entity"] = o.Entity
	toSerialize["records"] = o.Records
	toSerialize["total_count"] = o.TotalCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsJourneyListResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Entity     *ProductAnalyticsJourneyEntity `json:"entity"`
		Records    *[]map[string]interface{}      `json:"records"`
		TotalCount *int64                         `json:"total_count"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Entity == nil {
		return fmt.Errorf("required field entity missing")
	}
	if all.Records == nil {
		return fmt.Errorf("required field records missing")
	}
	if all.TotalCount == nil {
		return fmt.Errorf("required field total_count missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"entity", "records", "total_count"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Entity.IsValid() {
		hasInvalidField = true
	} else {
		o.Entity = *all.Entity
	}
	o.Records = *all.Records
	o.TotalCount = *all.TotalCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
