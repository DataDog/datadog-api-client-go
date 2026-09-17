// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneySearchGraphFilter A filter applied to a step, or a range of steps, of the journey graph.
type ProductAnalyticsJourneySearchGraphFilter struct {
	// The journey-level metric the graph filter applies to.
	Name ProductAnalyticsJourneySearchGraphFilterName `json:"name"`
	// Comparison operator applied to the graph filter value.
	Operator ProductAnalyticsJourneySearchGraphFilterOperator `json:"operator"`
	// A reference to a step, or a range of steps, in the journey.
	// Use a `node` target to name a single step, or a `path` target to name the range
	// between two steps.
	Target *ProductAnalyticsJourneyTarget `json:"target,omitempty"`
	// Value compared against the metric. Durations are expressed in milliseconds.
	Value int64 `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsJourneySearchGraphFilter instantiates a new ProductAnalyticsJourneySearchGraphFilter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsJourneySearchGraphFilter(name ProductAnalyticsJourneySearchGraphFilterName, operator ProductAnalyticsJourneySearchGraphFilterOperator, value int64) *ProductAnalyticsJourneySearchGraphFilter {
	this := ProductAnalyticsJourneySearchGraphFilter{}
	this.Name = name
	this.Operator = operator
	this.Value = value
	return &this
}

// NewProductAnalyticsJourneySearchGraphFilterWithDefaults instantiates a new ProductAnalyticsJourneySearchGraphFilter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsJourneySearchGraphFilterWithDefaults() *ProductAnalyticsJourneySearchGraphFilter {
	this := ProductAnalyticsJourneySearchGraphFilter{}
	return &this
}

// GetName returns the Name field value.
func (o *ProductAnalyticsJourneySearchGraphFilter) GetName() ProductAnalyticsJourneySearchGraphFilterName {
	if o == nil {
		var ret ProductAnalyticsJourneySearchGraphFilterName
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneySearchGraphFilter) GetNameOk() (*ProductAnalyticsJourneySearchGraphFilterName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ProductAnalyticsJourneySearchGraphFilter) SetName(v ProductAnalyticsJourneySearchGraphFilterName) {
	o.Name = v
}

// GetOperator returns the Operator field value.
func (o *ProductAnalyticsJourneySearchGraphFilter) GetOperator() ProductAnalyticsJourneySearchGraphFilterOperator {
	if o == nil {
		var ret ProductAnalyticsJourneySearchGraphFilterOperator
		return ret
	}
	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneySearchGraphFilter) GetOperatorOk() (*ProductAnalyticsJourneySearchGraphFilterOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value.
func (o *ProductAnalyticsJourneySearchGraphFilter) SetOperator(v ProductAnalyticsJourneySearchGraphFilterOperator) {
	o.Operator = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneySearchGraphFilter) GetTarget() ProductAnalyticsJourneyTarget {
	if o == nil || o.Target == nil {
		var ret ProductAnalyticsJourneyTarget
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneySearchGraphFilter) GetTargetOk() (*ProductAnalyticsJourneyTarget, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneySearchGraphFilter) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given ProductAnalyticsJourneyTarget and assigns it to the Target field.
func (o *ProductAnalyticsJourneySearchGraphFilter) SetTarget(v ProductAnalyticsJourneyTarget) {
	o.Target = &v
}

// GetValue returns the Value field value.
func (o *ProductAnalyticsJourneySearchGraphFilter) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneySearchGraphFilter) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *ProductAnalyticsJourneySearchGraphFilter) SetValue(v int64) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsJourneySearchGraphFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["operator"] = o.Operator
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsJourneySearchGraphFilter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name     *ProductAnalyticsJourneySearchGraphFilterName     `json:"name"`
		Operator *ProductAnalyticsJourneySearchGraphFilterOperator `json:"operator"`
		Target   *ProductAnalyticsJourneyTarget                    `json:"target,omitempty"`
		Value    *int64                                            `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Operator == nil {
		return fmt.Errorf("required field operator missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "operator", "target", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Name.IsValid() {
		hasInvalidField = true
	} else {
		o.Name = *all.Name
	}
	if !all.Operator.IsValid() {
		hasInvalidField = true
	} else {
		o.Operator = *all.Operator
	}
	o.Target = all.Target
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
