// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CompositeReferentialDataQueryComputeItems Test-only generated schema for the dashboard versioning proof of concept.
type CompositeReferentialDataQueryComputeItems struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Aggregation CompositeReferentialDataQueryComputeItemsAggregation `json:"aggregation"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Column *string `json:"column,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCompositeReferentialDataQueryComputeItems instantiates a new CompositeReferentialDataQueryComputeItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCompositeReferentialDataQueryComputeItems(aggregation CompositeReferentialDataQueryComputeItemsAggregation, name string) *CompositeReferentialDataQueryComputeItems {
	this := CompositeReferentialDataQueryComputeItems{}
	this.Aggregation = aggregation
	this.Name = name
	return &this
}

// NewCompositeReferentialDataQueryComputeItemsWithDefaults instantiates a new CompositeReferentialDataQueryComputeItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCompositeReferentialDataQueryComputeItemsWithDefaults() *CompositeReferentialDataQueryComputeItems {
	this := CompositeReferentialDataQueryComputeItems{}
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *CompositeReferentialDataQueryComputeItems) GetAggregation() CompositeReferentialDataQueryComputeItemsAggregation {
	if o == nil {
		var ret CompositeReferentialDataQueryComputeItemsAggregation
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *CompositeReferentialDataQueryComputeItems) GetAggregationOk() (*CompositeReferentialDataQueryComputeItemsAggregation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *CompositeReferentialDataQueryComputeItems) SetAggregation(v CompositeReferentialDataQueryComputeItemsAggregation) {
	o.Aggregation = v
}

// GetColumn returns the Column field value if set, zero value otherwise.
func (o *CompositeReferentialDataQueryComputeItems) GetColumn() string {
	if o == nil || o.Column == nil {
		var ret string
		return ret
	}
	return *o.Column
}

// GetColumnOk returns a tuple with the Column field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeReferentialDataQueryComputeItems) GetColumnOk() (*string, bool) {
	if o == nil || o.Column == nil {
		return nil, false
	}
	return o.Column, true
}

// HasColumn returns a boolean if a field has been set.
func (o *CompositeReferentialDataQueryComputeItems) HasColumn() bool {
	return o != nil && o.Column != nil
}

// SetColumn gets a reference to the given string and assigns it to the Column field.
func (o *CompositeReferentialDataQueryComputeItems) SetColumn(v string) {
	o.Column = &v
}

// GetName returns the Name field value.
func (o *CompositeReferentialDataQueryComputeItems) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CompositeReferentialDataQueryComputeItems) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CompositeReferentialDataQueryComputeItems) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CompositeReferentialDataQueryComputeItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregation"] = o.Aggregation
	if o.Column != nil {
		toSerialize["column"] = o.Column
	}
	toSerialize["name"] = o.Name
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CompositeReferentialDataQueryComputeItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *CompositeReferentialDataQueryComputeItemsAggregation `json:"aggregation"`
		Column      *string                                               `json:"column,omitempty"`
		Name        *string                                               `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Aggregation == nil {
		return fmt.Errorf("required field aggregation missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}

	hasInvalidField := false
	if !all.Aggregation.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregation = *all.Aggregation
	}
	o.Column = all.Column
	o.Name = *all.Name

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
