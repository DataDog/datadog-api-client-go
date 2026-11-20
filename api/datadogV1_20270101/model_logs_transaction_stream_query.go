// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsTransactionStreamQuery Test-only generated schema for the dashboard versioning proof of concept.
type LogsTransactionStreamQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	CalculatedFields []EventsCalculatedField `json:"calculated_fields,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Compute []LogsTransactionCompute `json:"compute,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CrossOrgUuids []string `json:"cross_org_uuids,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource LogsTransactionStreamQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	EndQueryString *string `json:"end_query_string,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy []interface{} `json:"group_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Indexes []string `json:"indexes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryString string `json:"query_string"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort *LogsTransactionStreamQuery20270101Sort `json:"sort,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	StartQueryString *string `json:"start_query_string,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Storage *StorageType `json:"storage,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Version *LogsTransactionStreamQueryVersion `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewLogsTransactionStreamQuery instantiates a new LogsTransactionStreamQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsTransactionStreamQuery(dataSource LogsTransactionStreamQueryDataSource, queryString string) *LogsTransactionStreamQuery {
	this := LogsTransactionStreamQuery{}
	this.DataSource = dataSource
	this.QueryString = queryString
	return &this
}

// NewLogsTransactionStreamQueryWithDefaults instantiates a new LogsTransactionStreamQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsTransactionStreamQueryWithDefaults() *LogsTransactionStreamQuery {
	this := LogsTransactionStreamQuery{}
	return &this
}

// GetCalculatedFields returns the CalculatedFields field value if set, zero value otherwise.
func (o *LogsTransactionStreamQuery) GetCalculatedFields() []EventsCalculatedField {
	if o == nil || o.CalculatedFields == nil {
		var ret []EventsCalculatedField
		return ret
	}
	return o.CalculatedFields
}

// GetCalculatedFieldsOk returns a tuple with the CalculatedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsTransactionStreamQuery) GetCalculatedFieldsOk() (*[]EventsCalculatedField, bool) {
	if o == nil || o.CalculatedFields == nil {
		return nil, false
	}
	return &o.CalculatedFields, true
}

// HasCalculatedFields returns a boolean if a field has been set.
func (o *LogsTransactionStreamQuery) HasCalculatedFields() bool {
	return o != nil && o.CalculatedFields != nil
}

// SetCalculatedFields gets a reference to the given []EventsCalculatedField and assigns it to the CalculatedFields field.
func (o *LogsTransactionStreamQuery) SetCalculatedFields(v []EventsCalculatedField) {
	o.CalculatedFields = v
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *LogsTransactionStreamQuery) GetCompute() []LogsTransactionCompute {
	if o == nil || o.Compute == nil {
		var ret []LogsTransactionCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsTransactionStreamQuery) GetComputeOk() (*[]LogsTransactionCompute, bool) {
	if o == nil || o.Compute == nil {
		return nil, false
	}
	return &o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *LogsTransactionStreamQuery) HasCompute() bool {
	return o != nil && o.Compute != nil
}

// SetCompute gets a reference to the given []LogsTransactionCompute and assigns it to the Compute field.
func (o *LogsTransactionStreamQuery) SetCompute(v []LogsTransactionCompute) {
	o.Compute = v
}

// GetCrossOrgUuids returns the CrossOrgUuids field value if set, zero value otherwise.
func (o *LogsTransactionStreamQuery) GetCrossOrgUuids() []string {
	if o == nil || o.CrossOrgUuids == nil {
		var ret []string
		return ret
	}
	return o.CrossOrgUuids
}

// GetCrossOrgUuidsOk returns a tuple with the CrossOrgUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsTransactionStreamQuery) GetCrossOrgUuidsOk() (*[]string, bool) {
	if o == nil || o.CrossOrgUuids == nil {
		return nil, false
	}
	return &o.CrossOrgUuids, true
}

// HasCrossOrgUuids returns a boolean if a field has been set.
func (o *LogsTransactionStreamQuery) HasCrossOrgUuids() bool {
	return o != nil && o.CrossOrgUuids != nil
}

// SetCrossOrgUuids gets a reference to the given []string and assigns it to the CrossOrgUuids field.
func (o *LogsTransactionStreamQuery) SetCrossOrgUuids(v []string) {
	o.CrossOrgUuids = v
}

// GetDataSource returns the DataSource field value.
func (o *LogsTransactionStreamQuery) GetDataSource() LogsTransactionStreamQueryDataSource {
	if o == nil {
		var ret LogsTransactionStreamQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *LogsTransactionStreamQuery) GetDataSourceOk() (*LogsTransactionStreamQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *LogsTransactionStreamQuery) SetDataSource(v LogsTransactionStreamQueryDataSource) {
	o.DataSource = v
}

// GetEndQueryString returns the EndQueryString field value if set, zero value otherwise.
func (o *LogsTransactionStreamQuery) GetEndQueryString() string {
	if o == nil || o.EndQueryString == nil {
		var ret string
		return ret
	}
	return *o.EndQueryString
}

// GetEndQueryStringOk returns a tuple with the EndQueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsTransactionStreamQuery) GetEndQueryStringOk() (*string, bool) {
	if o == nil || o.EndQueryString == nil {
		return nil, false
	}
	return o.EndQueryString, true
}

// HasEndQueryString returns a boolean if a field has been set.
func (o *LogsTransactionStreamQuery) HasEndQueryString() bool {
	return o != nil && o.EndQueryString != nil
}

// SetEndQueryString gets a reference to the given string and assigns it to the EndQueryString field.
func (o *LogsTransactionStreamQuery) SetEndQueryString(v string) {
	o.EndQueryString = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *LogsTransactionStreamQuery) GetGroupBy() []interface{} {
	if o == nil || o.GroupBy == nil {
		var ret []interface{}
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsTransactionStreamQuery) GetGroupByOk() (*[]interface{}, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *LogsTransactionStreamQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []interface{} and assigns it to the GroupBy field.
func (o *LogsTransactionStreamQuery) SetGroupBy(v []interface{}) {
	o.GroupBy = v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *LogsTransactionStreamQuery) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsTransactionStreamQuery) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *LogsTransactionStreamQuery) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *LogsTransactionStreamQuery) SetIndexes(v []string) {
	o.Indexes = v
}

// GetQueryString returns the QueryString field value.
func (o *LogsTransactionStreamQuery) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *LogsTransactionStreamQuery) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value.
func (o *LogsTransactionStreamQuery) SetQueryString(v string) {
	o.QueryString = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *LogsTransactionStreamQuery) GetSort() LogsTransactionStreamQuery20270101Sort {
	if o == nil || o.Sort == nil {
		var ret LogsTransactionStreamQuery20270101Sort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsTransactionStreamQuery) GetSortOk() (*LogsTransactionStreamQuery20270101Sort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *LogsTransactionStreamQuery) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given LogsTransactionStreamQuery20270101Sort and assigns it to the Sort field.
func (o *LogsTransactionStreamQuery) SetSort(v LogsTransactionStreamQuery20270101Sort) {
	o.Sort = &v
}

// GetStartQueryString returns the StartQueryString field value if set, zero value otherwise.
func (o *LogsTransactionStreamQuery) GetStartQueryString() string {
	if o == nil || o.StartQueryString == nil {
		var ret string
		return ret
	}
	return *o.StartQueryString
}

// GetStartQueryStringOk returns a tuple with the StartQueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsTransactionStreamQuery) GetStartQueryStringOk() (*string, bool) {
	if o == nil || o.StartQueryString == nil {
		return nil, false
	}
	return o.StartQueryString, true
}

// HasStartQueryString returns a boolean if a field has been set.
func (o *LogsTransactionStreamQuery) HasStartQueryString() bool {
	return o != nil && o.StartQueryString != nil
}

// SetStartQueryString gets a reference to the given string and assigns it to the StartQueryString field.
func (o *LogsTransactionStreamQuery) SetStartQueryString(v string) {
	o.StartQueryString = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *LogsTransactionStreamQuery) GetStorage() StorageType {
	if o == nil || o.Storage == nil {
		var ret StorageType
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsTransactionStreamQuery) GetStorageOk() (*StorageType, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *LogsTransactionStreamQuery) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given StorageType and assigns it to the Storage field.
func (o *LogsTransactionStreamQuery) SetStorage(v StorageType) {
	o.Storage = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *LogsTransactionStreamQuery) GetVersion() LogsTransactionStreamQueryVersion {
	if o == nil || o.Version == nil {
		var ret LogsTransactionStreamQueryVersion
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsTransactionStreamQuery) GetVersionOk() (*LogsTransactionStreamQueryVersion, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *LogsTransactionStreamQuery) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given LogsTransactionStreamQueryVersion and assigns it to the Version field.
func (o *LogsTransactionStreamQuery) SetVersion(v LogsTransactionStreamQueryVersion) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsTransactionStreamQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CalculatedFields != nil {
		toSerialize["calculated_fields"] = o.CalculatedFields
	}
	if o.Compute != nil {
		toSerialize["compute"] = o.Compute
	}
	if o.CrossOrgUuids != nil {
		toSerialize["cross_org_uuids"] = o.CrossOrgUuids
	}
	toSerialize["data_source"] = o.DataSource
	if o.EndQueryString != nil {
		toSerialize["end_query_string"] = o.EndQueryString
	}
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
	if o.StartQueryString != nil {
		toSerialize["start_query_string"] = o.StartQueryString
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsTransactionStreamQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CalculatedFields []EventsCalculatedField                 `json:"calculated_fields,omitempty"`
		Compute          []LogsTransactionCompute                `json:"compute,omitempty"`
		CrossOrgUuids    []string                                `json:"cross_org_uuids,omitempty"`
		DataSource       *LogsTransactionStreamQueryDataSource   `json:"data_source"`
		EndQueryString   *string                                 `json:"end_query_string,omitempty"`
		GroupBy          []interface{}                           `json:"group_by,omitempty"`
		Indexes          []string                                `json:"indexes,omitempty"`
		QueryString      *string                                 `json:"query_string"`
		Sort             *LogsTransactionStreamQuery20270101Sort `json:"sort,omitempty"`
		StartQueryString *string                                 `json:"start_query_string,omitempty"`
		Storage          *StorageType                            `json:"storage,omitempty"`
		Version          *LogsTransactionStreamQueryVersion      `json:"version,omitempty"`
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
	o.Compute = all.Compute
	o.CrossOrgUuids = all.CrossOrgUuids
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.EndQueryString = all.EndQueryString
	o.GroupBy = all.GroupBy
	o.Indexes = all.Indexes
	o.QueryString = *all.QueryString
	o.Sort = all.Sort
	o.StartQueryString = all.StartQueryString
	if all.Storage != nil && !all.Storage.IsValid() {
		hasInvalidField = true
	} else {
		o.Storage = all.Storage
	}
	if all.Version != nil && !all.Version.IsValid() {
		hasInvalidField = true
	} else {
		o.Version = all.Version
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
