// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
)

// ScalarColumn A single column in a scalar query response.
type ScalarColumn struct {
	// List of units.
	Unit []Unit `json:"unit,omitempty"`
	// Array of values for each group-by combination that results from one formula or query.
	Values []float64 `json:"values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewScalarColumn instantiates a new ScalarColumn object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScalarColumn() *ScalarColumn {
	this := ScalarColumn{}
	return &this
}

// NewScalarColumnWithDefaults instantiates a new ScalarColumn object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScalarColumnWithDefaults() *ScalarColumn {
	this := ScalarColumn{}
	return &this
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *ScalarColumn) GetUnit() []Unit {
	if o == nil || o.Unit == nil {
		var ret []Unit
		return ret
	}
	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalarColumn) GetUnitOk() (*[]Unit, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return &o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *ScalarColumn) HasUnit() bool {
	return o != nil && o.Unit != nil
}

// SetUnit gets a reference to the given []Unit and assigns it to the Unit field.
func (o *ScalarColumn) SetUnit(v []Unit) {
	o.Unit = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *ScalarColumn) GetValues() []float64 {
	if o == nil || o.Values == nil {
		var ret []float64
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalarColumn) GetValuesOk() (*[]float64, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *ScalarColumn) HasValues() bool {
	return o != nil && o.Values != nil
}

// SetValues gets a reference to the given []float64 and assigns it to the Values field.
func (o *ScalarColumn) SetValues(v []float64) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScalarColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScalarColumn) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Unit   []Unit    `json:"unit,omitempty"`
		Values []float64 `json:"values,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Unit = all.Unit
	o.Values = all.Values
	return nil
}
