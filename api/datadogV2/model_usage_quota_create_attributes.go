// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageQuotaCreateAttributes Attributes for creating or updating a usage quota by scope.
type UsageQuotaCreateAttributes struct {
	// Whether to actively block usage above the limit instead of only tracking or alerting on it.
	Enforced bool `json:"enforced"`
	// A namespace-specific key and value identifying what the quota applies to within an organization. The object must contain exactly one entry. Use `"*"` as the value for the default quota applied to entities without a specific quota, or omit the scope for an organization-wide quota. A specific value must identify an existing user handle in the caller's organization when `include_descendants` is false. When `include_descendants` is true, the handle must exist in the caller's organization or in at least one targeted descendant organization; the quota is then applied only to the organizations where that handle exists, and the request fails only if the handle exists in none of them.
	Scope map[string]string `json:"scope,omitempty"`
	// The quota limit to set in the usage units defined by the quota namespace. For an organization-wide quota (scope omitted), the limit must be greater than the usage already recorded in the current period.
	UsageLimit int64 `json:"usage_limit"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUsageQuotaCreateAttributes instantiates a new UsageQuotaCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageQuotaCreateAttributes(enforced bool, usageLimit int64) *UsageQuotaCreateAttributes {
	this := UsageQuotaCreateAttributes{}
	this.Enforced = enforced
	this.UsageLimit = usageLimit
	return &this
}

// NewUsageQuotaCreateAttributesWithDefaults instantiates a new UsageQuotaCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageQuotaCreateAttributesWithDefaults() *UsageQuotaCreateAttributes {
	this := UsageQuotaCreateAttributes{}
	return &this
}

// GetEnforced returns the Enforced field value.
func (o *UsageQuotaCreateAttributes) GetEnforced() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enforced
}

// GetEnforcedOk returns a tuple with the Enforced field value
// and a boolean to check if the value has been set.
func (o *UsageQuotaCreateAttributes) GetEnforcedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enforced, true
}

// SetEnforced sets field value.
func (o *UsageQuotaCreateAttributes) SetEnforced(v bool) {
	o.Enforced = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *UsageQuotaCreateAttributes) GetScope() map[string]string {
	if o == nil || o.Scope == nil {
		var ret map[string]string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageQuotaCreateAttributes) GetScopeOk() (*map[string]string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return &o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *UsageQuotaCreateAttributes) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given map[string]string and assigns it to the Scope field.
func (o *UsageQuotaCreateAttributes) SetScope(v map[string]string) {
	o.Scope = v
}

// GetUsageLimit returns the UsageLimit field value.
func (o *UsageQuotaCreateAttributes) GetUsageLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.UsageLimit
}

// GetUsageLimitOk returns a tuple with the UsageLimit field value
// and a boolean to check if the value has been set.
func (o *UsageQuotaCreateAttributes) GetUsageLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageLimit, true
}

// SetUsageLimit sets field value.
func (o *UsageQuotaCreateAttributes) SetUsageLimit(v int64) {
	o.UsageLimit = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageQuotaCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["enforced"] = o.Enforced
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	toSerialize["usage_limit"] = o.UsageLimit

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageQuotaCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enforced   *bool             `json:"enforced"`
		Scope      map[string]string `json:"scope,omitempty"`
		UsageLimit *int64            `json:"usage_limit"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enforced == nil {
		return fmt.Errorf("required field enforced missing")
	}
	if all.UsageLimit == nil {
		return fmt.Errorf("required field usage_limit missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enforced", "scope", "usage_limit"})
	} else {
		return err
	}
	o.Enforced = *all.Enforced
	o.Scope = all.Scope
	o.UsageLimit = *all.UsageLimit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
