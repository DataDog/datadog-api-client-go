// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsSankeyAggregatedNode A node aggregated into an "other" node.
type ProductAnalyticsSankeyAggregatedNode struct {
	//
	Id *string `json:"id,omitempty"`
	//
	IncomingValue *int64 `json:"incoming_value,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	OutgoingValue *int64 `json:"outgoing_value,omitempty"`
	//
	Type *ProductAnalyticsSankeyAggregatedNodeType `json:"type,omitempty"`
	//
	Value *int64 `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsSankeyAggregatedNode instantiates a new ProductAnalyticsSankeyAggregatedNode object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsSankeyAggregatedNode() *ProductAnalyticsSankeyAggregatedNode {
	this := ProductAnalyticsSankeyAggregatedNode{}
	return &this
}

// NewProductAnalyticsSankeyAggregatedNodeWithDefaults instantiates a new ProductAnalyticsSankeyAggregatedNode object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsSankeyAggregatedNodeWithDefaults() *ProductAnalyticsSankeyAggregatedNode {
	this := ProductAnalyticsSankeyAggregatedNode{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeyAggregatedNode) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyAggregatedNode) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeyAggregatedNode) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductAnalyticsSankeyAggregatedNode) SetId(v string) {
	o.Id = &v
}

// GetIncomingValue returns the IncomingValue field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeyAggregatedNode) GetIncomingValue() int64 {
	if o == nil || o.IncomingValue == nil {
		var ret int64
		return ret
	}
	return *o.IncomingValue
}

// GetIncomingValueOk returns a tuple with the IncomingValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyAggregatedNode) GetIncomingValueOk() (*int64, bool) {
	if o == nil || o.IncomingValue == nil {
		return nil, false
	}
	return o.IncomingValue, true
}

// HasIncomingValue returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeyAggregatedNode) HasIncomingValue() bool {
	return o != nil && o.IncomingValue != nil
}

// SetIncomingValue gets a reference to the given int64 and assigns it to the IncomingValue field.
func (o *ProductAnalyticsSankeyAggregatedNode) SetIncomingValue(v int64) {
	o.IncomingValue = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeyAggregatedNode) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyAggregatedNode) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeyAggregatedNode) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProductAnalyticsSankeyAggregatedNode) SetName(v string) {
	o.Name = &v
}

// GetOutgoingValue returns the OutgoingValue field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeyAggregatedNode) GetOutgoingValue() int64 {
	if o == nil || o.OutgoingValue == nil {
		var ret int64
		return ret
	}
	return *o.OutgoingValue
}

// GetOutgoingValueOk returns a tuple with the OutgoingValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyAggregatedNode) GetOutgoingValueOk() (*int64, bool) {
	if o == nil || o.OutgoingValue == nil {
		return nil, false
	}
	return o.OutgoingValue, true
}

// HasOutgoingValue returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeyAggregatedNode) HasOutgoingValue() bool {
	return o != nil && o.OutgoingValue != nil
}

// SetOutgoingValue gets a reference to the given int64 and assigns it to the OutgoingValue field.
func (o *ProductAnalyticsSankeyAggregatedNode) SetOutgoingValue(v int64) {
	o.OutgoingValue = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeyAggregatedNode) GetType() ProductAnalyticsSankeyAggregatedNodeType {
	if o == nil || o.Type == nil {
		var ret ProductAnalyticsSankeyAggregatedNodeType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyAggregatedNode) GetTypeOk() (*ProductAnalyticsSankeyAggregatedNodeType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeyAggregatedNode) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given ProductAnalyticsSankeyAggregatedNodeType and assigns it to the Type field.
func (o *ProductAnalyticsSankeyAggregatedNode) SetType(v ProductAnalyticsSankeyAggregatedNodeType) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeyAggregatedNode) GetValue() int64 {
	if o == nil || o.Value == nil {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyAggregatedNode) GetValueOk() (*int64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeyAggregatedNode) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *ProductAnalyticsSankeyAggregatedNode) SetValue(v int64) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsSankeyAggregatedNode) MarshalJSON() ([]byte, error) {
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
func (o *ProductAnalyticsSankeyAggregatedNode) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id            *string                                   `json:"id,omitempty"`
		IncomingValue *int64                                    `json:"incoming_value,omitempty"`
		Name          *string                                   `json:"name,omitempty"`
		OutgoingValue *int64                                    `json:"outgoing_value,omitempty"`
		Type          *ProductAnalyticsSankeyAggregatedNodeType `json:"type,omitempty"`
		Value         *int64                                    `json:"value,omitempty"`
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

	hasInvalidField := false
	o.Id = all.Id
	o.IncomingValue = all.IncomingValue
	o.Name = all.Name
	o.OutgoingValue = all.OutgoingValue
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
