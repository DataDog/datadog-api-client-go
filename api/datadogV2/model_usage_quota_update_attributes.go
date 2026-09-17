// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageQuotaUpdateAttributes Attributes to update on a usage quota. Omitting a property leaves its current value unchanged.
type UsageQuotaUpdateAttributes struct {
	// Whether to actively block usage above the limit. Omit this field to leave the current enforcement setting unchanged.
	Enforced datadog.NullableBool `json:"enforced,omitempty"`
	// The new quota limit in the usage units defined by the quota namespace. For an organization-wide quota (empty scope), the limit must be greater than the usage already recorded in the current period. Omit this field to leave the current limit unchanged.
	UsageLimit datadog.NullableInt64 `json:"usage_limit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUsageQuotaUpdateAttributes instantiates a new UsageQuotaUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageQuotaUpdateAttributes() *UsageQuotaUpdateAttributes {
	this := UsageQuotaUpdateAttributes{}
	return &this
}

// NewUsageQuotaUpdateAttributesWithDefaults instantiates a new UsageQuotaUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageQuotaUpdateAttributesWithDefaults() *UsageQuotaUpdateAttributes {
	this := UsageQuotaUpdateAttributes{}
	return &this
}

// GetEnforced returns the Enforced field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageQuotaUpdateAttributes) GetEnforced() bool {
	if o == nil || o.Enforced.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Enforced.Get()
}

// GetEnforcedOk returns a tuple with the Enforced field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageQuotaUpdateAttributes) GetEnforcedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enforced.Get(), o.Enforced.IsSet()
}

// HasEnforced returns a boolean if a field has been set.
func (o *UsageQuotaUpdateAttributes) HasEnforced() bool {
	return o != nil && o.Enforced.IsSet()
}

// SetEnforced gets a reference to the given datadog.NullableBool and assigns it to the Enforced field.
func (o *UsageQuotaUpdateAttributes) SetEnforced(v bool) {
	o.Enforced.Set(&v)
}

// SetEnforcedNil sets the value for Enforced to be an explicit nil.
func (o *UsageQuotaUpdateAttributes) SetEnforcedNil() {
	o.Enforced.Set(nil)
}

// UnsetEnforced ensures that no value is present for Enforced, not even an explicit nil.
func (o *UsageQuotaUpdateAttributes) UnsetEnforced() {
	o.Enforced.Unset()
}

// GetUsageLimit returns the UsageLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageQuotaUpdateAttributes) GetUsageLimit() int64 {
	if o == nil || o.UsageLimit.Get() == nil {
		var ret int64
		return ret
	}
	return *o.UsageLimit.Get()
}

// GetUsageLimitOk returns a tuple with the UsageLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageQuotaUpdateAttributes) GetUsageLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsageLimit.Get(), o.UsageLimit.IsSet()
}

// HasUsageLimit returns a boolean if a field has been set.
func (o *UsageQuotaUpdateAttributes) HasUsageLimit() bool {
	return o != nil && o.UsageLimit.IsSet()
}

// SetUsageLimit gets a reference to the given datadog.NullableInt64 and assigns it to the UsageLimit field.
func (o *UsageQuotaUpdateAttributes) SetUsageLimit(v int64) {
	o.UsageLimit.Set(&v)
}

// SetUsageLimitNil sets the value for UsageLimit to be an explicit nil.
func (o *UsageQuotaUpdateAttributes) SetUsageLimitNil() {
	o.UsageLimit.Set(nil)
}

// UnsetUsageLimit ensures that no value is present for UsageLimit, not even an explicit nil.
func (o *UsageQuotaUpdateAttributes) UnsetUsageLimit() {
	o.UsageLimit.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageQuotaUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Enforced.IsSet() {
		toSerialize["enforced"] = o.Enforced.Get()
	}
	if o.UsageLimit.IsSet() {
		toSerialize["usage_limit"] = o.UsageLimit.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageQuotaUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enforced   datadog.NullableBool  `json:"enforced,omitempty"`
		UsageLimit datadog.NullableInt64 `json:"usage_limit,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enforced", "usage_limit"})
	} else {
		return err
	}
	o.Enforced = all.Enforced
	o.UsageLimit = all.UsageLimit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
