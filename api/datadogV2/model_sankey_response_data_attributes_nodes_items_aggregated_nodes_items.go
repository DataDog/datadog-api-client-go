// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SankeyResponseDataAttributesNodesItemsAggregatedNodesItems
type SankeyResponseDataAttributesNodesItemsAggregatedNodesItems struct {
	//
	Id *string `json:"id,omitempty"`
	//
	IncomingValue *int64 `json:"incoming_value,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	OutgoingValue *int64 `json:"outgoing_value,omitempty"`
	//
	Type *string `json:"type,omitempty"`
	//
	Value *int64 `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSankeyResponseDataAttributesNodesItemsAggregatedNodesItems instantiates a new SankeyResponseDataAttributesNodesItemsAggregatedNodesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSankeyResponseDataAttributesNodesItemsAggregatedNodesItems() *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems {
	this := SankeyResponseDataAttributesNodesItemsAggregatedNodesItems{}
	return &this
}

// NewSankeyResponseDataAttributesNodesItemsAggregatedNodesItemsWithDefaults instantiates a new SankeyResponseDataAttributesNodesItemsAggregatedNodesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSankeyResponseDataAttributesNodesItemsAggregatedNodesItemsWithDefaults() *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems {
	this := SankeyResponseDataAttributesNodesItemsAggregatedNodesItems{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) SetId(v string) {
	o.Id = &v
}

// GetIncomingValue returns the IncomingValue field value if set, zero value otherwise.
func (o *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) GetIncomingValue() int64 {
	if o == nil || o.IncomingValue == nil {
		var ret int64
		return ret
	}
	return *o.IncomingValue
}

// GetIncomingValueOk returns a tuple with the IncomingValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) GetIncomingValueOk() (*int64, bool) {
	if o == nil || o.IncomingValue == nil {
		return nil, false
	}
	return o.IncomingValue, true
}

// HasIncomingValue returns a boolean if a field has been set.
func (o *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) HasIncomingValue() bool {
	return o != nil && o.IncomingValue != nil
}

// SetIncomingValue gets a reference to the given int64 and assigns it to the IncomingValue field.
func (o *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) SetIncomingValue(v int64) {
	o.IncomingValue = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) SetName(v string) {
	o.Name = &v
}

// GetOutgoingValue returns the OutgoingValue field value if set, zero value otherwise.
func (o *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) GetOutgoingValue() int64 {
	if o == nil || o.OutgoingValue == nil {
		var ret int64
		return ret
	}
	return *o.OutgoingValue
}

// GetOutgoingValueOk returns a tuple with the OutgoingValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) GetOutgoingValueOk() (*int64, bool) {
	if o == nil || o.OutgoingValue == nil {
		return nil, false
	}
	return o.OutgoingValue, true
}

// HasOutgoingValue returns a boolean if a field has been set.
func (o *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) HasOutgoingValue() bool {
	return o != nil && o.OutgoingValue != nil
}

// SetOutgoingValue gets a reference to the given int64 and assigns it to the OutgoingValue field.
func (o *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) SetOutgoingValue(v int64) {
	o.OutgoingValue = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) GetValue() int64 {
	if o == nil || o.Value == nil {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) GetValueOk() (*int64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) SetValue(v int64) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IncomingValue != nil {
		toSerialize["incoming_value"] = o.IncomingValue
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OutgoingValue != nil {
		toSerialize["outgoing_value"] = o.OutgoingValue
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
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
func (o *SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id            *string `json:"id,omitempty"`
		IncomingValue *int64  `json:"incoming_value,omitempty"`
		Name          *string `json:"name,omitempty"`
		OutgoingValue *int64  `json:"outgoing_value,omitempty"`
		Type          *string `json:"type,omitempty"`
		Value         *int64  `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "incoming_value", "name", "outgoing_value", "type", "value"})
	} else {
		return err
	}
	o.Id = all.Id
	o.IncomingValue = all.IncomingValue
	o.Name = all.Name
	o.OutgoingValue = all.OutgoingValue
	o.Type = all.Type
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
