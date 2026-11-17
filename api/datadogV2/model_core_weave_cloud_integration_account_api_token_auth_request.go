// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CoreWeaveCloudIntegrationAccountAPITokenAuthRequest API token authentication, the only method CoreWeave integration accounts support. Create the token in the CoreWeave Cloud console under Access Tokens; it starts with `CW-SECRET-`. The user who owns the token must have the Observability Viewer role, because tokens inherit their user's permissions.
type CoreWeaveCloudIntegrationAccountAPITokenAuthRequest struct {
	// Secret CoreWeave API access token used to authenticate with CoreWeave. Values that do not start with `CW-SECRET-` are rejected with a `422` response.
	ApiToken string `json:"api_token"`
	// The authentication method type.
	AuthType CoreWeaveCloudIntegrationAccountAPITokenAuthType `json:"auth_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCoreWeaveCloudIntegrationAccountAPITokenAuthRequest instantiates a new CoreWeaveCloudIntegrationAccountAPITokenAuthRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCoreWeaveCloudIntegrationAccountAPITokenAuthRequest(apiToken string, authType CoreWeaveCloudIntegrationAccountAPITokenAuthType) *CoreWeaveCloudIntegrationAccountAPITokenAuthRequest {
	this := CoreWeaveCloudIntegrationAccountAPITokenAuthRequest{}
	this.ApiToken = apiToken
	this.AuthType = authType
	return &this
}

// NewCoreWeaveCloudIntegrationAccountAPITokenAuthRequestWithDefaults instantiates a new CoreWeaveCloudIntegrationAccountAPITokenAuthRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCoreWeaveCloudIntegrationAccountAPITokenAuthRequestWithDefaults() *CoreWeaveCloudIntegrationAccountAPITokenAuthRequest {
	this := CoreWeaveCloudIntegrationAccountAPITokenAuthRequest{}
	var authType CoreWeaveCloudIntegrationAccountAPITokenAuthType = COREWEAVECLOUDINTEGRATIONACCOUNTAPITOKENAUTHTYPE_COREWEAVE_API_TOKEN
	this.AuthType = authType
	return &this
}

// GetApiToken returns the ApiToken field value.
func (o *CoreWeaveCloudIntegrationAccountAPITokenAuthRequest) GetApiToken() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApiToken
}

// GetApiTokenOk returns a tuple with the ApiToken field value
// and a boolean to check if the value has been set.
func (o *CoreWeaveCloudIntegrationAccountAPITokenAuthRequest) GetApiTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiToken, true
}

// SetApiToken sets field value.
func (o *CoreWeaveCloudIntegrationAccountAPITokenAuthRequest) SetApiToken(v string) {
	o.ApiToken = v
}

// GetAuthType returns the AuthType field value.
func (o *CoreWeaveCloudIntegrationAccountAPITokenAuthRequest) GetAuthType() CoreWeaveCloudIntegrationAccountAPITokenAuthType {
	if o == nil {
		var ret CoreWeaveCloudIntegrationAccountAPITokenAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *CoreWeaveCloudIntegrationAccountAPITokenAuthRequest) GetAuthTypeOk() (*CoreWeaveCloudIntegrationAccountAPITokenAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *CoreWeaveCloudIntegrationAccountAPITokenAuthRequest) SetAuthType(v CoreWeaveCloudIntegrationAccountAPITokenAuthType) {
	o.AuthType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CoreWeaveCloudIntegrationAccountAPITokenAuthRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["api_token"] = o.ApiToken
	toSerialize["auth_type"] = o.AuthType
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CoreWeaveCloudIntegrationAccountAPITokenAuthRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiToken *string                                           `json:"api_token"`
		AuthType *CoreWeaveCloudIntegrationAccountAPITokenAuthType `json:"auth_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApiToken == nil {
		return fmt.Errorf("required field api_token missing")
	}
	if all.AuthType == nil {
		return fmt.Errorf("required field auth_type missing")
	}

	hasInvalidField := false
	o.ApiToken = *all.ApiToken
	if !all.AuthType.IsValid() {
		hasInvalidField = true
	} else {
		o.AuthType = *all.AuthType
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
