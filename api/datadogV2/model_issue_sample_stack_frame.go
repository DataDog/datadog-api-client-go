// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueSampleStackFrame A stack frame.
type IssueSampleStackFrame struct {
	// Source code around a stack frame. Empty when the source code is unavailable.
	CodeSnippet *IssueSampleCodeSnippet `json:"code_snippet,omitempty"`
	// Location of a stack frame in a repository file. Empty when the source code is unavailable.
	GitFileLocation *IssueSampleGitFileLocation `json:"git_file_location,omitempty"`
	// Text of the frame, as it appears in the stack trace.
	Text *string `json:"text,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueSampleStackFrame instantiates a new IssueSampleStackFrame object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueSampleStackFrame() *IssueSampleStackFrame {
	this := IssueSampleStackFrame{}
	return &this
}

// NewIssueSampleStackFrameWithDefaults instantiates a new IssueSampleStackFrame object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueSampleStackFrameWithDefaults() *IssueSampleStackFrame {
	this := IssueSampleStackFrame{}
	return &this
}

// GetCodeSnippet returns the CodeSnippet field value if set, zero value otherwise.
func (o *IssueSampleStackFrame) GetCodeSnippet() IssueSampleCodeSnippet {
	if o == nil || o.CodeSnippet == nil {
		var ret IssueSampleCodeSnippet
		return ret
	}
	return *o.CodeSnippet
}

// GetCodeSnippetOk returns a tuple with the CodeSnippet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleStackFrame) GetCodeSnippetOk() (*IssueSampleCodeSnippet, bool) {
	if o == nil || o.CodeSnippet == nil {
		return nil, false
	}
	return o.CodeSnippet, true
}

// HasCodeSnippet returns a boolean if a field has been set.
func (o *IssueSampleStackFrame) HasCodeSnippet() bool {
	return o != nil && o.CodeSnippet != nil
}

// SetCodeSnippet gets a reference to the given IssueSampleCodeSnippet and assigns it to the CodeSnippet field.
func (o *IssueSampleStackFrame) SetCodeSnippet(v IssueSampleCodeSnippet) {
	o.CodeSnippet = &v
}

// GetGitFileLocation returns the GitFileLocation field value if set, zero value otherwise.
func (o *IssueSampleStackFrame) GetGitFileLocation() IssueSampleGitFileLocation {
	if o == nil || o.GitFileLocation == nil {
		var ret IssueSampleGitFileLocation
		return ret
	}
	return *o.GitFileLocation
}

// GetGitFileLocationOk returns a tuple with the GitFileLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleStackFrame) GetGitFileLocationOk() (*IssueSampleGitFileLocation, bool) {
	if o == nil || o.GitFileLocation == nil {
		return nil, false
	}
	return o.GitFileLocation, true
}

// HasGitFileLocation returns a boolean if a field has been set.
func (o *IssueSampleStackFrame) HasGitFileLocation() bool {
	return o != nil && o.GitFileLocation != nil
}

// SetGitFileLocation gets a reference to the given IssueSampleGitFileLocation and assigns it to the GitFileLocation field.
func (o *IssueSampleStackFrame) SetGitFileLocation(v IssueSampleGitFileLocation) {
	o.GitFileLocation = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *IssueSampleStackFrame) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleStackFrame) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *IssueSampleStackFrame) HasText() bool {
	return o != nil && o.Text != nil
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *IssueSampleStackFrame) SetText(v string) {
	o.Text = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueSampleStackFrame) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CodeSnippet != nil {
		toSerialize["code_snippet"] = o.CodeSnippet
	}
	if o.GitFileLocation != nil {
		toSerialize["git_file_location"] = o.GitFileLocation
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueSampleStackFrame) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CodeSnippet     *IssueSampleCodeSnippet     `json:"code_snippet,omitempty"`
		GitFileLocation *IssueSampleGitFileLocation `json:"git_file_location,omitempty"`
		Text            *string                     `json:"text,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"code_snippet", "git_file_location", "text"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CodeSnippet != nil && all.CodeSnippet.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CodeSnippet = all.CodeSnippet
	if all.GitFileLocation != nil && all.GitFileLocation.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.GitFileLocation = all.GitFileLocation
	o.Text = all.Text

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
