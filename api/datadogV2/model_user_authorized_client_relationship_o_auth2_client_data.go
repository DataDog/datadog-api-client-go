// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserAuthorizedClientRelationshipOAuth2ClientData Data identifying the OAuth2 client that was authorized.
type UserAuthorizedClientRelationshipOAuth2ClientData struct {
	// The ID of the OAuth2 client.
	Id string `json:"id"`
	// OAuth2 client resource type.
	Type UserAuthorizedClientRelationshipOAuth2ClientDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUserAuthorizedClientRelationshipOAuth2ClientData instantiates a new UserAuthorizedClientRelationshipOAuth2ClientData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserAuthorizedClientRelationshipOAuth2ClientData(id string, typeVar UserAuthorizedClientRelationshipOAuth2ClientDataType) *UserAuthorizedClientRelationshipOAuth2ClientData {
	this := UserAuthorizedClientRelationshipOAuth2ClientData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewUserAuthorizedClientRelationshipOAuth2ClientDataWithDefaults instantiates a new UserAuthorizedClientRelationshipOAuth2ClientData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserAuthorizedClientRelationshipOAuth2ClientDataWithDefaults() *UserAuthorizedClientRelationshipOAuth2ClientData {
	this := UserAuthorizedClientRelationshipOAuth2ClientData{}
	return &this
}

// GetId returns the Id field value.
func (o *UserAuthorizedClientRelationshipOAuth2ClientData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserAuthorizedClientRelationshipOAuth2ClientData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *UserAuthorizedClientRelationshipOAuth2ClientData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *UserAuthorizedClientRelationshipOAuth2ClientData) GetType() UserAuthorizedClientRelationshipOAuth2ClientDataType {
	if o == nil {
		var ret UserAuthorizedClientRelationshipOAuth2ClientDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserAuthorizedClientRelationshipOAuth2ClientData) GetTypeOk() (*UserAuthorizedClientRelationshipOAuth2ClientDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *UserAuthorizedClientRelationshipOAuth2ClientData) SetType(v UserAuthorizedClientRelationshipOAuth2ClientDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UserAuthorizedClientRelationshipOAuth2ClientData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserAuthorizedClientRelationshipOAuth2ClientData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string                                               `json:"id"`
		Type *UserAuthorizedClientRelationshipOAuth2ClientDataType `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
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
