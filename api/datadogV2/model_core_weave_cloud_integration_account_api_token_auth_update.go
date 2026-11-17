// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate API token authentication, the only method CoreWeave integration accounts support. Only the fields provided are changed; omit `api_token` to keep the stored one.
type CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate struct {
	// Secret CoreWeave API access token used to authenticate with CoreWeave. It cannot be cleared; sending `null` or an empty string is rejected with a `422` response.
	ApiToken *string `json:"api_token,omitempty"`
	// The authentication method type.
	AuthType CoreWeaveCloudIntegrationAccountAPITokenAuthType `json:"auth_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCoreWeaveCloudIntegrationAccountAPITokenAuthUpdate instantiates a new CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCoreWeaveCloudIntegrationAccountAPITokenAuthUpdate(authType CoreWeaveCloudIntegrationAccountAPITokenAuthType) *CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate {
	this := CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate{}
	this.AuthType = authType
	return &this
}

// NewCoreWeaveCloudIntegrationAccountAPITokenAuthUpdateWithDefaults instantiates a new CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCoreWeaveCloudIntegrationAccountAPITokenAuthUpdateWithDefaults() *CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate {
	this := CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate{}
	var authType CoreWeaveCloudIntegrationAccountAPITokenAuthType = COREWEAVECLOUDINTEGRATIONACCOUNTAPITOKENAUTHTYPE_COREWEAVE_API_TOKEN
	this.AuthType = authType
	return &this
}

// GetApiToken returns the ApiToken field value if set, zero value otherwise.
func (o *CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate) GetApiToken() string {
	if o == nil || o.ApiToken == nil {
		var ret string
		return ret
	}
	return *o.ApiToken
}

// GetApiTokenOk returns a tuple with the ApiToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate) GetApiTokenOk() (*string, bool) {
	if o == nil || o.ApiToken == nil {
		return nil, false
	}
	return o.ApiToken, true
}

// HasApiToken returns a boolean if a field has been set.
func (o *CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate) HasApiToken() bool {
	return o != nil && o.ApiToken != nil
}

// SetApiToken gets a reference to the given string and assigns it to the ApiToken field.
func (o *CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate) SetApiToken(v string) {
	o.ApiToken = &v
}

// GetAuthType returns the AuthType field value.
func (o *CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate) GetAuthType() CoreWeaveCloudIntegrationAccountAPITokenAuthType {
	if o == nil {
		var ret CoreWeaveCloudIntegrationAccountAPITokenAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate) GetAuthTypeOk() (*CoreWeaveCloudIntegrationAccountAPITokenAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate) SetAuthType(v CoreWeaveCloudIntegrationAccountAPITokenAuthType) {
	o.AuthType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApiToken != nil {
		toSerialize["api_token"] = o.ApiToken
	}
	toSerialize["auth_type"] = o.AuthType
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiToken *string                                           `json:"api_token,omitempty"`
		AuthType *CoreWeaveCloudIntegrationAccountAPITokenAuthType `json:"auth_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuthType == nil {
		return fmt.Errorf("required field auth_type missing")
	}

	hasInvalidField := false
	o.ApiToken = all.ApiToken
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
