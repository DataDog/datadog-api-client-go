// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DistributionBackendRequest20270101OneOf4455772896 Test-only generated schema for the dashboard versioning proof of concept.
type DistributionBackendRequest20270101OneOf4455772896 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ProcessQuery ProcessQueryDefinition `json:"process_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *RequestStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDistributionBackendRequest20270101OneOf4455772896 instantiates a new DistributionBackendRequest20270101OneOf4455772896 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDistributionBackendRequest20270101OneOf4455772896(processQuery ProcessQueryDefinition) *DistributionBackendRequest20270101OneOf4455772896 {
	this := DistributionBackendRequest20270101OneOf4455772896{}
	this.ProcessQuery = processQuery
	return &this
}

// NewDistributionBackendRequest20270101OneOf4455772896WithDefaults instantiates a new DistributionBackendRequest20270101OneOf4455772896 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDistributionBackendRequest20270101OneOf4455772896WithDefaults() *DistributionBackendRequest20270101OneOf4455772896 {
	this := DistributionBackendRequest20270101OneOf4455772896{}
	return &this
}

// GetProcessQuery returns the ProcessQuery field value.
func (o *DistributionBackendRequest20270101OneOf4455772896) GetProcessQuery() ProcessQueryDefinition {
	if o == nil {
		var ret ProcessQueryDefinition
		return ret
	}
	return o.ProcessQuery
}

// GetProcessQueryOk returns a tuple with the ProcessQuery field value
// and a boolean to check if the value has been set.
func (o *DistributionBackendRequest20270101OneOf4455772896) GetProcessQueryOk() (*ProcessQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessQuery, true
}

// SetProcessQuery sets field value.
func (o *DistributionBackendRequest20270101OneOf4455772896) SetProcessQuery(v ProcessQueryDefinition) {
	o.ProcessQuery = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *DistributionBackendRequest20270101OneOf4455772896) GetStyle() RequestStyle {
	if o == nil || o.Style == nil {
		var ret RequestStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionBackendRequest20270101OneOf4455772896) GetStyleOk() (*RequestStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *DistributionBackendRequest20270101OneOf4455772896) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given RequestStyle and assigns it to the Style field.
func (o *DistributionBackendRequest20270101OneOf4455772896) SetStyle(v RequestStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DistributionBackendRequest20270101OneOf4455772896) MarshalJSON() ([]byte, error) {
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
func (o *DistributionBackendRequest20270101OneOf4455772896) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ProcessQuery *ProcessQueryDefinition `json:"process_query"`
		Style        *RequestStyle           `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ProcessQuery == nil {
		return fmt.Errorf("required field process_query missing")
	}

	hasInvalidField := false
	if all.ProcessQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ProcessQuery = *all.ProcessQuery
	if all.Style != nil && all.Style.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Style = all.Style

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
