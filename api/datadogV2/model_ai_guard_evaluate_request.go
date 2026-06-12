// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AIGuardEvaluateRequest The evaluation request payload containing conversation messages and optional metadata.
type AIGuardEvaluateRequest struct {
	// The list of conversation messages to evaluate. Must contain at least one message.
	Messages []AIGuardMessage `json:"messages"`
	// Optional metadata providing context about the originating service and request.
	Meta *AIGuardMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAIGuardEvaluateRequest instantiates a new AIGuardEvaluateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAIGuardEvaluateRequest(messages []AIGuardMessage) *AIGuardEvaluateRequest {
	this := AIGuardEvaluateRequest{}
	this.Messages = messages
	return &this
}

// NewAIGuardEvaluateRequestWithDefaults instantiates a new AIGuardEvaluateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAIGuardEvaluateRequestWithDefaults() *AIGuardEvaluateRequest {
	this := AIGuardEvaluateRequest{}
	return &this
}

// GetMessages returns the Messages field value.
func (o *AIGuardEvaluateRequest) GetMessages() []AIGuardMessage {
	if o == nil {
		var ret []AIGuardMessage
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *AIGuardEvaluateRequest) GetMessagesOk() (*[]AIGuardMessage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Messages, true
}

// SetMessages sets field value.
func (o *AIGuardEvaluateRequest) SetMessages(v []AIGuardMessage) {
	o.Messages = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AIGuardEvaluateRequest) GetMeta() AIGuardMeta {
	if o == nil || o.Meta == nil {
		var ret AIGuardMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIGuardEvaluateRequest) GetMetaOk() (*AIGuardMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AIGuardEvaluateRequest) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given AIGuardMeta and assigns it to the Meta field.
func (o *AIGuardEvaluateRequest) SetMeta(v AIGuardMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AIGuardEvaluateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["messages"] = o.Messages
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AIGuardEvaluateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Messages *[]AIGuardMessage `json:"messages"`
		Meta     *AIGuardMeta      `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Messages == nil {
		return fmt.Errorf("required field messages missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"messages", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Messages = *all.Messages
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
