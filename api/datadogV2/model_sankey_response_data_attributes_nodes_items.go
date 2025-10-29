// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SankeyResponseDataAttributesNodesItems
type SankeyResponseDataAttributesNodesItems struct {
	//
	AggregatedNodes []SankeyResponseDataAttributesNodesItemsAggregatedNodesItems `json:"aggregated_nodes,omitempty"`
	//
	Column *int64 `json:"column,omitempty"`
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

// NewSankeyResponseDataAttributesNodesItems instantiates a new SankeyResponseDataAttributesNodesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSankeyResponseDataAttributesNodesItems() *SankeyResponseDataAttributesNodesItems {
	this := SankeyResponseDataAttributesNodesItems{}
	return &this
}

// NewSankeyResponseDataAttributesNodesItemsWithDefaults instantiates a new SankeyResponseDataAttributesNodesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSankeyResponseDataAttributesNodesItemsWithDefaults() *SankeyResponseDataAttributesNodesItems {
	this := SankeyResponseDataAttributesNodesItems{}
	return &this
}

// GetAggregatedNodes returns the AggregatedNodes field value if set, zero value otherwise.
func (o *SankeyResponseDataAttributesNodesItems) GetAggregatedNodes() []SankeyResponseDataAttributesNodesItemsAggregatedNodesItems {
	if o == nil || o.AggregatedNodes == nil {
		var ret []SankeyResponseDataAttributesNodesItemsAggregatedNodesItems
		return ret
	}
	return o.AggregatedNodes
}

// GetAggregatedNodesOk returns a tuple with the AggregatedNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyResponseDataAttributesNodesItems) GetAggregatedNodesOk() (*[]SankeyResponseDataAttributesNodesItemsAggregatedNodesItems, bool) {
	if o == nil || o.AggregatedNodes == nil {
		return nil, false
	}
	return &o.AggregatedNodes, true
}

// HasAggregatedNodes returns a boolean if a field has been set.
func (o *SankeyResponseDataAttributesNodesItems) HasAggregatedNodes() bool {
	return o != nil && o.AggregatedNodes != nil
}

// SetAggregatedNodes gets a reference to the given []SankeyResponseDataAttributesNodesItemsAggregatedNodesItems and assigns it to the AggregatedNodes field.
func (o *SankeyResponseDataAttributesNodesItems) SetAggregatedNodes(v []SankeyResponseDataAttributesNodesItemsAggregatedNodesItems) {
	o.AggregatedNodes = v
}

// GetColumn returns the Column field value if set, zero value otherwise.
func (o *SankeyResponseDataAttributesNodesItems) GetColumn() int64 {
	if o == nil || o.Column == nil {
		var ret int64
		return ret
	}
	return *o.Column
}

// GetColumnOk returns a tuple with the Column field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyResponseDataAttributesNodesItems) GetColumnOk() (*int64, bool) {
	if o == nil || o.Column == nil {
		return nil, false
	}
	return o.Column, true
}

// HasColumn returns a boolean if a field has been set.
func (o *SankeyResponseDataAttributesNodesItems) HasColumn() bool {
	return o != nil && o.Column != nil
}

// SetColumn gets a reference to the given int64 and assigns it to the Column field.
func (o *SankeyResponseDataAttributesNodesItems) SetColumn(v int64) {
	o.Column = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SankeyResponseDataAttributesNodesItems) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyResponseDataAttributesNodesItems) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SankeyResponseDataAttributesNodesItems) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SankeyResponseDataAttributesNodesItems) SetId(v string) {
	o.Id = &v
}

// GetIncomingValue returns the IncomingValue field value if set, zero value otherwise.
func (o *SankeyResponseDataAttributesNodesItems) GetIncomingValue() int64 {
	if o == nil || o.IncomingValue == nil {
		var ret int64
		return ret
	}
	return *o.IncomingValue
}

// GetIncomingValueOk returns a tuple with the IncomingValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyResponseDataAttributesNodesItems) GetIncomingValueOk() (*int64, bool) {
	if o == nil || o.IncomingValue == nil {
		return nil, false
	}
	return o.IncomingValue, true
}

// HasIncomingValue returns a boolean if a field has been set.
func (o *SankeyResponseDataAttributesNodesItems) HasIncomingValue() bool {
	return o != nil && o.IncomingValue != nil
}

// SetIncomingValue gets a reference to the given int64 and assigns it to the IncomingValue field.
func (o *SankeyResponseDataAttributesNodesItems) SetIncomingValue(v int64) {
	o.IncomingValue = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SankeyResponseDataAttributesNodesItems) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyResponseDataAttributesNodesItems) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SankeyResponseDataAttributesNodesItems) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SankeyResponseDataAttributesNodesItems) SetName(v string) {
	o.Name = &v
}

// GetOutgoingValue returns the OutgoingValue field value if set, zero value otherwise.
func (o *SankeyResponseDataAttributesNodesItems) GetOutgoingValue() int64 {
	if o == nil || o.OutgoingValue == nil {
		var ret int64
		return ret
	}
	return *o.OutgoingValue
}

// GetOutgoingValueOk returns a tuple with the OutgoingValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyResponseDataAttributesNodesItems) GetOutgoingValueOk() (*int64, bool) {
	if o == nil || o.OutgoingValue == nil {
		return nil, false
	}
	return o.OutgoingValue, true
}

// HasOutgoingValue returns a boolean if a field has been set.
func (o *SankeyResponseDataAttributesNodesItems) HasOutgoingValue() bool {
	return o != nil && o.OutgoingValue != nil
}

// SetOutgoingValue gets a reference to the given int64 and assigns it to the OutgoingValue field.
func (o *SankeyResponseDataAttributesNodesItems) SetOutgoingValue(v int64) {
	o.OutgoingValue = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SankeyResponseDataAttributesNodesItems) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyResponseDataAttributesNodesItems) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SankeyResponseDataAttributesNodesItems) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SankeyResponseDataAttributesNodesItems) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SankeyResponseDataAttributesNodesItems) GetValue() int64 {
	if o == nil || o.Value == nil {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyResponseDataAttributesNodesItems) GetValueOk() (*int64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SankeyResponseDataAttributesNodesItems) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *SankeyResponseDataAttributesNodesItems) SetValue(v int64) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SankeyResponseDataAttributesNodesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AggregatedNodes != nil {
		toSerialize["aggregated_nodes"] = o.AggregatedNodes
	}
	if o.Column != nil {
		toSerialize["column"] = o.Column
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
func (o *SankeyResponseDataAttributesNodesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AggregatedNodes []SankeyResponseDataAttributesNodesItemsAggregatedNodesItems `json:"aggregated_nodes,omitempty"`
		Column          *int64                                                       `json:"column,omitempty"`
		Id              *string                                                      `json:"id,omitempty"`
		IncomingValue   *int64                                                       `json:"incoming_value,omitempty"`
		Name            *string                                                      `json:"name,omitempty"`
		OutgoingValue   *int64                                                       `json:"outgoing_value,omitempty"`
		Type            *string                                                      `json:"type,omitempty"`
		Value           *int64                                                       `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregated_nodes", "column", "id", "incoming_value", "name", "outgoing_value", "type", "value"})
	} else {
		return err
	}
	o.AggregatedNodes = all.AggregatedNodes
	o.Column = all.Column
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
