// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudIntegrationAccountUpdateAttributes Updatable attributes of an Elastic Cloud integration account. Every field is optional; only the fields provided are changed.
type ElasticCloudIntegrationAccountUpdateAttributes struct {
	// Partial Elastic Cloud interface for updates. Exactly one interface variant is set, selected by its `type`.
	Interface *ElasticCloudInterfaceUpdate `json:"interface,omitempty"`
	// Human-readable name of the account.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticCloudIntegrationAccountUpdateAttributes instantiates a new ElasticCloudIntegrationAccountUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticCloudIntegrationAccountUpdateAttributes() *ElasticCloudIntegrationAccountUpdateAttributes {
	this := ElasticCloudIntegrationAccountUpdateAttributes{}
	return &this
}

// NewElasticCloudIntegrationAccountUpdateAttributesWithDefaults instantiates a new ElasticCloudIntegrationAccountUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticCloudIntegrationAccountUpdateAttributesWithDefaults() *ElasticCloudIntegrationAccountUpdateAttributes {
	this := ElasticCloudIntegrationAccountUpdateAttributes{}
	return &this
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *ElasticCloudIntegrationAccountUpdateAttributes) GetInterface() ElasticCloudInterfaceUpdate {
	if o == nil || o.Interface == nil {
		var ret ElasticCloudInterfaceUpdate
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudIntegrationAccountUpdateAttributes) GetInterfaceOk() (*ElasticCloudInterfaceUpdate, bool) {
	if o == nil || o.Interface == nil {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *ElasticCloudIntegrationAccountUpdateAttributes) HasInterface() bool {
	return o != nil && o.Interface != nil
}

// SetInterface gets a reference to the given ElasticCloudInterfaceUpdate and assigns it to the Interface field.
func (o *ElasticCloudIntegrationAccountUpdateAttributes) SetInterface(v ElasticCloudInterfaceUpdate) {
	o.Interface = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ElasticCloudIntegrationAccountUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudIntegrationAccountUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ElasticCloudIntegrationAccountUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ElasticCloudIntegrationAccountUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticCloudIntegrationAccountUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Interface != nil {
		toSerialize["interface"] = o.Interface
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticCloudIntegrationAccountUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Interface *ElasticCloudInterfaceUpdate `json:"interface,omitempty"`
		Name      *string                      `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"interface", "name"})
	} else {
		return err
	}
	o.Interface = all.Interface
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
