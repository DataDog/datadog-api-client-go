// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationHistoryEvent A single event in the investigation history.
type RemediationHistoryEvent struct {
	// The type of event.
	EventType *string `json:"event_type,omitempty"`
	// History event ID.
	Id *string `json:"id,omitempty"`
	// Opaque JSON event body, base64-encoded. Decode and parse according to event_type.
	Payload *string `json:"payload,omitempty"`
	// Event time in epoch milliseconds (64-bit integer encoded as a string).
	TimestampMs *string `json:"timestamp_ms,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRemediationHistoryEvent instantiates a new RemediationHistoryEvent object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRemediationHistoryEvent() *RemediationHistoryEvent {
	this := RemediationHistoryEvent{}
	return &this
}

// NewRemediationHistoryEventWithDefaults instantiates a new RemediationHistoryEvent object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRemediationHistoryEventWithDefaults() *RemediationHistoryEvent {
	this := RemediationHistoryEvent{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *RemediationHistoryEvent) GetEventType() string {
	if o == nil || o.EventType == nil {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationHistoryEvent) GetEventTypeOk() (*string, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *RemediationHistoryEvent) HasEventType() bool {
	return o != nil && o.EventType != nil
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *RemediationHistoryEvent) SetEventType(v string) {
	o.EventType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RemediationHistoryEvent) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationHistoryEvent) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RemediationHistoryEvent) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RemediationHistoryEvent) SetId(v string) {
	o.Id = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *RemediationHistoryEvent) GetPayload() string {
	if o == nil || o.Payload == nil {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationHistoryEvent) GetPayloadOk() (*string, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *RemediationHistoryEvent) HasPayload() bool {
	return o != nil && o.Payload != nil
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *RemediationHistoryEvent) SetPayload(v string) {
	o.Payload = &v
}

// GetTimestampMs returns the TimestampMs field value if set, zero value otherwise.
func (o *RemediationHistoryEvent) GetTimestampMs() string {
	if o == nil || o.TimestampMs == nil {
		var ret string
		return ret
	}
	return *o.TimestampMs
}

// GetTimestampMsOk returns a tuple with the TimestampMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationHistoryEvent) GetTimestampMsOk() (*string, bool) {
	if o == nil || o.TimestampMs == nil {
		return nil, false
	}
	return o.TimestampMs, true
}

// HasTimestampMs returns a boolean if a field has been set.
func (o *RemediationHistoryEvent) HasTimestampMs() bool {
	return o != nil && o.TimestampMs != nil
}

// SetTimestampMs gets a reference to the given string and assigns it to the TimestampMs field.
func (o *RemediationHistoryEvent) SetTimestampMs(v string) {
	o.TimestampMs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RemediationHistoryEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EventType != nil {
		toSerialize["event_type"] = o.EventType
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}
	if o.TimestampMs != nil {
		toSerialize["timestamp_ms"] = o.TimestampMs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RemediationHistoryEvent) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EventType   *string `json:"event_type,omitempty"`
		Id          *string `json:"id,omitempty"`
		Payload     *string `json:"payload,omitempty"`
		TimestampMs *string `json:"timestamp_ms,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"event_type", "id", "payload", "timestamp_ms"})
	} else {
		return err
	}
	o.EventType = all.EventType
	o.Id = all.Id
	o.Payload = all.Payload
	o.TimestampMs = all.TimestampMs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
