// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringOktaIntegrationCredentialsValidateAttributes The Okta credentials to validate against the external entity source.
type SecurityMonitoringOktaIntegrationCredentialsValidateAttributes struct {
	// The domain associated with the external entity source.
	Domain string `json:"domain"`
	// The source type for an Okta entity context sync.
	IntegrationType SecurityMonitoringIntegrationTypeOkta `json:"integration_type"`
	// Credentials for an Okta entity context sync.
	Secrets SecurityMonitoringIntegrationConfigOktaSecrets `json:"secrets"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringOktaIntegrationCredentialsValidateAttributes instantiates a new SecurityMonitoringOktaIntegrationCredentialsValidateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringOktaIntegrationCredentialsValidateAttributes(domain string, integrationType SecurityMonitoringIntegrationTypeOkta, secrets SecurityMonitoringIntegrationConfigOktaSecrets) *SecurityMonitoringOktaIntegrationCredentialsValidateAttributes {
	this := SecurityMonitoringOktaIntegrationCredentialsValidateAttributes{}
	this.Domain = domain
	this.IntegrationType = integrationType
	this.Secrets = secrets
	return &this
}

// NewSecurityMonitoringOktaIntegrationCredentialsValidateAttributesWithDefaults instantiates a new SecurityMonitoringOktaIntegrationCredentialsValidateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringOktaIntegrationCredentialsValidateAttributesWithDefaults() *SecurityMonitoringOktaIntegrationCredentialsValidateAttributes {
	this := SecurityMonitoringOktaIntegrationCredentialsValidateAttributes{}
	return &this
}

// GetDomain returns the Domain field value.
func (o *SecurityMonitoringOktaIntegrationCredentialsValidateAttributes) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringOktaIntegrationCredentialsValidateAttributes) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value.
func (o *SecurityMonitoringOktaIntegrationCredentialsValidateAttributes) SetDomain(v string) {
	o.Domain = v
}

// GetIntegrationType returns the IntegrationType field value.
func (o *SecurityMonitoringOktaIntegrationCredentialsValidateAttributes) GetIntegrationType() SecurityMonitoringIntegrationTypeOkta {
	if o == nil {
		var ret SecurityMonitoringIntegrationTypeOkta
		return ret
	}
	return o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringOktaIntegrationCredentialsValidateAttributes) GetIntegrationTypeOk() (*SecurityMonitoringIntegrationTypeOkta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationType, true
}

// SetIntegrationType sets field value.
func (o *SecurityMonitoringOktaIntegrationCredentialsValidateAttributes) SetIntegrationType(v SecurityMonitoringIntegrationTypeOkta) {
	o.IntegrationType = v
}

// GetSecrets returns the Secrets field value.
func (o *SecurityMonitoringOktaIntegrationCredentialsValidateAttributes) GetSecrets() SecurityMonitoringIntegrationConfigOktaSecrets {
	if o == nil {
		var ret SecurityMonitoringIntegrationConfigOktaSecrets
		return ret
	}
	return o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringOktaIntegrationCredentialsValidateAttributes) GetSecretsOk() (*SecurityMonitoringIntegrationConfigOktaSecrets, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secrets, true
}

// SetSecrets sets field value.
func (o *SecurityMonitoringOktaIntegrationCredentialsValidateAttributes) SetSecrets(v SecurityMonitoringIntegrationConfigOktaSecrets) {
	o.Secrets = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringOktaIntegrationCredentialsValidateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["domain"] = o.Domain
	toSerialize["integration_type"] = o.IntegrationType
	toSerialize["secrets"] = o.Secrets

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringOktaIntegrationCredentialsValidateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Domain          *string                                         `json:"domain"`
		IntegrationType *SecurityMonitoringIntegrationTypeOkta          `json:"integration_type"`
		Secrets         *SecurityMonitoringIntegrationConfigOktaSecrets `json:"secrets"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Domain == nil {
		return fmt.Errorf("required field domain missing")
	}
	if all.IntegrationType == nil {
		return fmt.Errorf("required field integration_type missing")
	}
	if all.Secrets == nil {
		return fmt.Errorf("required field secrets missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"domain", "integration_type", "secrets"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Domain = *all.Domain
	if !all.IntegrationType.IsValid() {
		hasInvalidField = true
	} else {
		o.IntegrationType = *all.IntegrationType
	}
	if all.Secrets.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Secrets = *all.Secrets

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
