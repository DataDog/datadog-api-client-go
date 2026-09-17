// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyComputedColumn A computed column added to each row. Requesting `first_conversion_timestamps` adds one
// `<node alias>_timestamp` key per step.
type ProductAnalyticsJourneyComputedColumn struct {
	// Name of a computed column to add to each row.
	Name ProductAnalyticsJourneyComputedColumnName `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsJourneyComputedColumn instantiates a new ProductAnalyticsJourneyComputedColumn object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsJourneyComputedColumn(name ProductAnalyticsJourneyComputedColumnName) *ProductAnalyticsJourneyComputedColumn {
	this := ProductAnalyticsJourneyComputedColumn{}
	this.Name = name
	return &this
}

// NewProductAnalyticsJourneyComputedColumnWithDefaults instantiates a new ProductAnalyticsJourneyComputedColumn object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsJourneyComputedColumnWithDefaults() *ProductAnalyticsJourneyComputedColumn {
	this := ProductAnalyticsJourneyComputedColumn{}
	return &this
}

// GetName returns the Name field value.
func (o *ProductAnalyticsJourneyComputedColumn) GetName() ProductAnalyticsJourneyComputedColumnName {
	if o == nil {
		var ret ProductAnalyticsJourneyComputedColumnName
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyComputedColumn) GetNameOk() (*ProductAnalyticsJourneyComputedColumnName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ProductAnalyticsJourneyComputedColumn) SetName(v ProductAnalyticsJourneyComputedColumnName) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsJourneyComputedColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsJourneyComputedColumn) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name *ProductAnalyticsJourneyComputedColumnName `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Name.IsValid() {
		hasInvalidField = true
	} else {
		o.Name = *all.Name
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
