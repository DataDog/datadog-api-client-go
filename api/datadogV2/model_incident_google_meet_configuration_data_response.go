// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentGoogleMeetConfigurationDataResponse Google Meet configuration data in a response.
type IncidentGoogleMeetConfigurationDataResponse struct {
	// Attributes of a Google Meet configuration.
	Attributes IncidentGoogleMeetConfigurationDataAttributesResponse `json:"attributes"`
	// The configuration identifier.
	Id uuid.UUID `json:"id"`
	// Relationships for a Google Meet configuration.
	Relationships *IncidentGoogleMeetConfigurationRelationships `json:"relationships,omitempty"`
	// Google Meet configuration resource type.
	Type IncidentGoogleMeetConfigurationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentGoogleMeetConfigurationDataResponse instantiates a new IncidentGoogleMeetConfigurationDataResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentGoogleMeetConfigurationDataResponse(attributes IncidentGoogleMeetConfigurationDataAttributesResponse, id uuid.UUID, typeVar IncidentGoogleMeetConfigurationType) *IncidentGoogleMeetConfigurationDataResponse {
	this := IncidentGoogleMeetConfigurationDataResponse{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewIncidentGoogleMeetConfigurationDataResponseWithDefaults instantiates a new IncidentGoogleMeetConfigurationDataResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentGoogleMeetConfigurationDataResponseWithDefaults() *IncidentGoogleMeetConfigurationDataResponse {
	this := IncidentGoogleMeetConfigurationDataResponse{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *IncidentGoogleMeetConfigurationDataResponse) GetAttributes() IncidentGoogleMeetConfigurationDataAttributesResponse {
	if o == nil {
		var ret IncidentGoogleMeetConfigurationDataAttributesResponse
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *IncidentGoogleMeetConfigurationDataResponse) GetAttributesOk() (*IncidentGoogleMeetConfigurationDataAttributesResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *IncidentGoogleMeetConfigurationDataResponse) SetAttributes(v IncidentGoogleMeetConfigurationDataAttributesResponse) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *IncidentGoogleMeetConfigurationDataResponse) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IncidentGoogleMeetConfigurationDataResponse) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *IncidentGoogleMeetConfigurationDataResponse) SetId(v uuid.UUID) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *IncidentGoogleMeetConfigurationDataResponse) GetRelationships() IncidentGoogleMeetConfigurationRelationships {
	if o == nil || o.Relationships == nil {
		var ret IncidentGoogleMeetConfigurationRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentGoogleMeetConfigurationDataResponse) GetRelationshipsOk() (*IncidentGoogleMeetConfigurationRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *IncidentGoogleMeetConfigurationDataResponse) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given IncidentGoogleMeetConfigurationRelationships and assigns it to the Relationships field.
func (o *IncidentGoogleMeetConfigurationDataResponse) SetRelationships(v IncidentGoogleMeetConfigurationRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *IncidentGoogleMeetConfigurationDataResponse) GetType() IncidentGoogleMeetConfigurationType {
	if o == nil {
		var ret IncidentGoogleMeetConfigurationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentGoogleMeetConfigurationDataResponse) GetTypeOk() (*IncidentGoogleMeetConfigurationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentGoogleMeetConfigurationDataResponse) SetType(v IncidentGoogleMeetConfigurationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentGoogleMeetConfigurationDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentGoogleMeetConfigurationDataResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *IncidentGoogleMeetConfigurationDataAttributesResponse `json:"attributes"`
		Id            *uuid.UUID                                             `json:"id"`
		Relationships *IncidentGoogleMeetConfigurationRelationships          `json:"relationships,omitempty"`
		Type          *IncidentGoogleMeetConfigurationType                   `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = all.Relationships
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
