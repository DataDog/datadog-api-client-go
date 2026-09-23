// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate Private Action Runner authentication. The runner holds the Databricks credentials, so this method carries no secrets. Only the fields provided are changed.
type DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate struct {
	// The authentication method type.
	AuthType DatabricksIntegrationAccountPrivateActionRunnerAuthType `json:"auth_type"`
	// Unique identifier of the Private Action Runner connection holding the credentials.
	ConnectionId *uuid.UUID `json:"connection_id,omitempty"`
	// Path of the credential inside the secret backend configured on the runner. Omit it to keep the stored path, send `null` or an empty string to remove it, or send a value to replace it.
	SecretPath datadog.NullableString `json:"secret_path,omitempty"`
	// Unique identifier of the user the Private Action Runner connection belongs to.
	UserUuid *uuid.UUID `json:"user_uuid,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDatabricksIntegrationAccountPrivateActionRunnerAuthUpdate instantiates a new DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabricksIntegrationAccountPrivateActionRunnerAuthUpdate(authType DatabricksIntegrationAccountPrivateActionRunnerAuthType) *DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate {
	this := DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate{}
	this.AuthType = authType
	return &this
}

// NewDatabricksIntegrationAccountPrivateActionRunnerAuthUpdateWithDefaults instantiates a new DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabricksIntegrationAccountPrivateActionRunnerAuthUpdateWithDefaults() *DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate {
	this := DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate{}
	var authType DatabricksIntegrationAccountPrivateActionRunnerAuthType = DATABRICKSINTEGRATIONACCOUNTPRIVATEACTIONRUNNERAUTHTYPE_PRIVATE_ACTION_RUNNER
	this.AuthType = authType
	return &this
}

// GetAuthType returns the AuthType field value.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate) GetAuthType() DatabricksIntegrationAccountPrivateActionRunnerAuthType {
	if o == nil {
		var ret DatabricksIntegrationAccountPrivateActionRunnerAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate) GetAuthTypeOk() (*DatabricksIntegrationAccountPrivateActionRunnerAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate) SetAuthType(v DatabricksIntegrationAccountPrivateActionRunnerAuthType) {
	o.AuthType = v
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate) GetConnectionId() uuid.UUID {
	if o == nil || o.ConnectionId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate) GetConnectionIdOk() (*uuid.UUID, bool) {
	if o == nil || o.ConnectionId == nil {
		return nil, false
	}
	return o.ConnectionId, true
}

// HasConnectionId returns a boolean if a field has been set.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate) HasConnectionId() bool {
	return o != nil && o.ConnectionId != nil
}

// SetConnectionId gets a reference to the given uuid.UUID and assigns it to the ConnectionId field.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate) SetConnectionId(v uuid.UUID) {
	o.ConnectionId = &v
}

// GetSecretPath returns the SecretPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate) GetSecretPath() string {
	if o == nil || o.SecretPath.Get() == nil {
		var ret string
		return ret
	}
	return *o.SecretPath.Get()
}

// GetSecretPathOk returns a tuple with the SecretPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate) GetSecretPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretPath.Get(), o.SecretPath.IsSet()
}

// HasSecretPath returns a boolean if a field has been set.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate) HasSecretPath() bool {
	return o != nil && o.SecretPath.IsSet()
}

// SetSecretPath gets a reference to the given datadog.NullableString and assigns it to the SecretPath field.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate) SetSecretPath(v string) {
	o.SecretPath.Set(&v)
}

// SetSecretPathNil sets the value for SecretPath to be an explicit nil.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate) SetSecretPathNil() {
	o.SecretPath.Set(nil)
}

// UnsetSecretPath ensures that no value is present for SecretPath, not even an explicit nil.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate) UnsetSecretPath() {
	o.SecretPath.Unset()
}

// GetUserUuid returns the UserUuid field value if set, zero value otherwise.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate) GetUserUuid() uuid.UUID {
	if o == nil || o.UserUuid == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.UserUuid
}

// GetUserUuidOk returns a tuple with the UserUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate) GetUserUuidOk() (*uuid.UUID, bool) {
	if o == nil || o.UserUuid == nil {
		return nil, false
	}
	return o.UserUuid, true
}

// HasUserUuid returns a boolean if a field has been set.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate) HasUserUuid() bool {
	return o != nil && o.UserUuid != nil
}

// SetUserUuid gets a reference to the given uuid.UUID and assigns it to the UserUuid field.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate) SetUserUuid(v uuid.UUID) {
	o.UserUuid = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth_type"] = o.AuthType
	if o.ConnectionId != nil {
		toSerialize["connection_id"] = o.ConnectionId
	}
	if o.SecretPath.IsSet() {
		toSerialize["secret_path"] = o.SecretPath.Get()
	}
	if o.UserUuid != nil {
		toSerialize["user_uuid"] = o.UserUuid
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthType     *DatabricksIntegrationAccountPrivateActionRunnerAuthType `json:"auth_type"`
		ConnectionId *uuid.UUID                                               `json:"connection_id,omitempty"`
		SecretPath   datadog.NullableString                                   `json:"secret_path,omitempty"`
		UserUuid     *uuid.UUID                                               `json:"user_uuid,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuthType == nil {
		return fmt.Errorf("required field auth_type missing")
	}

	hasInvalidField := false
	if !all.AuthType.IsValid() {
		hasInvalidField = true
	} else {
		o.AuthType = *all.AuthType
	}
	o.ConnectionId = all.ConnectionId
	o.SecretPath = all.SecretPath
	o.UserUuid = all.UserUuid

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
