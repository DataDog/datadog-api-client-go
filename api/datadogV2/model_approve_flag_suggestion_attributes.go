// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApproveFlagSuggestionAttributes Approval options for the suggestion.
type ApproveFlagSuggestionAttributes struct {
	// Whether to apply the change immediately. Set to `false` to approve without applying, then
	// call `ApplyFlagSuggestion` later.
	Apply *bool `json:"apply,omitempty"`
	// A comment explaining the approval.
	Comment *string `json:"comment,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApproveFlagSuggestionAttributes instantiates a new ApproveFlagSuggestionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApproveFlagSuggestionAttributes() *ApproveFlagSuggestionAttributes {
	this := ApproveFlagSuggestionAttributes{}
	var apply bool = true
	this.Apply = &apply
	return &this
}

// NewApproveFlagSuggestionAttributesWithDefaults instantiates a new ApproveFlagSuggestionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApproveFlagSuggestionAttributesWithDefaults() *ApproveFlagSuggestionAttributes {
	this := ApproveFlagSuggestionAttributes{}
	var apply bool = true
	this.Apply = &apply
	return &this
}

// GetApply returns the Apply field value if set, zero value otherwise.
func (o *ApproveFlagSuggestionAttributes) GetApply() bool {
	if o == nil || o.Apply == nil {
		var ret bool
		return ret
	}
	return *o.Apply
}

// GetApplyOk returns a tuple with the Apply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApproveFlagSuggestionAttributes) GetApplyOk() (*bool, bool) {
	if o == nil || o.Apply == nil {
		return nil, false
	}
	return o.Apply, true
}

// HasApply returns a boolean if a field has been set.
func (o *ApproveFlagSuggestionAttributes) HasApply() bool {
	return o != nil && o.Apply != nil
}

// SetApply gets a reference to the given bool and assigns it to the Apply field.
func (o *ApproveFlagSuggestionAttributes) SetApply(v bool) {
	o.Apply = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ApproveFlagSuggestionAttributes) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApproveFlagSuggestionAttributes) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ApproveFlagSuggestionAttributes) HasComment() bool {
	return o != nil && o.Comment != nil
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ApproveFlagSuggestionAttributes) SetComment(v string) {
	o.Comment = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApproveFlagSuggestionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Apply != nil {
		toSerialize["apply"] = o.Apply
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApproveFlagSuggestionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Apply   *bool   `json:"apply,omitempty"`
		Comment *string `json:"comment,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"apply", "comment"})
	} else {
		return err
	}
	o.Apply = all.Apply
	o.Comment = all.Comment

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
