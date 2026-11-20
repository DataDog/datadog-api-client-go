// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WatchdogAlertsQuery Test-only generated schema for the dashboard versioning proof of concept.
type WatchdogAlertsQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	CrossOrgUuids []string `json:"cross_org_uuids,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource WatchdogAlertsQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryString string `json:"query_string"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort *WatchdogAlertsQuery20270101Sort `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewWatchdogAlertsQuery instantiates a new WatchdogAlertsQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWatchdogAlertsQuery(dataSource WatchdogAlertsQueryDataSource, queryString string) *WatchdogAlertsQuery {
	this := WatchdogAlertsQuery{}
	this.DataSource = dataSource
	this.QueryString = queryString
	return &this
}

// NewWatchdogAlertsQueryWithDefaults instantiates a new WatchdogAlertsQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWatchdogAlertsQueryWithDefaults() *WatchdogAlertsQuery {
	this := WatchdogAlertsQuery{}
	return &this
}

// GetCrossOrgUuids returns the CrossOrgUuids field value if set, zero value otherwise.
func (o *WatchdogAlertsQuery) GetCrossOrgUuids() []string {
	if o == nil || o.CrossOrgUuids == nil {
		var ret []string
		return ret
	}
	return o.CrossOrgUuids
}

// GetCrossOrgUuidsOk returns a tuple with the CrossOrgUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchdogAlertsQuery) GetCrossOrgUuidsOk() (*[]string, bool) {
	if o == nil || o.CrossOrgUuids == nil {
		return nil, false
	}
	return &o.CrossOrgUuids, true
}

// HasCrossOrgUuids returns a boolean if a field has been set.
func (o *WatchdogAlertsQuery) HasCrossOrgUuids() bool {
	return o != nil && o.CrossOrgUuids != nil
}

// SetCrossOrgUuids gets a reference to the given []string and assigns it to the CrossOrgUuids field.
func (o *WatchdogAlertsQuery) SetCrossOrgUuids(v []string) {
	o.CrossOrgUuids = v
}

// GetDataSource returns the DataSource field value.
func (o *WatchdogAlertsQuery) GetDataSource() WatchdogAlertsQueryDataSource {
	if o == nil {
		var ret WatchdogAlertsQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *WatchdogAlertsQuery) GetDataSourceOk() (*WatchdogAlertsQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *WatchdogAlertsQuery) SetDataSource(v WatchdogAlertsQueryDataSource) {
	o.DataSource = v
}

// GetQueryString returns the QueryString field value.
func (o *WatchdogAlertsQuery) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *WatchdogAlertsQuery) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value.
func (o *WatchdogAlertsQuery) SetQueryString(v string) {
	o.QueryString = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *WatchdogAlertsQuery) GetSort() WatchdogAlertsQuery20270101Sort {
	if o == nil || o.Sort == nil {
		var ret WatchdogAlertsQuery20270101Sort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchdogAlertsQuery) GetSortOk() (*WatchdogAlertsQuery20270101Sort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *WatchdogAlertsQuery) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given WatchdogAlertsQuery20270101Sort and assigns it to the Sort field.
func (o *WatchdogAlertsQuery) SetSort(v WatchdogAlertsQuery20270101Sort) {
	o.Sort = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o WatchdogAlertsQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CrossOrgUuids != nil {
		toSerialize["cross_org_uuids"] = o.CrossOrgUuids
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["query_string"] = o.QueryString
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WatchdogAlertsQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CrossOrgUuids []string                         `json:"cross_org_uuids,omitempty"`
		DataSource    *WatchdogAlertsQueryDataSource   `json:"data_source"`
		QueryString   *string                          `json:"query_string"`
		Sort          *WatchdogAlertsQuery20270101Sort `json:"sort,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.QueryString == nil {
		return fmt.Errorf("required field query_string missing")
	}

	hasInvalidField := false
	o.CrossOrgUuids = all.CrossOrgUuids
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.QueryString = *all.QueryString
	o.Sort = all.Sort

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
