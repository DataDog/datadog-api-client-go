// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringIntegrationConfigCrowdStrikeSecrets Credentials for a CrowdStrike entity context sync.
type SecurityMonitoringIntegrationConfigCrowdStrikeSecrets struct {
	// The CrowdStrike API client ID.
	ClientId string `json:"client_id"`
	// The CrowdStrike API client secret.
	ClientSecret string `json:"client_secret"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringIntegrationConfigCrowdStrikeSecrets instantiates a new SecurityMonitoringIntegrationConfigCrowdStrikeSecrets object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringIntegrationConfigCrowdStrikeSecrets(clientId string, clientSecret string) *SecurityMonitoringIntegrationConfigCrowdStrikeSecrets {
	this := SecurityMonitoringIntegrationConfigCrowdStrikeSecrets{}
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	return &this
}

// NewSecurityMonitoringIntegrationConfigCrowdStrikeSecretsWithDefaults instantiates a new SecurityMonitoringIntegrationConfigCrowdStrikeSecrets object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringIntegrationConfigCrowdStrikeSecretsWithDefaults() *SecurityMonitoringIntegrationConfigCrowdStrikeSecrets {
	this := SecurityMonitoringIntegrationConfigCrowdStrikeSecrets{}
	return &this
}

// GetClientId returns the ClientId field value.
func (o *SecurityMonitoringIntegrationConfigCrowdStrikeSecrets) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigCrowdStrikeSecrets) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value.
func (o *SecurityMonitoringIntegrationConfigCrowdStrikeSecrets) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value.
func (o *SecurityMonitoringIntegrationConfigCrowdStrikeSecrets) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigCrowdStrikeSecrets) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value.
func (o *SecurityMonitoringIntegrationConfigCrowdStrikeSecrets) SetClientSecret(v string) {
	o.ClientSecret = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringIntegrationConfigCrowdStrikeSecrets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["client_id"] = o.ClientId
	toSerialize["client_secret"] = o.ClientSecret

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringIntegrationConfigCrowdStrikeSecrets) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClientId     *string `json:"client_id"`
		ClientSecret *string `json:"client_secret"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ClientId == nil {
		return fmt.Errorf("required field client_id missing")
	}
	if all.ClientSecret == nil {
		return fmt.Errorf("required field client_secret missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"client_id", "client_secret"})
	} else {
		return err
	}
	o.ClientId = *all.ClientId
	o.ClientSecret = *all.ClientSecret

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
