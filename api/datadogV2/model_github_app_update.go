// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GithubAppUpdate Update to GitHub App credentials.
type GithubAppUpdate struct {
	// The numeric ID of the GitHub App.
	GithubAppId *GithubAppUpdateGithubAppId `json:"github_app_id,omitempty"`
	// The installation ID of the GitHub App.
	GithubInstallationId *GithubAppUpdateGithubInstallationId `json:"github_installation_id,omitempty"`
	// Hostname of the GitHub Enterprise Server. Omit this field for github.com.
	Hostname *GithubAppUpdateHostname `json:"hostname,omitempty"`
	// PEM-encoded RSA private key for the GitHub App.
	PrivateKey *GithubAppUpdatePrivateKey `json:"private_key,omitempty"`
	// The credential type for `GithubApp`.
	Type GithubAppType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGithubAppUpdate instantiates a new GithubAppUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGithubAppUpdate(typeVar GithubAppType) *GithubAppUpdate {
	this := GithubAppUpdate{}
	this.Type = typeVar
	return &this
}

// NewGithubAppUpdateWithDefaults instantiates a new GithubAppUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGithubAppUpdateWithDefaults() *GithubAppUpdate {
	this := GithubAppUpdate{}
	return &this
}

// GetGithubAppId returns the GithubAppId field value if set, zero value otherwise.
func (o *GithubAppUpdate) GetGithubAppId() GithubAppUpdateGithubAppId {
	if o == nil || o.GithubAppId == nil {
		var ret GithubAppUpdateGithubAppId
		return ret
	}
	return *o.GithubAppId
}

// GetGithubAppIdOk returns a tuple with the GithubAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubAppUpdate) GetGithubAppIdOk() (*GithubAppUpdateGithubAppId, bool) {
	if o == nil || o.GithubAppId == nil {
		return nil, false
	}
	return o.GithubAppId, true
}

// HasGithubAppId returns a boolean if a field has been set.
func (o *GithubAppUpdate) HasGithubAppId() bool {
	return o != nil && o.GithubAppId != nil
}

// SetGithubAppId gets a reference to the given GithubAppUpdateGithubAppId and assigns it to the GithubAppId field.
func (o *GithubAppUpdate) SetGithubAppId(v GithubAppUpdateGithubAppId) {
	o.GithubAppId = &v
}

// GetGithubInstallationId returns the GithubInstallationId field value if set, zero value otherwise.
func (o *GithubAppUpdate) GetGithubInstallationId() GithubAppUpdateGithubInstallationId {
	if o == nil || o.GithubInstallationId == nil {
		var ret GithubAppUpdateGithubInstallationId
		return ret
	}
	return *o.GithubInstallationId
}

// GetGithubInstallationIdOk returns a tuple with the GithubInstallationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubAppUpdate) GetGithubInstallationIdOk() (*GithubAppUpdateGithubInstallationId, bool) {
	if o == nil || o.GithubInstallationId == nil {
		return nil, false
	}
	return o.GithubInstallationId, true
}

// HasGithubInstallationId returns a boolean if a field has been set.
func (o *GithubAppUpdate) HasGithubInstallationId() bool {
	return o != nil && o.GithubInstallationId != nil
}

// SetGithubInstallationId gets a reference to the given GithubAppUpdateGithubInstallationId and assigns it to the GithubInstallationId field.
func (o *GithubAppUpdate) SetGithubInstallationId(v GithubAppUpdateGithubInstallationId) {
	o.GithubInstallationId = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *GithubAppUpdate) GetHostname() GithubAppUpdateHostname {
	if o == nil || o.Hostname == nil {
		var ret GithubAppUpdateHostname
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubAppUpdate) GetHostnameOk() (*GithubAppUpdateHostname, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *GithubAppUpdate) HasHostname() bool {
	return o != nil && o.Hostname != nil
}

// SetHostname gets a reference to the given GithubAppUpdateHostname and assigns it to the Hostname field.
func (o *GithubAppUpdate) SetHostname(v GithubAppUpdateHostname) {
	o.Hostname = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *GithubAppUpdate) GetPrivateKey() GithubAppUpdatePrivateKey {
	if o == nil || o.PrivateKey == nil {
		var ret GithubAppUpdatePrivateKey
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubAppUpdate) GetPrivateKeyOk() (*GithubAppUpdatePrivateKey, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *GithubAppUpdate) HasPrivateKey() bool {
	return o != nil && o.PrivateKey != nil
}

// SetPrivateKey gets a reference to the given GithubAppUpdatePrivateKey and assigns it to the PrivateKey field.
func (o *GithubAppUpdate) SetPrivateKey(v GithubAppUpdatePrivateKey) {
	o.PrivateKey = &v
}

// GetType returns the Type field value.
func (o *GithubAppUpdate) GetType() GithubAppType {
	if o == nil {
		var ret GithubAppType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GithubAppUpdate) GetTypeOk() (*GithubAppType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *GithubAppUpdate) SetType(v GithubAppType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GithubAppUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.GithubAppId != nil {
		toSerialize["github_app_id"] = o.GithubAppId
	}
	if o.GithubInstallationId != nil {
		toSerialize["github_installation_id"] = o.GithubInstallationId
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
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
func (o *GithubAppUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GithubAppId          *GithubAppUpdateGithubAppId          `json:"github_app_id,omitempty"`
		GithubInstallationId *GithubAppUpdateGithubInstallationId `json:"github_installation_id,omitempty"`
		Hostname             *GithubAppUpdateHostname             `json:"hostname,omitempty"`
		PrivateKey           *GithubAppUpdatePrivateKey           `json:"private_key,omitempty"`
		Type                 *GithubAppType                       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"github_app_id", "github_installation_id", "hostname", "private_key", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.GithubAppId = all.GithubAppId
	o.GithubInstallationId = all.GithubInstallationId
	o.Hostname = all.Hostname
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
