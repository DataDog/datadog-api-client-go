// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatmapBackendRequest20270101OneOf4458139936 Test-only generated schema for the dashboard versioning proof of concept.
type HeatmapBackendRequest20270101OneOf4458139936 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ProcessQuery ProcessQueryDefinition_20270101 `json:"process_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *RequestStyle_20270101 `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatmapBackendRequest20270101OneOf4458139936 instantiates a new HeatmapBackendRequest20270101OneOf4458139936 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatmapBackendRequest20270101OneOf4458139936(processQuery ProcessQueryDefinition_20270101) *HeatmapBackendRequest20270101OneOf4458139936 {
	this := HeatmapBackendRequest20270101OneOf4458139936{}
	this.ProcessQuery = processQuery
	return &this
}

// NewHeatmapBackendRequest20270101OneOf4458139936WithDefaults instantiates a new HeatmapBackendRequest20270101OneOf4458139936 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatmapBackendRequest20270101OneOf4458139936WithDefaults() *HeatmapBackendRequest20270101OneOf4458139936 {
	this := HeatmapBackendRequest20270101OneOf4458139936{}
	return &this
}

// GetProcessQuery returns the ProcessQuery field value.
func (o *HeatmapBackendRequest20270101OneOf4458139936) GetProcessQuery() ProcessQueryDefinition_20270101 {
	if o == nil {
		var ret ProcessQueryDefinition_20270101
		return ret
	}
	return o.ProcessQuery
}

// GetProcessQueryOk returns a tuple with the ProcessQuery field value
// and a boolean to check if the value has been set.
func (o *HeatmapBackendRequest20270101OneOf4458139936) GetProcessQueryOk() (*ProcessQueryDefinition_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessQuery, true
}

// SetProcessQuery sets field value.
func (o *HeatmapBackendRequest20270101OneOf4458139936) SetProcessQuery(v ProcessQueryDefinition_20270101) {
	o.ProcessQuery = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *HeatmapBackendRequest20270101OneOf4458139936) GetStyle() RequestStyle_20270101 {
	if o == nil || o.Style == nil {
		var ret RequestStyle_20270101
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatmapBackendRequest20270101OneOf4458139936) GetStyleOk() (*RequestStyle_20270101, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *HeatmapBackendRequest20270101OneOf4458139936) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given RequestStyle_20270101 and assigns it to the Style field.
func (o *HeatmapBackendRequest20270101OneOf4458139936) SetStyle(v RequestStyle_20270101) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatmapBackendRequest20270101OneOf4458139936) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["process_query"] = o.ProcessQuery
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HeatmapBackendRequest20270101OneOf4458139936) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ProcessQuery *ProcessQueryDefinition_20270101 `json:"process_query"`
		Style        *RequestStyle_20270101           `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ProcessQuery == nil {
		return fmt.Errorf("required field process_query missing")
	}
	o.ProcessQuery = *all.ProcessQuery
	o.Style = all.Style

	return nil
}
