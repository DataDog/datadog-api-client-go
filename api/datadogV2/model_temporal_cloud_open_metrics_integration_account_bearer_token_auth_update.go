// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate Bearer token authentication. Use the API key of a Temporal Cloud service account with the Metrics Read-Only account-level role. Only the fields provided are changed; omit `token` to keep the stored one.
type TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate struct {
	// The authentication method type.
	AuthType TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthType `json:"auth_type"`
	// Secret token used to authenticate with Temporal Cloud.
	Token *string `json:"token,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate instantiates a new TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate(authType TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthType) *TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate {
	this := TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate{}
	this.AuthType = authType
	return &this
}

// NewTemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdateWithDefaults instantiates a new TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdateWithDefaults() *TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate {
	this := TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate{}
	var authType TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthType = TEMPORALCLOUDOPENMETRICSINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN
	this.AuthType = authType
	return &this
}

// GetAuthType returns the AuthType field value.
func (o *TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate) GetAuthType() TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthType {
	if o == nil {
		var ret TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate) GetAuthTypeOk() (*TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate) SetAuthType(v TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthType) {
	o.AuthType = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate) HasToken() bool {
	return o != nil && o.Token != nil
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate) SetToken(v string) {
	o.Token = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate) MarshalJSON() ([]byte, error) {
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
func (o *TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthType *TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthType `json:"auth_type"`
		Token    *string                                                        `json:"token,omitempty"`
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
