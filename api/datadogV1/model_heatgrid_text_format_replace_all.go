// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridTextFormatReplaceAll Replace the entire matching group label.
type HeatgridTextFormatReplaceAll struct {
	// Replace the entire matching group label.
	Type HeatgridTextFormatReplaceAllType `json:"type"`
	// Replacement string.
	With string `json:"with"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatgridTextFormatReplaceAll instantiates a new HeatgridTextFormatReplaceAll object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatgridTextFormatReplaceAll(typeVar HeatgridTextFormatReplaceAllType, with string) *HeatgridTextFormatReplaceAll {
	this := HeatgridTextFormatReplaceAll{}
	this.Type = typeVar
	this.With = with
	return &this
}

// NewHeatgridTextFormatReplaceAllWithDefaults instantiates a new HeatgridTextFormatReplaceAll object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatgridTextFormatReplaceAllWithDefaults() *HeatgridTextFormatReplaceAll {
	this := HeatgridTextFormatReplaceAll{}
	return &this
}

// GetType returns the Type field value.
func (o *HeatgridTextFormatReplaceAll) GetType() HeatgridTextFormatReplaceAllType {
	if o == nil {
		var ret HeatgridTextFormatReplaceAllType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HeatgridTextFormatReplaceAll) GetTypeOk() (*HeatgridTextFormatReplaceAllType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *HeatgridTextFormatReplaceAll) SetType(v HeatgridTextFormatReplaceAllType) {
	o.Type = v
}

// GetWith returns the With field value.
func (o *HeatgridTextFormatReplaceAll) GetWith() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.With
}

// GetWithOk returns a tuple with the With field value
// and a boolean to check if the value has been set.
func (o *HeatgridTextFormatReplaceAll) GetWithOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.With, true
}

// SetWith sets field value.
func (o *HeatgridTextFormatReplaceAll) SetWith(v string) {
	o.With = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatgridTextFormatReplaceAll) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	toSerialize["with"] = o.With
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HeatgridTextFormatReplaceAll) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type *HeatgridTextFormatReplaceAllType `json:"type"`
		With *string                           `json:"with"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.With == nil {
		return fmt.Errorf("required field with missing")
	}

	hasInvalidField := false
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.With = *all.With

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
