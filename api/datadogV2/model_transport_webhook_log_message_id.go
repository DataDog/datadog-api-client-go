// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TransportWebhookLogMessageId The message identifiers.
type TransportWebhookLogMessageId struct {
	// The RFC 5322 Message-ID.
	MessageId *string `json:"message_id,omitempty"`
	// The SMTP transaction identifier.
	SmtpId *string `json:"smtp_id,omitempty"`
	// The transport provider event identifier.
	TransportEventId *string `json:"transport_event_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTransportWebhookLogMessageId instantiates a new TransportWebhookLogMessageId object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTransportWebhookLogMessageId() *TransportWebhookLogMessageId {
	this := TransportWebhookLogMessageId{}
	return &this
}

// NewTransportWebhookLogMessageIdWithDefaults instantiates a new TransportWebhookLogMessageId object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTransportWebhookLogMessageIdWithDefaults() *TransportWebhookLogMessageId {
	this := TransportWebhookLogMessageId{}
	return &this
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *TransportWebhookLogMessageId) GetMessageId() string {
	if o == nil || o.MessageId == nil {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogMessageId) GetMessageIdOk() (*string, bool) {
	if o == nil || o.MessageId == nil {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *TransportWebhookLogMessageId) HasMessageId() bool {
	return o != nil && o.MessageId != nil
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *TransportWebhookLogMessageId) SetMessageId(v string) {
	o.MessageId = &v
}

// GetSmtpId returns the SmtpId field value if set, zero value otherwise.
func (o *TransportWebhookLogMessageId) GetSmtpId() string {
	if o == nil || o.SmtpId == nil {
		var ret string
		return ret
	}
	return *o.SmtpId
}

// GetSmtpIdOk returns a tuple with the SmtpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogMessageId) GetSmtpIdOk() (*string, bool) {
	if o == nil || o.SmtpId == nil {
		return nil, false
	}
	return o.SmtpId, true
}

// HasSmtpId returns a boolean if a field has been set.
func (o *TransportWebhookLogMessageId) HasSmtpId() bool {
	return o != nil && o.SmtpId != nil
}

// SetSmtpId gets a reference to the given string and assigns it to the SmtpId field.
func (o *TransportWebhookLogMessageId) SetSmtpId(v string) {
	o.SmtpId = &v
}

// GetTransportEventId returns the TransportEventId field value if set, zero value otherwise.
func (o *TransportWebhookLogMessageId) GetTransportEventId() string {
	if o == nil || o.TransportEventId == nil {
		var ret string
		return ret
	}
	return *o.TransportEventId
}

// GetTransportEventIdOk returns a tuple with the TransportEventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogMessageId) GetTransportEventIdOk() (*string, bool) {
	if o == nil || o.TransportEventId == nil {
		return nil, false
	}
	return o.TransportEventId, true
}

// HasTransportEventId returns a boolean if a field has been set.
func (o *TransportWebhookLogMessageId) HasTransportEventId() bool {
	return o != nil && o.TransportEventId != nil
}

// SetTransportEventId gets a reference to the given string and assigns it to the TransportEventId field.
func (o *TransportWebhookLogMessageId) SetTransportEventId(v string) {
	o.TransportEventId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TransportWebhookLogMessageId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.MessageId != nil {
		toSerialize["message_id"] = o.MessageId
	}
	if o.SmtpId != nil {
		toSerialize["smtp_id"] = o.SmtpId
	}
	if o.TransportEventId != nil {
		toSerialize["transport_event_id"] = o.TransportEventId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TransportWebhookLogMessageId) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MessageId        *string `json:"message_id,omitempty"`
		SmtpId           *string `json:"smtp_id,omitempty"`
		TransportEventId *string `json:"transport_event_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"message_id", "smtp_id", "transport_event_id"})
	} else {
		return err
	}
	o.MessageId = all.MessageId
	o.SmtpId = all.SmtpId
	o.TransportEventId = all.TransportEventId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
