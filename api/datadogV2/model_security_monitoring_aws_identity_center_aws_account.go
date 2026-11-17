// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringAWSIdentityCenterAWSAccount An integrated AWS account the AWS IAM Identity Center integration reads users from.
type SecurityMonitoringAWSIdentityCenterAWSAccount struct {
	// The ID of the AWS account integration configuration.
	AccountConfigId string `json:"account_config_id"`
	// The AWS account ID.
	AccountId string `json:"account_id"`
	// The name of the AWS account integration.
	AccountName string `json:"account_name"`
	// Whether AWS IAM Identity Center users were collected from this AWS account.
	CollectingUsers bool `json:"collecting_users"`
	// Whether resource collection is enabled for this AWS account.
	ResourceCollectionEnabled bool `json:"resource_collection_enabled"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringAWSIdentityCenterAWSAccount instantiates a new SecurityMonitoringAWSIdentityCenterAWSAccount object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringAWSIdentityCenterAWSAccount(accountConfigId string, accountId string, accountName string, collectingUsers bool, resourceCollectionEnabled bool) *SecurityMonitoringAWSIdentityCenterAWSAccount {
	this := SecurityMonitoringAWSIdentityCenterAWSAccount{}
	this.AccountConfigId = accountConfigId
	this.AccountId = accountId
	this.AccountName = accountName
	this.CollectingUsers = collectingUsers
	this.ResourceCollectionEnabled = resourceCollectionEnabled
	return &this
}

// NewSecurityMonitoringAWSIdentityCenterAWSAccountWithDefaults instantiates a new SecurityMonitoringAWSIdentityCenterAWSAccount object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringAWSIdentityCenterAWSAccountWithDefaults() *SecurityMonitoringAWSIdentityCenterAWSAccount {
	this := SecurityMonitoringAWSIdentityCenterAWSAccount{}
	return &this
}

// GetAccountConfigId returns the AccountConfigId field value.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccount) GetAccountConfigId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountConfigId
}

// GetAccountConfigIdOk returns a tuple with the AccountConfigId field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccount) GetAccountConfigIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountConfigId, true
}

// SetAccountConfigId sets field value.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccount) SetAccountConfigId(v string) {
	o.AccountConfigId = v
}

// GetAccountId returns the AccountId field value.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccount) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccount) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccount) SetAccountId(v string) {
	o.AccountId = v
}

// GetAccountName returns the AccountName field value.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccount) GetAccountName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccount) GetAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountName, true
}

// SetAccountName sets field value.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccount) SetAccountName(v string) {
	o.AccountName = v
}

// GetCollectingUsers returns the CollectingUsers field value.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccount) GetCollectingUsers() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.CollectingUsers
}

// GetCollectingUsersOk returns a tuple with the CollectingUsers field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccount) GetCollectingUsersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectingUsers, true
}

// SetCollectingUsers sets field value.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccount) SetCollectingUsers(v bool) {
	o.CollectingUsers = v
}

// GetResourceCollectionEnabled returns the ResourceCollectionEnabled field value.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccount) GetResourceCollectionEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.ResourceCollectionEnabled
}

// GetResourceCollectionEnabledOk returns a tuple with the ResourceCollectionEnabled field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccount) GetResourceCollectionEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceCollectionEnabled, true
}

// SetResourceCollectionEnabled sets field value.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccount) SetResourceCollectionEnabled(v bool) {
	o.ResourceCollectionEnabled = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringAWSIdentityCenterAWSAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_config_id"] = o.AccountConfigId
	toSerialize["account_id"] = o.AccountId
	toSerialize["account_name"] = o.AccountName
	toSerialize["collecting_users"] = o.CollectingUsers
	toSerialize["resource_collection_enabled"] = o.ResourceCollectionEnabled

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccount) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountConfigId           *string `json:"account_config_id"`
		AccountId                 *string `json:"account_id"`
		AccountName               *string `json:"account_name"`
		CollectingUsers           *bool   `json:"collecting_users"`
		ResourceCollectionEnabled *bool   `json:"resource_collection_enabled"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountConfigId == nil {
		return fmt.Errorf("required field account_config_id missing")
	}
	if all.AccountId == nil {
		return fmt.Errorf("required field account_id missing")
	}
	if all.AccountName == nil {
		return fmt.Errorf("required field account_name missing")
	}
	if all.CollectingUsers == nil {
		return fmt.Errorf("required field collecting_users missing")
	}
	if all.ResourceCollectionEnabled == nil {
		return fmt.Errorf("required field resource_collection_enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_config_id", "account_id", "account_name", "collecting_users", "resource_collection_enabled"})
	} else {
		return err
	}
	o.AccountConfigId = *all.AccountConfigId
	o.AccountId = *all.AccountId
	o.AccountName = *all.AccountName
	o.CollectingUsers = *all.CollectingUsers
	o.ResourceCollectionEnabled = *all.ResourceCollectionEnabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
