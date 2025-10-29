// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SankeyResponseDataAttributesLinksItems
type SankeyResponseDataAttributesLinksItems struct {
	//
	Column *int64 `json:"column,omitempty"`
	//
	Id *string `json:"id,omitempty"`
	//
	Source *string `json:"source,omitempty"`
	//
	Target *string `json:"target,omitempty"`
	//
	Value *int64 `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSankeyResponseDataAttributesLinksItems instantiates a new SankeyResponseDataAttributesLinksItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSankeyResponseDataAttributesLinksItems() *SankeyResponseDataAttributesLinksItems {
	this := SankeyResponseDataAttributesLinksItems{}
	return &this
}

// NewSankeyResponseDataAttributesLinksItemsWithDefaults instantiates a new SankeyResponseDataAttributesLinksItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSankeyResponseDataAttributesLinksItemsWithDefaults() *SankeyResponseDataAttributesLinksItems {
	this := SankeyResponseDataAttributesLinksItems{}
	return &this
}

// GetColumn returns the Column field value if set, zero value otherwise.
func (o *SankeyResponseDataAttributesLinksItems) GetColumn() int64 {
	if o == nil || o.Column == nil {
		var ret int64
		return ret
	}
	return *o.Column
}

// GetColumnOk returns a tuple with the Column field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyResponseDataAttributesLinksItems) GetColumnOk() (*int64, bool) {
	if o == nil || o.Column == nil {
		return nil, false
	}
	return o.Column, true
}

// HasColumn returns a boolean if a field has been set.
func (o *SankeyResponseDataAttributesLinksItems) HasColumn() bool {
	return o != nil && o.Column != nil
}

// SetColumn gets a reference to the given int64 and assigns it to the Column field.
func (o *SankeyResponseDataAttributesLinksItems) SetColumn(v int64) {
	o.Column = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SankeyResponseDataAttributesLinksItems) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyResponseDataAttributesLinksItems) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SankeyResponseDataAttributesLinksItems) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SankeyResponseDataAttributesLinksItems) SetId(v string) {
	o.Id = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *SankeyResponseDataAttributesLinksItems) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyResponseDataAttributesLinksItems) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *SankeyResponseDataAttributesLinksItems) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *SankeyResponseDataAttributesLinksItems) SetSource(v string) {
	o.Source = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *SankeyResponseDataAttributesLinksItems) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyResponseDataAttributesLinksItems) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *SankeyResponseDataAttributesLinksItems) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *SankeyResponseDataAttributesLinksItems) SetTarget(v string) {
	o.Target = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SankeyResponseDataAttributesLinksItems) GetValue() int64 {
	if o == nil || o.Value == nil {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyResponseDataAttributesLinksItems) GetValueOk() (*int64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SankeyResponseDataAttributesLinksItems) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *SankeyResponseDataAttributesLinksItems) SetValue(v int64) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SankeyResponseDataAttributesLinksItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Column != nil {
		toSerialize["column"] = o.Column
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SankeyResponseDataAttributesLinksItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Column *int64  `json:"column,omitempty"`
		Id     *string `json:"id,omitempty"`
		Source *string `json:"source,omitempty"`
		Target *string `json:"target,omitempty"`
		Value  *int64  `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"column", "id", "source", "target", "value"})
	} else {
		return err
	}
	o.Column = all.Column
	o.Id = all.Id
	o.Source = all.Source
	o.Target = all.Target
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
