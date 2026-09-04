// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationAccountPrivateActionRunnerAuthRequest Private Action Runner authentication. The runner holds the credentials, so this method carries no secrets; `connection_id` and `user_uuid` must be provided on every submission.
type IntegrationAccountPrivateActionRunnerAuthRequest struct {
	// The authentication method type.
	AuthType IntegrationAccountPrivateActionRunnerAuthType `json:"auth_type"`
	// Unique identifier of the Private Action Runner connection holding the credentials.
	ConnectionId string `json:"connection_id"`
	// Path of the credential inside the secret backend configured on the runner.
	SecretPath *string `json:"secret_path,omitempty"`
	// Unique identifier of the user the Private Action Runner connection belongs to.
	UserUuid string `json:"user_uuid"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewIntegrationAccountPrivateActionRunnerAuthRequest instantiates a new IntegrationAccountPrivateActionRunnerAuthRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegrationAccountPrivateActionRunnerAuthRequest(authType IntegrationAccountPrivateActionRunnerAuthType, connectionId string, userUuid string) *IntegrationAccountPrivateActionRunnerAuthRequest {
	this := IntegrationAccountPrivateActionRunnerAuthRequest{}
	this.AuthType = authType
	this.ConnectionId = connectionId
	this.UserUuid = userUuid
	return &this
}

// NewIntegrationAccountPrivateActionRunnerAuthRequestWithDefaults instantiates a new IntegrationAccountPrivateActionRunnerAuthRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegrationAccountPrivateActionRunnerAuthRequestWithDefaults() *IntegrationAccountPrivateActionRunnerAuthRequest {
	this := IntegrationAccountPrivateActionRunnerAuthRequest{}
	var authType IntegrationAccountPrivateActionRunnerAuthType = INTEGRATIONACCOUNTPRIVATEACTIONRUNNERAUTHTYPE_PRIVATE_ACTION_RUNNER
	this.AuthType = authType
	return &this
}

// GetAuthType returns the AuthType field value.
func (o *IntegrationAccountPrivateActionRunnerAuthRequest) GetAuthType() IntegrationAccountPrivateActionRunnerAuthType {
	if o == nil {
		var ret IntegrationAccountPrivateActionRunnerAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *IntegrationAccountPrivateActionRunnerAuthRequest) GetAuthTypeOk() (*IntegrationAccountPrivateActionRunnerAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *IntegrationAccountPrivateActionRunnerAuthRequest) SetAuthType(v IntegrationAccountPrivateActionRunnerAuthType) {
	o.AuthType = v
}

// GetConnectionId returns the ConnectionId field value.
func (o *IntegrationAccountPrivateActionRunnerAuthRequest) GetConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *IntegrationAccountPrivateActionRunnerAuthRequest) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value.
func (o *IntegrationAccountPrivateActionRunnerAuthRequest) SetConnectionId(v string) {
	o.ConnectionId = v
}

// GetSecretPath returns the SecretPath field value if set, zero value otherwise.
func (o *IntegrationAccountPrivateActionRunnerAuthRequest) GetSecretPath() string {
	if o == nil || o.SecretPath == nil {
		var ret string
		return ret
	}
	return *o.SecretPath
}

// GetSecretPathOk returns a tuple with the SecretPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationAccountPrivateActionRunnerAuthRequest) GetSecretPathOk() (*string, bool) {
	if o == nil || o.SecretPath == nil {
		return nil, false
	}
	return o.SecretPath, true
}

// HasSecretPath returns a boolean if a field has been set.
func (o *IntegrationAccountPrivateActionRunnerAuthRequest) HasSecretPath() bool {
	return o != nil && o.SecretPath != nil
}

// SetSecretPath gets a reference to the given string and assigns it to the SecretPath field.
func (o *IntegrationAccountPrivateActionRunnerAuthRequest) SetSecretPath(v string) {
	o.SecretPath = &v
}

// GetUserUuid returns the UserUuid field value.
func (o *IntegrationAccountPrivateActionRunnerAuthRequest) GetUserUuid() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserUuid
}

// GetUserUuidOk returns a tuple with the UserUuid field value
// and a boolean to check if the value has been set.
func (o *IntegrationAccountPrivateActionRunnerAuthRequest) GetUserUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserUuid, true
}

// SetUserUuid sets field value.
func (o *IntegrationAccountPrivateActionRunnerAuthRequest) SetUserUuid(v string) {
	o.UserUuid = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IntegrationAccountPrivateActionRunnerAuthRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth_type"] = o.AuthType
	toSerialize["connection_id"] = o.ConnectionId
	if o.SecretPath != nil {
		toSerialize["secret_path"] = o.SecretPath
	}
	toSerialize["user_uuid"] = o.UserUuid
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IntegrationAccountPrivateActionRunnerAuthRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthType     *IntegrationAccountPrivateActionRunnerAuthType `json:"auth_type"`
		ConnectionId *string                                        `json:"connection_id"`
		SecretPath   *string                                        `json:"secret_path,omitempty"`
		UserUuid     *string                                        `json:"user_uuid"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuthType == nil {
		return fmt.Errorf("required field auth_type missing")
	}
	if all.ConnectionId == nil {
		return fmt.Errorf("required field connection_id missing")
	}
	if all.UserUuid == nil {
		return fmt.Errorf("required field user_uuid missing")
	}

	hasInvalidField := false
	if !all.AuthType.IsValid() {
		hasInvalidField = true
	} else {
		o.AuthType = *all.AuthType
	}
	o.ConnectionId = *all.ConnectionId
	o.SecretPath = all.SecretPath
	o.UserUuid = *all.UserUuid

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
