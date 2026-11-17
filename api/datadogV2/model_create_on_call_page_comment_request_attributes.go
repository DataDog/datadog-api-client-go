// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateOnCallPageCommentRequestAttributes Attributes for creating a comment on an On-Call Page.
type CreateOnCallPageCommentRequestAttributes struct {
	// The content of the comment.
	Body string `json:"body"`
	// The IDs of the users mentioned in the comment.
	MentionedUserIds []uuid.UUID `json:"mentioned_user_ids,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateOnCallPageCommentRequestAttributes instantiates a new CreateOnCallPageCommentRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateOnCallPageCommentRequestAttributes(body string) *CreateOnCallPageCommentRequestAttributes {
	this := CreateOnCallPageCommentRequestAttributes{}
	this.Body = body
	return &this
}

// NewCreateOnCallPageCommentRequestAttributesWithDefaults instantiates a new CreateOnCallPageCommentRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateOnCallPageCommentRequestAttributesWithDefaults() *CreateOnCallPageCommentRequestAttributes {
	this := CreateOnCallPageCommentRequestAttributes{}
	return &this
}

// GetBody returns the Body field value.
func (o *CreateOnCallPageCommentRequestAttributes) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *CreateOnCallPageCommentRequestAttributes) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value.
func (o *CreateOnCallPageCommentRequestAttributes) SetBody(v string) {
	o.Body = v
}

// GetMentionedUserIds returns the MentionedUserIds field value if set, zero value otherwise.
func (o *CreateOnCallPageCommentRequestAttributes) GetMentionedUserIds() []uuid.UUID {
	if o == nil || o.MentionedUserIds == nil {
		var ret []uuid.UUID
		return ret
	}
	return o.MentionedUserIds
}

// GetMentionedUserIdsOk returns a tuple with the MentionedUserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOnCallPageCommentRequestAttributes) GetMentionedUserIdsOk() (*[]uuid.UUID, bool) {
	if o == nil || o.MentionedUserIds == nil {
		return nil, false
	}
	return &o.MentionedUserIds, true
}

// HasMentionedUserIds returns a boolean if a field has been set.
func (o *CreateOnCallPageCommentRequestAttributes) HasMentionedUserIds() bool {
	return o != nil && o.MentionedUserIds != nil
}

// SetMentionedUserIds gets a reference to the given []uuid.UUID and assigns it to the MentionedUserIds field.
func (o *CreateOnCallPageCommentRequestAttributes) SetMentionedUserIds(v []uuid.UUID) {
	o.MentionedUserIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateOnCallPageCommentRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["body"] = o.Body
	if o.MentionedUserIds != nil {
		toSerialize["mentioned_user_ids"] = o.MentionedUserIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateOnCallPageCommentRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Body             *string     `json:"body"`
		MentionedUserIds []uuid.UUID `json:"mentioned_user_ids,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Body == nil {
		return fmt.Errorf("required field body missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"body", "mentioned_user_ids"})
	} else {
		return err
	}
	o.Body = *all.Body
	o.MentionedUserIds = all.MentionedUserIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
