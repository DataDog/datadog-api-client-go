// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceRepositoryInfoRequestAttributes Attributes for the service repository info request.
type ServiceRepositoryInfoRequestAttributes struct {
	// The debug ID of a JavaScript source map, in UUID format. Alternative to `service` and `version`.
	DebugId *uuid.UUID `json:"debug_id,omitempty"`
	// The name of the service. Required with `version` when `debug_id` is not set.
	Service *string `json:"service,omitempty"`
	// The version of the service. Required with `service` when `debug_id` is not set.
	Version *string `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceRepositoryInfoRequestAttributes instantiates a new ServiceRepositoryInfoRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceRepositoryInfoRequestAttributes() *ServiceRepositoryInfoRequestAttributes {
	this := ServiceRepositoryInfoRequestAttributes{}
	return &this
}

// NewServiceRepositoryInfoRequestAttributesWithDefaults instantiates a new ServiceRepositoryInfoRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceRepositoryInfoRequestAttributesWithDefaults() *ServiceRepositoryInfoRequestAttributes {
	this := ServiceRepositoryInfoRequestAttributes{}
	return &this
}

// GetDebugId returns the DebugId field value if set, zero value otherwise.
func (o *ServiceRepositoryInfoRequestAttributes) GetDebugId() uuid.UUID {
	if o == nil || o.DebugId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.DebugId
}

// GetDebugIdOk returns a tuple with the DebugId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRepositoryInfoRequestAttributes) GetDebugIdOk() (*uuid.UUID, bool) {
	if o == nil || o.DebugId == nil {
		return nil, false
	}
	return o.DebugId, true
}

// HasDebugId returns a boolean if a field has been set.
func (o *ServiceRepositoryInfoRequestAttributes) HasDebugId() bool {
	return o != nil && o.DebugId != nil
}

// SetDebugId gets a reference to the given uuid.UUID and assigns it to the DebugId field.
func (o *ServiceRepositoryInfoRequestAttributes) SetDebugId(v uuid.UUID) {
	o.DebugId = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ServiceRepositoryInfoRequestAttributes) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRepositoryInfoRequestAttributes) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ServiceRepositoryInfoRequestAttributes) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *ServiceRepositoryInfoRequestAttributes) SetService(v string) {
	o.Service = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ServiceRepositoryInfoRequestAttributes) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRepositoryInfoRequestAttributes) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ServiceRepositoryInfoRequestAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ServiceRepositoryInfoRequestAttributes) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceRepositoryInfoRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DebugId != nil {
		toSerialize["debug_id"] = o.DebugId
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceRepositoryInfoRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DebugId *uuid.UUID `json:"debug_id,omitempty"`
		Service *string    `json:"service,omitempty"`
		Version *string    `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"debug_id", "service", "version"})
	} else {
		return err
	}
	o.DebugId = all.DebugId
	o.Service = all.Service
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
