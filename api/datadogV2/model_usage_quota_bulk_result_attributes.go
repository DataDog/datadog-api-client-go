// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageQuotaBulkResultAttributes Attributes of a usage quota bulk write result. On success, all fields except `error` are present. On failure, only `error` is present and the other fields are omitted.
type UsageQuotaBulkResultAttributes struct {
	// Whether usage above the limit is actively blocked instead of only tracked or alerted on. Omitted if this item failed to write.
	Enforced *bool `json:"enforced,omitempty"`
	// An error message describing why this item failed to write. Omitted if this item was written successfully.
	Error *string `json:"error,omitempty"`
	// The public ID of the organization that owns the quota. Omitted if this item failed to write.
	OrgPublicId *string `json:"org_public_id,omitempty"`
	// A namespace-specific key and value identifying what the quota applies to within an organization. The object contains exactly one entry. A value of `"*"` identifies the default quota applied to entities without a specific quota. This field is omitted for an organization-wide quota.
	Scope map[string]string `json:"scope,omitempty"`
	// The quota limit in the usage units defined by the quota namespace. May be fractional for quotas configured before public writes required whole units. Omitted if this item failed to write.
	UsageLimit *float64 `json:"usage_limit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUsageQuotaBulkResultAttributes instantiates a new UsageQuotaBulkResultAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageQuotaBulkResultAttributes() *UsageQuotaBulkResultAttributes {
	this := UsageQuotaBulkResultAttributes{}
	return &this
}

// NewUsageQuotaBulkResultAttributesWithDefaults instantiates a new UsageQuotaBulkResultAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageQuotaBulkResultAttributesWithDefaults() *UsageQuotaBulkResultAttributes {
	this := UsageQuotaBulkResultAttributes{}
	return &this
}

// GetEnforced returns the Enforced field value if set, zero value otherwise.
func (o *UsageQuotaBulkResultAttributes) GetEnforced() bool {
	if o == nil || o.Enforced == nil {
		var ret bool
		return ret
	}
	return *o.Enforced
}

// GetEnforcedOk returns a tuple with the Enforced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageQuotaBulkResultAttributes) GetEnforcedOk() (*bool, bool) {
	if o == nil || o.Enforced == nil {
		return nil, false
	}
	return o.Enforced, true
}

// HasEnforced returns a boolean if a field has been set.
func (o *UsageQuotaBulkResultAttributes) HasEnforced() bool {
	return o != nil && o.Enforced != nil
}

// SetEnforced gets a reference to the given bool and assigns it to the Enforced field.
func (o *UsageQuotaBulkResultAttributes) SetEnforced(v bool) {
	o.Enforced = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *UsageQuotaBulkResultAttributes) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageQuotaBulkResultAttributes) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *UsageQuotaBulkResultAttributes) HasError() bool {
	return o != nil && o.Error != nil
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *UsageQuotaBulkResultAttributes) SetError(v string) {
	o.Error = &v
}

// GetOrgPublicId returns the OrgPublicId field value if set, zero value otherwise.
func (o *UsageQuotaBulkResultAttributes) GetOrgPublicId() string {
	if o == nil || o.OrgPublicId == nil {
		var ret string
		return ret
	}
	return *o.OrgPublicId
}

// GetOrgPublicIdOk returns a tuple with the OrgPublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageQuotaBulkResultAttributes) GetOrgPublicIdOk() (*string, bool) {
	if o == nil || o.OrgPublicId == nil {
		return nil, false
	}
	return o.OrgPublicId, true
}

// HasOrgPublicId returns a boolean if a field has been set.
func (o *UsageQuotaBulkResultAttributes) HasOrgPublicId() bool {
	return o != nil && o.OrgPublicId != nil
}

// SetOrgPublicId gets a reference to the given string and assigns it to the OrgPublicId field.
func (o *UsageQuotaBulkResultAttributes) SetOrgPublicId(v string) {
	o.OrgPublicId = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *UsageQuotaBulkResultAttributes) GetScope() map[string]string {
	if o == nil || o.Scope == nil {
		var ret map[string]string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageQuotaBulkResultAttributes) GetScopeOk() (*map[string]string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return &o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *UsageQuotaBulkResultAttributes) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given map[string]string and assigns it to the Scope field.
func (o *UsageQuotaBulkResultAttributes) SetScope(v map[string]string) {
	o.Scope = v
}

// GetUsageLimit returns the UsageLimit field value if set, zero value otherwise.
func (o *UsageQuotaBulkResultAttributes) GetUsageLimit() float64 {
	if o == nil || o.UsageLimit == nil {
		var ret float64
		return ret
	}
	return *o.UsageLimit
}

// GetUsageLimitOk returns a tuple with the UsageLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageQuotaBulkResultAttributes) GetUsageLimitOk() (*float64, bool) {
	if o == nil || o.UsageLimit == nil {
		return nil, false
	}
	return o.UsageLimit, true
}

// HasUsageLimit returns a boolean if a field has been set.
func (o *UsageQuotaBulkResultAttributes) HasUsageLimit() bool {
	return o != nil && o.UsageLimit != nil
}

// SetUsageLimit gets a reference to the given float64 and assigns it to the UsageLimit field.
func (o *UsageQuotaBulkResultAttributes) SetUsageLimit(v float64) {
	o.UsageLimit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageQuotaBulkResultAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Enforced != nil {
		toSerialize["enforced"] = o.Enforced
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.OrgPublicId != nil {
		toSerialize["org_public_id"] = o.OrgPublicId
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.UsageLimit != nil {
		toSerialize["usage_limit"] = o.UsageLimit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageQuotaBulkResultAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enforced    *bool             `json:"enforced,omitempty"`
		Error       *string           `json:"error,omitempty"`
		OrgPublicId *string           `json:"org_public_id,omitempty"`
		Scope       map[string]string `json:"scope,omitempty"`
		UsageLimit  *float64          `json:"usage_limit,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enforced", "error", "org_public_id", "scope", "usage_limit"})
	} else {
		return err
	}
	o.Enforced = all.Enforced
	o.Error = all.Error
	o.OrgPublicId = all.OrgPublicId
	o.Scope = all.Scope
	o.UsageLimit = all.UsageLimit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
