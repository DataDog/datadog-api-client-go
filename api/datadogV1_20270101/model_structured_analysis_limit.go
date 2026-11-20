// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StructuredAnalysisLimit Test-only generated schema for the dashboard versioning proof of concept.
type StructuredAnalysisLimit struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit float64 `json:"limit"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type StructuredAnalysisLimitType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewStructuredAnalysisLimit instantiates a new StructuredAnalysisLimit object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStructuredAnalysisLimit(limit float64, typeVar StructuredAnalysisLimitType) *StructuredAnalysisLimit {
	this := StructuredAnalysisLimit{}
	this.Limit = limit
	this.Type = typeVar
	return &this
}

// NewStructuredAnalysisLimitWithDefaults instantiates a new StructuredAnalysisLimit object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStructuredAnalysisLimitWithDefaults() *StructuredAnalysisLimit {
	this := StructuredAnalysisLimit{}
	return &this
}

// GetLimit returns the Limit field value.
func (o *StructuredAnalysisLimit) GetLimit() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisLimit) GetLimitOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value.
func (o *StructuredAnalysisLimit) SetLimit(v float64) {
	o.Limit = v
}

// GetType returns the Type field value.
func (o *StructuredAnalysisLimit) GetType() StructuredAnalysisLimitType {
	if o == nil {
		var ret StructuredAnalysisLimitType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisLimit) GetTypeOk() (*StructuredAnalysisLimitType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *StructuredAnalysisLimit) SetType(v StructuredAnalysisLimitType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StructuredAnalysisLimit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["limit"] = o.Limit
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StructuredAnalysisLimit) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Limit *float64                     `json:"limit"`
		Type  *StructuredAnalysisLimitType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Limit == nil {
		return fmt.Errorf("required field limit missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Limit = *all.Limit
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
