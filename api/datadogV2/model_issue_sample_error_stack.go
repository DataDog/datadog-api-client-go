// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueSampleErrorStack Parsed stack trace of the error. Only the sample picked by the `git_metadata` bias has frames.
type IssueSampleErrorStack struct {
	// Stack frames, from the top of the stack.
	Frames []IssueSampleStackFrame `json:"frames,omitempty"`
	// Programming language used to parse the stack trace.
	Language *string `json:"language,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueSampleErrorStack instantiates a new IssueSampleErrorStack object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueSampleErrorStack() *IssueSampleErrorStack {
	this := IssueSampleErrorStack{}
	return &this
}

// NewIssueSampleErrorStackWithDefaults instantiates a new IssueSampleErrorStack object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueSampleErrorStackWithDefaults() *IssueSampleErrorStack {
	this := IssueSampleErrorStack{}
	return &this
}

// GetFrames returns the Frames field value if set, zero value otherwise.
func (o *IssueSampleErrorStack) GetFrames() []IssueSampleStackFrame {
	if o == nil || o.Frames == nil {
		var ret []IssueSampleStackFrame
		return ret
	}
	return o.Frames
}

// GetFramesOk returns a tuple with the Frames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleErrorStack) GetFramesOk() (*[]IssueSampleStackFrame, bool) {
	if o == nil || o.Frames == nil {
		return nil, false
	}
	return &o.Frames, true
}

// HasFrames returns a boolean if a field has been set.
func (o *IssueSampleErrorStack) HasFrames() bool {
	return o != nil && o.Frames != nil
}

// SetFrames gets a reference to the given []IssueSampleStackFrame and assigns it to the Frames field.
func (o *IssueSampleErrorStack) SetFrames(v []IssueSampleStackFrame) {
	o.Frames = v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *IssueSampleErrorStack) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleErrorStack) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *IssueSampleErrorStack) HasLanguage() bool {
	return o != nil && o.Language != nil
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *IssueSampleErrorStack) SetLanguage(v string) {
	o.Language = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueSampleErrorStack) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Frames != nil {
		toSerialize["frames"] = o.Frames
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueSampleErrorStack) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Frames   []IssueSampleStackFrame `json:"frames,omitempty"`
		Language *string                 `json:"language,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"frames", "language"})
	} else {
		return err
	}
	o.Frames = all.Frames
	o.Language = all.Language

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
