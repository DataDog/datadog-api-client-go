// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOCountSpec A count-based SLI specification.
type SLOCountSpec struct {
	// A count-based SLI specification, composed of three parts: the good events formula, the total events formula,
	// and the involved queries.
	Count SLOCountCondition `json:"count"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSLOCountSpec instantiates a new SLOCountSpec object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSLOCountSpec(count SLOCountCondition) *SLOCountSpec {
	this := SLOCountSpec{}
	this.Count = count
	return &this
}

// NewSLOCountSpecWithDefaults instantiates a new SLOCountSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSLOCountSpecWithDefaults() *SLOCountSpec {
	this := SLOCountSpec{}
	return &this
}

// GetCount returns the Count field value.
func (o *SLOCountSpec) GetCount() SLOCountCondition {
	if o == nil {
		var ret SLOCountCondition
		return ret
	}
	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *SLOCountSpec) GetCountOk() (*SLOCountCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *SLOCountSpec) SetCount(v SLOCountCondition) {
	o.Count = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SLOCountSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["count"] = o.Count
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SLOCountSpec) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Count *SLOCountCondition `json:"count"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Count == nil {
		return fmt.Errorf("required field count missing")
	}

	hasInvalidField := false
	if all.Count.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Count = *all.Count

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
