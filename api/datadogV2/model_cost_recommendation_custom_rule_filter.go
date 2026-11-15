// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostRecommendationCustomRuleFilter A tag filter. A rule applies to a resource only when every filter of the rule matches.
type CostRecommendationCustomRuleFilter struct {
	// How the filter compares the tag. `is`, `is_not`, `contains`, `does_not_contain`, `like` (glob), and `not_like`
	// compare the tag values with `values`. `all` matches resources that have the tag and `untagged` matches resources
	// that do not, so both ignore `values`.
	FilterType CostRecommendationCustomRuleFilterType `json:"filter_type"`
	// The tag key the filter compares.
	Tag string `json:"tag"`
	// Tag values the filter compares. Required unless `filter_type` is `all` or `untagged`.
	Values []string `json:"values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostRecommendationCustomRuleFilter instantiates a new CostRecommendationCustomRuleFilter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostRecommendationCustomRuleFilter(filterType CostRecommendationCustomRuleFilterType, tag string) *CostRecommendationCustomRuleFilter {
	this := CostRecommendationCustomRuleFilter{}
	this.FilterType = filterType
	this.Tag = tag
	return &this
}

// NewCostRecommendationCustomRuleFilterWithDefaults instantiates a new CostRecommendationCustomRuleFilter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostRecommendationCustomRuleFilterWithDefaults() *CostRecommendationCustomRuleFilter {
	this := CostRecommendationCustomRuleFilter{}
	return &this
}

// GetFilterType returns the FilterType field value.
func (o *CostRecommendationCustomRuleFilter) GetFilterType() CostRecommendationCustomRuleFilterType {
	if o == nil {
		var ret CostRecommendationCustomRuleFilterType
		return ret
	}
	return o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationCustomRuleFilter) GetFilterTypeOk() (*CostRecommendationCustomRuleFilterType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterType, true
}

// SetFilterType sets field value.
func (o *CostRecommendationCustomRuleFilter) SetFilterType(v CostRecommendationCustomRuleFilterType) {
	o.FilterType = v
}

// GetTag returns the Tag field value.
func (o *CostRecommendationCustomRuleFilter) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationCustomRuleFilter) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value.
func (o *CostRecommendationCustomRuleFilter) SetTag(v string) {
	o.Tag = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *CostRecommendationCustomRuleFilter) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationCustomRuleFilter) GetValuesOk() (*[]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *CostRecommendationCustomRuleFilter) HasValues() bool {
	return o != nil && o.Values != nil
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *CostRecommendationCustomRuleFilter) SetValues(v []string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostRecommendationCustomRuleFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["filter_type"] = o.FilterType
	toSerialize["tag"] = o.Tag
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostRecommendationCustomRuleFilter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FilterType *CostRecommendationCustomRuleFilterType `json:"filter_type"`
		Tag        *string                                 `json:"tag"`
		Values     []string                                `json:"values,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FilterType == nil {
		return fmt.Errorf("required field filter_type missing")
	}
	if all.Tag == nil {
		return fmt.Errorf("required field tag missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"filter_type", "tag", "values"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.FilterType.IsValid() {
		hasInvalidField = true
	} else {
		o.FilterType = *all.FilterType
	}
	o.Tag = *all.Tag
	o.Values = all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
