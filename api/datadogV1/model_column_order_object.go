// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ColumnOrderObject Order criteria
type ColumnOrderObject struct {
	// Index of the sorted column
	Index *float64 `json:"index,omitempty"`
	// Name of the sorted column
	Name *string `json:"name,omitempty"`
	// Order object
	Order ColumnOrderObjectOrder `json:"order"`
	// type of column
	Type ColumnOrderObjectType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewColumnOrderObject instantiates a new ColumnOrderObject object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewColumnOrderObject(order ColumnOrderObjectOrder, typeVar ColumnOrderObjectType) *ColumnOrderObject {
	this := ColumnOrderObject{}
	this.Order = order
	this.Type = typeVar
	return &this
}

// NewColumnOrderObjectWithDefaults instantiates a new ColumnOrderObject object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewColumnOrderObjectWithDefaults() *ColumnOrderObject {
	this := ColumnOrderObject{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *ColumnOrderObject) GetIndex() float64 {
	if o == nil || o.Index == nil {
		var ret float64
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnOrderObject) GetIndexOk() (*float64, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *ColumnOrderObject) HasIndex() bool {
	return o != nil && o.Index != nil
}

// SetIndex gets a reference to the given float64 and assigns it to the Index field.
func (o *ColumnOrderObject) SetIndex(v float64) {
	o.Index = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ColumnOrderObject) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnOrderObject) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ColumnOrderObject) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ColumnOrderObject) SetName(v string) {
	o.Name = &v
}

// GetOrder returns the Order field value.
func (o *ColumnOrderObject) GetOrder() ColumnOrderObjectOrder {
	if o == nil {
		var ret ColumnOrderObjectOrder
		return ret
	}
	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *ColumnOrderObject) GetOrderOk() (*ColumnOrderObjectOrder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value.
func (o *ColumnOrderObject) SetOrder(v ColumnOrderObjectOrder) {
	o.Order = v
}

// GetType returns the Type field value.
func (o *ColumnOrderObject) GetType() ColumnOrderObjectType {
	if o == nil {
		var ret ColumnOrderObjectType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ColumnOrderObject) GetTypeOk() (*ColumnOrderObjectType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ColumnOrderObject) SetType(v ColumnOrderObjectType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ColumnOrderObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["order"] = o.Order
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ColumnOrderObject) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Index *float64                `json:"index,omitempty"`
		Name  *string                 `json:"name,omitempty"`
		Order *ColumnOrderObjectOrder `json:"order"`
		Type  *ColumnOrderObjectType  `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Order == nil {
		return fmt.Errorf("required field order missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"index", "name", "order", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Index = all.Index
	o.Name = all.Name
	if !all.Order.IsValid() {
		hasInvalidField = true
	} else {
		o.Order = *all.Order
	}
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
