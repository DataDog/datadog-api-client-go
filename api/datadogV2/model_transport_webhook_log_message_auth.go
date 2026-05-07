// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TransportWebhookLogMessageAuth The message authentication details.
type TransportWebhookLogMessageAuth struct {
	// The TLS version or negotiation information.
	DeliveredWithTls *string `json:"delivered_with_tls,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTransportWebhookLogMessageAuth instantiates a new TransportWebhookLogMessageAuth object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTransportWebhookLogMessageAuth() *TransportWebhookLogMessageAuth {
	this := TransportWebhookLogMessageAuth{}
	return &this
}

// NewTransportWebhookLogMessageAuthWithDefaults instantiates a new TransportWebhookLogMessageAuth object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTransportWebhookLogMessageAuthWithDefaults() *TransportWebhookLogMessageAuth {
	this := TransportWebhookLogMessageAuth{}
	return &this
}

// GetDeliveredWithTls returns the DeliveredWithTls field value if set, zero value otherwise.
func (o *TransportWebhookLogMessageAuth) GetDeliveredWithTls() string {
	if o == nil || o.DeliveredWithTls == nil {
		var ret string
		return ret
	}
	return *o.DeliveredWithTls
}

// GetDeliveredWithTlsOk returns a tuple with the DeliveredWithTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogMessageAuth) GetDeliveredWithTlsOk() (*string, bool) {
	if o == nil || o.DeliveredWithTls == nil {
		return nil, false
	}
	return o.DeliveredWithTls, true
}

// HasDeliveredWithTls returns a boolean if a field has been set.
func (o *TransportWebhookLogMessageAuth) HasDeliveredWithTls() bool {
	return o != nil && o.DeliveredWithTls != nil
}

// SetDeliveredWithTls gets a reference to the given string and assigns it to the DeliveredWithTls field.
func (o *TransportWebhookLogMessageAuth) SetDeliveredWithTls(v string) {
	o.DeliveredWithTls = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TransportWebhookLogMessageAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DeliveredWithTls != nil {
		toSerialize["delivered_with_tls"] = o.DeliveredWithTls
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TransportWebhookLogMessageAuth) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DeliveredWithTls *string `json:"delivered_with_tls,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"delivered_with_tls"})
	} else {
		return err
	}
	o.DeliveredWithTls = all.DeliveredWithTls

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
