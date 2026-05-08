// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TransportWebhookLogNetworkIp The IP address information.
type TransportWebhookLogNetworkIp struct {
	// Per-IP attribute records, each pairing an IP address with the providers that observed it.
	Attributes []TransportWebhookLogIpAttribute `json:"attributes,omitempty"`
	// The list of IP addresses.
	List []string `json:"list,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTransportWebhookLogNetworkIp instantiates a new TransportWebhookLogNetworkIp object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTransportWebhookLogNetworkIp() *TransportWebhookLogNetworkIp {
	this := TransportWebhookLogNetworkIp{}
	return &this
}

// NewTransportWebhookLogNetworkIpWithDefaults instantiates a new TransportWebhookLogNetworkIp object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTransportWebhookLogNetworkIpWithDefaults() *TransportWebhookLogNetworkIp {
	this := TransportWebhookLogNetworkIp{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TransportWebhookLogNetworkIp) GetAttributes() []TransportWebhookLogIpAttribute {
	if o == nil || o.Attributes == nil {
		var ret []TransportWebhookLogIpAttribute
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogNetworkIp) GetAttributesOk() (*[]TransportWebhookLogIpAttribute, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TransportWebhookLogNetworkIp) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given []TransportWebhookLogIpAttribute and assigns it to the Attributes field.
func (o *TransportWebhookLogNetworkIp) SetAttributes(v []TransportWebhookLogIpAttribute) {
	o.Attributes = v
}

// GetList returns the List field value if set, zero value otherwise.
func (o *TransportWebhookLogNetworkIp) GetList() []string {
	if o == nil || o.List == nil {
		var ret []string
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogNetworkIp) GetListOk() (*[]string, bool) {
	if o == nil || o.List == nil {
		return nil, false
	}
	return &o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *TransportWebhookLogNetworkIp) HasList() bool {
	return o != nil && o.List != nil
}

// SetList gets a reference to the given []string and assigns it to the List field.
func (o *TransportWebhookLogNetworkIp) SetList(v []string) {
	o.List = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TransportWebhookLogNetworkIp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.List != nil {
		toSerialize["list"] = o.List
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TransportWebhookLogNetworkIp) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes []TransportWebhookLogIpAttribute `json:"attributes,omitempty"`
		List       []string                         `json:"list,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "list"})
	} else {
		return err
	}
	o.Attributes = all.Attributes
	o.List = all.List

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
