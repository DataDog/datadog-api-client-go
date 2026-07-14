// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardAvailableValuesEventsQuery Query for available values using an events-based data source (spans, logs, or rum).
type DashboardAvailableValuesEventsQuery struct {
	// The events-based data source for the query.
	DataSource DashboardAvailableValuesEventsDataSource `json:"data_source"`
	// The fields to group by in the query.
	GroupBy []DashboardAvailableValuesEventsQueryGroupByItems `json:"group_by"`
	// The search filter for the query.
	Search DashboardAvailableValuesEventsQuerySearch `json:"search"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDashboardAvailableValuesEventsQuery instantiates a new DashboardAvailableValuesEventsQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardAvailableValuesEventsQuery(dataSource DashboardAvailableValuesEventsDataSource, groupBy []DashboardAvailableValuesEventsQueryGroupByItems, search DashboardAvailableValuesEventsQuerySearch) *DashboardAvailableValuesEventsQuery {
	this := DashboardAvailableValuesEventsQuery{}
	this.DataSource = dataSource
	this.GroupBy = groupBy
	this.Search = search
	return &this
}

// NewDashboardAvailableValuesEventsQueryWithDefaults instantiates a new DashboardAvailableValuesEventsQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardAvailableValuesEventsQueryWithDefaults() *DashboardAvailableValuesEventsQuery {
	this := DashboardAvailableValuesEventsQuery{}
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *DashboardAvailableValuesEventsQuery) GetDataSource() DashboardAvailableValuesEventsDataSource {
	if o == nil {
		var ret DashboardAvailableValuesEventsDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *DashboardAvailableValuesEventsQuery) GetDataSourceOk() (*DashboardAvailableValuesEventsDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *DashboardAvailableValuesEventsQuery) SetDataSource(v DashboardAvailableValuesEventsDataSource) {
	o.DataSource = v
}

// GetGroupBy returns the GroupBy field value.
func (o *DashboardAvailableValuesEventsQuery) GetGroupBy() []DashboardAvailableValuesEventsQueryGroupByItems {
	if o == nil {
		var ret []DashboardAvailableValuesEventsQueryGroupByItems
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value
// and a boolean to check if the value has been set.
func (o *DashboardAvailableValuesEventsQuery) GetGroupByOk() (*[]DashboardAvailableValuesEventsQueryGroupByItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// SetGroupBy sets field value.
func (o *DashboardAvailableValuesEventsQuery) SetGroupBy(v []DashboardAvailableValuesEventsQueryGroupByItems) {
	o.GroupBy = v
}

// GetSearch returns the Search field value.
func (o *DashboardAvailableValuesEventsQuery) GetSearch() DashboardAvailableValuesEventsQuerySearch {
	if o == nil {
		var ret DashboardAvailableValuesEventsQuerySearch
		return ret
	}
	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *DashboardAvailableValuesEventsQuery) GetSearchOk() (*DashboardAvailableValuesEventsQuerySearch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value.
func (o *DashboardAvailableValuesEventsQuery) SetSearch(v DashboardAvailableValuesEventsQuerySearch) {
	o.Search = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardAvailableValuesEventsQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["group_by"] = o.GroupBy
	toSerialize["search"] = o.Search
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardAvailableValuesEventsQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource *DashboardAvailableValuesEventsDataSource          `json:"data_source"`
		GroupBy    *[]DashboardAvailableValuesEventsQueryGroupByItems `json:"group_by"`
		Search     *DashboardAvailableValuesEventsQuerySearch         `json:"search"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.GroupBy == nil {
		return fmt.Errorf("required field group_by missing")
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
	o.GroupBy = *all.GroupBy
	if all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = *all.Search

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
