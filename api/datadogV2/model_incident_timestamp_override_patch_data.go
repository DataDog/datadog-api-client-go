// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTimestampOverridePatchData Data for patching an incident timestamp override.
type IncidentTimestampOverridePatchData struct {
	// Attributes for patching an incident timestamp override.
	Attributes IncidentTimestampOverridePatchAttributes `json:"attributes"`
	// The JSON:API type for timestamp overrides.
	Type IncidentsTimestampOverridesType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentTimestampOverridePatchData instantiates a new IncidentTimestampOverridePatchData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentTimestampOverridePatchData(attributes IncidentTimestampOverridePatchAttributes, typeVar IncidentsTimestampOverridesType) *IncidentTimestampOverridePatchData {
	this := IncidentTimestampOverridePatchData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewIncidentTimestampOverridePatchDataWithDefaults instantiates a new IncidentTimestampOverridePatchData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentTimestampOverridePatchDataWithDefaults() *IncidentTimestampOverridePatchData {
	this := IncidentTimestampOverridePatchData{}
	var typeVar IncidentsTimestampOverridesType = INCIDENTSTIMESTAMPOVERRIDESTYPE_INCIDENTS_TIMESTAMP_OVERRIDES
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *IncidentTimestampOverridePatchData) GetAttributes() IncidentTimestampOverridePatchAttributes {
	if o == nil {
		var ret IncidentTimestampOverridePatchAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *IncidentTimestampOverridePatchData) GetAttributesOk() (*IncidentTimestampOverridePatchAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *IncidentTimestampOverridePatchData) SetAttributes(v IncidentTimestampOverridePatchAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *IncidentTimestampOverridePatchData) GetType() IncidentsTimestampOverridesType {
	if o == nil {
		var ret IncidentsTimestampOverridesType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentTimestampOverridePatchData) GetTypeOk() (*IncidentsTimestampOverridesType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentTimestampOverridePatchData) SetType(v IncidentsTimestampOverridesType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentTimestampOverridePatchData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentTimestampOverridePatchData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *IncidentTimestampOverridePatchAttributes `json:"attributes"`
		Type       *IncidentsTimestampOverridesType          `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
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
