// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksIntegrationAccountPrivateActionRunnerAuthResponse The Private Action Runner authentication method configured on the account.
type DatabricksIntegrationAccountPrivateActionRunnerAuthResponse struct {
	// The authentication method type.
	AuthType DatabricksIntegrationAccountPrivateActionRunnerAuthType `json:"auth_type"`
	// Unique identifier of the Private Action Runner connection holding the credentials.
	ConnectionId string `json:"connection_id"`
	// Path of the credential inside the secret backend configured on the runner.
	SecretPath *string `json:"secret_path,omitempty"`
	// Unique identifier of the user the Private Action Runner connection belongs to.
	UserUuid string `json:"user_uuid"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatabricksIntegrationAccountPrivateActionRunnerAuthResponse instantiates a new DatabricksIntegrationAccountPrivateActionRunnerAuthResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabricksIntegrationAccountPrivateActionRunnerAuthResponse(authType DatabricksIntegrationAccountPrivateActionRunnerAuthType, connectionId string, userUuid string) *DatabricksIntegrationAccountPrivateActionRunnerAuthResponse {
	this := DatabricksIntegrationAccountPrivateActionRunnerAuthResponse{}
	this.AuthType = authType
	this.ConnectionId = connectionId
	this.UserUuid = userUuid
	return &this
}

// NewDatabricksIntegrationAccountPrivateActionRunnerAuthResponseWithDefaults instantiates a new DatabricksIntegrationAccountPrivateActionRunnerAuthResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabricksIntegrationAccountPrivateActionRunnerAuthResponseWithDefaults() *DatabricksIntegrationAccountPrivateActionRunnerAuthResponse {
	this := DatabricksIntegrationAccountPrivateActionRunnerAuthResponse{}
	var authType DatabricksIntegrationAccountPrivateActionRunnerAuthType = DATABRICKSINTEGRATIONACCOUNTPRIVATEACTIONRUNNERAUTHTYPE_PRIVATE_ACTION_RUNNER
	this.AuthType = authType
	return &this
}

// GetAuthType returns the AuthType field value.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthResponse) GetAuthType() DatabricksIntegrationAccountPrivateActionRunnerAuthType {
	if o == nil {
		var ret DatabricksIntegrationAccountPrivateActionRunnerAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthResponse) GetAuthTypeOk() (*DatabricksIntegrationAccountPrivateActionRunnerAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthResponse) SetAuthType(v DatabricksIntegrationAccountPrivateActionRunnerAuthType) {
	o.AuthType = v
}

// GetConnectionId returns the ConnectionId field value.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthResponse) GetConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthResponse) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthResponse) SetConnectionId(v string) {
	o.ConnectionId = v
}

// GetSecretPath returns the SecretPath field value if set, zero value otherwise.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthResponse) GetSecretPath() string {
	if o == nil || o.SecretPath == nil {
		var ret string
		return ret
	}
	return *o.SecretPath
}

// GetSecretPathOk returns a tuple with the SecretPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthResponse) GetSecretPathOk() (*string, bool) {
	if o == nil || o.SecretPath == nil {
		return nil, false
	}
	return o.SecretPath, true
}

// HasSecretPath returns a boolean if a field has been set.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthResponse) HasSecretPath() bool {
	return o != nil && o.SecretPath != nil
}

// SetSecretPath gets a reference to the given string and assigns it to the SecretPath field.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthResponse) SetSecretPath(v string) {
	o.SecretPath = &v
}

// GetUserUuid returns the UserUuid field value.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthResponse) GetUserUuid() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserUuid
}

// GetUserUuidOk returns a tuple with the UserUuid field value
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthResponse) GetUserUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserUuid, true
}

// SetUserUuid sets field value.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthResponse) SetUserUuid(v string) {
	o.UserUuid = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabricksIntegrationAccountPrivateActionRunnerAuthResponse) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatabricksIntegrationAccountPrivateActionRunnerAuthResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthType     *DatabricksIntegrationAccountPrivateActionRunnerAuthType `json:"auth_type"`
		ConnectionId *string                                                  `json:"connection_id"`
		SecretPath   *string                                                  `json:"secret_path,omitempty"`
		UserUuid     *string                                                  `json:"user_uuid"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth_type", "connection_id", "secret_path", "user_uuid"})
	} else {
		return err
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

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
