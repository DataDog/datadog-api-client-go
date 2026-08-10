// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudMonitoringInterfaceUpdate Partial Elastic Cloud monitoring interface for updates.
type ElasticCloudMonitoringInterfaceUpdate struct {
	// Authentication methods supported by the Elastic Cloud interface. Exactly one is set, selected by its `type`.
	Authentication *ElasticCloudAuthentication `json:"authentication,omitempty"`
	// Dataflows for the Elastic Cloud monitoring interface.
	Dataflows []ElasticCloudDataflow `json:"dataflows,omitempty"`
	// Partial Elastic Cloud interface settings for updates.
	Settings *ElasticCloudSettingsUpdate `json:"settings,omitempty"`
	// Interface discriminator for the Elastic Cloud monitoring interface.
	Type ElasticCloudMonitoringInterfaceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticCloudMonitoringInterfaceUpdate instantiates a new ElasticCloudMonitoringInterfaceUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticCloudMonitoringInterfaceUpdate(typeVar ElasticCloudMonitoringInterfaceType) *ElasticCloudMonitoringInterfaceUpdate {
	this := ElasticCloudMonitoringInterfaceUpdate{}
	this.Type = typeVar
	return &this
}

// NewElasticCloudMonitoringInterfaceUpdateWithDefaults instantiates a new ElasticCloudMonitoringInterfaceUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticCloudMonitoringInterfaceUpdateWithDefaults() *ElasticCloudMonitoringInterfaceUpdate {
	this := ElasticCloudMonitoringInterfaceUpdate{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *ElasticCloudMonitoringInterfaceUpdate) GetAuthentication() ElasticCloudAuthentication {
	if o == nil || o.Authentication == nil {
		var ret ElasticCloudAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudMonitoringInterfaceUpdate) GetAuthenticationOk() (*ElasticCloudAuthentication, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *ElasticCloudMonitoringInterfaceUpdate) HasAuthentication() bool {
	return o != nil && o.Authentication != nil
}

// SetAuthentication gets a reference to the given ElasticCloudAuthentication and assigns it to the Authentication field.
func (o *ElasticCloudMonitoringInterfaceUpdate) SetAuthentication(v ElasticCloudAuthentication) {
	o.Authentication = &v
}

// GetDataflows returns the Dataflows field value if set, zero value otherwise.
func (o *ElasticCloudMonitoringInterfaceUpdate) GetDataflows() []ElasticCloudDataflow {
	if o == nil || o.Dataflows == nil {
		var ret []ElasticCloudDataflow
		return ret
	}
	return o.Dataflows
}

// GetDataflowsOk returns a tuple with the Dataflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudMonitoringInterfaceUpdate) GetDataflowsOk() (*[]ElasticCloudDataflow, bool) {
	if o == nil || o.Dataflows == nil {
		return nil, false
	}
	return &o.Dataflows, true
}

// HasDataflows returns a boolean if a field has been set.
func (o *ElasticCloudMonitoringInterfaceUpdate) HasDataflows() bool {
	return o != nil && o.Dataflows != nil
}

// SetDataflows gets a reference to the given []ElasticCloudDataflow and assigns it to the Dataflows field.
func (o *ElasticCloudMonitoringInterfaceUpdate) SetDataflows(v []ElasticCloudDataflow) {
	o.Dataflows = v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *ElasticCloudMonitoringInterfaceUpdate) GetSettings() ElasticCloudSettingsUpdate {
	if o == nil || o.Settings == nil {
		var ret ElasticCloudSettingsUpdate
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudMonitoringInterfaceUpdate) GetSettingsOk() (*ElasticCloudSettingsUpdate, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *ElasticCloudMonitoringInterfaceUpdate) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given ElasticCloudSettingsUpdate and assigns it to the Settings field.
func (o *ElasticCloudMonitoringInterfaceUpdate) SetSettings(v ElasticCloudSettingsUpdate) {
	o.Settings = &v
}

// GetType returns the Type field value.
func (o *ElasticCloudMonitoringInterfaceUpdate) GetType() ElasticCloudMonitoringInterfaceType {
	if o == nil {
		var ret ElasticCloudMonitoringInterfaceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ElasticCloudMonitoringInterfaceUpdate) GetTypeOk() (*ElasticCloudMonitoringInterfaceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ElasticCloudMonitoringInterfaceUpdate) SetType(v ElasticCloudMonitoringInterfaceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticCloudMonitoringInterfaceUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Authentication != nil {
		toSerialize["authentication"] = o.Authentication
	}
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
func (o *ElasticCloudMonitoringInterfaceUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *ElasticCloudAuthentication          `json:"authentication,omitempty"`
		Dataflows      []ElasticCloudDataflow               `json:"dataflows,omitempty"`
		Settings       *ElasticCloudSettingsUpdate          `json:"settings,omitempty"`
		Type           *ElasticCloudMonitoringInterfaceType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	o.Authentication = all.Authentication
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
