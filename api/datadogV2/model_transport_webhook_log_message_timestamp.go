// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TransportWebhookLogMessageTimestamp The message delivery timing information.
type TransportWebhookLogMessageTimestamp struct {
	// The Unix timestamp of the event.
	EventTimestamp *float64 `json:"event_timestamp,omitempty"`
	// The total delivery time in seconds.
	Lifetime *float64 `json:"lifetime,omitempty"`
	// Number of seconds the message spent in the delivery queue.
	QueueTime *float64 `json:"queue_time,omitempty"`
	// The scheduled delivery time as a Unix timestamp.
	ScheduledTime *float64 `json:"scheduled_time,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTransportWebhookLogMessageTimestamp instantiates a new TransportWebhookLogMessageTimestamp object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTransportWebhookLogMessageTimestamp() *TransportWebhookLogMessageTimestamp {
	this := TransportWebhookLogMessageTimestamp{}
	return &this
}

// NewTransportWebhookLogMessageTimestampWithDefaults instantiates a new TransportWebhookLogMessageTimestamp object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTransportWebhookLogMessageTimestampWithDefaults() *TransportWebhookLogMessageTimestamp {
	this := TransportWebhookLogMessageTimestamp{}
	return &this
}

// GetEventTimestamp returns the EventTimestamp field value if set, zero value otherwise.
func (o *TransportWebhookLogMessageTimestamp) GetEventTimestamp() float64 {
	if o == nil || o.EventTimestamp == nil {
		var ret float64
		return ret
	}
	return *o.EventTimestamp
}

// GetEventTimestampOk returns a tuple with the EventTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogMessageTimestamp) GetEventTimestampOk() (*float64, bool) {
	if o == nil || o.EventTimestamp == nil {
		return nil, false
	}
	return o.EventTimestamp, true
}

// HasEventTimestamp returns a boolean if a field has been set.
func (o *TransportWebhookLogMessageTimestamp) HasEventTimestamp() bool {
	return o != nil && o.EventTimestamp != nil
}

// SetEventTimestamp gets a reference to the given float64 and assigns it to the EventTimestamp field.
func (o *TransportWebhookLogMessageTimestamp) SetEventTimestamp(v float64) {
	o.EventTimestamp = &v
}

// GetLifetime returns the Lifetime field value if set, zero value otherwise.
func (o *TransportWebhookLogMessageTimestamp) GetLifetime() float64 {
	if o == nil || o.Lifetime == nil {
		var ret float64
		return ret
	}
	return *o.Lifetime
}

// GetLifetimeOk returns a tuple with the Lifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogMessageTimestamp) GetLifetimeOk() (*float64, bool) {
	if o == nil || o.Lifetime == nil {
		return nil, false
	}
	return o.Lifetime, true
}

// HasLifetime returns a boolean if a field has been set.
func (o *TransportWebhookLogMessageTimestamp) HasLifetime() bool {
	return o != nil && o.Lifetime != nil
}

// SetLifetime gets a reference to the given float64 and assigns it to the Lifetime field.
func (o *TransportWebhookLogMessageTimestamp) SetLifetime(v float64) {
	o.Lifetime = &v
}

// GetQueueTime returns the QueueTime field value if set, zero value otherwise.
func (o *TransportWebhookLogMessageTimestamp) GetQueueTime() float64 {
	if o == nil || o.QueueTime == nil {
		var ret float64
		return ret
	}
	return *o.QueueTime
}

// GetQueueTimeOk returns a tuple with the QueueTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogMessageTimestamp) GetQueueTimeOk() (*float64, bool) {
	if o == nil || o.QueueTime == nil {
		return nil, false
	}
	return o.QueueTime, true
}

// HasQueueTime returns a boolean if a field has been set.
func (o *TransportWebhookLogMessageTimestamp) HasQueueTime() bool {
	return o != nil && o.QueueTime != nil
}

// SetQueueTime gets a reference to the given float64 and assigns it to the QueueTime field.
func (o *TransportWebhookLogMessageTimestamp) SetQueueTime(v float64) {
	o.QueueTime = &v
}

// GetScheduledTime returns the ScheduledTime field value if set, zero value otherwise.
func (o *TransportWebhookLogMessageTimestamp) GetScheduledTime() float64 {
	if o == nil || o.ScheduledTime == nil {
		var ret float64
		return ret
	}
	return *o.ScheduledTime
}

// GetScheduledTimeOk returns a tuple with the ScheduledTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogMessageTimestamp) GetScheduledTimeOk() (*float64, bool) {
	if o == nil || o.ScheduledTime == nil {
		return nil, false
	}
	return o.ScheduledTime, true
}

// HasScheduledTime returns a boolean if a field has been set.
func (o *TransportWebhookLogMessageTimestamp) HasScheduledTime() bool {
	return o != nil && o.ScheduledTime != nil
}

// SetScheduledTime gets a reference to the given float64 and assigns it to the ScheduledTime field.
func (o *TransportWebhookLogMessageTimestamp) SetScheduledTime(v float64) {
	o.ScheduledTime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TransportWebhookLogMessageTimestamp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EventTimestamp != nil {
		toSerialize["event_timestamp"] = o.EventTimestamp
	}
	if o.Lifetime != nil {
		toSerialize["lifetime"] = o.Lifetime
	}
	if o.QueueTime != nil {
		toSerialize["queue_time"] = o.QueueTime
	}
	if o.ScheduledTime != nil {
		toSerialize["scheduled_time"] = o.ScheduledTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TransportWebhookLogMessageTimestamp) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EventTimestamp *float64 `json:"event_timestamp,omitempty"`
		Lifetime       *float64 `json:"lifetime,omitempty"`
		QueueTime      *float64 `json:"queue_time,omitempty"`
		ScheduledTime  *float64 `json:"scheduled_time,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"event_timestamp", "lifetime", "queue_time", "scheduled_time"})
	} else {
		return err
	}
	o.EventTimestamp = all.EventTimestamp
	o.Lifetime = all.Lifetime
	o.QueueTime = all.QueueTime
	o.ScheduledTime = all.ScheduledTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
