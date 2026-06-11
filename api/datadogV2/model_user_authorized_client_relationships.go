// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserAuthorizedClientRelationships Relationships for a user authorized client.
type UserAuthorizedClientRelationships struct {
	// Relationship to the OAuth2 client that was authorized.
	Oauth2Client UserAuthorizedClientRelationshipOAuth2Client `json:"oauth2_client"`
	// Relationship to the scopes granted to the OAuth2 client.
	Scopes UserAuthorizedClientRelationshipScopes `json:"scopes"`
	// Relationship to the user who granted this authorization.
	User UserAuthorizedClientRelationshipUser `json:"user"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUserAuthorizedClientRelationships instantiates a new UserAuthorizedClientRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserAuthorizedClientRelationships(oauth2Client UserAuthorizedClientRelationshipOAuth2Client, scopes UserAuthorizedClientRelationshipScopes, user UserAuthorizedClientRelationshipUser) *UserAuthorizedClientRelationships {
	this := UserAuthorizedClientRelationships{}
	this.Oauth2Client = oauth2Client
	this.Scopes = scopes
	this.User = user
	return &this
}

// NewUserAuthorizedClientRelationshipsWithDefaults instantiates a new UserAuthorizedClientRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserAuthorizedClientRelationshipsWithDefaults() *UserAuthorizedClientRelationships {
	this := UserAuthorizedClientRelationships{}
	return &this
}

// GetOauth2Client returns the Oauth2Client field value.
func (o *UserAuthorizedClientRelationships) GetOauth2Client() UserAuthorizedClientRelationshipOAuth2Client {
	if o == nil {
		var ret UserAuthorizedClientRelationshipOAuth2Client
		return ret
	}
	return o.Oauth2Client
}

// GetOauth2ClientOk returns a tuple with the Oauth2Client field value
// and a boolean to check if the value has been set.
func (o *UserAuthorizedClientRelationships) GetOauth2ClientOk() (*UserAuthorizedClientRelationshipOAuth2Client, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Oauth2Client, true
}

// SetOauth2Client sets field value.
func (o *UserAuthorizedClientRelationships) SetOauth2Client(v UserAuthorizedClientRelationshipOAuth2Client) {
	o.Oauth2Client = v
}

// GetScopes returns the Scopes field value.
func (o *UserAuthorizedClientRelationships) GetScopes() UserAuthorizedClientRelationshipScopes {
	if o == nil {
		var ret UserAuthorizedClientRelationshipScopes
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *UserAuthorizedClientRelationships) GetScopesOk() (*UserAuthorizedClientRelationshipScopes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scopes, true
}

// SetScopes sets field value.
func (o *UserAuthorizedClientRelationships) SetScopes(v UserAuthorizedClientRelationshipScopes) {
	o.Scopes = v
}

// GetUser returns the User field value.
func (o *UserAuthorizedClientRelationships) GetUser() UserAuthorizedClientRelationshipUser {
	if o == nil {
		var ret UserAuthorizedClientRelationshipUser
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *UserAuthorizedClientRelationships) GetUserOk() (*UserAuthorizedClientRelationshipUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value.
func (o *UserAuthorizedClientRelationships) SetUser(v UserAuthorizedClientRelationshipUser) {
	o.User = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UserAuthorizedClientRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["oauth2_client"] = o.Oauth2Client
	toSerialize["scopes"] = o.Scopes
	toSerialize["user"] = o.User

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserAuthorizedClientRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Oauth2Client *UserAuthorizedClientRelationshipOAuth2Client `json:"oauth2_client"`
		Scopes       *UserAuthorizedClientRelationshipScopes       `json:"scopes"`
		User         *UserAuthorizedClientRelationshipUser         `json:"user"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Oauth2Client == nil {
		return fmt.Errorf("required field oauth2_client missing")
	}
	if all.Scopes == nil {
		return fmt.Errorf("required field scopes missing")
	}
	if all.User == nil {
		return fmt.Errorf("required field user missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"oauth2_client", "scopes", "user"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Oauth2Client.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Oauth2Client = *all.Oauth2Client
	if all.Scopes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Scopes = *all.Scopes
	if all.User.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.User = *all.User

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
