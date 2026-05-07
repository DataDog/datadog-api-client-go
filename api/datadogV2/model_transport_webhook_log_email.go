// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TransportWebhookLogEmail The email address details.
type TransportWebhookLogEmail struct {
	// The recipient email address.
	Address *string `json:"address,omitempty"`
	// The recipient domain.
	Domain *string `json:"domain,omitempty"`
	// The email subject line.
	Subject *string `json:"subject,omitempty"`
	// Email categorization tags applied by the transport provider (for example, "transactional", "marketing").
	Type []string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTransportWebhookLogEmail instantiates a new TransportWebhookLogEmail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTransportWebhookLogEmail() *TransportWebhookLogEmail {
	this := TransportWebhookLogEmail{}
	return &this
}

// NewTransportWebhookLogEmailWithDefaults instantiates a new TransportWebhookLogEmail object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTransportWebhookLogEmailWithDefaults() *TransportWebhookLogEmail {
	this := TransportWebhookLogEmail{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *TransportWebhookLogEmail) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogEmail) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *TransportWebhookLogEmail) HasAddress() bool {
	return o != nil && o.Address != nil
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *TransportWebhookLogEmail) SetAddress(v string) {
	o.Address = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *TransportWebhookLogEmail) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogEmail) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *TransportWebhookLogEmail) HasDomain() bool {
	return o != nil && o.Domain != nil
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *TransportWebhookLogEmail) SetDomain(v string) {
	o.Domain = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *TransportWebhookLogEmail) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogEmail) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *TransportWebhookLogEmail) HasSubject() bool {
	return o != nil && o.Subject != nil
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *TransportWebhookLogEmail) SetSubject(v string) {
	o.Subject = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TransportWebhookLogEmail) GetType() []string {
	if o == nil || o.Type == nil {
		var ret []string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogEmail) GetTypeOk() (*[]string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TransportWebhookLogEmail) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given []string and assigns it to the Type field.
func (o *TransportWebhookLogEmail) SetType(v []string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TransportWebhookLogEmail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TransportWebhookLogEmail) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Address *string  `json:"address,omitempty"`
		Domain  *string  `json:"domain,omitempty"`
		Subject *string  `json:"subject,omitempty"`
		Type    []string `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"address", "domain", "subject", "type"})
	} else {
		return err
	}
	o.Address = all.Address
	o.Domain = all.Domain
	o.Subject = all.Subject
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
