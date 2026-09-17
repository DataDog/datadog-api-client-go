// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageQuotaResponseAttributes Attributes of a usage quota.
type UsageQuotaResponseAttributes struct {
	// Whether usage above the limit is actively blocked instead of only tracked or alerted on.
	Enforced bool `json:"enforced"`
	// The public ID of the organization that owns the quota.
	OrgPublicId string `json:"org_public_id"`
	// A namespace-specific key and value identifying what the quota applies to within an organization. The object contains exactly one entry. A value of `"*"` identifies the default quota applied to entities without a specific quota. This field is omitted for an organization-wide quota.
	Scope map[string]string `json:"scope,omitempty"`
	// The quota limit in the usage units defined by the quota namespace. May be fractional for quotas configured before public writes required whole units.
	UsageLimit float64 `json:"usage_limit"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUsageQuotaResponseAttributes instantiates a new UsageQuotaResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageQuotaResponseAttributes(enforced bool, orgPublicId string, usageLimit float64) *UsageQuotaResponseAttributes {
	this := UsageQuotaResponseAttributes{}
	this.Enforced = enforced
	this.OrgPublicId = orgPublicId
	this.UsageLimit = usageLimit
	return &this
}

// NewUsageQuotaResponseAttributesWithDefaults instantiates a new UsageQuotaResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageQuotaResponseAttributesWithDefaults() *UsageQuotaResponseAttributes {
	this := UsageQuotaResponseAttributes{}
	return &this
}

// GetEnforced returns the Enforced field value.
func (o *UsageQuotaResponseAttributes) GetEnforced() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enforced
}

// GetEnforcedOk returns a tuple with the Enforced field value
// and a boolean to check if the value has been set.
func (o *UsageQuotaResponseAttributes) GetEnforcedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enforced, true
}

// SetEnforced sets field value.
func (o *UsageQuotaResponseAttributes) SetEnforced(v bool) {
	o.Enforced = v
}

// GetOrgPublicId returns the OrgPublicId field value.
func (o *UsageQuotaResponseAttributes) GetOrgPublicId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgPublicId
}

// GetOrgPublicIdOk returns a tuple with the OrgPublicId field value
// and a boolean to check if the value has been set.
func (o *UsageQuotaResponseAttributes) GetOrgPublicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgPublicId, true
}

// SetOrgPublicId sets field value.
func (o *UsageQuotaResponseAttributes) SetOrgPublicId(v string) {
	o.OrgPublicId = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *UsageQuotaResponseAttributes) GetScope() map[string]string {
	if o == nil || o.Scope == nil {
		var ret map[string]string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageQuotaResponseAttributes) GetScopeOk() (*map[string]string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return &o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *UsageQuotaResponseAttributes) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given map[string]string and assigns it to the Scope field.
func (o *UsageQuotaResponseAttributes) SetScope(v map[string]string) {
	o.Scope = v
}

// GetUsageLimit returns the UsageLimit field value.
func (o *UsageQuotaResponseAttributes) GetUsageLimit() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.UsageLimit
}

// GetUsageLimitOk returns a tuple with the UsageLimit field value
// and a boolean to check if the value has been set.
func (o *UsageQuotaResponseAttributes) GetUsageLimitOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageLimit, true
}

// SetUsageLimit sets field value.
func (o *UsageQuotaResponseAttributes) SetUsageLimit(v float64) {
	o.UsageLimit = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageQuotaResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["enforced"] = o.Enforced
	toSerialize["org_public_id"] = o.OrgPublicId
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
func (o *UsageQuotaResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enforced    *bool             `json:"enforced"`
		OrgPublicId *string           `json:"org_public_id"`
		Scope       map[string]string `json:"scope,omitempty"`
		UsageLimit  *float64          `json:"usage_limit"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enforced == nil {
		return fmt.Errorf("required field enforced missing")
	}
	if all.OrgPublicId == nil {
		return fmt.Errorf("required field org_public_id missing")
	}
	if all.UsageLimit == nil {
		return fmt.Errorf("required field usage_limit missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enforced", "org_public_id", "scope", "usage_limit"})
	} else {
		return err
	}
	o.Enforced = *all.Enforced
	o.OrgPublicId = *all.OrgPublicId
	o.Scope = all.Scope
	o.UsageLimit = *all.UsageLimit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
