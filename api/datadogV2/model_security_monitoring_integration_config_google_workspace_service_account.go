// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount The Google Cloud service account JSON used to authenticate against the Google Workspace Admin SDK. Additional keys beyond those documented are preserved.
type SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount struct {
	// The service account client email.
	ClientEmail string `json:"client_email"`
	// The service account private key.
	PrivateKey string `json:"private_key"`
	// The Google Cloud project ID that owns the service account.
	ProjectId string `json:"project_id"`
	// The service account type. Must be `service_account`.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount instantiates a new SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount(clientEmail string, privateKey string, projectId string, typeVar string) *SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount {
	this := SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount{}
	this.ClientEmail = clientEmail
	this.PrivateKey = privateKey
	this.ProjectId = projectId
	this.Type = typeVar
	return &this
}

// NewSecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccountWithDefaults instantiates a new SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccountWithDefaults() *SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount {
	this := SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount{}
	return &this
}

// GetClientEmail returns the ClientEmail field value.
func (o *SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount) GetClientEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientEmail
}

// GetClientEmailOk returns a tuple with the ClientEmail field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount) GetClientEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientEmail, true
}

// SetClientEmail sets field value.
func (o *SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount) SetClientEmail(v string) {
	o.ClientEmail = v
}

// GetPrivateKey returns the PrivateKey field value.
func (o *SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount) GetPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value.
func (o *SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount) SetPrivateKey(v string) {
	o.PrivateKey = v
}

// GetProjectId returns the ProjectId field value.
func (o *SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value.
func (o *SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount) SetProjectId(v string) {
	o.ProjectId = v
}

// GetType returns the Type field value.
func (o *SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["client_email"] = o.ClientEmail
	toSerialize["private_key"] = o.PrivateKey
	toSerialize["project_id"] = o.ProjectId
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClientEmail *string `json:"client_email"`
		PrivateKey  *string `json:"private_key"`
		ProjectId   *string `json:"project_id"`
		Type        *string `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ClientEmail == nil {
		return fmt.Errorf("required field client_email missing")
	}
	if all.PrivateKey == nil {
		return fmt.Errorf("required field private_key missing")
	}
	if all.ProjectId == nil {
		return fmt.Errorf("required field project_id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"client_email", "private_key", "project_id", "type"})
	} else {
		return err
	}
	o.ClientEmail = *all.ClientEmail
	o.PrivateKey = *all.PrivateKey
	o.ProjectId = *all.ProjectId
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
