// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudIntegration Elastic Cloud integration configuration.
type ElasticCloudIntegration struct {
	// Elastic Cloud interface (source-type). Exactly one interface variant is set, selected by its `type`.
	Interface ElasticCloudInterface `json:"interface"`
	// Integration discriminator for Elastic Cloud.
	Type ElasticCloudIntegrationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticCloudIntegration instantiates a new ElasticCloudIntegration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticCloudIntegration(interfaceVar ElasticCloudInterface, typeVar ElasticCloudIntegrationType) *ElasticCloudIntegration {
	this := ElasticCloudIntegration{}
	this.Interface = interfaceVar
	this.Type = typeVar
	return &this
}

// NewElasticCloudIntegrationWithDefaults instantiates a new ElasticCloudIntegration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticCloudIntegrationWithDefaults() *ElasticCloudIntegration {
	this := ElasticCloudIntegration{}
	return &this
}

// GetInterface returns the Interface field value.
func (o *ElasticCloudIntegration) GetInterface() ElasticCloudInterface {
	if o == nil {
		var ret ElasticCloudInterface
		return ret
	}
	return o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value
// and a boolean to check if the value has been set.
func (o *ElasticCloudIntegration) GetInterfaceOk() (*ElasticCloudInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interface, true
}

// SetInterface sets field value.
func (o *ElasticCloudIntegration) SetInterface(v ElasticCloudInterface) {
	o.Interface = v
}

// GetType returns the Type field value.
func (o *ElasticCloudIntegration) GetType() ElasticCloudIntegrationType {
	if o == nil {
		var ret ElasticCloudIntegrationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ElasticCloudIntegration) GetTypeOk() (*ElasticCloudIntegrationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ElasticCloudIntegration) SetType(v ElasticCloudIntegrationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticCloudIntegration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["interface"] = o.Interface
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticCloudIntegration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Interface *ElasticCloudInterface       `json:"interface"`
		Type      *ElasticCloudIntegrationType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Interface == nil {
		return fmt.Errorf("required field interface missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"interface", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Interface = *all.Interface
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
