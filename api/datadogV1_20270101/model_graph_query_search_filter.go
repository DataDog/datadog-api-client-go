// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GraphQuerySearchFilter Test-only generated schema for the dashboard versioning proof of concept.
type GraphQuerySearchFilter struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Operator GraphQuerySearchFilterOperator `json:"operator"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Target *GraphTarget `json:"target,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Value float64 `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewGraphQuerySearchFilter instantiates a new GraphQuerySearchFilter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGraphQuerySearchFilter(name string, operator GraphQuerySearchFilterOperator, value float64) *GraphQuerySearchFilter {
	this := GraphQuerySearchFilter{}
	this.Name = name
	this.Operator = operator
	this.Value = value
	return &this
}

// NewGraphQuerySearchFilterWithDefaults instantiates a new GraphQuerySearchFilter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGraphQuerySearchFilterWithDefaults() *GraphQuerySearchFilter {
	this := GraphQuerySearchFilter{}
	return &this
}

// GetName returns the Name field value.
func (o *GraphQuerySearchFilter) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GraphQuerySearchFilter) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *GraphQuerySearchFilter) SetName(v string) {
	o.Name = v
}

// GetOperator returns the Operator field value.
func (o *GraphQuerySearchFilter) GetOperator() GraphQuerySearchFilterOperator {
	if o == nil {
		var ret GraphQuerySearchFilterOperator
		return ret
	}
	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *GraphQuerySearchFilter) GetOperatorOk() (*GraphQuerySearchFilterOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value.
func (o *GraphQuerySearchFilter) SetOperator(v GraphQuerySearchFilterOperator) {
	o.Operator = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *GraphQuerySearchFilter) GetTarget() GraphTarget {
	if o == nil || o.Target == nil {
		var ret GraphTarget
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphQuerySearchFilter) GetTargetOk() (*GraphTarget, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *GraphQuerySearchFilter) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given GraphTarget and assigns it to the Target field.
func (o *GraphQuerySearchFilter) SetTarget(v GraphTarget) {
	o.Target = &v
}

// GetValue returns the Value field value.
func (o *GraphQuerySearchFilter) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GraphQuerySearchFilter) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *GraphQuerySearchFilter) SetValue(v float64) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GraphQuerySearchFilter) MarshalJSON() ([]byte, error) {
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
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GraphQuerySearchFilter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name     *string                         `json:"name"`
		Operator *GraphQuerySearchFilterOperator `json:"operator"`
		Target   *GraphTarget                    `json:"target,omitempty"`
		Value    *float64                        `json:"value"`
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

	hasInvalidField := false
	o.Name = *all.Name
	if !all.Operator.IsValid() {
		hasInvalidField = true
	} else {
		o.Operator = *all.Operator
	}
	o.Target = all.Target
	o.Value = *all.Value

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
