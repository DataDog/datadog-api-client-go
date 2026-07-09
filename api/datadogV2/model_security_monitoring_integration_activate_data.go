// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringIntegrationActivateData The configuration overrides for the integration to activate.
type SecurityMonitoringIntegrationActivateData struct {
	// Overrides applied when activating the integration. All fields are optional.
	Attributes *SecurityMonitoringIntegrationActivateAttributes `json:"attributes,omitempty"`
	// The type of the resource. The value should always be `activate_entra_id_request`.
	Type *SecurityMonitoringIntegrationActivateResourceType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringIntegrationActivateData instantiates a new SecurityMonitoringIntegrationActivateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringIntegrationActivateData() *SecurityMonitoringIntegrationActivateData {
	this := SecurityMonitoringIntegrationActivateData{}
	var typeVar SecurityMonitoringIntegrationActivateResourceType = SECURITYMONITORINGINTEGRATIONACTIVATERESOURCETYPE_ACTIVATE_ENTRA_ID_REQUEST
	this.Type = &typeVar
	return &this
}

// NewSecurityMonitoringIntegrationActivateDataWithDefaults instantiates a new SecurityMonitoringIntegrationActivateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringIntegrationActivateDataWithDefaults() *SecurityMonitoringIntegrationActivateData {
	this := SecurityMonitoringIntegrationActivateData{}
	var typeVar SecurityMonitoringIntegrationActivateResourceType = SECURITYMONITORINGINTEGRATIONACTIVATERESOURCETYPE_ACTIVATE_ENTRA_ID_REQUEST
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SecurityMonitoringIntegrationActivateData) GetAttributes() SecurityMonitoringIntegrationActivateAttributes {
	if o == nil || o.Attributes == nil {
		var ret SecurityMonitoringIntegrationActivateAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationActivateData) GetAttributesOk() (*SecurityMonitoringIntegrationActivateAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SecurityMonitoringIntegrationActivateData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given SecurityMonitoringIntegrationActivateAttributes and assigns it to the Attributes field.
func (o *SecurityMonitoringIntegrationActivateData) SetAttributes(v SecurityMonitoringIntegrationActivateAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SecurityMonitoringIntegrationActivateData) GetType() SecurityMonitoringIntegrationActivateResourceType {
	if o == nil || o.Type == nil {
		var ret SecurityMonitoringIntegrationActivateResourceType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationActivateData) GetTypeOk() (*SecurityMonitoringIntegrationActivateResourceType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SecurityMonitoringIntegrationActivateData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given SecurityMonitoringIntegrationActivateResourceType and assigns it to the Type field.
func (o *SecurityMonitoringIntegrationActivateData) SetType(v SecurityMonitoringIntegrationActivateResourceType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringIntegrationActivateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringIntegrationActivateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *SecurityMonitoringIntegrationActivateAttributes   `json:"attributes,omitempty"`
		Type       *SecurityMonitoringIntegrationActivateResourceType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
