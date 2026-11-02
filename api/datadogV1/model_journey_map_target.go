// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JourneyMapTarget The pair of journey steps the paths are visualized between.
type JourneyMapTarget struct {
	// Key of the end step in `search.node_objects`.
	End string `json:"end"`
	// Key of the start step in `search.node_objects`.
	Start string `json:"start"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewJourneyMapTarget instantiates a new JourneyMapTarget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJourneyMapTarget(end string, start string) *JourneyMapTarget {
	this := JourneyMapTarget{}
	this.End = end
	this.Start = start
	return &this
}

// NewJourneyMapTargetWithDefaults instantiates a new JourneyMapTarget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJourneyMapTargetWithDefaults() *JourneyMapTarget {
	this := JourneyMapTarget{}
	return &this
}

// GetEnd returns the End field value.
func (o *JourneyMapTarget) GetEnd() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *JourneyMapTarget) GetEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value.
func (o *JourneyMapTarget) SetEnd(v string) {
	o.End = v
}

// GetStart returns the Start field value.
func (o *JourneyMapTarget) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *JourneyMapTarget) GetStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value.
func (o *JourneyMapTarget) SetStart(v string) {
	o.Start = v
}

// MarshalJSON serializes the struct using spec logic.
func (o JourneyMapTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["end"] = o.End
	toSerialize["start"] = o.Start
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *JourneyMapTarget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		End   *string `json:"end"`
		Start *string `json:"start"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.End == nil {
		return fmt.Errorf("required field end missing")
	}
	if all.Start == nil {
		return fmt.Errorf("required field start missing")
	}
	o.End = *all.End
	o.Start = *all.Start

	return nil
}
