// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationGoogleSecurityOperationsDestinationAuth Google Security Operations destination authentication.
type CustomDestinationGoogleSecurityOperationsDestinationAuth struct {
	// The Google Security Operations client email.
	ClientEmail string `json:"client_email"`
	// The Google Security Operations client ID. This field is not returned by the API.
	ClientId string `json:"client_id"`
	// The Google Security Operations private key. This field is not returned by the API.
	PrivateKey string `json:"private_key"`
	// The Google Security Operations private key ID. This field is not returned by the API.
	PrivateKeyId string `json:"private_key_id"`
	// Google Security Operations project ID.
	ProjectId string `json:"project_id"`
	// Type of the Google Security Operations destination authentication.
	Type CustomDestinationGoogleSecurityOperationsDestinationAuthType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomDestinationGoogleSecurityOperationsDestinationAuth instantiates a new CustomDestinationGoogleSecurityOperationsDestinationAuth object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomDestinationGoogleSecurityOperationsDestinationAuth(clientEmail string, clientId string, privateKey string, privateKeyId string, projectId string, typeVar CustomDestinationGoogleSecurityOperationsDestinationAuthType) *CustomDestinationGoogleSecurityOperationsDestinationAuth {
	this := CustomDestinationGoogleSecurityOperationsDestinationAuth{}
	this.ClientEmail = clientEmail
	this.ClientId = clientId
	this.PrivateKey = privateKey
	this.PrivateKeyId = privateKeyId
	this.ProjectId = projectId
	this.Type = typeVar
	return &this
}

// NewCustomDestinationGoogleSecurityOperationsDestinationAuthWithDefaults instantiates a new CustomDestinationGoogleSecurityOperationsDestinationAuth object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomDestinationGoogleSecurityOperationsDestinationAuthWithDefaults() *CustomDestinationGoogleSecurityOperationsDestinationAuth {
	this := CustomDestinationGoogleSecurityOperationsDestinationAuth{}
	var typeVar CustomDestinationGoogleSecurityOperationsDestinationAuthType = CUSTOMDESTINATIONGOOGLESECURITYOPERATIONSDESTINATIONAUTHTYPE_GCP_PRIVATE_KEY
	this.Type = typeVar
	return &this
}

// GetClientEmail returns the ClientEmail field value.
func (o *CustomDestinationGoogleSecurityOperationsDestinationAuth) GetClientEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientEmail
}

// GetClientEmailOk returns a tuple with the ClientEmail field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationGoogleSecurityOperationsDestinationAuth) GetClientEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientEmail, true
}

// SetClientEmail sets field value.
func (o *CustomDestinationGoogleSecurityOperationsDestinationAuth) SetClientEmail(v string) {
	o.ClientEmail = v
}

// GetClientId returns the ClientId field value.
func (o *CustomDestinationGoogleSecurityOperationsDestinationAuth) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationGoogleSecurityOperationsDestinationAuth) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value.
func (o *CustomDestinationGoogleSecurityOperationsDestinationAuth) SetClientId(v string) {
	o.ClientId = v
}

// GetPrivateKey returns the PrivateKey field value.
func (o *CustomDestinationGoogleSecurityOperationsDestinationAuth) GetPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationGoogleSecurityOperationsDestinationAuth) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value.
func (o *CustomDestinationGoogleSecurityOperationsDestinationAuth) SetPrivateKey(v string) {
	o.PrivateKey = v
}

// GetPrivateKeyId returns the PrivateKeyId field value.
func (o *CustomDestinationGoogleSecurityOperationsDestinationAuth) GetPrivateKeyId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PrivateKeyId
}

// GetPrivateKeyIdOk returns a tuple with the PrivateKeyId field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationGoogleSecurityOperationsDestinationAuth) GetPrivateKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKeyId, true
}

// SetPrivateKeyId sets field value.
func (o *CustomDestinationGoogleSecurityOperationsDestinationAuth) SetPrivateKeyId(v string) {
	o.PrivateKeyId = v
}

// GetProjectId returns the ProjectId field value.
func (o *CustomDestinationGoogleSecurityOperationsDestinationAuth) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationGoogleSecurityOperationsDestinationAuth) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value.
func (o *CustomDestinationGoogleSecurityOperationsDestinationAuth) SetProjectId(v string) {
	o.ProjectId = v
}

// GetType returns the Type field value.
func (o *CustomDestinationGoogleSecurityOperationsDestinationAuth) GetType() CustomDestinationGoogleSecurityOperationsDestinationAuthType {
	if o == nil {
		var ret CustomDestinationGoogleSecurityOperationsDestinationAuthType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationGoogleSecurityOperationsDestinationAuth) GetTypeOk() (*CustomDestinationGoogleSecurityOperationsDestinationAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CustomDestinationGoogleSecurityOperationsDestinationAuth) SetType(v CustomDestinationGoogleSecurityOperationsDestinationAuthType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomDestinationGoogleSecurityOperationsDestinationAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["client_email"] = o.ClientEmail
	toSerialize["client_id"] = o.ClientId
	toSerialize["private_key"] = o.PrivateKey
	toSerialize["private_key_id"] = o.PrivateKeyId
	toSerialize["project_id"] = o.ProjectId
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomDestinationGoogleSecurityOperationsDestinationAuth) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClientEmail  *string                                                       `json:"client_email"`
		ClientId     *string                                                       `json:"client_id"`
		PrivateKey   *string                                                       `json:"private_key"`
		PrivateKeyId *string                                                       `json:"private_key_id"`
		ProjectId    *string                                                       `json:"project_id"`
		Type         *CustomDestinationGoogleSecurityOperationsDestinationAuthType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ClientEmail == nil {
		return fmt.Errorf("required field client_email missing")
	}
	if all.ClientId == nil {
		return fmt.Errorf("required field client_id missing")
	}
	if all.PrivateKey == nil {
		return fmt.Errorf("required field private_key missing")
	}
	if all.PrivateKeyId == nil {
		return fmt.Errorf("required field private_key_id missing")
	}
	if all.ProjectId == nil {
		return fmt.Errorf("required field project_id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"client_email", "client_id", "private_key", "private_key_id", "project_id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ClientEmail = *all.ClientEmail
	o.ClientId = *all.ClientId
	o.PrivateKey = *all.PrivateKey
	o.PrivateKeyId = *all.PrivateKeyId
	o.ProjectId = *all.ProjectId
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
