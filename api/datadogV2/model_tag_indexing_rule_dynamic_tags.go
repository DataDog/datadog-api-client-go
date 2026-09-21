// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagIndexingRuleDynamicTags Options for dynamic tag indexing applied per metric, such as tags filtered by query usage.
//
// Before a tag key is dropped by this rule, two grace period conditions must be met:
//
// 1. The metric must be submitted for at least as long as the selected window.
// 2. A tag key must have been submitted for at least 15 days.
//
// Any metric or tag key that does not meet these conditions are excluded from this
// indexing rule. The `exclude_not_*` fields require `exclude_tags_mode` to be set to `true`.
type TagIndexingRuleDynamicTags struct {
	// Tags that have not been queried within this window are excluded from indexing. Maximum of `7776000` (90 days).
	ExcludeNotQueriedWindowSeconds *int64 `json:"exclude_not_queried_window_seconds,omitempty"`
	// Tags not used in any dashboards,  monitors, notebooks, or SLOs are excluded from indexing.
	ExcludeNotUsedInAssets *bool `json:"exclude_not_used_in_assets,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagIndexingRuleDynamicTags instantiates a new TagIndexingRuleDynamicTags object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagIndexingRuleDynamicTags() *TagIndexingRuleDynamicTags {
	this := TagIndexingRuleDynamicTags{}
	return &this
}

// NewTagIndexingRuleDynamicTagsWithDefaults instantiates a new TagIndexingRuleDynamicTags object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagIndexingRuleDynamicTagsWithDefaults() *TagIndexingRuleDynamicTags {
	this := TagIndexingRuleDynamicTags{}
	return &this
}

// GetExcludeNotQueriedWindowSeconds returns the ExcludeNotQueriedWindowSeconds field value if set, zero value otherwise.
func (o *TagIndexingRuleDynamicTags) GetExcludeNotQueriedWindowSeconds() int64 {
	if o == nil || o.ExcludeNotQueriedWindowSeconds == nil {
		var ret int64
		return ret
	}
	return *o.ExcludeNotQueriedWindowSeconds
}

// GetExcludeNotQueriedWindowSecondsOk returns a tuple with the ExcludeNotQueriedWindowSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleDynamicTags) GetExcludeNotQueriedWindowSecondsOk() (*int64, bool) {
	if o == nil || o.ExcludeNotQueriedWindowSeconds == nil {
		return nil, false
	}
	return o.ExcludeNotQueriedWindowSeconds, true
}

// HasExcludeNotQueriedWindowSeconds returns a boolean if a field has been set.
func (o *TagIndexingRuleDynamicTags) HasExcludeNotQueriedWindowSeconds() bool {
	return o != nil && o.ExcludeNotQueriedWindowSeconds != nil
}

// SetExcludeNotQueriedWindowSeconds gets a reference to the given int64 and assigns it to the ExcludeNotQueriedWindowSeconds field.
func (o *TagIndexingRuleDynamicTags) SetExcludeNotQueriedWindowSeconds(v int64) {
	o.ExcludeNotQueriedWindowSeconds = &v
}

// GetExcludeNotUsedInAssets returns the ExcludeNotUsedInAssets field value if set, zero value otherwise.
func (o *TagIndexingRuleDynamicTags) GetExcludeNotUsedInAssets() bool {
	if o == nil || o.ExcludeNotUsedInAssets == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeNotUsedInAssets
}

// GetExcludeNotUsedInAssetsOk returns a tuple with the ExcludeNotUsedInAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleDynamicTags) GetExcludeNotUsedInAssetsOk() (*bool, bool) {
	if o == nil || o.ExcludeNotUsedInAssets == nil {
		return nil, false
	}
	return o.ExcludeNotUsedInAssets, true
}

// HasExcludeNotUsedInAssets returns a boolean if a field has been set.
func (o *TagIndexingRuleDynamicTags) HasExcludeNotUsedInAssets() bool {
	return o != nil && o.ExcludeNotUsedInAssets != nil
}

// SetExcludeNotUsedInAssets gets a reference to the given bool and assigns it to the ExcludeNotUsedInAssets field.
func (o *TagIndexingRuleDynamicTags) SetExcludeNotUsedInAssets(v bool) {
	o.ExcludeNotUsedInAssets = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagIndexingRuleDynamicTags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExcludeNotQueriedWindowSeconds != nil {
		toSerialize["exclude_not_queried_window_seconds"] = o.ExcludeNotQueriedWindowSeconds
	}
	if o.ExcludeNotUsedInAssets != nil {
		toSerialize["exclude_not_used_in_assets"] = o.ExcludeNotUsedInAssets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TagIndexingRuleDynamicTags) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExcludeNotQueriedWindowSeconds *int64 `json:"exclude_not_queried_window_seconds,omitempty"`
		ExcludeNotUsedInAssets         *bool  `json:"exclude_not_used_in_assets,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"exclude_not_queried_window_seconds", "exclude_not_used_in_assets"})
	} else {
		return err
	}
	o.ExcludeNotQueriedWindowSeconds = all.ExcludeNotQueriedWindowSeconds
	o.ExcludeNotUsedInAssets = all.ExcludeNotUsedInAssets

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
