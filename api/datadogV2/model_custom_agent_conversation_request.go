// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomAgentConversationRequest
type CustomAgentConversationRequest struct {
	// The conversation ID to continue an existing conversation.
	ConversationId *string `json:"conversationId,omitempty"`
	// Optional JSON schema to structure the output.
	OutputSchema interface{} `json:"outputSchema,omitempty"`
	// The user's prompt for the conversation.
	UserPrompt string `json:"userPrompt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomAgentConversationRequest instantiates a new CustomAgentConversationRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomAgentConversationRequest(userPrompt string) *CustomAgentConversationRequest {
	this := CustomAgentConversationRequest{}
	this.UserPrompt = userPrompt
	return &this
}

// NewCustomAgentConversationRequestWithDefaults instantiates a new CustomAgentConversationRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomAgentConversationRequestWithDefaults() *CustomAgentConversationRequest {
	this := CustomAgentConversationRequest{}
	return &this
}

// GetConversationId returns the ConversationId field value if set, zero value otherwise.
func (o *CustomAgentConversationRequest) GetConversationId() string {
	if o == nil || o.ConversationId == nil {
		var ret string
		return ret
	}
	return *o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAgentConversationRequest) GetConversationIdOk() (*string, bool) {
	if o == nil || o.ConversationId == nil {
		return nil, false
	}
	return o.ConversationId, true
}

// HasConversationId returns a boolean if a field has been set.
func (o *CustomAgentConversationRequest) HasConversationId() bool {
	return o != nil && o.ConversationId != nil
}

// SetConversationId gets a reference to the given string and assigns it to the ConversationId field.
func (o *CustomAgentConversationRequest) SetConversationId(v string) {
	o.ConversationId = &v
}

// GetOutputSchema returns the OutputSchema field value if set, zero value otherwise.
func (o *CustomAgentConversationRequest) GetOutputSchema() interface{} {
	if o == nil || o.OutputSchema == nil {
		var ret interface{}
		return ret
	}
	return o.OutputSchema
}

// GetOutputSchemaOk returns a tuple with the OutputSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAgentConversationRequest) GetOutputSchemaOk() (*interface{}, bool) {
	if o == nil || o.OutputSchema == nil {
		return nil, false
	}
	return &o.OutputSchema, true
}

// HasOutputSchema returns a boolean if a field has been set.
func (o *CustomAgentConversationRequest) HasOutputSchema() bool {
	return o != nil && o.OutputSchema != nil
}

// SetOutputSchema gets a reference to the given interface{} and assigns it to the OutputSchema field.
func (o *CustomAgentConversationRequest) SetOutputSchema(v interface{}) {
	o.OutputSchema = v
}

// GetUserPrompt returns the UserPrompt field value.
func (o *CustomAgentConversationRequest) GetUserPrompt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserPrompt
}

// GetUserPromptOk returns a tuple with the UserPrompt field value
// and a boolean to check if the value has been set.
func (o *CustomAgentConversationRequest) GetUserPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserPrompt, true
}

// SetUserPrompt sets field value.
func (o *CustomAgentConversationRequest) SetUserPrompt(v string) {
	o.UserPrompt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomAgentConversationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConversationId != nil {
		toSerialize["conversationId"] = o.ConversationId
	}
	if o.OutputSchema != nil {
		toSerialize["outputSchema"] = o.OutputSchema
	}
	toSerialize["userPrompt"] = o.UserPrompt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomAgentConversationRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConversationId *string     `json:"conversationId,omitempty"`
		OutputSchema   interface{} `json:"outputSchema,omitempty"`
		UserPrompt     *string     `json:"userPrompt"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.UserPrompt == nil {
		return fmt.Errorf("required field userPrompt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"conversationId", "outputSchema", "userPrompt"})
	} else {
		return err
	}
	o.ConversationId = all.ConversationId
	o.OutputSchema = all.OutputSchema
	o.UserPrompt = *all.UserPrompt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
