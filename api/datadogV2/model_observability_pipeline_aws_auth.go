// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAwsAuth AWS authentication credentials used for accessing AWS services such as S3.
// If omitted, the system’s default credentials are used (for example, the IAM role and environment variables).
type ObservabilityPipelineAwsAuth struct {
	// The Amazon Resource Name (ARN) of the role to assume.
	AssumeRole datadog.NullableString `json:"assume_role,omitempty"`
	// A unique identifier for cross-account role assumption.
	ExternalId datadog.NullableString `json:"external_id,omitempty"`
	// A session identifier used for logging and tracing the assumed role session.
	SessionName datadog.NullableString `json:"session_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineAwsAuth instantiates a new ObservabilityPipelineAwsAuth object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineAwsAuth() *ObservabilityPipelineAwsAuth {
	this := ObservabilityPipelineAwsAuth{}
	return &this
}

// NewObservabilityPipelineAwsAuthWithDefaults instantiates a new ObservabilityPipelineAwsAuth object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineAwsAuthWithDefaults() *ObservabilityPipelineAwsAuth {
	this := ObservabilityPipelineAwsAuth{}
	return &this
}

// GetAssumeRole returns the AssumeRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObservabilityPipelineAwsAuth) GetAssumeRole() string {
	if o == nil || o.AssumeRole.Get() == nil {
		var ret string
		return ret
	}
	return *o.AssumeRole.Get()
}

// GetAssumeRoleOk returns a tuple with the AssumeRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ObservabilityPipelineAwsAuth) GetAssumeRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssumeRole.Get(), o.AssumeRole.IsSet()
}

// HasAssumeRole returns a boolean if a field has been set.
func (o *ObservabilityPipelineAwsAuth) HasAssumeRole() bool {
	return o != nil && o.AssumeRole.IsSet()
}

// SetAssumeRole gets a reference to the given datadog.NullableString and assigns it to the AssumeRole field.
func (o *ObservabilityPipelineAwsAuth) SetAssumeRole(v string) {
	o.AssumeRole.Set(&v)
}

// SetAssumeRoleNil sets the value for AssumeRole to be an explicit nil.
func (o *ObservabilityPipelineAwsAuth) SetAssumeRoleNil() {
	o.AssumeRole.Set(nil)
}

// UnsetAssumeRole ensures that no value is present for AssumeRole, not even an explicit nil.
func (o *ObservabilityPipelineAwsAuth) UnsetAssumeRole() {
	o.AssumeRole.Unset()
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObservabilityPipelineAwsAuth) GetExternalId() string {
	if o == nil || o.ExternalId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ObservabilityPipelineAwsAuth) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// HasExternalId returns a boolean if a field has been set.
func (o *ObservabilityPipelineAwsAuth) HasExternalId() bool {
	return o != nil && o.ExternalId.IsSet()
}

// SetExternalId gets a reference to the given datadog.NullableString and assigns it to the ExternalId field.
func (o *ObservabilityPipelineAwsAuth) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}

// SetExternalIdNil sets the value for ExternalId to be an explicit nil.
func (o *ObservabilityPipelineAwsAuth) SetExternalIdNil() {
	o.ExternalId.Set(nil)
}

// UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil.
func (o *ObservabilityPipelineAwsAuth) UnsetExternalId() {
	o.ExternalId.Unset()
}

// GetSessionName returns the SessionName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObservabilityPipelineAwsAuth) GetSessionName() string {
	if o == nil || o.SessionName.Get() == nil {
		var ret string
		return ret
	}
	return *o.SessionName.Get()
}

// GetSessionNameOk returns a tuple with the SessionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ObservabilityPipelineAwsAuth) GetSessionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionName.Get(), o.SessionName.IsSet()
}

// HasSessionName returns a boolean if a field has been set.
func (o *ObservabilityPipelineAwsAuth) HasSessionName() bool {
	return o != nil && o.SessionName.IsSet()
}

// SetSessionName gets a reference to the given datadog.NullableString and assigns it to the SessionName field.
func (o *ObservabilityPipelineAwsAuth) SetSessionName(v string) {
	o.SessionName.Set(&v)
}

// SetSessionNameNil sets the value for SessionName to be an explicit nil.
func (o *ObservabilityPipelineAwsAuth) SetSessionNameNil() {
	o.SessionName.Set(nil)
}

// UnsetSessionName ensures that no value is present for SessionName, not even an explicit nil.
func (o *ObservabilityPipelineAwsAuth) UnsetSessionName() {
	o.SessionName.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineAwsAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AssumeRole.IsSet() {
		toSerialize["assume_role"] = o.AssumeRole.Get()
	}
	if o.ExternalId.IsSet() {
		toSerialize["external_id"] = o.ExternalId.Get()
	}
	if o.SessionName.IsSet() {
		toSerialize["session_name"] = o.SessionName.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineAwsAuth) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssumeRole  datadog.NullableString `json:"assume_role,omitempty"`
		ExternalId  datadog.NullableString `json:"external_id,omitempty"`
		SessionName datadog.NullableString `json:"session_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assume_role", "external_id", "session_name"})
	} else {
		return err
	}
	o.AssumeRole = all.AssumeRole
	o.ExternalId = all.ExternalId
	o.SessionName = all.SessionName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
