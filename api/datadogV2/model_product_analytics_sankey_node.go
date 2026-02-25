// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsSankeyNode A node in the Sankey diagram representing a page or aggregation.
type ProductAnalyticsSankeyNode struct {
	// Nodes aggregated into this node (for "other" type).
	AggregatedNodes []ProductAnalyticsSankeyAggregatedNode `json:"aggregated_nodes,omitempty"`
	// The step column (0-indexed).
	Column *int64 `json:"column,omitempty"`
	//
	DropoffValue *int64 `json:"dropoff_value,omitempty"`
	//
	Id *string `json:"id,omitempty"`
	//
	IncomingValue *int64 `json:"incoming_value,omitempty"`
	// The page name.
	Name *string `json:"name,omitempty"`
	//
	OutgoingValue *int64 `json:"outgoing_value,omitempty"`
	// Node type.
	Type *ProductAnalyticsSankeyNodeType `json:"type,omitempty"`
	// The number of sessions through this node.
	Value *int64 `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsSankeyNode instantiates a new ProductAnalyticsSankeyNode object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsSankeyNode() *ProductAnalyticsSankeyNode {
	this := ProductAnalyticsSankeyNode{}
	return &this
}

// NewProductAnalyticsSankeyNodeWithDefaults instantiates a new ProductAnalyticsSankeyNode object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsSankeyNodeWithDefaults() *ProductAnalyticsSankeyNode {
	this := ProductAnalyticsSankeyNode{}
	return &this
}

// GetAggregatedNodes returns the AggregatedNodes field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeyNode) GetAggregatedNodes() []ProductAnalyticsSankeyAggregatedNode {
	if o == nil || o.AggregatedNodes == nil {
		var ret []ProductAnalyticsSankeyAggregatedNode
		return ret
	}
	return o.AggregatedNodes
}

// GetAggregatedNodesOk returns a tuple with the AggregatedNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyNode) GetAggregatedNodesOk() (*[]ProductAnalyticsSankeyAggregatedNode, bool) {
	if o == nil || o.AggregatedNodes == nil {
		return nil, false
	}
	return &o.AggregatedNodes, true
}

// HasAggregatedNodes returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeyNode) HasAggregatedNodes() bool {
	return o != nil && o.AggregatedNodes != nil
}

// SetAggregatedNodes gets a reference to the given []ProductAnalyticsSankeyAggregatedNode and assigns it to the AggregatedNodes field.
func (o *ProductAnalyticsSankeyNode) SetAggregatedNodes(v []ProductAnalyticsSankeyAggregatedNode) {
	o.AggregatedNodes = v
}

// GetColumn returns the Column field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeyNode) GetColumn() int64 {
	if o == nil || o.Column == nil {
		var ret int64
		return ret
	}
	return *o.Column
}

// GetColumnOk returns a tuple with the Column field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyNode) GetColumnOk() (*int64, bool) {
	if o == nil || o.Column == nil {
		return nil, false
	}
	return o.Column, true
}

// HasColumn returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeyNode) HasColumn() bool {
	return o != nil && o.Column != nil
}

// SetColumn gets a reference to the given int64 and assigns it to the Column field.
func (o *ProductAnalyticsSankeyNode) SetColumn(v int64) {
	o.Column = &v
}

// GetDropoffValue returns the DropoffValue field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeyNode) GetDropoffValue() int64 {
	if o == nil || o.DropoffValue == nil {
		var ret int64
		return ret
	}
	return *o.DropoffValue
}

// GetDropoffValueOk returns a tuple with the DropoffValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyNode) GetDropoffValueOk() (*int64, bool) {
	if o == nil || o.DropoffValue == nil {
		return nil, false
	}
	return o.DropoffValue, true
}

// HasDropoffValue returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeyNode) HasDropoffValue() bool {
	return o != nil && o.DropoffValue != nil
}

// SetDropoffValue gets a reference to the given int64 and assigns it to the DropoffValue field.
func (o *ProductAnalyticsSankeyNode) SetDropoffValue(v int64) {
	o.DropoffValue = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeyNode) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyNode) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeyNode) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductAnalyticsSankeyNode) SetId(v string) {
	o.Id = &v
}

// GetIncomingValue returns the IncomingValue field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeyNode) GetIncomingValue() int64 {
	if o == nil || o.IncomingValue == nil {
		var ret int64
		return ret
	}
	return *o.IncomingValue
}

// GetIncomingValueOk returns a tuple with the IncomingValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyNode) GetIncomingValueOk() (*int64, bool) {
	if o == nil || o.IncomingValue == nil {
		return nil, false
	}
	return o.IncomingValue, true
}

// HasIncomingValue returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeyNode) HasIncomingValue() bool {
	return o != nil && o.IncomingValue != nil
}

// SetIncomingValue gets a reference to the given int64 and assigns it to the IncomingValue field.
func (o *ProductAnalyticsSankeyNode) SetIncomingValue(v int64) {
	o.IncomingValue = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeyNode) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyNode) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeyNode) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProductAnalyticsSankeyNode) SetName(v string) {
	o.Name = &v
}

// GetOutgoingValue returns the OutgoingValue field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeyNode) GetOutgoingValue() int64 {
	if o == nil || o.OutgoingValue == nil {
		var ret int64
		return ret
	}
	return *o.OutgoingValue
}

// GetOutgoingValueOk returns a tuple with the OutgoingValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyNode) GetOutgoingValueOk() (*int64, bool) {
	if o == nil || o.OutgoingValue == nil {
		return nil, false
	}
	return o.OutgoingValue, true
}

// HasOutgoingValue returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeyNode) HasOutgoingValue() bool {
	return o != nil && o.OutgoingValue != nil
}

// SetOutgoingValue gets a reference to the given int64 and assigns it to the OutgoingValue field.
func (o *ProductAnalyticsSankeyNode) SetOutgoingValue(v int64) {
	o.OutgoingValue = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeyNode) GetType() ProductAnalyticsSankeyNodeType {
	if o == nil || o.Type == nil {
		var ret ProductAnalyticsSankeyNodeType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyNode) GetTypeOk() (*ProductAnalyticsSankeyNodeType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeyNode) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given ProductAnalyticsSankeyNodeType and assigns it to the Type field.
func (o *ProductAnalyticsSankeyNode) SetType(v ProductAnalyticsSankeyNodeType) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeyNode) GetValue() int64 {
	if o == nil || o.Value == nil {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyNode) GetValueOk() (*int64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeyNode) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *ProductAnalyticsSankeyNode) SetValue(v int64) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsSankeyNode) MarshalJSON() ([]byte, error) {
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
	if o.DropoffValue != nil {
		toSerialize["dropoff_value"] = o.DropoffValue
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
func (o *ProductAnalyticsSankeyNode) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AggregatedNodes []ProductAnalyticsSankeyAggregatedNode `json:"aggregated_nodes,omitempty"`
		Column          *int64                                 `json:"column,omitempty"`
		DropoffValue    *int64                                 `json:"dropoff_value,omitempty"`
		Id              *string                                `json:"id,omitempty"`
		IncomingValue   *int64                                 `json:"incoming_value,omitempty"`
		Name            *string                                `json:"name,omitempty"`
		OutgoingValue   *int64                                 `json:"outgoing_value,omitempty"`
		Type            *ProductAnalyticsSankeyNodeType        `json:"type,omitempty"`
		Value           *int64                                 `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregated_nodes", "column", "dropoff_value", "id", "incoming_value", "name", "outgoing_value", "type", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AggregatedNodes = all.AggregatedNodes
	o.Column = all.Column
	o.DropoffValue = all.DropoffValue
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
