// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdatedServiceAccessToken Datadog access token returned by the update endpoint.
type UpdatedServiceAccessToken struct {
	// Attributes of an access token.
	Attributes *ServiceAccessTokenAttributes `json:"attributes,omitempty"`
	// ID of the access token.
	Id string `json:"id"`
	// Resources related to the access token.
	Relationships *UpdatedServiceAccessTokenRelationships `json:"relationships,omitempty"`
	// Service access tokens resource type.
	Type ServiceAccessTokensType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdatedServiceAccessToken instantiates a new UpdatedServiceAccessToken object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdatedServiceAccessToken(id string, typeVar ServiceAccessTokensType) *UpdatedServiceAccessToken {
	this := UpdatedServiceAccessToken{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewUpdatedServiceAccessTokenWithDefaults instantiates a new UpdatedServiceAccessToken object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdatedServiceAccessTokenWithDefaults() *UpdatedServiceAccessToken {
	this := UpdatedServiceAccessToken{}
	var typeVar ServiceAccessTokensType = SERVICEACCESSTOKENSTYPE_SERVICE_ACCESS_TOKENS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UpdatedServiceAccessToken) GetAttributes() ServiceAccessTokenAttributes {
	if o == nil || o.Attributes == nil {
		var ret ServiceAccessTokenAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatedServiceAccessToken) GetAttributesOk() (*ServiceAccessTokenAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UpdatedServiceAccessToken) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given ServiceAccessTokenAttributes and assigns it to the Attributes field.
func (o *UpdatedServiceAccessToken) SetAttributes(v ServiceAccessTokenAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *UpdatedServiceAccessToken) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdatedServiceAccessToken) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *UpdatedServiceAccessToken) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *UpdatedServiceAccessToken) GetRelationships() UpdatedServiceAccessTokenRelationships {
	if o == nil || o.Relationships == nil {
		var ret UpdatedServiceAccessTokenRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatedServiceAccessToken) GetRelationshipsOk() (*UpdatedServiceAccessTokenRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *UpdatedServiceAccessToken) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given UpdatedServiceAccessTokenRelationships and assigns it to the Relationships field.
func (o *UpdatedServiceAccessToken) SetRelationships(v UpdatedServiceAccessTokenRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *UpdatedServiceAccessToken) GetType() ServiceAccessTokensType {
	if o == nil {
		var ret ServiceAccessTokensType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdatedServiceAccessToken) GetTypeOk() (*ServiceAccessTokensType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *UpdatedServiceAccessToken) SetType(v ServiceAccessTokensType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdatedServiceAccessToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
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
func (o *UpdatedServiceAccessToken) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *ServiceAccessTokenAttributes           `json:"attributes,omitempty"`
		Id            *string                                 `json:"id"`
		Relationships *UpdatedServiceAccessTokenRelationships `json:"relationships,omitempty"`
		Type          *ServiceAccessTokensType                `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
