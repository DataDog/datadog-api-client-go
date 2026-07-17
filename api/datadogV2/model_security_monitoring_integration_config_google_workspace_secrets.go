// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringIntegrationConfigGoogleWorkspaceSecrets Credentials for a Google Workspace entity context sync.
type SecurityMonitoringIntegrationConfigGoogleWorkspaceSecrets struct {
	// The admin email to impersonate for domain-wide delegation.
	AdminEmail *string `json:"admin_email,omitempty"`
	// The Google Cloud service account JSON used to authenticate against the Google Workspace Admin SDK. Additional keys beyond those documented are preserved.
	ServiceAccountJson SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount `json:"service_account_json"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringIntegrationConfigGoogleWorkspaceSecrets instantiates a new SecurityMonitoringIntegrationConfigGoogleWorkspaceSecrets object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringIntegrationConfigGoogleWorkspaceSecrets(serviceAccountJson SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount) *SecurityMonitoringIntegrationConfigGoogleWorkspaceSecrets {
	this := SecurityMonitoringIntegrationConfigGoogleWorkspaceSecrets{}
	this.ServiceAccountJson = serviceAccountJson
	return &this
}

// NewSecurityMonitoringIntegrationConfigGoogleWorkspaceSecretsWithDefaults instantiates a new SecurityMonitoringIntegrationConfigGoogleWorkspaceSecrets object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringIntegrationConfigGoogleWorkspaceSecretsWithDefaults() *SecurityMonitoringIntegrationConfigGoogleWorkspaceSecrets {
	this := SecurityMonitoringIntegrationConfigGoogleWorkspaceSecrets{}
	return &this
}

// GetAdminEmail returns the AdminEmail field value if set, zero value otherwise.
func (o *SecurityMonitoringIntegrationConfigGoogleWorkspaceSecrets) GetAdminEmail() string {
	if o == nil || o.AdminEmail == nil {
		var ret string
		return ret
	}
	return *o.AdminEmail
}

// GetAdminEmailOk returns a tuple with the AdminEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigGoogleWorkspaceSecrets) GetAdminEmailOk() (*string, bool) {
	if o == nil || o.AdminEmail == nil {
		return nil, false
	}
	return o.AdminEmail, true
}

// HasAdminEmail returns a boolean if a field has been set.
func (o *SecurityMonitoringIntegrationConfigGoogleWorkspaceSecrets) HasAdminEmail() bool {
	return o != nil && o.AdminEmail != nil
}

// SetAdminEmail gets a reference to the given string and assigns it to the AdminEmail field.
func (o *SecurityMonitoringIntegrationConfigGoogleWorkspaceSecrets) SetAdminEmail(v string) {
	o.AdminEmail = &v
}

// GetServiceAccountJson returns the ServiceAccountJson field value.
func (o *SecurityMonitoringIntegrationConfigGoogleWorkspaceSecrets) GetServiceAccountJson() SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount {
	if o == nil {
		var ret SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount
		return ret
	}
	return o.ServiceAccountJson
}

// GetServiceAccountJsonOk returns a tuple with the ServiceAccountJson field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigGoogleWorkspaceSecrets) GetServiceAccountJsonOk() (*SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceAccountJson, true
}

// SetServiceAccountJson sets field value.
func (o *SecurityMonitoringIntegrationConfigGoogleWorkspaceSecrets) SetServiceAccountJson(v SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount) {
	o.ServiceAccountJson = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringIntegrationConfigGoogleWorkspaceSecrets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AdminEmail != nil {
		toSerialize["admin_email"] = o.AdminEmail
	}
	toSerialize["service_account_json"] = o.ServiceAccountJson

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringIntegrationConfigGoogleWorkspaceSecrets) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AdminEmail         *string                                                           `json:"admin_email,omitempty"`
		ServiceAccountJson *SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount `json:"service_account_json"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ServiceAccountJson == nil {
		return fmt.Errorf("required field service_account_json missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"admin_email", "service_account_json"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AdminEmail = all.AdminEmail
	if all.ServiceAccountJson.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ServiceAccountJson = *all.ServiceAccountJson

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
