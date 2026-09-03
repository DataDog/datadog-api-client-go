// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksIntegrationAccountPatAuthResponse The Databricks personal access token authentication method configured on the account. Deprecated: migrate these accounts to `databricks-oauth` or `private-action-runner`.
//
// Deprecated: This model is deprecated.
type DatabricksIntegrationAccountPatAuthResponse struct {
	// The authentication method type.
	AuthType DatabricksIntegrationAccountPatAuthType `json:"auth_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatabricksIntegrationAccountPatAuthResponse instantiates a new DatabricksIntegrationAccountPatAuthResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabricksIntegrationAccountPatAuthResponse(authType DatabricksIntegrationAccountPatAuthType) *DatabricksIntegrationAccountPatAuthResponse {
	this := DatabricksIntegrationAccountPatAuthResponse{}
	this.AuthType = authType
	return &this
}

// NewDatabricksIntegrationAccountPatAuthResponseWithDefaults instantiates a new DatabricksIntegrationAccountPatAuthResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabricksIntegrationAccountPatAuthResponseWithDefaults() *DatabricksIntegrationAccountPatAuthResponse {
	this := DatabricksIntegrationAccountPatAuthResponse{}
	var authType DatabricksIntegrationAccountPatAuthType = DATABRICKSINTEGRATIONACCOUNTPATAUTHTYPE_PAT
	this.AuthType = authType
	return &this
}

// GetAuthType returns the AuthType field value.
func (o *DatabricksIntegrationAccountPatAuthResponse) GetAuthType() DatabricksIntegrationAccountPatAuthType {
	if o == nil {
		var ret DatabricksIntegrationAccountPatAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationAccountPatAuthResponse) GetAuthTypeOk() (*DatabricksIntegrationAccountPatAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *DatabricksIntegrationAccountPatAuthResponse) SetAuthType(v DatabricksIntegrationAccountPatAuthType) {
	o.AuthType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabricksIntegrationAccountPatAuthResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth_type"] = o.AuthType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatabricksIntegrationAccountPatAuthResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthType *DatabricksIntegrationAccountPatAuthType `json:"auth_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuthType == nil {
		return fmt.Errorf("required field auth_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth_type"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.AuthType.IsValid() {
		hasInvalidField = true
	} else {
		o.AuthType = *all.AuthType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
