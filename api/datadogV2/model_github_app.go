// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GithubApp GitHub App credentials.
type GithubApp struct {
	// The numeric ID of the GitHub App.
	GithubAppId GithubAppGithubAppId `json:"github_app_id"`
	// The installation ID of the GitHub App.
	GithubInstallationId GithubAppGithubInstallationId `json:"github_installation_id"`
	// Hostname of the GitHub Enterprise Server. Omit this field for github.com.
	Hostname *GithubAppHostname `json:"hostname,omitempty"`
	// PEM-encoded RSA private key for the GitHub App.
	PrivateKey GithubAppPrivateKey `json:"private_key"`
	// The credential type for `GithubApp`.
	Type GithubAppType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGithubApp instantiates a new GithubApp object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGithubApp(githubAppId GithubAppGithubAppId, githubInstallationId GithubAppGithubInstallationId, privateKey GithubAppPrivateKey, typeVar GithubAppType) *GithubApp {
	this := GithubApp{}
	this.GithubAppId = githubAppId
	this.GithubInstallationId = githubInstallationId
	this.PrivateKey = privateKey
	this.Type = typeVar
	return &this
}

// NewGithubAppWithDefaults instantiates a new GithubApp object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGithubAppWithDefaults() *GithubApp {
	this := GithubApp{}
	return &this
}

// GetGithubAppId returns the GithubAppId field value.
func (o *GithubApp) GetGithubAppId() GithubAppGithubAppId {
	if o == nil {
		var ret GithubAppGithubAppId
		return ret
	}
	return o.GithubAppId
}

// GetGithubAppIdOk returns a tuple with the GithubAppId field value
// and a boolean to check if the value has been set.
func (o *GithubApp) GetGithubAppIdOk() (*GithubAppGithubAppId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GithubAppId, true
}

// SetGithubAppId sets field value.
func (o *GithubApp) SetGithubAppId(v GithubAppGithubAppId) {
	o.GithubAppId = v
}

// GetGithubInstallationId returns the GithubInstallationId field value.
func (o *GithubApp) GetGithubInstallationId() GithubAppGithubInstallationId {
	if o == nil {
		var ret GithubAppGithubInstallationId
		return ret
	}
	return o.GithubInstallationId
}

// GetGithubInstallationIdOk returns a tuple with the GithubInstallationId field value
// and a boolean to check if the value has been set.
func (o *GithubApp) GetGithubInstallationIdOk() (*GithubAppGithubInstallationId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GithubInstallationId, true
}

// SetGithubInstallationId sets field value.
func (o *GithubApp) SetGithubInstallationId(v GithubAppGithubInstallationId) {
	o.GithubInstallationId = v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *GithubApp) GetHostname() GithubAppHostname {
	if o == nil || o.Hostname == nil {
		var ret GithubAppHostname
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubApp) GetHostnameOk() (*GithubAppHostname, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *GithubApp) HasHostname() bool {
	return o != nil && o.Hostname != nil
}

// SetHostname gets a reference to the given GithubAppHostname and assigns it to the Hostname field.
func (o *GithubApp) SetHostname(v GithubAppHostname) {
	o.Hostname = &v
}

// GetPrivateKey returns the PrivateKey field value.
func (o *GithubApp) GetPrivateKey() GithubAppPrivateKey {
	if o == nil {
		var ret GithubAppPrivateKey
		return ret
	}
	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *GithubApp) GetPrivateKeyOk() (*GithubAppPrivateKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value.
func (o *GithubApp) SetPrivateKey(v GithubAppPrivateKey) {
	o.PrivateKey = v
}

// GetType returns the Type field value.
func (o *GithubApp) GetType() GithubAppType {
	if o == nil {
		var ret GithubAppType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GithubApp) GetTypeOk() (*GithubAppType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *GithubApp) SetType(v GithubAppType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GithubApp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["github_app_id"] = o.GithubAppId
	toSerialize["github_installation_id"] = o.GithubInstallationId
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	toSerialize["private_key"] = o.PrivateKey
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GithubApp) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GithubAppId          *GithubAppGithubAppId          `json:"github_app_id"`
		GithubInstallationId *GithubAppGithubInstallationId `json:"github_installation_id"`
		Hostname             *GithubAppHostname             `json:"hostname,omitempty"`
		PrivateKey           *GithubAppPrivateKey           `json:"private_key"`
		Type                 *GithubAppType                 `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.GithubAppId == nil {
		return fmt.Errorf("required field github_app_id missing")
	}
	if all.GithubInstallationId == nil {
		return fmt.Errorf("required field github_installation_id missing")
	}
	if all.PrivateKey == nil {
		return fmt.Errorf("required field private_key missing")
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
	o.GithubAppId = *all.GithubAppId
	o.GithubInstallationId = *all.GithubInstallationId
	o.Hostname = all.Hostname
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
