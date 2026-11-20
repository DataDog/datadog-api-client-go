// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GraphTargetPath Test-only generated schema for the dashboard versioning proof of concept.
type GraphTargetPath struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	End *string `json:"end,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Start string `json:"start"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type GraphTargetPathType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewGraphTargetPath instantiates a new GraphTargetPath object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGraphTargetPath(start string, typeVar GraphTargetPathType) *GraphTargetPath {
	this := GraphTargetPath{}
	this.Start = start
	this.Type = typeVar
	return &this
}

// NewGraphTargetPathWithDefaults instantiates a new GraphTargetPath object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGraphTargetPathWithDefaults() *GraphTargetPath {
	this := GraphTargetPath{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *GraphTargetPath) GetEnd() string {
	if o == nil || o.End == nil {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphTargetPath) GetEndOk() (*string, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *GraphTargetPath) HasEnd() bool {
	return o != nil && o.End != nil
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *GraphTargetPath) SetEnd(v string) {
	o.End = &v
}

// GetStart returns the Start field value.
func (o *GraphTargetPath) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *GraphTargetPath) GetStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value.
func (o *GraphTargetPath) SetStart(v string) {
	o.Start = v
}

// GetType returns the Type field value.
func (o *GraphTargetPath) GetType() GraphTargetPathType {
	if o == nil {
		var ret GraphTargetPathType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GraphTargetPath) GetTypeOk() (*GraphTargetPathType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *GraphTargetPath) SetType(v GraphTargetPathType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GraphTargetPath) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	toSerialize["start"] = o.Start
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GraphTargetPath) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		End   *string              `json:"end,omitempty"`
		Start *string              `json:"start"`
		Type  *GraphTargetPathType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Start == nil {
		return fmt.Errorf("required field start missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.End = all.End
	o.Start = *all.Start
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
