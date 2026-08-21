// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsFormulaJourneyQuery Query definition for a journey timeseries request.
type ProductAnalyticsFormulaJourneyQuery struct {
	// Defines the metric computed over the journey.
	Compute ProductAnalyticsGraphQueryCompute `json:"compute"`
	// Segments the results by the values of one or more facets.
	GroupBy []ProductAnalyticsGraphQueryGroupBy `json:"group_by,omitempty"`
	// Caller-defined identifier echoed back in the results.
	QueryId *string `json:"query_id,omitempty"`
	// Defines the steps of the journey and the filters applied to it.
	Search ProductAnalyticsJourneySearch `json:"search"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsFormulaJourneyQuery instantiates a new ProductAnalyticsFormulaJourneyQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsFormulaJourneyQuery(compute ProductAnalyticsGraphQueryCompute, search ProductAnalyticsJourneySearch) *ProductAnalyticsFormulaJourneyQuery {
	this := ProductAnalyticsFormulaJourneyQuery{}
	this.Compute = compute
	this.Search = search
	return &this
}

// NewProductAnalyticsFormulaJourneyQueryWithDefaults instantiates a new ProductAnalyticsFormulaJourneyQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsFormulaJourneyQueryWithDefaults() *ProductAnalyticsFormulaJourneyQuery {
	this := ProductAnalyticsFormulaJourneyQuery{}
	return &this
}

// GetCompute returns the Compute field value.
func (o *ProductAnalyticsFormulaJourneyQuery) GetCompute() ProductAnalyticsGraphQueryCompute {
	if o == nil {
		var ret ProductAnalyticsGraphQueryCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFormulaJourneyQuery) GetComputeOk() (*ProductAnalyticsGraphQueryCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *ProductAnalyticsFormulaJourneyQuery) SetCompute(v ProductAnalyticsGraphQueryCompute) {
	o.Compute = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *ProductAnalyticsFormulaJourneyQuery) GetGroupBy() []ProductAnalyticsGraphQueryGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []ProductAnalyticsGraphQueryGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFormulaJourneyQuery) GetGroupByOk() (*[]ProductAnalyticsGraphQueryGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *ProductAnalyticsFormulaJourneyQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []ProductAnalyticsGraphQueryGroupBy and assigns it to the GroupBy field.
func (o *ProductAnalyticsFormulaJourneyQuery) SetGroupBy(v []ProductAnalyticsGraphQueryGroupBy) {
	o.GroupBy = v
}

// GetQueryId returns the QueryId field value if set, zero value otherwise.
func (o *ProductAnalyticsFormulaJourneyQuery) GetQueryId() string {
	if o == nil || o.QueryId == nil {
		var ret string
		return ret
	}
	return *o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFormulaJourneyQuery) GetQueryIdOk() (*string, bool) {
	if o == nil || o.QueryId == nil {
		return nil, false
	}
	return o.QueryId, true
}

// HasQueryId returns a boolean if a field has been set.
func (o *ProductAnalyticsFormulaJourneyQuery) HasQueryId() bool {
	return o != nil && o.QueryId != nil
}

// SetQueryId gets a reference to the given string and assigns it to the QueryId field.
func (o *ProductAnalyticsFormulaJourneyQuery) SetQueryId(v string) {
	o.QueryId = &v
}

// GetSearch returns the Search field value.
func (o *ProductAnalyticsFormulaJourneyQuery) GetSearch() ProductAnalyticsJourneySearch {
	if o == nil {
		var ret ProductAnalyticsJourneySearch
		return ret
	}
	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFormulaJourneyQuery) GetSearchOk() (*ProductAnalyticsJourneySearch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value.
func (o *ProductAnalyticsFormulaJourneyQuery) SetSearch(v ProductAnalyticsJourneySearch) {
	o.Search = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsFormulaJourneyQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["compute"] = o.Compute
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	if o.QueryId != nil {
		toSerialize["query_id"] = o.QueryId
	}
	toSerialize["search"] = o.Search

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsFormulaJourneyQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute *ProductAnalyticsGraphQueryCompute  `json:"compute"`
		GroupBy []ProductAnalyticsGraphQueryGroupBy `json:"group_by,omitempty"`
		QueryId *string                             `json:"query_id,omitempty"`
		Search  *ProductAnalyticsJourneySearch      `json:"search"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Compute == nil {
		return fmt.Errorf("required field compute missing")
	}
	if all.Search == nil {
		return fmt.Errorf("required field search missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"compute", "group_by", "query_id", "search"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Compute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compute = *all.Compute
	o.GroupBy = all.GroupBy
	o.QueryId = all.QueryId
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
