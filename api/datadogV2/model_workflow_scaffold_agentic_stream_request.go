// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowScaffoldAgenticStreamRequest
type WorkflowScaffoldAgenticStreamRequest struct {
	// Previous chat messages for iterative workflow generation.
	ChatHistory []ChatMessage `json:"chatHistory,omitempty"`
	// The existing workflow specification to modify.
	ExistingWorkflow interface{} `json:"existingWorkflow,omitempty"`
	// The previous action taken in the workflow generation.
	PreviousAction *string `json:"previousAction,omitempty"`
	//
	UserContext *UserContext `json:"userContext,omitempty"`
	// The user's prompt for generating or modifying the workflow.
	UserPrompt string `json:"userPrompt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWorkflowScaffoldAgenticStreamRequest instantiates a new WorkflowScaffoldAgenticStreamRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWorkflowScaffoldAgenticStreamRequest(userPrompt string) *WorkflowScaffoldAgenticStreamRequest {
	this := WorkflowScaffoldAgenticStreamRequest{}
	this.UserPrompt = userPrompt
	return &this
}

// NewWorkflowScaffoldAgenticStreamRequestWithDefaults instantiates a new WorkflowScaffoldAgenticStreamRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWorkflowScaffoldAgenticStreamRequestWithDefaults() *WorkflowScaffoldAgenticStreamRequest {
	this := WorkflowScaffoldAgenticStreamRequest{}
	return &this
}

// GetChatHistory returns the ChatHistory field value if set, zero value otherwise.
func (o *WorkflowScaffoldAgenticStreamRequest) GetChatHistory() []ChatMessage {
	if o == nil || o.ChatHistory == nil {
		var ret []ChatMessage
		return ret
	}
	return o.ChatHistory
}

// GetChatHistoryOk returns a tuple with the ChatHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowScaffoldAgenticStreamRequest) GetChatHistoryOk() (*[]ChatMessage, bool) {
	if o == nil || o.ChatHistory == nil {
		return nil, false
	}
	return &o.ChatHistory, true
}

// HasChatHistory returns a boolean if a field has been set.
func (o *WorkflowScaffoldAgenticStreamRequest) HasChatHistory() bool {
	return o != nil && o.ChatHistory != nil
}

// SetChatHistory gets a reference to the given []ChatMessage and assigns it to the ChatHistory field.
func (o *WorkflowScaffoldAgenticStreamRequest) SetChatHistory(v []ChatMessage) {
	o.ChatHistory = v
}

// GetExistingWorkflow returns the ExistingWorkflow field value if set, zero value otherwise.
func (o *WorkflowScaffoldAgenticStreamRequest) GetExistingWorkflow() interface{} {
	if o == nil || o.ExistingWorkflow == nil {
		var ret interface{}
		return ret
	}
	return o.ExistingWorkflow
}

// GetExistingWorkflowOk returns a tuple with the ExistingWorkflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowScaffoldAgenticStreamRequest) GetExistingWorkflowOk() (*interface{}, bool) {
	if o == nil || o.ExistingWorkflow == nil {
		return nil, false
	}
	return &o.ExistingWorkflow, true
}

// HasExistingWorkflow returns a boolean if a field has been set.
func (o *WorkflowScaffoldAgenticStreamRequest) HasExistingWorkflow() bool {
	return o != nil && o.ExistingWorkflow != nil
}

// SetExistingWorkflow gets a reference to the given interface{} and assigns it to the ExistingWorkflow field.
func (o *WorkflowScaffoldAgenticStreamRequest) SetExistingWorkflow(v interface{}) {
	o.ExistingWorkflow = v
}

// GetPreviousAction returns the PreviousAction field value if set, zero value otherwise.
func (o *WorkflowScaffoldAgenticStreamRequest) GetPreviousAction() string {
	if o == nil || o.PreviousAction == nil {
		var ret string
		return ret
	}
	return *o.PreviousAction
}

// GetPreviousActionOk returns a tuple with the PreviousAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowScaffoldAgenticStreamRequest) GetPreviousActionOk() (*string, bool) {
	if o == nil || o.PreviousAction == nil {
		return nil, false
	}
	return o.PreviousAction, true
}

// HasPreviousAction returns a boolean if a field has been set.
func (o *WorkflowScaffoldAgenticStreamRequest) HasPreviousAction() bool {
	return o != nil && o.PreviousAction != nil
}

// SetPreviousAction gets a reference to the given string and assigns it to the PreviousAction field.
func (o *WorkflowScaffoldAgenticStreamRequest) SetPreviousAction(v string) {
	o.PreviousAction = &v
}

// GetUserContext returns the UserContext field value if set, zero value otherwise.
func (o *WorkflowScaffoldAgenticStreamRequest) GetUserContext() UserContext {
	if o == nil || o.UserContext == nil {
		var ret UserContext
		return ret
	}
	return *o.UserContext
}

// GetUserContextOk returns a tuple with the UserContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowScaffoldAgenticStreamRequest) GetUserContextOk() (*UserContext, bool) {
	if o == nil || o.UserContext == nil {
		return nil, false
	}
	return o.UserContext, true
}

// HasUserContext returns a boolean if a field has been set.
func (o *WorkflowScaffoldAgenticStreamRequest) HasUserContext() bool {
	return o != nil && o.UserContext != nil
}

// SetUserContext gets a reference to the given UserContext and assigns it to the UserContext field.
func (o *WorkflowScaffoldAgenticStreamRequest) SetUserContext(v UserContext) {
	o.UserContext = &v
}

// GetUserPrompt returns the UserPrompt field value.
func (o *WorkflowScaffoldAgenticStreamRequest) GetUserPrompt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserPrompt
}

// GetUserPromptOk returns a tuple with the UserPrompt field value
// and a boolean to check if the value has been set.
func (o *WorkflowScaffoldAgenticStreamRequest) GetUserPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserPrompt, true
}

// SetUserPrompt sets field value.
func (o *WorkflowScaffoldAgenticStreamRequest) SetUserPrompt(v string) {
	o.UserPrompt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WorkflowScaffoldAgenticStreamRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ChatHistory != nil {
		toSerialize["chatHistory"] = o.ChatHistory
	}
	if o.ExistingWorkflow != nil {
		toSerialize["existingWorkflow"] = o.ExistingWorkflow
	}
	if o.PreviousAction != nil {
		toSerialize["previousAction"] = o.PreviousAction
	}
	if o.UserContext != nil {
		toSerialize["userContext"] = o.UserContext
	}
	toSerialize["userPrompt"] = o.UserPrompt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WorkflowScaffoldAgenticStreamRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChatHistory      []ChatMessage `json:"chatHistory,omitempty"`
		ExistingWorkflow interface{}   `json:"existingWorkflow,omitempty"`
		PreviousAction   *string       `json:"previousAction,omitempty"`
		UserContext      *UserContext  `json:"userContext,omitempty"`
		UserPrompt       *string       `json:"userPrompt"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.UserPrompt == nil {
		return fmt.Errorf("required field userPrompt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"chatHistory", "existingWorkflow", "previousAction", "userContext", "userPrompt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ChatHistory = all.ChatHistory
	o.ExistingWorkflow = all.ExistingWorkflow
	o.PreviousAction = all.PreviousAction
	if all.UserContext != nil && all.UserContext.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UserContext = all.UserContext
	o.UserPrompt = *all.UserPrompt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
