// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudWorkloadSecurityAgentRuleActionNetworkFilter The network filter action applied on the network traffic matching the rule.
type CloudWorkloadSecurityAgentRuleActionNetworkFilter struct {
	// The filter expression of the network filter action.
	Filter *string `json:"filter,omitempty"`
	// The policy of the network filter action.
	Policy *string `json:"policy,omitempty"`
	// The scope of the network filter action.
	Scope *string `json:"scope,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudWorkloadSecurityAgentRuleActionNetworkFilter instantiates a new CloudWorkloadSecurityAgentRuleActionNetworkFilter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudWorkloadSecurityAgentRuleActionNetworkFilter() *CloudWorkloadSecurityAgentRuleActionNetworkFilter {
	this := CloudWorkloadSecurityAgentRuleActionNetworkFilter{}
	return &this
}

// NewCloudWorkloadSecurityAgentRuleActionNetworkFilterWithDefaults instantiates a new CloudWorkloadSecurityAgentRuleActionNetworkFilter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudWorkloadSecurityAgentRuleActionNetworkFilterWithDefaults() *CloudWorkloadSecurityAgentRuleActionNetworkFilter {
	this := CloudWorkloadSecurityAgentRuleActionNetworkFilter{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleActionNetworkFilter) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleActionNetworkFilter) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleActionNetworkFilter) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *CloudWorkloadSecurityAgentRuleActionNetworkFilter) SetFilter(v string) {
	o.Filter = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleActionNetworkFilter) GetPolicy() string {
	if o == nil || o.Policy == nil {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleActionNetworkFilter) GetPolicyOk() (*string, bool) {
	if o == nil || o.Policy == nil {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleActionNetworkFilter) HasPolicy() bool {
	return o != nil && o.Policy != nil
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *CloudWorkloadSecurityAgentRuleActionNetworkFilter) SetPolicy(v string) {
	o.Policy = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleActionNetworkFilter) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleActionNetworkFilter) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleActionNetworkFilter) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *CloudWorkloadSecurityAgentRuleActionNetworkFilter) SetScope(v string) {
	o.Scope = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudWorkloadSecurityAgentRuleActionNetworkFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Policy != nil {
		toSerialize["policy"] = o.Policy
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudWorkloadSecurityAgentRuleActionNetworkFilter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filter *string `json:"filter,omitempty"`
		Policy *string `json:"policy,omitempty"`
		Scope  *string `json:"scope,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"filter", "policy", "scope"})
	} else {
		return err
	}
	o.Filter = all.Filter
	o.Policy = all.Policy
	o.Scope = all.Scope

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
