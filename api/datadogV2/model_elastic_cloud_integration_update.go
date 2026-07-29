// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudIntegrationUpdate Partial Elastic Cloud integration configuration for updates.
type ElasticCloudIntegrationUpdate struct {
	// Partial Elastic Cloud interface for updates. Exactly one interface variant is set, selected by its `type`.
	Interface *ElasticCloudInterfaceUpdate `json:"interface,omitempty"`
	// Integration discriminator for Elastic Cloud.
	Type ElasticCloudIntegrationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticCloudIntegrationUpdate instantiates a new ElasticCloudIntegrationUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticCloudIntegrationUpdate(typeVar ElasticCloudIntegrationType) *ElasticCloudIntegrationUpdate {
	this := ElasticCloudIntegrationUpdate{}
	this.Type = typeVar
	return &this
}

// NewElasticCloudIntegrationUpdateWithDefaults instantiates a new ElasticCloudIntegrationUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticCloudIntegrationUpdateWithDefaults() *ElasticCloudIntegrationUpdate {
	this := ElasticCloudIntegrationUpdate{}
	return &this
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *ElasticCloudIntegrationUpdate) GetInterface() ElasticCloudInterfaceUpdate {
	if o == nil || o.Interface == nil {
		var ret ElasticCloudInterfaceUpdate
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudIntegrationUpdate) GetInterfaceOk() (*ElasticCloudInterfaceUpdate, bool) {
	if o == nil || o.Interface == nil {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *ElasticCloudIntegrationUpdate) HasInterface() bool {
	return o != nil && o.Interface != nil
}

// SetInterface gets a reference to the given ElasticCloudInterfaceUpdate and assigns it to the Interface field.
func (o *ElasticCloudIntegrationUpdate) SetInterface(v ElasticCloudInterfaceUpdate) {
	o.Interface = &v
}

// GetType returns the Type field value.
func (o *ElasticCloudIntegrationUpdate) GetType() ElasticCloudIntegrationType {
	if o == nil {
		var ret ElasticCloudIntegrationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ElasticCloudIntegrationUpdate) GetTypeOk() (*ElasticCloudIntegrationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ElasticCloudIntegrationUpdate) SetType(v ElasticCloudIntegrationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticCloudIntegrationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Interface != nil {
		toSerialize["interface"] = o.Interface
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticCloudIntegrationUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Interface *ElasticCloudInterfaceUpdate `json:"interface,omitempty"`
		Type      *ElasticCloudIntegrationType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	o.Interface = all.Interface
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
