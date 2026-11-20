// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StructuredAnalysisFilter Test-only generated schema for the dashboard versioning proof of concept.
type StructuredAnalysisFilter struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Filter string `json:"filter"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type StructuredAnalysisFilterType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewStructuredAnalysisFilter instantiates a new StructuredAnalysisFilter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStructuredAnalysisFilter(filter string, typeVar StructuredAnalysisFilterType) *StructuredAnalysisFilter {
	this := StructuredAnalysisFilter{}
	this.Filter = filter
	this.Type = typeVar
	return &this
}

// NewStructuredAnalysisFilterWithDefaults instantiates a new StructuredAnalysisFilter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStructuredAnalysisFilterWithDefaults() *StructuredAnalysisFilter {
	this := StructuredAnalysisFilter{}
	return &this
}

// GetFilter returns the Filter field value.
func (o *StructuredAnalysisFilter) GetFilter() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisFilter) GetFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value.
func (o *StructuredAnalysisFilter) SetFilter(v string) {
	o.Filter = v
}

// GetType returns the Type field value.
func (o *StructuredAnalysisFilter) GetType() StructuredAnalysisFilterType {
	if o == nil {
		var ret StructuredAnalysisFilterType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisFilter) GetTypeOk() (*StructuredAnalysisFilterType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *StructuredAnalysisFilter) SetType(v StructuredAnalysisFilterType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StructuredAnalysisFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["filter"] = o.Filter
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StructuredAnalysisFilter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filter *string                       `json:"filter"`
		Type   *StructuredAnalysisFilterType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Filter == nil {
		return fmt.Errorf("required field filter missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Filter = *all.Filter
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
