// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AgenticEvent
type AgenticEvent struct {
	// The message content.
	Message string `json:"message"`
	// Additional payload data for the event.
	Payload interface{} `json:"payload,omitempty"`
	// The type of agentic event.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAgenticEvent instantiates a new AgenticEvent object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAgenticEvent(message string, typeVar string) *AgenticEvent {
	this := AgenticEvent{}
	this.Message = message
	this.Type = typeVar
	return &this
}

// NewAgenticEventWithDefaults instantiates a new AgenticEvent object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAgenticEventWithDefaults() *AgenticEvent {
	this := AgenticEvent{}
	return &this
}

// GetMessage returns the Message field value.
func (o *AgenticEvent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AgenticEvent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *AgenticEvent) SetMessage(v string) {
	o.Message = v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *AgenticEvent) GetPayload() interface{} {
	if o == nil || o.Payload == nil {
		var ret interface{}
		return ret
	}
	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgenticEvent) GetPayloadOk() (*interface{}, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return &o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *AgenticEvent) HasPayload() bool {
	return o != nil && o.Payload != nil
}

// SetPayload gets a reference to the given interface{} and assigns it to the Payload field.
func (o *AgenticEvent) SetPayload(v interface{}) {
	o.Payload = v
}

// GetType returns the Type field value.
func (o *AgenticEvent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AgenticEvent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AgenticEvent) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AgenticEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["message"] = o.Message
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AgenticEvent) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Message *string     `json:"message"`
		Payload interface{} `json:"payload,omitempty"`
		Type    *string     `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"message", "payload", "type"})
	} else {
		return err
	}
	o.Message = *all.Message
	o.Payload = all.Payload
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
