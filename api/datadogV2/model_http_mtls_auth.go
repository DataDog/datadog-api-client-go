// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HTTPMtlsAuth The definition of the `HTTPMtlsAuth` object.
type HTTPMtlsAuth struct {
	// Certificate of authority used to sign the request.
	Certificate string `json:"certificate"`
	// Private key used for the MTLS handshake
	PrivateKey string `json:"private_key"`
	// The definition of the `HTTPMtlsAuth` object.
	Type HTTPMtlsAuthType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHTTPMtlsAuth instantiates a new HTTPMtlsAuth object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHTTPMtlsAuth(certificate string, privateKey string, typeVar HTTPMtlsAuthType) *HTTPMtlsAuth {
	this := HTTPMtlsAuth{}
	this.Certificate = certificate
	this.PrivateKey = privateKey
	this.Type = typeVar
	return &this
}

// NewHTTPMtlsAuthWithDefaults instantiates a new HTTPMtlsAuth object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHTTPMtlsAuthWithDefaults() *HTTPMtlsAuth {
	this := HTTPMtlsAuth{}
	return &this
}

// GetCertificate returns the Certificate field value.
func (o *HTTPMtlsAuth) GetCertificate() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value
// and a boolean to check if the value has been set.
func (o *HTTPMtlsAuth) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Certificate, true
}

// SetCertificate sets field value.
func (o *HTTPMtlsAuth) SetCertificate(v string) {
	o.Certificate = v
}

// GetPrivateKey returns the PrivateKey field value.
func (o *HTTPMtlsAuth) GetPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *HTTPMtlsAuth) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value.
func (o *HTTPMtlsAuth) SetPrivateKey(v string) {
	o.PrivateKey = v
}

// GetType returns the Type field value.
func (o *HTTPMtlsAuth) GetType() HTTPMtlsAuthType {
	if o == nil {
		var ret HTTPMtlsAuthType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HTTPMtlsAuth) GetTypeOk() (*HTTPMtlsAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *HTTPMtlsAuth) SetType(v HTTPMtlsAuthType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HTTPMtlsAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["certificate"] = o.Certificate
	toSerialize["private_key"] = o.PrivateKey
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HTTPMtlsAuth) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Certificate *string           `json:"certificate"`
		PrivateKey  *string           `json:"private_key"`
		Type        *HTTPMtlsAuthType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Certificate == nil {
		return fmt.Errorf("required field certificate missing")
	}
	if all.PrivateKey == nil {
		return fmt.Errorf("required field private_key missing")
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
	o.Certificate = *all.Certificate
	o.PrivateKey = *all.PrivateKey
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
