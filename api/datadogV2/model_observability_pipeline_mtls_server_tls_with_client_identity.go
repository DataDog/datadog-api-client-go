// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineMtlsServerTlsWithClientIdentity Configuration for enabling TLS encryption between the pipeline component and external connecting clients, or mutual TLS when `verify_certificate` is `true`. Mutual TLS also enables an optional allowlist of client certificate identities.
type ObservabilityPipelineMtlsServerTlsWithClientIdentity struct {
	// Path to the Certificate Authority (CA) file used to validate connecting clients' TLS certificates.
	CaFile *string `json:"ca_file,omitempty"`
	// Restricts which clients may connect, based on the identities in their TLS certificate. Checks run during the TLS handshake, after the client certificate chain is validated against `ca_file`, and never relax that validation.
	// Requires `verify_certificate: true` and `ca_file`. Set exactly one of `allowlist_file` or `allowlist`.
	// When set, connections are closed after a maximum age of 300 seconds so that clients removed from the allowlist are disconnected.
	// Requires Observability Pipelines Worker 2.27 or later.
	ClientIdentity *ObservabilityPipelineClientIdentity `json:"client_identity,omitempty"`
	// Path to the TLS server certificate file used to identify the pipeline component to connecting clients.
	CrtFile string `json:"crt_file"`
	// Path to the private key file associated with the TLS server certificate.
	KeyFile *string `json:"key_file,omitempty"`
	// Name of the environment variable or secret that holds the passphrase for the private key file.
	KeyPassKey *string `json:"key_pass_key,omitempty"`
	// When `true`, requires client connections to present a valid certificate, enabling mutual TLS authentication.
	VerifyCertificate *bool `json:"verify_certificate,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineMtlsServerTlsWithClientIdentity instantiates a new ObservabilityPipelineMtlsServerTlsWithClientIdentity object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineMtlsServerTlsWithClientIdentity(crtFile string) *ObservabilityPipelineMtlsServerTlsWithClientIdentity {
	this := ObservabilityPipelineMtlsServerTlsWithClientIdentity{}
	this.CrtFile = crtFile
	return &this
}

// NewObservabilityPipelineMtlsServerTlsWithClientIdentityWithDefaults instantiates a new ObservabilityPipelineMtlsServerTlsWithClientIdentity object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineMtlsServerTlsWithClientIdentityWithDefaults() *ObservabilityPipelineMtlsServerTlsWithClientIdentity {
	this := ObservabilityPipelineMtlsServerTlsWithClientIdentity{}
	return &this
}

// GetCaFile returns the CaFile field value if set, zero value otherwise.
func (o *ObservabilityPipelineMtlsServerTlsWithClientIdentity) GetCaFile() string {
	if o == nil || o.CaFile == nil {
		var ret string
		return ret
	}
	return *o.CaFile
}

// GetCaFileOk returns a tuple with the CaFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMtlsServerTlsWithClientIdentity) GetCaFileOk() (*string, bool) {
	if o == nil || o.CaFile == nil {
		return nil, false
	}
	return o.CaFile, true
}

// HasCaFile returns a boolean if a field has been set.
func (o *ObservabilityPipelineMtlsServerTlsWithClientIdentity) HasCaFile() bool {
	return o != nil && o.CaFile != nil
}

// SetCaFile gets a reference to the given string and assigns it to the CaFile field.
func (o *ObservabilityPipelineMtlsServerTlsWithClientIdentity) SetCaFile(v string) {
	o.CaFile = &v
}

// GetClientIdentity returns the ClientIdentity field value if set, zero value otherwise.
func (o *ObservabilityPipelineMtlsServerTlsWithClientIdentity) GetClientIdentity() ObservabilityPipelineClientIdentity {
	if o == nil || o.ClientIdentity == nil {
		var ret ObservabilityPipelineClientIdentity
		return ret
	}
	return *o.ClientIdentity
}

// GetClientIdentityOk returns a tuple with the ClientIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMtlsServerTlsWithClientIdentity) GetClientIdentityOk() (*ObservabilityPipelineClientIdentity, bool) {
	if o == nil || o.ClientIdentity == nil {
		return nil, false
	}
	return o.ClientIdentity, true
}

// HasClientIdentity returns a boolean if a field has been set.
func (o *ObservabilityPipelineMtlsServerTlsWithClientIdentity) HasClientIdentity() bool {
	return o != nil && o.ClientIdentity != nil
}

// SetClientIdentity gets a reference to the given ObservabilityPipelineClientIdentity and assigns it to the ClientIdentity field.
func (o *ObservabilityPipelineMtlsServerTlsWithClientIdentity) SetClientIdentity(v ObservabilityPipelineClientIdentity) {
	o.ClientIdentity = &v
}

// GetCrtFile returns the CrtFile field value.
func (o *ObservabilityPipelineMtlsServerTlsWithClientIdentity) GetCrtFile() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CrtFile
}

// GetCrtFileOk returns a tuple with the CrtFile field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMtlsServerTlsWithClientIdentity) GetCrtFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CrtFile, true
}

// SetCrtFile sets field value.
func (o *ObservabilityPipelineMtlsServerTlsWithClientIdentity) SetCrtFile(v string) {
	o.CrtFile = v
}

// GetKeyFile returns the KeyFile field value if set, zero value otherwise.
func (o *ObservabilityPipelineMtlsServerTlsWithClientIdentity) GetKeyFile() string {
	if o == nil || o.KeyFile == nil {
		var ret string
		return ret
	}
	return *o.KeyFile
}

// GetKeyFileOk returns a tuple with the KeyFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMtlsServerTlsWithClientIdentity) GetKeyFileOk() (*string, bool) {
	if o == nil || o.KeyFile == nil {
		return nil, false
	}
	return o.KeyFile, true
}

// HasKeyFile returns a boolean if a field has been set.
func (o *ObservabilityPipelineMtlsServerTlsWithClientIdentity) HasKeyFile() bool {
	return o != nil && o.KeyFile != nil
}

// SetKeyFile gets a reference to the given string and assigns it to the KeyFile field.
func (o *ObservabilityPipelineMtlsServerTlsWithClientIdentity) SetKeyFile(v string) {
	o.KeyFile = &v
}

// GetKeyPassKey returns the KeyPassKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineMtlsServerTlsWithClientIdentity) GetKeyPassKey() string {
	if o == nil || o.KeyPassKey == nil {
		var ret string
		return ret
	}
	return *o.KeyPassKey
}

// GetKeyPassKeyOk returns a tuple with the KeyPassKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMtlsServerTlsWithClientIdentity) GetKeyPassKeyOk() (*string, bool) {
	if o == nil || o.KeyPassKey == nil {
		return nil, false
	}
	return o.KeyPassKey, true
}

// HasKeyPassKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineMtlsServerTlsWithClientIdentity) HasKeyPassKey() bool {
	return o != nil && o.KeyPassKey != nil
}

// SetKeyPassKey gets a reference to the given string and assigns it to the KeyPassKey field.
func (o *ObservabilityPipelineMtlsServerTlsWithClientIdentity) SetKeyPassKey(v string) {
	o.KeyPassKey = &v
}

// GetVerifyCertificate returns the VerifyCertificate field value if set, zero value otherwise.
func (o *ObservabilityPipelineMtlsServerTlsWithClientIdentity) GetVerifyCertificate() bool {
	if o == nil || o.VerifyCertificate == nil {
		var ret bool
		return ret
	}
	return *o.VerifyCertificate
}

// GetVerifyCertificateOk returns a tuple with the VerifyCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMtlsServerTlsWithClientIdentity) GetVerifyCertificateOk() (*bool, bool) {
	if o == nil || o.VerifyCertificate == nil {
		return nil, false
	}
	return o.VerifyCertificate, true
}

// HasVerifyCertificate returns a boolean if a field has been set.
func (o *ObservabilityPipelineMtlsServerTlsWithClientIdentity) HasVerifyCertificate() bool {
	return o != nil && o.VerifyCertificate != nil
}

// SetVerifyCertificate gets a reference to the given bool and assigns it to the VerifyCertificate field.
func (o *ObservabilityPipelineMtlsServerTlsWithClientIdentity) SetVerifyCertificate(v bool) {
	o.VerifyCertificate = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineMtlsServerTlsWithClientIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CaFile != nil {
		toSerialize["ca_file"] = o.CaFile
	}
	if o.ClientIdentity != nil {
		toSerialize["client_identity"] = o.ClientIdentity
	}
	toSerialize["crt_file"] = o.CrtFile
	if o.KeyFile != nil {
		toSerialize["key_file"] = o.KeyFile
	}
	if o.KeyPassKey != nil {
		toSerialize["key_pass_key"] = o.KeyPassKey
	}
	if o.VerifyCertificate != nil {
		toSerialize["verify_certificate"] = o.VerifyCertificate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineMtlsServerTlsWithClientIdentity) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CaFile            *string                              `json:"ca_file,omitempty"`
		ClientIdentity    *ObservabilityPipelineClientIdentity `json:"client_identity,omitempty"`
		CrtFile           *string                              `json:"crt_file"`
		KeyFile           *string                              `json:"key_file,omitempty"`
		KeyPassKey        *string                              `json:"key_pass_key,omitempty"`
		VerifyCertificate *bool                                `json:"verify_certificate,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CrtFile == nil {
		return fmt.Errorf("required field crt_file missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ca_file", "client_identity", "crt_file", "key_file", "key_pass_key", "verify_certificate"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CaFile = all.CaFile
	if all.ClientIdentity != nil && all.ClientIdentity.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ClientIdentity = all.ClientIdentity
	o.CrtFile = *all.CrtFile
	o.KeyFile = all.KeyFile
	o.KeyPassKey = all.KeyPassKey
	o.VerifyCertificate = all.VerifyCertificate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
