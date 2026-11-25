// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineClientIdentity Restricts which clients may connect, based on the identities in their TLS certificate. Checks run during the TLS handshake, after the client certificate chain is validated against `ca_file`, and never relax that validation.
// Requires `verify_certificate: true` and `ca_file`. Set exactly one of `allowlist_file` or `allowlist`.
// When set, connections are closed after a maximum age of 300 seconds so that clients removed from the allowlist are disconnected.
// Requires Observability Pipelines Worker 2.27 or later.
type ObservabilityPipelineClientIdentity struct {
	// An allowlist of client identities. A client is allowed when any rule matches.
	Allowlist *ObservabilityPipelineClientIdentityAllowlist `json:"allowlist,omitempty"`
	// Path to an allowlist file, relative to the Worker's data directory. The Worker reloads the file when it changes. Must not contain parent directory (`..`) references.
	AllowlistFile *string `json:"allowlist_file,omitempty"`
	// What happens to a client that no rule allows. `enforce` rejects the TLS handshake. `audit` accepts the client and reports that it would have been rejected.
	Mode *ObservabilityPipelineClientIdentityMode `json:"mode,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineClientIdentity instantiates a new ObservabilityPipelineClientIdentity object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineClientIdentity() *ObservabilityPipelineClientIdentity {
	this := ObservabilityPipelineClientIdentity{}
	var mode ObservabilityPipelineClientIdentityMode = OBSERVABILITYPIPELINECLIENTIDENTITYMODE_ENFORCE
	this.Mode = &mode
	return &this
}

// NewObservabilityPipelineClientIdentityWithDefaults instantiates a new ObservabilityPipelineClientIdentity object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineClientIdentityWithDefaults() *ObservabilityPipelineClientIdentity {
	this := ObservabilityPipelineClientIdentity{}
	var mode ObservabilityPipelineClientIdentityMode = OBSERVABILITYPIPELINECLIENTIDENTITYMODE_ENFORCE
	this.Mode = &mode
	return &this
}

// GetAllowlist returns the Allowlist field value if set, zero value otherwise.
func (o *ObservabilityPipelineClientIdentity) GetAllowlist() ObservabilityPipelineClientIdentityAllowlist {
	if o == nil || o.Allowlist == nil {
		var ret ObservabilityPipelineClientIdentityAllowlist
		return ret
	}
	return *o.Allowlist
}

// GetAllowlistOk returns a tuple with the Allowlist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineClientIdentity) GetAllowlistOk() (*ObservabilityPipelineClientIdentityAllowlist, bool) {
	if o == nil || o.Allowlist == nil {
		return nil, false
	}
	return o.Allowlist, true
}

// HasAllowlist returns a boolean if a field has been set.
func (o *ObservabilityPipelineClientIdentity) HasAllowlist() bool {
	return o != nil && o.Allowlist != nil
}

// SetAllowlist gets a reference to the given ObservabilityPipelineClientIdentityAllowlist and assigns it to the Allowlist field.
func (o *ObservabilityPipelineClientIdentity) SetAllowlist(v ObservabilityPipelineClientIdentityAllowlist) {
	o.Allowlist = &v
}

// GetAllowlistFile returns the AllowlistFile field value if set, zero value otherwise.
func (o *ObservabilityPipelineClientIdentity) GetAllowlistFile() string {
	if o == nil || o.AllowlistFile == nil {
		var ret string
		return ret
	}
	return *o.AllowlistFile
}

// GetAllowlistFileOk returns a tuple with the AllowlistFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineClientIdentity) GetAllowlistFileOk() (*string, bool) {
	if o == nil || o.AllowlistFile == nil {
		return nil, false
	}
	return o.AllowlistFile, true
}

// HasAllowlistFile returns a boolean if a field has been set.
func (o *ObservabilityPipelineClientIdentity) HasAllowlistFile() bool {
	return o != nil && o.AllowlistFile != nil
}

// SetAllowlistFile gets a reference to the given string and assigns it to the AllowlistFile field.
func (o *ObservabilityPipelineClientIdentity) SetAllowlistFile(v string) {
	o.AllowlistFile = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ObservabilityPipelineClientIdentity) GetMode() ObservabilityPipelineClientIdentityMode {
	if o == nil || o.Mode == nil {
		var ret ObservabilityPipelineClientIdentityMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineClientIdentity) GetModeOk() (*ObservabilityPipelineClientIdentityMode, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ObservabilityPipelineClientIdentity) HasMode() bool {
	return o != nil && o.Mode != nil
}

// SetMode gets a reference to the given ObservabilityPipelineClientIdentityMode and assigns it to the Mode field.
func (o *ObservabilityPipelineClientIdentity) SetMode(v ObservabilityPipelineClientIdentityMode) {
	o.Mode = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineClientIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Allowlist != nil {
		toSerialize["allowlist"] = o.Allowlist
	}
	if o.AllowlistFile != nil {
		toSerialize["allowlist_file"] = o.AllowlistFile
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineClientIdentity) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Allowlist     *ObservabilityPipelineClientIdentityAllowlist `json:"allowlist,omitempty"`
		AllowlistFile *string                                       `json:"allowlist_file,omitempty"`
		Mode          *ObservabilityPipelineClientIdentityMode      `json:"mode,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allowlist", "allowlist_file", "mode"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Allowlist != nil && all.Allowlist.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Allowlist = all.Allowlist
	o.AllowlistFile = all.AllowlistFile
	if all.Mode != nil && !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = all.Mode
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
