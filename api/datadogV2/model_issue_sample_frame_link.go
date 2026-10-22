// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueSampleFrameLink Source code location of a stack frame.
type IssueSampleFrameLink struct {
	// Location of a stack frame in a repository file. Empty when the source code is unavailable.
	CodeLocation IssueSampleGitFileLocation `json:"code_location"`
	// Index of the frame in the stack trace, starting at 0.
	FrameId int64 `json:"frame_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueSampleFrameLink instantiates a new IssueSampleFrameLink object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueSampleFrameLink(codeLocation IssueSampleGitFileLocation, frameId int64) *IssueSampleFrameLink {
	this := IssueSampleFrameLink{}
	this.CodeLocation = codeLocation
	this.FrameId = frameId
	return &this
}

// NewIssueSampleFrameLinkWithDefaults instantiates a new IssueSampleFrameLink object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueSampleFrameLinkWithDefaults() *IssueSampleFrameLink {
	this := IssueSampleFrameLink{}
	return &this
}

// GetCodeLocation returns the CodeLocation field value.
func (o *IssueSampleFrameLink) GetCodeLocation() IssueSampleGitFileLocation {
	if o == nil {
		var ret IssueSampleGitFileLocation
		return ret
	}
	return o.CodeLocation
}

// GetCodeLocationOk returns a tuple with the CodeLocation field value
// and a boolean to check if the value has been set.
func (o *IssueSampleFrameLink) GetCodeLocationOk() (*IssueSampleGitFileLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CodeLocation, true
}

// SetCodeLocation sets field value.
func (o *IssueSampleFrameLink) SetCodeLocation(v IssueSampleGitFileLocation) {
	o.CodeLocation = v
}

// GetFrameId returns the FrameId field value.
func (o *IssueSampleFrameLink) GetFrameId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.FrameId
}

// GetFrameIdOk returns a tuple with the FrameId field value
// and a boolean to check if the value has been set.
func (o *IssueSampleFrameLink) GetFrameIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FrameId, true
}

// SetFrameId sets field value.
func (o *IssueSampleFrameLink) SetFrameId(v int64) {
	o.FrameId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueSampleFrameLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["code_location"] = o.CodeLocation
	toSerialize["frame_id"] = o.FrameId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueSampleFrameLink) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CodeLocation *IssueSampleGitFileLocation `json:"code_location"`
		FrameId      *int64                      `json:"frame_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CodeLocation == nil {
		return fmt.Errorf("required field code_location missing")
	}
	if all.FrameId == nil {
		return fmt.Errorf("required field frame_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"code_location", "frame_id"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CodeLocation.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CodeLocation = *all.CodeLocation
	o.FrameId = *all.FrameId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
