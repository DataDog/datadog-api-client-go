// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StructuredAnalysisQuery Test-only generated schema for the dashboard versioning proof of concept.
type StructuredAnalysisQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	SourceDataset string `json:"source_dataset"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Transformations []StructuredAnalysisTransformation `json:"transformations"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type StructuredAnalysisQueryType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewStructuredAnalysisQuery instantiates a new StructuredAnalysisQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStructuredAnalysisQuery(sourceDataset string, transformations []StructuredAnalysisTransformation, typeVar StructuredAnalysisQueryType) *StructuredAnalysisQuery {
	this := StructuredAnalysisQuery{}
	this.SourceDataset = sourceDataset
	this.Transformations = transformations
	this.Type = typeVar
	return &this
}

// NewStructuredAnalysisQueryWithDefaults instantiates a new StructuredAnalysisQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStructuredAnalysisQueryWithDefaults() *StructuredAnalysisQuery {
	this := StructuredAnalysisQuery{}
	return &this
}

// GetSourceDataset returns the SourceDataset field value.
func (o *StructuredAnalysisQuery) GetSourceDataset() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SourceDataset
}

// GetSourceDatasetOk returns a tuple with the SourceDataset field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisQuery) GetSourceDatasetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceDataset, true
}

// SetSourceDataset sets field value.
func (o *StructuredAnalysisQuery) SetSourceDataset(v string) {
	o.SourceDataset = v
}

// GetTransformations returns the Transformations field value.
func (o *StructuredAnalysisQuery) GetTransformations() []StructuredAnalysisTransformation {
	if o == nil {
		var ret []StructuredAnalysisTransformation
		return ret
	}
	return o.Transformations
}

// GetTransformationsOk returns a tuple with the Transformations field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisQuery) GetTransformationsOk() (*[]StructuredAnalysisTransformation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transformations, true
}

// SetTransformations sets field value.
func (o *StructuredAnalysisQuery) SetTransformations(v []StructuredAnalysisTransformation) {
	o.Transformations = v
}

// GetType returns the Type field value.
func (o *StructuredAnalysisQuery) GetType() StructuredAnalysisQueryType {
	if o == nil {
		var ret StructuredAnalysisQueryType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisQuery) GetTypeOk() (*StructuredAnalysisQueryType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *StructuredAnalysisQuery) SetType(v StructuredAnalysisQueryType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StructuredAnalysisQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["source_dataset"] = o.SourceDataset
	toSerialize["transformations"] = o.Transformations
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StructuredAnalysisQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SourceDataset   *string                             `json:"source_dataset"`
		Transformations *[]StructuredAnalysisTransformation `json:"transformations"`
		Type            *StructuredAnalysisQueryType        `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SourceDataset == nil {
		return fmt.Errorf("required field source_dataset missing")
	}
	if all.Transformations == nil {
		return fmt.Errorf("required field transformations missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.SourceDataset = *all.SourceDataset
	o.Transformations = *all.Transformations
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
