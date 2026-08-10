// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudCcmInterface Elastic Cloud CCM (Cloud Cost Management) interface configuration.
type ElasticCloudCcmInterface struct {
	// Authentication methods supported by the Elastic Cloud CCM interface. Exactly one is set, selected by its `type`.
	Authentication ElasticCloudCcmAuthentication `json:"authentication"`
	// Dataflows for the Elastic Cloud CCM interface.
	Dataflows []ElasticCloudCcmDataflow `json:"dataflows,omitempty"`
	// Elastic Cloud CCM interface settings.
	Settings *ElasticCloudCcmSettings `json:"settings,omitempty"`
	// Interface discriminator for the Elastic Cloud CCM interface.
	Type ElasticCloudCcmInterfaceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticCloudCcmInterface instantiates a new ElasticCloudCcmInterface object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticCloudCcmInterface(authentication ElasticCloudCcmAuthentication, typeVar ElasticCloudCcmInterfaceType) *ElasticCloudCcmInterface {
	this := ElasticCloudCcmInterface{}
	this.Authentication = authentication
	this.Type = typeVar
	return &this
}

// NewElasticCloudCcmInterfaceWithDefaults instantiates a new ElasticCloudCcmInterface object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticCloudCcmInterfaceWithDefaults() *ElasticCloudCcmInterface {
	this := ElasticCloudCcmInterface{}
	return &this
}

// GetAuthentication returns the Authentication field value.
func (o *ElasticCloudCcmInterface) GetAuthentication() ElasticCloudCcmAuthentication {
	if o == nil {
		var ret ElasticCloudCcmAuthentication
		return ret
	}
	return o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value
// and a boolean to check if the value has been set.
func (o *ElasticCloudCcmInterface) GetAuthenticationOk() (*ElasticCloudCcmAuthentication, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authentication, true
}

// SetAuthentication sets field value.
func (o *ElasticCloudCcmInterface) SetAuthentication(v ElasticCloudCcmAuthentication) {
	o.Authentication = v
}

// GetDataflows returns the Dataflows field value if set, zero value otherwise.
func (o *ElasticCloudCcmInterface) GetDataflows() []ElasticCloudCcmDataflow {
	if o == nil || o.Dataflows == nil {
		var ret []ElasticCloudCcmDataflow
		return ret
	}
	return o.Dataflows
}

// GetDataflowsOk returns a tuple with the Dataflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudCcmInterface) GetDataflowsOk() (*[]ElasticCloudCcmDataflow, bool) {
	if o == nil || o.Dataflows == nil {
		return nil, false
	}
	return &o.Dataflows, true
}

// HasDataflows returns a boolean if a field has been set.
func (o *ElasticCloudCcmInterface) HasDataflows() bool {
	return o != nil && o.Dataflows != nil
}

// SetDataflows gets a reference to the given []ElasticCloudCcmDataflow and assigns it to the Dataflows field.
func (o *ElasticCloudCcmInterface) SetDataflows(v []ElasticCloudCcmDataflow) {
	o.Dataflows = v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *ElasticCloudCcmInterface) GetSettings() ElasticCloudCcmSettings {
	if o == nil || o.Settings == nil {
		var ret ElasticCloudCcmSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudCcmInterface) GetSettingsOk() (*ElasticCloudCcmSettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *ElasticCloudCcmInterface) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given ElasticCloudCcmSettings and assigns it to the Settings field.
func (o *ElasticCloudCcmInterface) SetSettings(v ElasticCloudCcmSettings) {
	o.Settings = &v
}

// GetType returns the Type field value.
func (o *ElasticCloudCcmInterface) GetType() ElasticCloudCcmInterfaceType {
	if o == nil {
		var ret ElasticCloudCcmInterfaceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ElasticCloudCcmInterface) GetTypeOk() (*ElasticCloudCcmInterfaceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ElasticCloudCcmInterface) SetType(v ElasticCloudCcmInterfaceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticCloudCcmInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["authentication"] = o.Authentication
	if o.Dataflows != nil {
		toSerialize["dataflows"] = o.Dataflows
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticCloudCcmInterface) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *ElasticCloudCcmAuthentication `json:"authentication"`
		Dataflows      []ElasticCloudCcmDataflow      `json:"dataflows,omitempty"`
		Settings       *ElasticCloudCcmSettings       `json:"settings,omitempty"`
		Type           *ElasticCloudCcmInterfaceType  `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Authentication == nil {
		return fmt.Errorf("required field authentication missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"authentication", "dataflows", "settings", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Authentication = *all.Authentication
	o.Dataflows = all.Dataflows
	if all.Settings != nil && all.Settings.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Settings = all.Settings
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
