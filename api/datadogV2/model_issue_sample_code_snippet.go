// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueSampleCodeSnippet Source code around a stack frame. Empty when the source code is unavailable.
type IssueSampleCodeSnippet struct {
	// Lines of source code.
	Lines []string `json:"lines,omitempty"`
	// Line number of the first line of the snippet, starting at 1.
	Offset *int64 `json:"offset,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueSampleCodeSnippet instantiates a new IssueSampleCodeSnippet object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueSampleCodeSnippet() *IssueSampleCodeSnippet {
	this := IssueSampleCodeSnippet{}
	return &this
}

// NewIssueSampleCodeSnippetWithDefaults instantiates a new IssueSampleCodeSnippet object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueSampleCodeSnippetWithDefaults() *IssueSampleCodeSnippet {
	this := IssueSampleCodeSnippet{}
	return &this
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *IssueSampleCodeSnippet) GetLines() []string {
	if o == nil || o.Lines == nil {
		var ret []string
		return ret
	}
	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleCodeSnippet) GetLinesOk() (*[]string, bool) {
	if o == nil || o.Lines == nil {
		return nil, false
	}
	return &o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *IssueSampleCodeSnippet) HasLines() bool {
	return o != nil && o.Lines != nil
}

// SetLines gets a reference to the given []string and assigns it to the Lines field.
func (o *IssueSampleCodeSnippet) SetLines(v []string) {
	o.Lines = v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *IssueSampleCodeSnippet) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleCodeSnippet) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *IssueSampleCodeSnippet) HasOffset() bool {
	return o != nil && o.Offset != nil
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *IssueSampleCodeSnippet) SetOffset(v int64) {
	o.Offset = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueSampleCodeSnippet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Lines != nil {
		toSerialize["lines"] = o.Lines
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueSampleCodeSnippet) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Lines  []string `json:"lines,omitempty"`
		Offset *int64   `json:"offset,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"lines", "offset"})
	} else {
		return err
	}
	o.Lines = all.Lines
	o.Offset = all.Offset

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
