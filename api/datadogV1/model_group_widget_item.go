// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GroupWidgetItem A widget defined inside a group.
type GroupWidgetItem struct {
	// Same as the regular widget definition but without group.
	Definition GroupWidgetItemDefinition `json:"definition"`
	// ID of the widget.
	Id *int64 `json:"id,omitempty"`
	// The layout for a widget on a `free` or **new dashboard layout** dashboard.
	Layout *WidgetLayout `json:"layout,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewGroupWidgetItem instantiates a new GroupWidgetItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGroupWidgetItem(definition GroupWidgetItemDefinition) *GroupWidgetItem {
	this := GroupWidgetItem{}
	this.Definition = definition
	return &this
}

// NewGroupWidgetItemWithDefaults instantiates a new GroupWidgetItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGroupWidgetItemWithDefaults() *GroupWidgetItem {
	this := GroupWidgetItem{}
	return &this
}

// GetDefinition returns the Definition field value.
func (o *GroupWidgetItem) GetDefinition() GroupWidgetItemDefinition {
	if o == nil {
		var ret GroupWidgetItemDefinition
		return ret
	}
	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *GroupWidgetItem) GetDefinitionOk() (*GroupWidgetItemDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definition, true
}

// SetDefinition sets field value.
func (o *GroupWidgetItem) SetDefinition(v GroupWidgetItemDefinition) {
	o.Definition = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupWidgetItem) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupWidgetItem) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupWidgetItem) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GroupWidgetItem) SetId(v int64) {
	o.Id = &v
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *GroupWidgetItem) GetLayout() WidgetLayout {
	if o == nil || o.Layout == nil {
		var ret WidgetLayout
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupWidgetItem) GetLayoutOk() (*WidgetLayout, bool) {
	if o == nil || o.Layout == nil {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *GroupWidgetItem) HasLayout() bool {
	return o != nil && o.Layout != nil
}

// SetLayout gets a reference to the given WidgetLayout and assigns it to the Layout field.
func (o *GroupWidgetItem) SetLayout(v WidgetLayout) {
	o.Layout = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GroupWidgetItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["definition"] = o.Definition
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Layout != nil {
		toSerialize["layout"] = o.Layout
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GroupWidgetItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Definition *GroupWidgetItemDefinition `json:"definition"`
		Id         *int64                     `json:"id,omitempty"`
		Layout     *WidgetLayout              `json:"layout,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Definition == nil {
		return fmt.Errorf("required field definition missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"definition", "id", "layout"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Definition = *all.Definition
	o.Id = all.Id
	if all.Layout != nil && all.Layout.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Layout = all.Layout

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
