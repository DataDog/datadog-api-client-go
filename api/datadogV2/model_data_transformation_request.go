// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DataTransformationRequest
type DataTransformationRequest struct {
	// Previous chat messages for iterative interaction.
	ChatHistory []ChatHistoryItem `json:"chatHistory,omitempty"`
	//
	Context *DataTransformationContext `json:"context,omitempty"`
	// The programming language for the transformation.
	Language *DataTransformationLanguage `json:"language,omitempty"`
	// Whether to stream the response.
	Stream *bool `json:"stream,omitempty"`
	// The user's prompt describing the desired transformation.
	UserPrompt string `json:"userPrompt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataTransformationRequest instantiates a new DataTransformationRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataTransformationRequest(userPrompt string) *DataTransformationRequest {
	this := DataTransformationRequest{}
	this.UserPrompt = userPrompt
	return &this
}

// NewDataTransformationRequestWithDefaults instantiates a new DataTransformationRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataTransformationRequestWithDefaults() *DataTransformationRequest {
	this := DataTransformationRequest{}
	return &this
}

// GetChatHistory returns the ChatHistory field value if set, zero value otherwise.
func (o *DataTransformationRequest) GetChatHistory() []ChatHistoryItem {
	if o == nil || o.ChatHistory == nil {
		var ret []ChatHistoryItem
		return ret
	}
	return o.ChatHistory
}

// GetChatHistoryOk returns a tuple with the ChatHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataTransformationRequest) GetChatHistoryOk() (*[]ChatHistoryItem, bool) {
	if o == nil || o.ChatHistory == nil {
		return nil, false
	}
	return &o.ChatHistory, true
}

// HasChatHistory returns a boolean if a field has been set.
func (o *DataTransformationRequest) HasChatHistory() bool {
	return o != nil && o.ChatHistory != nil
}

// SetChatHistory gets a reference to the given []ChatHistoryItem and assigns it to the ChatHistory field.
func (o *DataTransformationRequest) SetChatHistory(v []ChatHistoryItem) {
	o.ChatHistory = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *DataTransformationRequest) GetContext() DataTransformationContext {
	if o == nil || o.Context == nil {
		var ret DataTransformationContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataTransformationRequest) GetContextOk() (*DataTransformationContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *DataTransformationRequest) HasContext() bool {
	return o != nil && o.Context != nil
}

// SetContext gets a reference to the given DataTransformationContext and assigns it to the Context field.
func (o *DataTransformationRequest) SetContext(v DataTransformationContext) {
	o.Context = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *DataTransformationRequest) GetLanguage() DataTransformationLanguage {
	if o == nil || o.Language == nil {
		var ret DataTransformationLanguage
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataTransformationRequest) GetLanguageOk() (*DataTransformationLanguage, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *DataTransformationRequest) HasLanguage() bool {
	return o != nil && o.Language != nil
}

// SetLanguage gets a reference to the given DataTransformationLanguage and assigns it to the Language field.
func (o *DataTransformationRequest) SetLanguage(v DataTransformationLanguage) {
	o.Language = &v
}

// GetStream returns the Stream field value if set, zero value otherwise.
func (o *DataTransformationRequest) GetStream() bool {
	if o == nil || o.Stream == nil {
		var ret bool
		return ret
	}
	return *o.Stream
}

// GetStreamOk returns a tuple with the Stream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataTransformationRequest) GetStreamOk() (*bool, bool) {
	if o == nil || o.Stream == nil {
		return nil, false
	}
	return o.Stream, true
}

// HasStream returns a boolean if a field has been set.
func (o *DataTransformationRequest) HasStream() bool {
	return o != nil && o.Stream != nil
}

// SetStream gets a reference to the given bool and assigns it to the Stream field.
func (o *DataTransformationRequest) SetStream(v bool) {
	o.Stream = &v
}

// GetUserPrompt returns the UserPrompt field value.
func (o *DataTransformationRequest) GetUserPrompt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserPrompt
}

// GetUserPromptOk returns a tuple with the UserPrompt field value
// and a boolean to check if the value has been set.
func (o *DataTransformationRequest) GetUserPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserPrompt, true
}

// SetUserPrompt sets field value.
func (o *DataTransformationRequest) SetUserPrompt(v string) {
	o.UserPrompt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DataTransformationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ChatHistory != nil {
		toSerialize["chatHistory"] = o.ChatHistory
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.Stream != nil {
		toSerialize["stream"] = o.Stream
	}
	toSerialize["userPrompt"] = o.UserPrompt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataTransformationRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChatHistory []ChatHistoryItem           `json:"chatHistory,omitempty"`
		Context     *DataTransformationContext  `json:"context,omitempty"`
		Language    *DataTransformationLanguage `json:"language,omitempty"`
		Stream      *bool                       `json:"stream,omitempty"`
		UserPrompt  *string                     `json:"userPrompt"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.UserPrompt == nil {
		return fmt.Errorf("required field userPrompt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"chatHistory", "context", "language", "stream", "userPrompt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ChatHistory = all.ChatHistory
	if all.Context != nil && all.Context.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Context = all.Context
	if all.Language != nil && !all.Language.IsValid() {
		hasInvalidField = true
	} else {
		o.Language = all.Language
	}
	o.Stream = all.Stream
	o.UserPrompt = *all.UserPrompt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
