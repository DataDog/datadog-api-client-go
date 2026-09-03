// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksIntegrationAccountPatAuthUpdate Databricks personal access token authentication. Deprecated: accepted only on accounts that already use it, and never on creation. Use `databricks-oauth` or `private-action-runner` instead. Omit `token` to keep the stored one.
//
// Deprecated: This model is deprecated.
type DatabricksIntegrationAccountPatAuthUpdate struct {
	// The authentication method type.
	AuthType DatabricksIntegrationAccountPatAuthType `json:"auth_type"`
	// Secret Databricks personal access token.
	Token *string `json:"token,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDatabricksIntegrationAccountPatAuthUpdate instantiates a new DatabricksIntegrationAccountPatAuthUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabricksIntegrationAccountPatAuthUpdate(authType DatabricksIntegrationAccountPatAuthType) *DatabricksIntegrationAccountPatAuthUpdate {
	this := DatabricksIntegrationAccountPatAuthUpdate{}
	this.AuthType = authType
	return &this
}

// NewDatabricksIntegrationAccountPatAuthUpdateWithDefaults instantiates a new DatabricksIntegrationAccountPatAuthUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabricksIntegrationAccountPatAuthUpdateWithDefaults() *DatabricksIntegrationAccountPatAuthUpdate {
	this := DatabricksIntegrationAccountPatAuthUpdate{}
	var authType DatabricksIntegrationAccountPatAuthType = DATABRICKSINTEGRATIONACCOUNTPATAUTHTYPE_PAT
	this.AuthType = authType
	return &this
}

// GetAuthType returns the AuthType field value.
func (o *DatabricksIntegrationAccountPatAuthUpdate) GetAuthType() DatabricksIntegrationAccountPatAuthType {
	if o == nil {
		var ret DatabricksIntegrationAccountPatAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationAccountPatAuthUpdate) GetAuthTypeOk() (*DatabricksIntegrationAccountPatAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *DatabricksIntegrationAccountPatAuthUpdate) SetAuthType(v DatabricksIntegrationAccountPatAuthType) {
	o.AuthType = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DatabricksIntegrationAccountPatAuthUpdate) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationAccountPatAuthUpdate) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DatabricksIntegrationAccountPatAuthUpdate) HasToken() bool {
	return o != nil && o.Token != nil
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DatabricksIntegrationAccountPatAuthUpdate) SetToken(v string) {
	o.Token = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabricksIntegrationAccountPatAuthUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth_type"] = o.AuthType
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatabricksIntegrationAccountPatAuthUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthType *DatabricksIntegrationAccountPatAuthType `json:"auth_type"`
		Token    *string                                  `json:"token,omitempty"`
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
	o.Token = all.Token

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
