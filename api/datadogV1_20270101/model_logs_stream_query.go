// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsStreamQuery Test-only generated schema for the dashboard versioning proof of concept.
type LogsStreamQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	CalculatedFields []EventsCalculatedField `json:"calculated_fields,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Context *EventsContext `json:"context,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CrossOrgUuids []string `json:"cross_org_uuids,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource LogsStreamQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Extractions []EventsDataExtraction `json:"extractions,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Indexes []string `json:"indexes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryString string `json:"query_string"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort *LogsStreamQuery20270101Sort `json:"sort,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Storage *StorageType `json:"storage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewLogsStreamQuery instantiates a new LogsStreamQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsStreamQuery(dataSource LogsStreamQueryDataSource, queryString string) *LogsStreamQuery {
	this := LogsStreamQuery{}
	this.DataSource = dataSource
	this.QueryString = queryString
	return &this
}

// NewLogsStreamQueryWithDefaults instantiates a new LogsStreamQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsStreamQueryWithDefaults() *LogsStreamQuery {
	this := LogsStreamQuery{}
	return &this
}

// GetCalculatedFields returns the CalculatedFields field value if set, zero value otherwise.
func (o *LogsStreamQuery) GetCalculatedFields() []EventsCalculatedField {
	if o == nil || o.CalculatedFields == nil {
		var ret []EventsCalculatedField
		return ret
	}
	return o.CalculatedFields
}

// GetCalculatedFieldsOk returns a tuple with the CalculatedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsStreamQuery) GetCalculatedFieldsOk() (*[]EventsCalculatedField, bool) {
	if o == nil || o.CalculatedFields == nil {
		return nil, false
	}
	return &o.CalculatedFields, true
}

// HasCalculatedFields returns a boolean if a field has been set.
func (o *LogsStreamQuery) HasCalculatedFields() bool {
	return o != nil && o.CalculatedFields != nil
}

// SetCalculatedFields gets a reference to the given []EventsCalculatedField and assigns it to the CalculatedFields field.
func (o *LogsStreamQuery) SetCalculatedFields(v []EventsCalculatedField) {
	o.CalculatedFields = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *LogsStreamQuery) GetContext() EventsContext {
	if o == nil || o.Context == nil {
		var ret EventsContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsStreamQuery) GetContextOk() (*EventsContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *LogsStreamQuery) HasContext() bool {
	return o != nil && o.Context != nil
}

// SetContext gets a reference to the given EventsContext and assigns it to the Context field.
func (o *LogsStreamQuery) SetContext(v EventsContext) {
	o.Context = &v
}

// GetCrossOrgUuids returns the CrossOrgUuids field value if set, zero value otherwise.
func (o *LogsStreamQuery) GetCrossOrgUuids() []string {
	if o == nil || o.CrossOrgUuids == nil {
		var ret []string
		return ret
	}
	return o.CrossOrgUuids
}

// GetCrossOrgUuidsOk returns a tuple with the CrossOrgUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsStreamQuery) GetCrossOrgUuidsOk() (*[]string, bool) {
	if o == nil || o.CrossOrgUuids == nil {
		return nil, false
	}
	return &o.CrossOrgUuids, true
}

// HasCrossOrgUuids returns a boolean if a field has been set.
func (o *LogsStreamQuery) HasCrossOrgUuids() bool {
	return o != nil && o.CrossOrgUuids != nil
}

// SetCrossOrgUuids gets a reference to the given []string and assigns it to the CrossOrgUuids field.
func (o *LogsStreamQuery) SetCrossOrgUuids(v []string) {
	o.CrossOrgUuids = v
}

// GetDataSource returns the DataSource field value.
func (o *LogsStreamQuery) GetDataSource() LogsStreamQueryDataSource {
	if o == nil {
		var ret LogsStreamQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *LogsStreamQuery) GetDataSourceOk() (*LogsStreamQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *LogsStreamQuery) SetDataSource(v LogsStreamQueryDataSource) {
	o.DataSource = v
}

// GetExtractions returns the Extractions field value if set, zero value otherwise.
func (o *LogsStreamQuery) GetExtractions() []EventsDataExtraction {
	if o == nil || o.Extractions == nil {
		var ret []EventsDataExtraction
		return ret
	}
	return o.Extractions
}

// GetExtractionsOk returns a tuple with the Extractions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsStreamQuery) GetExtractionsOk() (*[]EventsDataExtraction, bool) {
	if o == nil || o.Extractions == nil {
		return nil, false
	}
	return &o.Extractions, true
}

// HasExtractions returns a boolean if a field has been set.
func (o *LogsStreamQuery) HasExtractions() bool {
	return o != nil && o.Extractions != nil
}

// SetExtractions gets a reference to the given []EventsDataExtraction and assigns it to the Extractions field.
func (o *LogsStreamQuery) SetExtractions(v []EventsDataExtraction) {
	o.Extractions = v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *LogsStreamQuery) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsStreamQuery) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *LogsStreamQuery) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *LogsStreamQuery) SetIndexes(v []string) {
	o.Indexes = v
}

// GetQueryString returns the QueryString field value.
func (o *LogsStreamQuery) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *LogsStreamQuery) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value.
func (o *LogsStreamQuery) SetQueryString(v string) {
	o.QueryString = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *LogsStreamQuery) GetSort() LogsStreamQuery20270101Sort {
	if o == nil || o.Sort == nil {
		var ret LogsStreamQuery20270101Sort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsStreamQuery) GetSortOk() (*LogsStreamQuery20270101Sort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *LogsStreamQuery) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given LogsStreamQuery20270101Sort and assigns it to the Sort field.
func (o *LogsStreamQuery) SetSort(v LogsStreamQuery20270101Sort) {
	o.Sort = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *LogsStreamQuery) GetStorage() StorageType {
	if o == nil || o.Storage == nil {
		var ret StorageType
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsStreamQuery) GetStorageOk() (*StorageType, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *LogsStreamQuery) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given StorageType and assigns it to the Storage field.
func (o *LogsStreamQuery) SetStorage(v StorageType) {
	o.Storage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsStreamQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CalculatedFields != nil {
		toSerialize["calculated_fields"] = o.CalculatedFields
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.CrossOrgUuids != nil {
		toSerialize["cross_org_uuids"] = o.CrossOrgUuids
	}
	toSerialize["data_source"] = o.DataSource
	if o.Extractions != nil {
		toSerialize["extractions"] = o.Extractions
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
func (o *LogsStreamQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CalculatedFields []EventsCalculatedField      `json:"calculated_fields,omitempty"`
		Context          *EventsContext               `json:"context,omitempty"`
		CrossOrgUuids    []string                     `json:"cross_org_uuids,omitempty"`
		DataSource       *LogsStreamQueryDataSource   `json:"data_source"`
		Extractions      []EventsDataExtraction       `json:"extractions,omitempty"`
		Indexes          []string                     `json:"indexes,omitempty"`
		QueryString      *string                      `json:"query_string"`
		Sort             *LogsStreamQuery20270101Sort `json:"sort,omitempty"`
		Storage          *StorageType                 `json:"storage,omitempty"`
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
	if all.Context != nil && all.Context.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Context = all.Context
	o.CrossOrgUuids = all.CrossOrgUuids
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Extractions = all.Extractions
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
