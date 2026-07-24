// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PostmortemTemplateResponseRelationships Relationships of a postmortem template returned in a response.
type PostmortemTemplateResponseRelationships struct {
	// Relationship to the incident type this template belongs to.
	IncidentType *PostmortemTemplateIncidentTypeRelationship `json:"incident_type,omitempty"`
	// Relationship to a user.
	LastModifiedByUser *PostmortemTemplateUserRelationship `json:"last_modified_by_user,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostmortemTemplateResponseRelationships instantiates a new PostmortemTemplateResponseRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostmortemTemplateResponseRelationships() *PostmortemTemplateResponseRelationships {
	this := PostmortemTemplateResponseRelationships{}
	return &this
}

// NewPostmortemTemplateResponseRelationshipsWithDefaults instantiates a new PostmortemTemplateResponseRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostmortemTemplateResponseRelationshipsWithDefaults() *PostmortemTemplateResponseRelationships {
	this := PostmortemTemplateResponseRelationships{}
	return &this
}

// GetIncidentType returns the IncidentType field value if set, zero value otherwise.
func (o *PostmortemTemplateResponseRelationships) GetIncidentType() PostmortemTemplateIncidentTypeRelationship {
	if o == nil || o.IncidentType == nil {
		var ret PostmortemTemplateIncidentTypeRelationship
		return ret
	}
	return *o.IncidentType
}

// GetIncidentTypeOk returns a tuple with the IncidentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostmortemTemplateResponseRelationships) GetIncidentTypeOk() (*PostmortemTemplateIncidentTypeRelationship, bool) {
	if o == nil || o.IncidentType == nil {
		return nil, false
	}
	return o.IncidentType, true
}

// HasIncidentType returns a boolean if a field has been set.
func (o *PostmortemTemplateResponseRelationships) HasIncidentType() bool {
	return o != nil && o.IncidentType != nil
}

// SetIncidentType gets a reference to the given PostmortemTemplateIncidentTypeRelationship and assigns it to the IncidentType field.
func (o *PostmortemTemplateResponseRelationships) SetIncidentType(v PostmortemTemplateIncidentTypeRelationship) {
	o.IncidentType = &v
}

// GetLastModifiedByUser returns the LastModifiedByUser field value if set, zero value otherwise.
func (o *PostmortemTemplateResponseRelationships) GetLastModifiedByUser() PostmortemTemplateUserRelationship {
	if o == nil || o.LastModifiedByUser == nil {
		var ret PostmortemTemplateUserRelationship
		return ret
	}
	return *o.LastModifiedByUser
}

// GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostmortemTemplateResponseRelationships) GetLastModifiedByUserOk() (*PostmortemTemplateUserRelationship, bool) {
	if o == nil || o.LastModifiedByUser == nil {
		return nil, false
	}
	return o.LastModifiedByUser, true
}

// HasLastModifiedByUser returns a boolean if a field has been set.
func (o *PostmortemTemplateResponseRelationships) HasLastModifiedByUser() bool {
	return o != nil && o.LastModifiedByUser != nil
}

// SetLastModifiedByUser gets a reference to the given PostmortemTemplateUserRelationship and assigns it to the LastModifiedByUser field.
func (o *PostmortemTemplateResponseRelationships) SetLastModifiedByUser(v PostmortemTemplateUserRelationship) {
	o.LastModifiedByUser = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostmortemTemplateResponseRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IncidentType != nil {
		toSerialize["incident_type"] = o.IncidentType
	}
	if o.LastModifiedByUser != nil {
		toSerialize["last_modified_by_user"] = o.LastModifiedByUser
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostmortemTemplateResponseRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncidentType       *PostmortemTemplateIncidentTypeRelationship `json:"incident_type,omitempty"`
		LastModifiedByUser *PostmortemTemplateUserRelationship         `json:"last_modified_by_user,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"incident_type", "last_modified_by_user"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.IncidentType != nil && all.IncidentType.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.IncidentType = all.IncidentType
	if all.LastModifiedByUser != nil && all.LastModifiedByUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LastModifiedByUser = all.LastModifiedByUser

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
