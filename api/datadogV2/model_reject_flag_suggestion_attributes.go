// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RejectFlagSuggestionAttributes Rejection options for the suggestion.
type RejectFlagSuggestionAttributes struct {
	// A comment explaining the rejection.
	Comment *string `json:"comment,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRejectFlagSuggestionAttributes instantiates a new RejectFlagSuggestionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRejectFlagSuggestionAttributes() *RejectFlagSuggestionAttributes {
	this := RejectFlagSuggestionAttributes{}
	return &this
}

// NewRejectFlagSuggestionAttributesWithDefaults instantiates a new RejectFlagSuggestionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRejectFlagSuggestionAttributesWithDefaults() *RejectFlagSuggestionAttributes {
	this := RejectFlagSuggestionAttributes{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *RejectFlagSuggestionAttributes) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RejectFlagSuggestionAttributes) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *RejectFlagSuggestionAttributes) HasComment() bool {
	return o != nil && o.Comment != nil
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *RejectFlagSuggestionAttributes) SetComment(v string) {
	o.Comment = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RejectFlagSuggestionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
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
func (o *RejectFlagSuggestionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Comment *string `json:"comment,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"comment"})
	} else {
		return err
	}
	o.Comment = all.Comment

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
