// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsEventQuery Test-only generated schema for the dashboard versioning proof of concept.
type ProductAnalyticsEventQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource ProductAnalyticsEventQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Search ProductAnalyticsEventQuerySearch `json:"search"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewProductAnalyticsEventQuery instantiates a new ProductAnalyticsEventQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsEventQuery(dataSource ProductAnalyticsEventQueryDataSource, search ProductAnalyticsEventQuerySearch) *ProductAnalyticsEventQuery {
	this := ProductAnalyticsEventQuery{}
	this.DataSource = dataSource
	this.Search = search
	return &this
}

// NewProductAnalyticsEventQueryWithDefaults instantiates a new ProductAnalyticsEventQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsEventQueryWithDefaults() *ProductAnalyticsEventQuery {
	this := ProductAnalyticsEventQuery{}
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *ProductAnalyticsEventQuery) GetDataSource() ProductAnalyticsEventQueryDataSource {
	if o == nil {
		var ret ProductAnalyticsEventQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsEventQuery) GetDataSourceOk() (*ProductAnalyticsEventQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *ProductAnalyticsEventQuery) SetDataSource(v ProductAnalyticsEventQueryDataSource) {
	o.DataSource = v
}

// GetSearch returns the Search field value.
func (o *ProductAnalyticsEventQuery) GetSearch() ProductAnalyticsEventQuerySearch {
	if o == nil {
		var ret ProductAnalyticsEventQuerySearch
		return ret
	}
	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsEventQuery) GetSearchOk() (*ProductAnalyticsEventQuerySearch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value.
func (o *ProductAnalyticsEventQuery) SetSearch(v ProductAnalyticsEventQuerySearch) {
	o.Search = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsEventQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["search"] = o.Search
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsEventQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource *ProductAnalyticsEventQueryDataSource `json:"data_source"`
		Search     *ProductAnalyticsEventQuerySearch     `json:"search"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Search == nil {
		return fmt.Errorf("required field search missing")
	}

	hasInvalidField := false
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	if all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = *all.Search

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
