// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TransportWebhookLogNetwork The network information for the event.
type TransportWebhookLogNetwork struct {
	// The IP address information.
	Ip *TransportWebhookLogNetworkIp `json:"ip,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTransportWebhookLogNetwork instantiates a new TransportWebhookLogNetwork object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTransportWebhookLogNetwork() *TransportWebhookLogNetwork {
	this := TransportWebhookLogNetwork{}
	return &this
}

// NewTransportWebhookLogNetworkWithDefaults instantiates a new TransportWebhookLogNetwork object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTransportWebhookLogNetworkWithDefaults() *TransportWebhookLogNetwork {
	this := TransportWebhookLogNetwork{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *TransportWebhookLogNetwork) GetIp() TransportWebhookLogNetworkIp {
	if o == nil || o.Ip == nil {
		var ret TransportWebhookLogNetworkIp
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogNetwork) GetIpOk() (*TransportWebhookLogNetworkIp, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *TransportWebhookLogNetwork) HasIp() bool {
	return o != nil && o.Ip != nil
}

// SetIp gets a reference to the given TransportWebhookLogNetworkIp and assigns it to the Ip field.
func (o *TransportWebhookLogNetwork) SetIp(v TransportWebhookLogNetworkIp) {
	o.Ip = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TransportWebhookLogNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TransportWebhookLogNetwork) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Ip *TransportWebhookLogNetworkIp `json:"ip,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ip"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Ip != nil && all.Ip.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Ip = all.Ip

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
