// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventsPlatformListStreamQuery Test-only generated schema for the dashboard versioning proof of concept.
type EventsPlatformListStreamQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	CrossOrgUuids []string `json:"cross_org_uuids,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource EventPlatformStreamDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Indexes []string `json:"indexes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	QuerySnapshot *bool `json:"query_snapshot,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryString string `json:"query_string"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort *EventsPlatformListStreamQuery20270101Sort `json:"sort,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Storage *StorageType `json:"storage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEventsPlatformListStreamQuery instantiates a new EventsPlatformListStreamQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventsPlatformListStreamQuery(dataSource EventPlatformStreamDataSource, queryString string) *EventsPlatformListStreamQuery {
	this := EventsPlatformListStreamQuery{}
	this.DataSource = dataSource
	this.QueryString = queryString
	return &this
}

// NewEventsPlatformListStreamQueryWithDefaults instantiates a new EventsPlatformListStreamQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventsPlatformListStreamQueryWithDefaults() *EventsPlatformListStreamQuery {
	this := EventsPlatformListStreamQuery{}
	return &this
}

// GetCrossOrgUuids returns the CrossOrgUuids field value if set, zero value otherwise.
func (o *EventsPlatformListStreamQuery) GetCrossOrgUuids() []string {
	if o == nil || o.CrossOrgUuids == nil {
		var ret []string
		return ret
	}
	return o.CrossOrgUuids
}

// GetCrossOrgUuidsOk returns a tuple with the CrossOrgUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsPlatformListStreamQuery) GetCrossOrgUuidsOk() (*[]string, bool) {
	if o == nil || o.CrossOrgUuids == nil {
		return nil, false
	}
	return &o.CrossOrgUuids, true
}

// HasCrossOrgUuids returns a boolean if a field has been set.
func (o *EventsPlatformListStreamQuery) HasCrossOrgUuids() bool {
	return o != nil && o.CrossOrgUuids != nil
}

// SetCrossOrgUuids gets a reference to the given []string and assigns it to the CrossOrgUuids field.
func (o *EventsPlatformListStreamQuery) SetCrossOrgUuids(v []string) {
	o.CrossOrgUuids = v
}

// GetDataSource returns the DataSource field value.
func (o *EventsPlatformListStreamQuery) GetDataSource() EventPlatformStreamDataSource {
	if o == nil {
		var ret EventPlatformStreamDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *EventsPlatformListStreamQuery) GetDataSourceOk() (*EventPlatformStreamDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *EventsPlatformListStreamQuery) SetDataSource(v EventPlatformStreamDataSource) {
	o.DataSource = v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *EventsPlatformListStreamQuery) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsPlatformListStreamQuery) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *EventsPlatformListStreamQuery) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *EventsPlatformListStreamQuery) SetIndexes(v []string) {
	o.Indexes = v
}

// GetQuerySnapshot returns the QuerySnapshot field value if set, zero value otherwise.
func (o *EventsPlatformListStreamQuery) GetQuerySnapshot() bool {
	if o == nil || o.QuerySnapshot == nil {
		var ret bool
		return ret
	}
	return *o.QuerySnapshot
}

// GetQuerySnapshotOk returns a tuple with the QuerySnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsPlatformListStreamQuery) GetQuerySnapshotOk() (*bool, bool) {
	if o == nil || o.QuerySnapshot == nil {
		return nil, false
	}
	return o.QuerySnapshot, true
}

// HasQuerySnapshot returns a boolean if a field has been set.
func (o *EventsPlatformListStreamQuery) HasQuerySnapshot() bool {
	return o != nil && o.QuerySnapshot != nil
}

// SetQuerySnapshot gets a reference to the given bool and assigns it to the QuerySnapshot field.
func (o *EventsPlatformListStreamQuery) SetQuerySnapshot(v bool) {
	o.QuerySnapshot = &v
}

// GetQueryString returns the QueryString field value.
func (o *EventsPlatformListStreamQuery) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *EventsPlatformListStreamQuery) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value.
func (o *EventsPlatformListStreamQuery) SetQueryString(v string) {
	o.QueryString = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *EventsPlatformListStreamQuery) GetSort() EventsPlatformListStreamQuery20270101Sort {
	if o == nil || o.Sort == nil {
		var ret EventsPlatformListStreamQuery20270101Sort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsPlatformListStreamQuery) GetSortOk() (*EventsPlatformListStreamQuery20270101Sort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *EventsPlatformListStreamQuery) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given EventsPlatformListStreamQuery20270101Sort and assigns it to the Sort field.
func (o *EventsPlatformListStreamQuery) SetSort(v EventsPlatformListStreamQuery20270101Sort) {
	o.Sort = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *EventsPlatformListStreamQuery) GetStorage() StorageType {
	if o == nil || o.Storage == nil {
		var ret StorageType
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsPlatformListStreamQuery) GetStorageOk() (*StorageType, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *EventsPlatformListStreamQuery) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given StorageType and assigns it to the Storage field.
func (o *EventsPlatformListStreamQuery) SetStorage(v StorageType) {
	o.Storage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventsPlatformListStreamQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CrossOrgUuids != nil {
		toSerialize["cross_org_uuids"] = o.CrossOrgUuids
	}
	toSerialize["data_source"] = o.DataSource
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	if o.QuerySnapshot != nil {
		toSerialize["query_snapshot"] = o.QuerySnapshot
	}
	toSerialize["query_string"] = o.QueryString
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventsPlatformListStreamQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CrossOrgUuids []string                                   `json:"cross_org_uuids,omitempty"`
		DataSource    *EventPlatformStreamDataSource             `json:"data_source"`
		Indexes       []string                                   `json:"indexes,omitempty"`
		QuerySnapshot *bool                                      `json:"query_snapshot,omitempty"`
		QueryString   *string                                    `json:"query_string"`
		Sort          *EventsPlatformListStreamQuery20270101Sort `json:"sort,omitempty"`
		Storage       *StorageType                               `json:"storage,omitempty"`
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
	o.Indexes = all.Indexes
	o.QuerySnapshot = all.QuerySnapshot
	o.QueryString = *all.QueryString
	o.Sort = all.Sort
	if all.Storage != nil && !all.Storage.IsValid() {
		hasInvalidField = true
	} else {
		o.Storage = all.Storage
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
