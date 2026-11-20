// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventStreamQuery Test-only generated schema for the dashboard versioning proof of concept.
type EventStreamQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	CrossOrgUuids []string `json:"cross_org_uuids,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource EventStreamQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	EventSize *EventStreamQueryEventSize `json:"event_size,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Indexes []string `json:"indexes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryString string `json:"query_string"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort *EventStreamQuery20270101Sort `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEventStreamQuery instantiates a new EventStreamQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventStreamQuery(dataSource EventStreamQueryDataSource, queryString string) *EventStreamQuery {
	this := EventStreamQuery{}
	this.DataSource = dataSource
	this.QueryString = queryString
	return &this
}

// NewEventStreamQueryWithDefaults instantiates a new EventStreamQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventStreamQueryWithDefaults() *EventStreamQuery {
	this := EventStreamQuery{}
	return &this
}

// GetCrossOrgUuids returns the CrossOrgUuids field value if set, zero value otherwise.
func (o *EventStreamQuery) GetCrossOrgUuids() []string {
	if o == nil || o.CrossOrgUuids == nil {
		var ret []string
		return ret
	}
	return o.CrossOrgUuids
}

// GetCrossOrgUuidsOk returns a tuple with the CrossOrgUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventStreamQuery) GetCrossOrgUuidsOk() (*[]string, bool) {
	if o == nil || o.CrossOrgUuids == nil {
		return nil, false
	}
	return &o.CrossOrgUuids, true
}

// HasCrossOrgUuids returns a boolean if a field has been set.
func (o *EventStreamQuery) HasCrossOrgUuids() bool {
	return o != nil && o.CrossOrgUuids != nil
}

// SetCrossOrgUuids gets a reference to the given []string and assigns it to the CrossOrgUuids field.
func (o *EventStreamQuery) SetCrossOrgUuids(v []string) {
	o.CrossOrgUuids = v
}

// GetDataSource returns the DataSource field value.
func (o *EventStreamQuery) GetDataSource() EventStreamQueryDataSource {
	if o == nil {
		var ret EventStreamQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *EventStreamQuery) GetDataSourceOk() (*EventStreamQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *EventStreamQuery) SetDataSource(v EventStreamQueryDataSource) {
	o.DataSource = v
}

// GetEventSize returns the EventSize field value if set, zero value otherwise.
func (o *EventStreamQuery) GetEventSize() EventStreamQueryEventSize {
	if o == nil || o.EventSize == nil {
		var ret EventStreamQueryEventSize
		return ret
	}
	return *o.EventSize
}

// GetEventSizeOk returns a tuple with the EventSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventStreamQuery) GetEventSizeOk() (*EventStreamQueryEventSize, bool) {
	if o == nil || o.EventSize == nil {
		return nil, false
	}
	return o.EventSize, true
}

// HasEventSize returns a boolean if a field has been set.
func (o *EventStreamQuery) HasEventSize() bool {
	return o != nil && o.EventSize != nil
}

// SetEventSize gets a reference to the given EventStreamQueryEventSize and assigns it to the EventSize field.
func (o *EventStreamQuery) SetEventSize(v EventStreamQueryEventSize) {
	o.EventSize = &v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *EventStreamQuery) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventStreamQuery) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *EventStreamQuery) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *EventStreamQuery) SetIndexes(v []string) {
	o.Indexes = v
}

// GetQueryString returns the QueryString field value.
func (o *EventStreamQuery) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *EventStreamQuery) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value.
func (o *EventStreamQuery) SetQueryString(v string) {
	o.QueryString = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *EventStreamQuery) GetSort() EventStreamQuery20270101Sort {
	if o == nil || o.Sort == nil {
		var ret EventStreamQuery20270101Sort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventStreamQuery) GetSortOk() (*EventStreamQuery20270101Sort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *EventStreamQuery) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given EventStreamQuery20270101Sort and assigns it to the Sort field.
func (o *EventStreamQuery) SetSort(v EventStreamQuery20270101Sort) {
	o.Sort = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventStreamQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CrossOrgUuids != nil {
		toSerialize["cross_org_uuids"] = o.CrossOrgUuids
	}
	toSerialize["data_source"] = o.DataSource
	if o.EventSize != nil {
		toSerialize["event_size"] = o.EventSize
	}
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	toSerialize["query_string"] = o.QueryString
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventStreamQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CrossOrgUuids []string                      `json:"cross_org_uuids,omitempty"`
		DataSource    *EventStreamQueryDataSource   `json:"data_source"`
		EventSize     *EventStreamQueryEventSize    `json:"event_size,omitempty"`
		Indexes       []string                      `json:"indexes,omitempty"`
		QueryString   *string                       `json:"query_string"`
		Sort          *EventStreamQuery20270101Sort `json:"sort,omitempty"`
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
	if all.EventSize != nil && !all.EventSize.IsValid() {
		hasInvalidField = true
	} else {
		o.EventSize = all.EventSize
	}
	o.Indexes = all.Indexes
	o.QueryString = *all.QueryString
	o.Sort = all.Sort

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
