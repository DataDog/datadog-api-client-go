// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventsContext Test-only generated schema for the dashboard versioning proof of concept.
type EventsContext struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ContextEvent *string `json:"context_event,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ToEvent *string `json:"to_event,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEventsContext instantiates a new EventsContext object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventsContext() *EventsContext {
	this := EventsContext{}
	return &this
}

// NewEventsContextWithDefaults instantiates a new EventsContext object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventsContextWithDefaults() *EventsContext {
	this := EventsContext{}
	return &this
}

// GetContextEvent returns the ContextEvent field value if set, zero value otherwise.
func (o *EventsContext) GetContextEvent() string {
	if o == nil || o.ContextEvent == nil {
		var ret string
		return ret
	}
	return *o.ContextEvent
}

// GetContextEventOk returns a tuple with the ContextEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsContext) GetContextEventOk() (*string, bool) {
	if o == nil || o.ContextEvent == nil {
		return nil, false
	}
	return o.ContextEvent, true
}

// HasContextEvent returns a boolean if a field has been set.
func (o *EventsContext) HasContextEvent() bool {
	return o != nil && o.ContextEvent != nil
}

// SetContextEvent gets a reference to the given string and assigns it to the ContextEvent field.
func (o *EventsContext) SetContextEvent(v string) {
	o.ContextEvent = &v
}

// GetToEvent returns the ToEvent field value if set, zero value otherwise.
func (o *EventsContext) GetToEvent() string {
	if o == nil || o.ToEvent == nil {
		var ret string
		return ret
	}
	return *o.ToEvent
}

// GetToEventOk returns a tuple with the ToEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsContext) GetToEventOk() (*string, bool) {
	if o == nil || o.ToEvent == nil {
		return nil, false
	}
	return o.ToEvent, true
}

// HasToEvent returns a boolean if a field has been set.
func (o *EventsContext) HasToEvent() bool {
	return o != nil && o.ToEvent != nil
}

// SetToEvent gets a reference to the given string and assigns it to the ToEvent field.
func (o *EventsContext) SetToEvent(v string) {
	o.ToEvent = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventsContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ContextEvent != nil {
		toSerialize["context_event"] = o.ContextEvent
	}
	if o.ToEvent != nil {
		toSerialize["to_event"] = o.ToEvent
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventsContext) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ContextEvent *string `json:"context_event,omitempty"`
		ToEvent      *string `json:"to_event,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.ContextEvent = all.ContextEvent
	o.ToEvent = all.ToEvent

	return nil
}
