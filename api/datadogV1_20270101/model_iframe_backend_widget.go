// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IframeBackendWidget Test-only generated schema for the dashboard versioning proof of concept.
type IframeBackendWidget struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Definition IframeBackendDefinition `json:"definition"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Id *float64 `json:"id,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Layout *WidgetLayout `json:"layout,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewIframeBackendWidget instantiates a new IframeBackendWidget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIframeBackendWidget(definition IframeBackendDefinition) *IframeBackendWidget {
	this := IframeBackendWidget{}
	this.Definition = definition
	return &this
}

// NewIframeBackendWidgetWithDefaults instantiates a new IframeBackendWidget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIframeBackendWidgetWithDefaults() *IframeBackendWidget {
	this := IframeBackendWidget{}
	return &this
}

// GetDefinition returns the Definition field value.
func (o *IframeBackendWidget) GetDefinition() IframeBackendDefinition {
	if o == nil {
		var ret IframeBackendDefinition
		return ret
	}
	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *IframeBackendWidget) GetDefinitionOk() (*IframeBackendDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definition, true
}

// SetDefinition sets field value.
func (o *IframeBackendWidget) SetDefinition(v IframeBackendDefinition) {
	o.Definition = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IframeBackendWidget) GetId() float64 {
	if o == nil || o.Id == nil {
		var ret float64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IframeBackendWidget) GetIdOk() (*float64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IframeBackendWidget) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given float64 and assigns it to the Id field.
func (o *IframeBackendWidget) SetId(v float64) {
	o.Id = &v
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *IframeBackendWidget) GetLayout() WidgetLayout {
	if o == nil || o.Layout == nil {
		var ret WidgetLayout
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IframeBackendWidget) GetLayoutOk() (*WidgetLayout, bool) {
	if o == nil || o.Layout == nil {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *IframeBackendWidget) HasLayout() bool {
	return o != nil && o.Layout != nil
}

// SetLayout gets a reference to the given WidgetLayout and assigns it to the Layout field.
func (o *IframeBackendWidget) SetLayout(v WidgetLayout) {
	o.Layout = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IframeBackendWidget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["definition"] = o.Definition
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Layout != nil {
		toSerialize["layout"] = o.Layout
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IframeBackendWidget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Definition *IframeBackendDefinition `json:"definition"`
		Id         *float64                 `json:"id,omitempty"`
		Layout     *WidgetLayout            `json:"layout,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Definition == nil {
		return fmt.Errorf("required field definition missing")
	}

	hasInvalidField := false
	if all.Definition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Definition = *all.Definition
	o.Id = all.Id
	o.Layout = all.Layout

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
