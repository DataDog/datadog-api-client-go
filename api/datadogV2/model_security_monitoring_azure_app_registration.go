// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringAzureAppRegistration An Azure App Registration discovered for the organization.
type SecurityMonitoringAzureAppRegistration struct {
	// The client ID of the App Registration.
	ClientId string `json:"client_id"`
	// The number of errors encountered while crawling resources for this App Registration.
	ErrorCount int64 `json:"error_count"`
	// Whether resource collection is enabled for this App Registration.
	ResourceCollectionEnabled bool `json:"resource_collection_enabled"`
	// The number of Azure subscriptions associated with this App Registration.
	SubscriptionCount int64 `json:"subscription_count"`
	// The Azure tenant ID of the App Registration.
	TenantId string `json:"tenant_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringAzureAppRegistration instantiates a new SecurityMonitoringAzureAppRegistration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringAzureAppRegistration(clientId string, errorCount int64, resourceCollectionEnabled bool, subscriptionCount int64, tenantId string) *SecurityMonitoringAzureAppRegistration {
	this := SecurityMonitoringAzureAppRegistration{}
	this.ClientId = clientId
	this.ErrorCount = errorCount
	this.ResourceCollectionEnabled = resourceCollectionEnabled
	this.SubscriptionCount = subscriptionCount
	this.TenantId = tenantId
	return &this
}

// NewSecurityMonitoringAzureAppRegistrationWithDefaults instantiates a new SecurityMonitoringAzureAppRegistration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringAzureAppRegistrationWithDefaults() *SecurityMonitoringAzureAppRegistration {
	this := SecurityMonitoringAzureAppRegistration{}
	return &this
}

// GetClientId returns the ClientId field value.
func (o *SecurityMonitoringAzureAppRegistration) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringAzureAppRegistration) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value.
func (o *SecurityMonitoringAzureAppRegistration) SetClientId(v string) {
	o.ClientId = v
}

// GetErrorCount returns the ErrorCount field value.
func (o *SecurityMonitoringAzureAppRegistration) GetErrorCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ErrorCount
}

// GetErrorCountOk returns a tuple with the ErrorCount field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringAzureAppRegistration) GetErrorCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCount, true
}

// SetErrorCount sets field value.
func (o *SecurityMonitoringAzureAppRegistration) SetErrorCount(v int64) {
	o.ErrorCount = v
}

// GetResourceCollectionEnabled returns the ResourceCollectionEnabled field value.
func (o *SecurityMonitoringAzureAppRegistration) GetResourceCollectionEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.ResourceCollectionEnabled
}

// GetResourceCollectionEnabledOk returns a tuple with the ResourceCollectionEnabled field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringAzureAppRegistration) GetResourceCollectionEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceCollectionEnabled, true
}

// SetResourceCollectionEnabled sets field value.
func (o *SecurityMonitoringAzureAppRegistration) SetResourceCollectionEnabled(v bool) {
	o.ResourceCollectionEnabled = v
}

// GetSubscriptionCount returns the SubscriptionCount field value.
func (o *SecurityMonitoringAzureAppRegistration) GetSubscriptionCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SubscriptionCount
}

// GetSubscriptionCountOk returns a tuple with the SubscriptionCount field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringAzureAppRegistration) GetSubscriptionCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionCount, true
}

// SetSubscriptionCount sets field value.
func (o *SecurityMonitoringAzureAppRegistration) SetSubscriptionCount(v int64) {
	o.SubscriptionCount = v
}

// GetTenantId returns the TenantId field value.
func (o *SecurityMonitoringAzureAppRegistration) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringAzureAppRegistration) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value.
func (o *SecurityMonitoringAzureAppRegistration) SetTenantId(v string) {
	o.TenantId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringAzureAppRegistration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["client_id"] = o.ClientId
	toSerialize["error_count"] = o.ErrorCount
	toSerialize["resource_collection_enabled"] = o.ResourceCollectionEnabled
	toSerialize["subscription_count"] = o.SubscriptionCount
	toSerialize["tenant_id"] = o.TenantId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringAzureAppRegistration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClientId                  *string `json:"client_id"`
		ErrorCount                *int64  `json:"error_count"`
		ResourceCollectionEnabled *bool   `json:"resource_collection_enabled"`
		SubscriptionCount         *int64  `json:"subscription_count"`
		TenantId                  *string `json:"tenant_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ClientId == nil {
		return fmt.Errorf("required field client_id missing")
	}
	if all.ErrorCount == nil {
		return fmt.Errorf("required field error_count missing")
	}
	if all.ResourceCollectionEnabled == nil {
		return fmt.Errorf("required field resource_collection_enabled missing")
	}
	if all.SubscriptionCount == nil {
		return fmt.Errorf("required field subscription_count missing")
	}
	if all.TenantId == nil {
		return fmt.Errorf("required field tenant_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"client_id", "error_count", "resource_collection_enabled", "subscription_count", "tenant_id"})
	} else {
		return err
	}
	o.ClientId = *all.ClientId
	o.ErrorCount = *all.ErrorCount
	o.ResourceCollectionEnabled = *all.ResourceCollectionEnabled
	o.SubscriptionCount = *all.SubscriptionCount
	o.TenantId = *all.TenantId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
