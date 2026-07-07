// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgAuthorizedClientRelationships Relationships for an org authorized client.
type OrgAuthorizedClientRelationships struct {
	// Relationship to the OAuth2 client for this org authorized client.
	Oauth2Client OrgAuthorizedClientRelationshipOAuth2Client `json:"oauth2_client"`
	// Relationship to the user authorized clients for this org authorized client.
	UserAuthorizedClients OrgAuthorizedClientRelationshipUserAuthorizedClients `json:"user_authorized_clients"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgAuthorizedClientRelationships instantiates a new OrgAuthorizedClientRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgAuthorizedClientRelationships(oauth2Client OrgAuthorizedClientRelationshipOAuth2Client, userAuthorizedClients OrgAuthorizedClientRelationshipUserAuthorizedClients) *OrgAuthorizedClientRelationships {
	this := OrgAuthorizedClientRelationships{}
	this.Oauth2Client = oauth2Client
	this.UserAuthorizedClients = userAuthorizedClients
	return &this
}

// NewOrgAuthorizedClientRelationshipsWithDefaults instantiates a new OrgAuthorizedClientRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgAuthorizedClientRelationshipsWithDefaults() *OrgAuthorizedClientRelationships {
	this := OrgAuthorizedClientRelationships{}
	return &this
}

// GetOauth2Client returns the Oauth2Client field value.
func (o *OrgAuthorizedClientRelationships) GetOauth2Client() OrgAuthorizedClientRelationshipOAuth2Client {
	if o == nil {
		var ret OrgAuthorizedClientRelationshipOAuth2Client
		return ret
	}
	return o.Oauth2Client
}

// GetOauth2ClientOk returns a tuple with the Oauth2Client field value
// and a boolean to check if the value has been set.
func (o *OrgAuthorizedClientRelationships) GetOauth2ClientOk() (*OrgAuthorizedClientRelationshipOAuth2Client, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Oauth2Client, true
}

// SetOauth2Client sets field value.
func (o *OrgAuthorizedClientRelationships) SetOauth2Client(v OrgAuthorizedClientRelationshipOAuth2Client) {
	o.Oauth2Client = v
}

// GetUserAuthorizedClients returns the UserAuthorizedClients field value.
func (o *OrgAuthorizedClientRelationships) GetUserAuthorizedClients() OrgAuthorizedClientRelationshipUserAuthorizedClients {
	if o == nil {
		var ret OrgAuthorizedClientRelationshipUserAuthorizedClients
		return ret
	}
	return o.UserAuthorizedClients
}

// GetUserAuthorizedClientsOk returns a tuple with the UserAuthorizedClients field value
// and a boolean to check if the value has been set.
func (o *OrgAuthorizedClientRelationships) GetUserAuthorizedClientsOk() (*OrgAuthorizedClientRelationshipUserAuthorizedClients, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserAuthorizedClients, true
}

// SetUserAuthorizedClients sets field value.
func (o *OrgAuthorizedClientRelationships) SetUserAuthorizedClients(v OrgAuthorizedClientRelationshipUserAuthorizedClients) {
	o.UserAuthorizedClients = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgAuthorizedClientRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["oauth2_client"] = o.Oauth2Client
	toSerialize["user_authorized_clients"] = o.UserAuthorizedClients

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgAuthorizedClientRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Oauth2Client          *OrgAuthorizedClientRelationshipOAuth2Client          `json:"oauth2_client"`
		UserAuthorizedClients *OrgAuthorizedClientRelationshipUserAuthorizedClients `json:"user_authorized_clients"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Oauth2Client == nil {
		return fmt.Errorf("required field oauth2_client missing")
	}
	if all.UserAuthorizedClients == nil {
		return fmt.Errorf("required field user_authorized_clients missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"oauth2_client", "user_authorized_clients"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Oauth2Client.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Oauth2Client = *all.Oauth2Client
	if all.UserAuthorizedClients.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UserAuthorizedClients = *all.UserAuthorizedClients

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
