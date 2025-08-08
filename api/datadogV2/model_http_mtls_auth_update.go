// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HTTPMtlsAuthUpdate The definition of the `HTTPMtlsAuth` object.
type HTTPMtlsAuthUpdate struct {
	// Certificate of authority used to sign the request.
	Certificate *string `json:"certificate,omitempty"`
	// Private key used for the MTLS handshake
	PrivateKey *string `json:"private_key,omitempty"`
	// The definition of the `HTTPMtlsAuth` object.
	Type HTTPMtlsAuthType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHTTPMtlsAuthUpdate instantiates a new HTTPMtlsAuthUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHTTPMtlsAuthUpdate(typeVar HTTPMtlsAuthType) *HTTPMtlsAuthUpdate {
	this := HTTPMtlsAuthUpdate{}
	this.Type = typeVar
	return &this
}

// NewHTTPMtlsAuthUpdateWithDefaults instantiates a new HTTPMtlsAuthUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHTTPMtlsAuthUpdateWithDefaults() *HTTPMtlsAuthUpdate {
	this := HTTPMtlsAuthUpdate{}
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *HTTPMtlsAuthUpdate) GetCertificate() string {
	if o == nil || o.Certificate == nil {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPMtlsAuthUpdate) GetCertificateOk() (*string, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *HTTPMtlsAuthUpdate) HasCertificate() bool {
	return o != nil && o.Certificate != nil
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *HTTPMtlsAuthUpdate) SetCertificate(v string) {
	o.Certificate = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *HTTPMtlsAuthUpdate) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPMtlsAuthUpdate) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *HTTPMtlsAuthUpdate) HasPrivateKey() bool {
	return o != nil && o.PrivateKey != nil
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *HTTPMtlsAuthUpdate) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetType returns the Type field value.
func (o *HTTPMtlsAuthUpdate) GetType() HTTPMtlsAuthType {
	if o == nil {
		var ret HTTPMtlsAuthType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HTTPMtlsAuthUpdate) GetTypeOk() (*HTTPMtlsAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *HTTPMtlsAuthUpdate) SetType(v HTTPMtlsAuthType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HTTPMtlsAuthUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Certificate != nil {
		toSerialize["certificate"] = o.Certificate
	}
	if o.PrivateKey != nil {
		toSerialize["private_key"] = o.PrivateKey
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HTTPMtlsAuthUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Certificate *string           `json:"certificate,omitempty"`
		PrivateKey  *string           `json:"private_key,omitempty"`
		Type        *HTTPMtlsAuthType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"certificate", "private_key", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Certificate = all.Certificate
	o.PrivateKey = all.PrivateKey
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
