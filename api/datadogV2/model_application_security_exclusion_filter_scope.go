// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityExclusionFilterScope The scope of the exclusion filter.
type ApplicationSecurityExclusionFilterScope struct {
	// The environment scope for the exclusion filter.
	Env *string `json:"env,omitempty"`
	// The service scope for the exclusion filter.
	Service *string `json:"service,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityExclusionFilterScope instantiates a new ApplicationSecurityExclusionFilterScope object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityExclusionFilterScope() *ApplicationSecurityExclusionFilterScope {
	this := ApplicationSecurityExclusionFilterScope{}
	return &this
}

// NewApplicationSecurityExclusionFilterScopeWithDefaults instantiates a new ApplicationSecurityExclusionFilterScope object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityExclusionFilterScopeWithDefaults() *ApplicationSecurityExclusionFilterScope {
	this := ApplicationSecurityExclusionFilterScope{}
	return &this
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterScope) GetEnv() string {
	if o == nil || o.Env == nil {
		var ret string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterScope) GetEnvOk() (*string, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterScope) HasEnv() bool {
	return o != nil && o.Env != nil
}

// SetEnv gets a reference to the given string and assigns it to the Env field.
func (o *ApplicationSecurityExclusionFilterScope) SetEnv(v string) {
	o.Env = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterScope) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterScope) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterScope) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *ApplicationSecurityExclusionFilterScope) SetService(v string) {
	o.Service = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityExclusionFilterScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationSecurityExclusionFilterScope) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Env     *string `json:"env,omitempty"`
		Service *string `json:"service,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"env", "service"})
	} else {
		return err
	}
	o.Env = all.Env
	o.Service = all.Service

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
