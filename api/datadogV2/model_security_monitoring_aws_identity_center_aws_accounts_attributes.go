// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringAWSIdentityCenterAWSAccountsAttributes The attributes of the AWS IAM Identity Center AWS account prerequisites.
type SecurityMonitoringAWSIdentityCenterAWSAccountsAttributes struct {
	// The integrated AWS accounts that users are read from, including AWS Organization management accounts
	// and any accounts that users were collected from.
	AwsAccounts []SecurityMonitoringAWSIdentityCenterAWSAccount `json:"aws_accounts"`
	// Whether AWS IAM Identity Center users were collected for the organization.
	HasValidPrerequisite bool `json:"has_valid_prerequisite"`
	// The missing setup step when `has_valid_prerequisite` is `false`.
	Reason *string `json:"reason,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringAWSIdentityCenterAWSAccountsAttributes instantiates a new SecurityMonitoringAWSIdentityCenterAWSAccountsAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringAWSIdentityCenterAWSAccountsAttributes(awsAccounts []SecurityMonitoringAWSIdentityCenterAWSAccount, hasValidPrerequisite bool) *SecurityMonitoringAWSIdentityCenterAWSAccountsAttributes {
	this := SecurityMonitoringAWSIdentityCenterAWSAccountsAttributes{}
	this.AwsAccounts = awsAccounts
	this.HasValidPrerequisite = hasValidPrerequisite
	return &this
}

// NewSecurityMonitoringAWSIdentityCenterAWSAccountsAttributesWithDefaults instantiates a new SecurityMonitoringAWSIdentityCenterAWSAccountsAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringAWSIdentityCenterAWSAccountsAttributesWithDefaults() *SecurityMonitoringAWSIdentityCenterAWSAccountsAttributes {
	this := SecurityMonitoringAWSIdentityCenterAWSAccountsAttributes{}
	return &this
}

// GetAwsAccounts returns the AwsAccounts field value.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccountsAttributes) GetAwsAccounts() []SecurityMonitoringAWSIdentityCenterAWSAccount {
	if o == nil {
		var ret []SecurityMonitoringAWSIdentityCenterAWSAccount
		return ret
	}
	return o.AwsAccounts
}

// GetAwsAccountsOk returns a tuple with the AwsAccounts field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccountsAttributes) GetAwsAccountsOk() (*[]SecurityMonitoringAWSIdentityCenterAWSAccount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsAccounts, true
}

// SetAwsAccounts sets field value.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccountsAttributes) SetAwsAccounts(v []SecurityMonitoringAWSIdentityCenterAWSAccount) {
	o.AwsAccounts = v
}

// GetHasValidPrerequisite returns the HasValidPrerequisite field value.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccountsAttributes) GetHasValidPrerequisite() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HasValidPrerequisite
}

// GetHasValidPrerequisiteOk returns a tuple with the HasValidPrerequisite field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccountsAttributes) GetHasValidPrerequisiteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasValidPrerequisite, true
}

// SetHasValidPrerequisite sets field value.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccountsAttributes) SetHasValidPrerequisite(v bool) {
	o.HasValidPrerequisite = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccountsAttributes) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccountsAttributes) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccountsAttributes) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccountsAttributes) SetReason(v string) {
	o.Reason = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringAWSIdentityCenterAWSAccountsAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aws_accounts"] = o.AwsAccounts
	toSerialize["has_valid_prerequisite"] = o.HasValidPrerequisite
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringAWSIdentityCenterAWSAccountsAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AwsAccounts          *[]SecurityMonitoringAWSIdentityCenterAWSAccount `json:"aws_accounts"`
		HasValidPrerequisite *bool                                            `json:"has_valid_prerequisite"`
		Reason               *string                                          `json:"reason,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AwsAccounts == nil {
		return fmt.Errorf("required field aws_accounts missing")
	}
	if all.HasValidPrerequisite == nil {
		return fmt.Errorf("required field has_valid_prerequisite missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aws_accounts", "has_valid_prerequisite", "reason"})
	} else {
		return err
	}
	o.AwsAccounts = *all.AwsAccounts
	o.HasValidPrerequisite = *all.HasValidPrerequisite
	o.Reason = all.Reason

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
