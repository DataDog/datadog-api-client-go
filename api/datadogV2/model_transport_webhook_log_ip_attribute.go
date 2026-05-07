// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TransportWebhookLogIpAttribute An IP attribute with its sources.
type TransportWebhookLogIpAttribute struct {
	// The IP address.
	Ip *string `json:"ip,omitempty"`
	// The transport providers or systems that reported this IP address.
	Source []string `json:"source,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTransportWebhookLogIpAttribute instantiates a new TransportWebhookLogIpAttribute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTransportWebhookLogIpAttribute() *TransportWebhookLogIpAttribute {
	this := TransportWebhookLogIpAttribute{}
	return &this
}

// NewTransportWebhookLogIpAttributeWithDefaults instantiates a new TransportWebhookLogIpAttribute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTransportWebhookLogIpAttributeWithDefaults() *TransportWebhookLogIpAttribute {
	this := TransportWebhookLogIpAttribute{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *TransportWebhookLogIpAttribute) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogIpAttribute) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *TransportWebhookLogIpAttribute) HasIp() bool {
	return o != nil && o.Ip != nil
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *TransportWebhookLogIpAttribute) SetIp(v string) {
	o.Ip = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *TransportWebhookLogIpAttribute) GetSource() []string {
	if o == nil || o.Source == nil {
		var ret []string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogIpAttribute) GetSourceOk() (*[]string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return &o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *TransportWebhookLogIpAttribute) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given []string and assigns it to the Source field.
func (o *TransportWebhookLogIpAttribute) SetSource(v []string) {
	o.Source = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TransportWebhookLogIpAttribute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TransportWebhookLogIpAttribute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Ip     *string  `json:"ip,omitempty"`
		Source []string `json:"source,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ip", "source"})
	} else {
		return err
	}
	o.Ip = all.Ip
	o.Source = all.Source

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
