// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentGoogleChatConfigurationDataRequest Google Chat configuration data in a create request.
type IncidentGoogleChatConfigurationDataRequest struct {
	// Attributes for creating a Google Chat configuration.
	Attributes IncidentGoogleChatConfigurationDataAttributesRequest `json:"attributes"`
	// Relationships for a Google Chat configuration create request.
	Relationships IncidentGoogleChatConfigurationRelationshipsRequest `json:"relationships"`
	// Google Chat configuration resource type.
	Type IncidentGoogleChatConfigurationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentGoogleChatConfigurationDataRequest instantiates a new IncidentGoogleChatConfigurationDataRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentGoogleChatConfigurationDataRequest(attributes IncidentGoogleChatConfigurationDataAttributesRequest, relationships IncidentGoogleChatConfigurationRelationshipsRequest, typeVar IncidentGoogleChatConfigurationType) *IncidentGoogleChatConfigurationDataRequest {
	this := IncidentGoogleChatConfigurationDataRequest{}
	this.Attributes = attributes
	this.Relationships = relationships
	this.Type = typeVar
	return &this
}

// NewIncidentGoogleChatConfigurationDataRequestWithDefaults instantiates a new IncidentGoogleChatConfigurationDataRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentGoogleChatConfigurationDataRequestWithDefaults() *IncidentGoogleChatConfigurationDataRequest {
	this := IncidentGoogleChatConfigurationDataRequest{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *IncidentGoogleChatConfigurationDataRequest) GetAttributes() IncidentGoogleChatConfigurationDataAttributesRequest {
	if o == nil {
		var ret IncidentGoogleChatConfigurationDataAttributesRequest
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *IncidentGoogleChatConfigurationDataRequest) GetAttributesOk() (*IncidentGoogleChatConfigurationDataAttributesRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *IncidentGoogleChatConfigurationDataRequest) SetAttributes(v IncidentGoogleChatConfigurationDataAttributesRequest) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value.
func (o *IncidentGoogleChatConfigurationDataRequest) GetRelationships() IncidentGoogleChatConfigurationRelationshipsRequest {
	if o == nil {
		var ret IncidentGoogleChatConfigurationRelationshipsRequest
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *IncidentGoogleChatConfigurationDataRequest) GetRelationshipsOk() (*IncidentGoogleChatConfigurationRelationshipsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value.
func (o *IncidentGoogleChatConfigurationDataRequest) SetRelationships(v IncidentGoogleChatConfigurationRelationshipsRequest) {
	o.Relationships = v
}

// GetType returns the Type field value.
func (o *IncidentGoogleChatConfigurationDataRequest) GetType() IncidentGoogleChatConfigurationType {
	if o == nil {
		var ret IncidentGoogleChatConfigurationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentGoogleChatConfigurationDataRequest) GetTypeOk() (*IncidentGoogleChatConfigurationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentGoogleChatConfigurationDataRequest) SetType(v IncidentGoogleChatConfigurationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentGoogleChatConfigurationDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentGoogleChatConfigurationDataRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *IncidentGoogleChatConfigurationDataAttributesRequest `json:"attributes"`
		Relationships *IncidentGoogleChatConfigurationRelationshipsRequest  `json:"relationships"`
		Type          *IncidentGoogleChatConfigurationType                  `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Relationships == nil {
		return fmt.Errorf("required field relationships missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	if all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = *all.Relationships
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
