// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// ServiceDefinitionV2Opsgenie Opsgenie integration for the service.
type ServiceDefinitionV2Opsgenie struct {
	// Opsgenie instance region.
	Region *ServiceDefinitionV2OpsgenieRegion `json:"region,omitempty"`
	// Opsgenie service id.
	ServiceId string `json:"service-id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewServiceDefinitionV2Opsgenie instantiates a new ServiceDefinitionV2Opsgenie object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceDefinitionV2Opsgenie(serviceId string) *ServiceDefinitionV2Opsgenie {
	this := ServiceDefinitionV2Opsgenie{}
	this.ServiceId = serviceId
	return &this
}

// NewServiceDefinitionV2OpsgenieWithDefaults instantiates a new ServiceDefinitionV2Opsgenie object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceDefinitionV2OpsgenieWithDefaults() *ServiceDefinitionV2Opsgenie {
	this := ServiceDefinitionV2Opsgenie{}
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Opsgenie) GetRegion() ServiceDefinitionV2OpsgenieRegion {
	if o == nil || o.Region == nil {
		var ret ServiceDefinitionV2OpsgenieRegion
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Opsgenie) GetRegionOk() (*ServiceDefinitionV2OpsgenieRegion, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Opsgenie) HasRegion() bool {
	return o != nil && o.Region != nil
}

// SetRegion gets a reference to the given ServiceDefinitionV2OpsgenieRegion and assigns it to the Region field.
func (o *ServiceDefinitionV2Opsgenie) SetRegion(v ServiceDefinitionV2OpsgenieRegion) {
	o.Region = &v
}

// GetServiceId returns the ServiceId field value.
func (o *ServiceDefinitionV2Opsgenie) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Opsgenie) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value.
func (o *ServiceDefinitionV2Opsgenie) SetServiceId(v string) {
	o.ServiceId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceDefinitionV2Opsgenie) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	toSerialize["service-id"] = o.ServiceId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceDefinitionV2Opsgenie) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		ServiceId *string `json:"service-id"`
	}{}
	all := struct {
		Region    *ServiceDefinitionV2OpsgenieRegion `json:"region,omitempty"`
		ServiceId string                             `json:"service-id"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.ServiceId == nil {
		return fmt.Errorf("required field service-id missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.Region; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Region = all.Region
	o.ServiceId = all.ServiceId
	return nil
}
