// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyFunnelQuery Query definition for a journey funnel request.
type ProductAnalyticsJourneyFunnelQuery struct {
	// Defines the metric computed at each funnel step.
	Compute *ProductAnalyticsJourneyFunnelCompute `json:"compute,omitempty"`
	// Segments the funnel by the values of one or more facets.
	GroupBy []ProductAnalyticsGraphQueryGroupBy `json:"group_by,omitempty"`
	// Defines the steps of the journey and the filters applied to it.
	Search ProductAnalyticsJourneySearch `json:"search"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsJourneyFunnelQuery instantiates a new ProductAnalyticsJourneyFunnelQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsJourneyFunnelQuery(search ProductAnalyticsJourneySearch) *ProductAnalyticsJourneyFunnelQuery {
	this := ProductAnalyticsJourneyFunnelQuery{}
	this.Search = search
	return &this
}

// NewProductAnalyticsJourneyFunnelQueryWithDefaults instantiates a new ProductAnalyticsJourneyFunnelQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsJourneyFunnelQueryWithDefaults() *ProductAnalyticsJourneyFunnelQuery {
	this := ProductAnalyticsJourneyFunnelQuery{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyFunnelQuery) GetCompute() ProductAnalyticsJourneyFunnelCompute {
	if o == nil || o.Compute == nil {
		var ret ProductAnalyticsJourneyFunnelCompute
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyFunnelQuery) GetComputeOk() (*ProductAnalyticsJourneyFunnelCompute, bool) {
	if o == nil || o.Compute == nil {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyFunnelQuery) HasCompute() bool {
	return o != nil && o.Compute != nil
}

// SetCompute gets a reference to the given ProductAnalyticsJourneyFunnelCompute and assigns it to the Compute field.
func (o *ProductAnalyticsJourneyFunnelQuery) SetCompute(v ProductAnalyticsJourneyFunnelCompute) {
	o.Compute = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyFunnelQuery) GetGroupBy() []ProductAnalyticsGraphQueryGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []ProductAnalyticsGraphQueryGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyFunnelQuery) GetGroupByOk() (*[]ProductAnalyticsGraphQueryGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyFunnelQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []ProductAnalyticsGraphQueryGroupBy and assigns it to the GroupBy field.
func (o *ProductAnalyticsJourneyFunnelQuery) SetGroupBy(v []ProductAnalyticsGraphQueryGroupBy) {
	o.GroupBy = v
}

// GetSearch returns the Search field value.
func (o *ProductAnalyticsJourneyFunnelQuery) GetSearch() ProductAnalyticsJourneySearch {
	if o == nil {
		var ret ProductAnalyticsJourneySearch
		return ret
	}
	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyFunnelQuery) GetSearchOk() (*ProductAnalyticsJourneySearch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value.
func (o *ProductAnalyticsJourneyFunnelQuery) SetSearch(v ProductAnalyticsJourneySearch) {
	o.Search = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsJourneyFunnelQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Compute != nil {
		toSerialize["compute"] = o.Compute
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	toSerialize["search"] = o.Search

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsJourneyFunnelQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute *ProductAnalyticsJourneyFunnelCompute `json:"compute,omitempty"`
		GroupBy []ProductAnalyticsGraphQueryGroupBy   `json:"group_by,omitempty"`
		Search  *ProductAnalyticsJourneySearch        `json:"search"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Search == nil {
		return fmt.Errorf("required field search missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"compute", "group_by", "search"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Compute != nil && all.Compute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compute = all.Compute
	o.GroupBy = all.GroupBy
	if all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = *all.Search

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
