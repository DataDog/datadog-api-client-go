// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TransportWebhookLogMessageResponse The SMTP response information.
type TransportWebhookLogMessageResponse struct {
	// The enhanced SMTP status code.
	EnhancedSmtpCode *string `json:"enhanced_smtp_code,omitempty"`
	// The SMTP response message.
	Reason *string `json:"reason,omitempty"`
	// The SMTP status code.
	SmtpCode *string `json:"smtp_code,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTransportWebhookLogMessageResponse instantiates a new TransportWebhookLogMessageResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTransportWebhookLogMessageResponse() *TransportWebhookLogMessageResponse {
	this := TransportWebhookLogMessageResponse{}
	return &this
}

// NewTransportWebhookLogMessageResponseWithDefaults instantiates a new TransportWebhookLogMessageResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTransportWebhookLogMessageResponseWithDefaults() *TransportWebhookLogMessageResponse {
	this := TransportWebhookLogMessageResponse{}
	return &this
}

// GetEnhancedSmtpCode returns the EnhancedSmtpCode field value if set, zero value otherwise.
func (o *TransportWebhookLogMessageResponse) GetEnhancedSmtpCode() string {
	if o == nil || o.EnhancedSmtpCode == nil {
		var ret string
		return ret
	}
	return *o.EnhancedSmtpCode
}

// GetEnhancedSmtpCodeOk returns a tuple with the EnhancedSmtpCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogMessageResponse) GetEnhancedSmtpCodeOk() (*string, bool) {
	if o == nil || o.EnhancedSmtpCode == nil {
		return nil, false
	}
	return o.EnhancedSmtpCode, true
}

// HasEnhancedSmtpCode returns a boolean if a field has been set.
func (o *TransportWebhookLogMessageResponse) HasEnhancedSmtpCode() bool {
	return o != nil && o.EnhancedSmtpCode != nil
}

// SetEnhancedSmtpCode gets a reference to the given string and assigns it to the EnhancedSmtpCode field.
func (o *TransportWebhookLogMessageResponse) SetEnhancedSmtpCode(v string) {
	o.EnhancedSmtpCode = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *TransportWebhookLogMessageResponse) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogMessageResponse) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *TransportWebhookLogMessageResponse) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *TransportWebhookLogMessageResponse) SetReason(v string) {
	o.Reason = &v
}

// GetSmtpCode returns the SmtpCode field value if set, zero value otherwise.
func (o *TransportWebhookLogMessageResponse) GetSmtpCode() string {
	if o == nil || o.SmtpCode == nil {
		var ret string
		return ret
	}
	return *o.SmtpCode
}

// GetSmtpCodeOk returns a tuple with the SmtpCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogMessageResponse) GetSmtpCodeOk() (*string, bool) {
	if o == nil || o.SmtpCode == nil {
		return nil, false
	}
	return o.SmtpCode, true
}

// HasSmtpCode returns a boolean if a field has been set.
func (o *TransportWebhookLogMessageResponse) HasSmtpCode() bool {
	return o != nil && o.SmtpCode != nil
}

// SetSmtpCode gets a reference to the given string and assigns it to the SmtpCode field.
func (o *TransportWebhookLogMessageResponse) SetSmtpCode(v string) {
	o.SmtpCode = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TransportWebhookLogMessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EnhancedSmtpCode != nil {
		toSerialize["enhanced_smtp_code"] = o.EnhancedSmtpCode
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.SmtpCode != nil {
		toSerialize["smtp_code"] = o.SmtpCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TransportWebhookLogMessageResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EnhancedSmtpCode *string `json:"enhanced_smtp_code,omitempty"`
		Reason           *string `json:"reason,omitempty"`
		SmtpCode         *string `json:"smtp_code,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enhanced_smtp_code", "reason", "smtp_code"})
	} else {
		return err
	}
	o.EnhancedSmtpCode = all.EnhancedSmtpCode
	o.Reason = all.Reason
	o.SmtpCode = all.SmtpCode

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
