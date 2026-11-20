// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ComputeTableColumnCompute Test-only generated schema for the dashboard versioning proof of concept.
type ComputeTableColumnCompute struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Aggregation string `json:"aggregation"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Filter *string `json:"filter,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Modifiers []TableColumnModifier `json:"modifiers,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Property *string `json:"property,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query string `json:"query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ScalarAggregation *ScalarAggregator `json:"scalar_aggregation,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewComputeTableColumnCompute instantiates a new ComputeTableColumnCompute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComputeTableColumnCompute(aggregation string, query string) *ComputeTableColumnCompute {
	this := ComputeTableColumnCompute{}
	this.Aggregation = aggregation
	this.Query = query
	return &this
}

// NewComputeTableColumnComputeWithDefaults instantiates a new ComputeTableColumnCompute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComputeTableColumnComputeWithDefaults() *ComputeTableColumnCompute {
	this := ComputeTableColumnCompute{}
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *ComputeTableColumnCompute) GetAggregation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *ComputeTableColumnCompute) GetAggregationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *ComputeTableColumnCompute) SetAggregation(v string) {
	o.Aggregation = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *ComputeTableColumnCompute) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeTableColumnCompute) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *ComputeTableColumnCompute) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *ComputeTableColumnCompute) SetFilter(v string) {
	o.Filter = &v
}

// GetModifiers returns the Modifiers field value if set, zero value otherwise.
func (o *ComputeTableColumnCompute) GetModifiers() []TableColumnModifier {
	if o == nil || o.Modifiers == nil {
		var ret []TableColumnModifier
		return ret
	}
	return o.Modifiers
}

// GetModifiersOk returns a tuple with the Modifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeTableColumnCompute) GetModifiersOk() (*[]TableColumnModifier, bool) {
	if o == nil || o.Modifiers == nil {
		return nil, false
	}
	return &o.Modifiers, true
}

// HasModifiers returns a boolean if a field has been set.
func (o *ComputeTableColumnCompute) HasModifiers() bool {
	return o != nil && o.Modifiers != nil
}

// SetModifiers gets a reference to the given []TableColumnModifier and assigns it to the Modifiers field.
func (o *ComputeTableColumnCompute) SetModifiers(v []TableColumnModifier) {
	o.Modifiers = v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *ComputeTableColumnCompute) GetProperty() string {
	if o == nil || o.Property == nil {
		var ret string
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeTableColumnCompute) GetPropertyOk() (*string, bool) {
	if o == nil || o.Property == nil {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *ComputeTableColumnCompute) HasProperty() bool {
	return o != nil && o.Property != nil
}

// SetProperty gets a reference to the given string and assigns it to the Property field.
func (o *ComputeTableColumnCompute) SetProperty(v string) {
	o.Property = &v
}

// GetQuery returns the Query field value.
func (o *ComputeTableColumnCompute) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ComputeTableColumnCompute) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *ComputeTableColumnCompute) SetQuery(v string) {
	o.Query = v
}

// GetScalarAggregation returns the ScalarAggregation field value if set, zero value otherwise.
func (o *ComputeTableColumnCompute) GetScalarAggregation() ScalarAggregator {
	if o == nil || o.ScalarAggregation == nil {
		var ret ScalarAggregator
		return ret
	}
	return *o.ScalarAggregation
}

// GetScalarAggregationOk returns a tuple with the ScalarAggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeTableColumnCompute) GetScalarAggregationOk() (*ScalarAggregator, bool) {
	if o == nil || o.ScalarAggregation == nil {
		return nil, false
	}
	return o.ScalarAggregation, true
}

// HasScalarAggregation returns a boolean if a field has been set.
func (o *ComputeTableColumnCompute) HasScalarAggregation() bool {
	return o != nil && o.ScalarAggregation != nil
}

// SetScalarAggregation gets a reference to the given ScalarAggregator and assigns it to the ScalarAggregation field.
func (o *ComputeTableColumnCompute) SetScalarAggregation(v ScalarAggregator) {
	o.ScalarAggregation = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComputeTableColumnCompute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregation"] = o.Aggregation
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Modifiers != nil {
		toSerialize["modifiers"] = o.Modifiers
	}
	if o.Property != nil {
		toSerialize["property"] = o.Property
	}
	toSerialize["query"] = o.Query
	if o.ScalarAggregation != nil {
		toSerialize["scalar_aggregation"] = o.ScalarAggregation
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComputeTableColumnCompute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation       *string               `json:"aggregation"`
		Filter            *string               `json:"filter,omitempty"`
		Modifiers         []TableColumnModifier `json:"modifiers,omitempty"`
		Property          *string               `json:"property,omitempty"`
		Query             *string               `json:"query"`
		ScalarAggregation *ScalarAggregator     `json:"scalar_aggregation,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Aggregation == nil {
		return fmt.Errorf("required field aggregation missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}

	hasInvalidField := false
	o.Aggregation = *all.Aggregation
	o.Filter = all.Filter
	o.Modifiers = all.Modifiers
	o.Property = all.Property
	o.Query = *all.Query
	if all.ScalarAggregation != nil && !all.ScalarAggregation.IsValid() {
		hasInvalidField = true
	} else {
		o.ScalarAggregation = all.ScalarAggregation
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
