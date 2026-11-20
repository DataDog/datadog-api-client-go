// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsPatternStreamQuery Test-only generated schema for the dashboard versioning proof of concept.
type LogsPatternStreamQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	CalculatedFields []EventsCalculatedField `json:"calculated_fields,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ClusteringPatternFieldPath *string `json:"clustering_pattern_field_path,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CrossOrgUuids []string `json:"cross_org_uuids,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource LogsPatternStreamQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy []LogsPatternGroupBy `json:"group_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Indexes []string `json:"indexes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryString string `json:"query_string"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort *LogsPatternStreamQuery20270101Sort `json:"sort,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Storage *StorageType `json:"storage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewLogsPatternStreamQuery instantiates a new LogsPatternStreamQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsPatternStreamQuery(dataSource LogsPatternStreamQueryDataSource, queryString string) *LogsPatternStreamQuery {
	this := LogsPatternStreamQuery{}
	this.DataSource = dataSource
	this.QueryString = queryString
	return &this
}

// NewLogsPatternStreamQueryWithDefaults instantiates a new LogsPatternStreamQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsPatternStreamQueryWithDefaults() *LogsPatternStreamQuery {
	this := LogsPatternStreamQuery{}
	return &this
}

// GetCalculatedFields returns the CalculatedFields field value if set, zero value otherwise.
func (o *LogsPatternStreamQuery) GetCalculatedFields() []EventsCalculatedField {
	if o == nil || o.CalculatedFields == nil {
		var ret []EventsCalculatedField
		return ret
	}
	return o.CalculatedFields
}

// GetCalculatedFieldsOk returns a tuple with the CalculatedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsPatternStreamQuery) GetCalculatedFieldsOk() (*[]EventsCalculatedField, bool) {
	if o == nil || o.CalculatedFields == nil {
		return nil, false
	}
	return &o.CalculatedFields, true
}

// HasCalculatedFields returns a boolean if a field has been set.
func (o *LogsPatternStreamQuery) HasCalculatedFields() bool {
	return o != nil && o.CalculatedFields != nil
}

// SetCalculatedFields gets a reference to the given []EventsCalculatedField and assigns it to the CalculatedFields field.
func (o *LogsPatternStreamQuery) SetCalculatedFields(v []EventsCalculatedField) {
	o.CalculatedFields = v
}

// GetClusteringPatternFieldPath returns the ClusteringPatternFieldPath field value if set, zero value otherwise.
func (o *LogsPatternStreamQuery) GetClusteringPatternFieldPath() string {
	if o == nil || o.ClusteringPatternFieldPath == nil {
		var ret string
		return ret
	}
	return *o.ClusteringPatternFieldPath
}

// GetClusteringPatternFieldPathOk returns a tuple with the ClusteringPatternFieldPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsPatternStreamQuery) GetClusteringPatternFieldPathOk() (*string, bool) {
	if o == nil || o.ClusteringPatternFieldPath == nil {
		return nil, false
	}
	return o.ClusteringPatternFieldPath, true
}

// HasClusteringPatternFieldPath returns a boolean if a field has been set.
func (o *LogsPatternStreamQuery) HasClusteringPatternFieldPath() bool {
	return o != nil && o.ClusteringPatternFieldPath != nil
}

// SetClusteringPatternFieldPath gets a reference to the given string and assigns it to the ClusteringPatternFieldPath field.
func (o *LogsPatternStreamQuery) SetClusteringPatternFieldPath(v string) {
	o.ClusteringPatternFieldPath = &v
}

// GetCrossOrgUuids returns the CrossOrgUuids field value if set, zero value otherwise.
func (o *LogsPatternStreamQuery) GetCrossOrgUuids() []string {
	if o == nil || o.CrossOrgUuids == nil {
		var ret []string
		return ret
	}
	return o.CrossOrgUuids
}

// GetCrossOrgUuidsOk returns a tuple with the CrossOrgUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsPatternStreamQuery) GetCrossOrgUuidsOk() (*[]string, bool) {
	if o == nil || o.CrossOrgUuids == nil {
		return nil, false
	}
	return &o.CrossOrgUuids, true
}

// HasCrossOrgUuids returns a boolean if a field has been set.
func (o *LogsPatternStreamQuery) HasCrossOrgUuids() bool {
	return o != nil && o.CrossOrgUuids != nil
}

// SetCrossOrgUuids gets a reference to the given []string and assigns it to the CrossOrgUuids field.
func (o *LogsPatternStreamQuery) SetCrossOrgUuids(v []string) {
	o.CrossOrgUuids = v
}

// GetDataSource returns the DataSource field value.
func (o *LogsPatternStreamQuery) GetDataSource() LogsPatternStreamQueryDataSource {
	if o == nil {
		var ret LogsPatternStreamQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *LogsPatternStreamQuery) GetDataSourceOk() (*LogsPatternStreamQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *LogsPatternStreamQuery) SetDataSource(v LogsPatternStreamQueryDataSource) {
	o.DataSource = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *LogsPatternStreamQuery) GetGroupBy() []LogsPatternGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []LogsPatternGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsPatternStreamQuery) GetGroupByOk() (*[]LogsPatternGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *LogsPatternStreamQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []LogsPatternGroupBy and assigns it to the GroupBy field.
func (o *LogsPatternStreamQuery) SetGroupBy(v []LogsPatternGroupBy) {
	o.GroupBy = v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *LogsPatternStreamQuery) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsPatternStreamQuery) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *LogsPatternStreamQuery) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *LogsPatternStreamQuery) SetIndexes(v []string) {
	o.Indexes = v
}

// GetQueryString returns the QueryString field value.
func (o *LogsPatternStreamQuery) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *LogsPatternStreamQuery) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value.
func (o *LogsPatternStreamQuery) SetQueryString(v string) {
	o.QueryString = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *LogsPatternStreamQuery) GetSort() LogsPatternStreamQuery20270101Sort {
	if o == nil || o.Sort == nil {
		var ret LogsPatternStreamQuery20270101Sort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsPatternStreamQuery) GetSortOk() (*LogsPatternStreamQuery20270101Sort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *LogsPatternStreamQuery) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given LogsPatternStreamQuery20270101Sort and assigns it to the Sort field.
func (o *LogsPatternStreamQuery) SetSort(v LogsPatternStreamQuery20270101Sort) {
	o.Sort = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *LogsPatternStreamQuery) GetStorage() StorageType {
	if o == nil || o.Storage == nil {
		var ret StorageType
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsPatternStreamQuery) GetStorageOk() (*StorageType, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *LogsPatternStreamQuery) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given StorageType and assigns it to the Storage field.
func (o *LogsPatternStreamQuery) SetStorage(v StorageType) {
	o.Storage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsPatternStreamQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CalculatedFields != nil {
		toSerialize["calculated_fields"] = o.CalculatedFields
	}
	if o.ClusteringPatternFieldPath != nil {
		toSerialize["clustering_pattern_field_path"] = o.ClusteringPatternFieldPath
	}
	if o.CrossOrgUuids != nil {
		toSerialize["cross_org_uuids"] = o.CrossOrgUuids
	}
	toSerialize["data_source"] = o.DataSource
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
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
func (o *LogsPatternStreamQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CalculatedFields           []EventsCalculatedField             `json:"calculated_fields,omitempty"`
		ClusteringPatternFieldPath *string                             `json:"clustering_pattern_field_path,omitempty"`
		CrossOrgUuids              []string                            `json:"cross_org_uuids,omitempty"`
		DataSource                 *LogsPatternStreamQueryDataSource   `json:"data_source"`
		GroupBy                    []LogsPatternGroupBy                `json:"group_by,omitempty"`
		Indexes                    []string                            `json:"indexes,omitempty"`
		QueryString                *string                             `json:"query_string"`
		Sort                       *LogsPatternStreamQuery20270101Sort `json:"sort,omitempty"`
		Storage                    *StorageType                        `json:"storage,omitempty"`
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
	o.CalculatedFields = all.CalculatedFields
	o.ClusteringPatternFieldPath = all.ClusteringPatternFieldPath
	o.CrossOrgUuids = all.CrossOrgUuids
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.GroupBy = all.GroupBy
	o.Indexes = all.Indexes
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
