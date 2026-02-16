// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PickActionRequest
type PickActionRequest struct {
	// The client type for action filtering.
	Client *ClientType `json:"client,omitempty"`
	// The number of relevant actions to return.
	NumberOfRelevantActions int64 `json:"number_of_relevant_actions"`
	// The stability level for action filtering.
	Stability *StabilityLevel `json:"stability,omitempty"`
	// The user's prompt to find relevant actions.
	UserPrompt string `json:"user_prompt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPickActionRequest instantiates a new PickActionRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPickActionRequest(numberOfRelevantActions int64, userPrompt string) *PickActionRequest {
	this := PickActionRequest{}
	this.NumberOfRelevantActions = numberOfRelevantActions
	this.UserPrompt = userPrompt
	return &this
}

// NewPickActionRequestWithDefaults instantiates a new PickActionRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPickActionRequestWithDefaults() *PickActionRequest {
	this := PickActionRequest{}
	return &this
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *PickActionRequest) GetClient() ClientType {
	if o == nil || o.Client == nil {
		var ret ClientType
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PickActionRequest) GetClientOk() (*ClientType, bool) {
	if o == nil || o.Client == nil {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *PickActionRequest) HasClient() bool {
	return o != nil && o.Client != nil
}

// SetClient gets a reference to the given ClientType and assigns it to the Client field.
func (o *PickActionRequest) SetClient(v ClientType) {
	o.Client = &v
}

// GetNumberOfRelevantActions returns the NumberOfRelevantActions field value.
func (o *PickActionRequest) GetNumberOfRelevantActions() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.NumberOfRelevantActions
}

// GetNumberOfRelevantActionsOk returns a tuple with the NumberOfRelevantActions field value
// and a boolean to check if the value has been set.
func (o *PickActionRequest) GetNumberOfRelevantActionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfRelevantActions, true
}

// SetNumberOfRelevantActions sets field value.
func (o *PickActionRequest) SetNumberOfRelevantActions(v int64) {
	o.NumberOfRelevantActions = v
}

// GetStability returns the Stability field value if set, zero value otherwise.
func (o *PickActionRequest) GetStability() StabilityLevel {
	if o == nil || o.Stability == nil {
		var ret StabilityLevel
		return ret
	}
	return *o.Stability
}

// GetStabilityOk returns a tuple with the Stability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PickActionRequest) GetStabilityOk() (*StabilityLevel, bool) {
	if o == nil || o.Stability == nil {
		return nil, false
	}
	return o.Stability, true
}

// HasStability returns a boolean if a field has been set.
func (o *PickActionRequest) HasStability() bool {
	return o != nil && o.Stability != nil
}

// SetStability gets a reference to the given StabilityLevel and assigns it to the Stability field.
func (o *PickActionRequest) SetStability(v StabilityLevel) {
	o.Stability = &v
}

// GetUserPrompt returns the UserPrompt field value.
func (o *PickActionRequest) GetUserPrompt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserPrompt
}

// GetUserPromptOk returns a tuple with the UserPrompt field value
// and a boolean to check if the value has been set.
func (o *PickActionRequest) GetUserPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserPrompt, true
}

// SetUserPrompt sets field value.
func (o *PickActionRequest) SetUserPrompt(v string) {
	o.UserPrompt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PickActionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Client != nil {
		toSerialize["client"] = o.Client
	}
	toSerialize["number_of_relevant_actions"] = o.NumberOfRelevantActions
	if o.Stability != nil {
		toSerialize["stability"] = o.Stability
	}
	toSerialize["user_prompt"] = o.UserPrompt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PickActionRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Client                  *ClientType     `json:"client,omitempty"`
		NumberOfRelevantActions *int64          `json:"number_of_relevant_actions"`
		Stability               *StabilityLevel `json:"stability,omitempty"`
		UserPrompt              *string         `json:"user_prompt"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.NumberOfRelevantActions == nil {
		return fmt.Errorf("required field number_of_relevant_actions missing")
	}
	if all.UserPrompt == nil {
		return fmt.Errorf("required field user_prompt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"client", "number_of_relevant_actions", "stability", "user_prompt"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Client != nil && !all.Client.IsValid() {
		hasInvalidField = true
	} else {
		o.Client = all.Client
	}
	o.NumberOfRelevantActions = *all.NumberOfRelevantActions
	if all.Stability != nil && !all.Stability.IsValid() {
		hasInvalidField = true
	} else {
		o.Stability = all.Stability
	}
	o.UserPrompt = *all.UserPrompt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
