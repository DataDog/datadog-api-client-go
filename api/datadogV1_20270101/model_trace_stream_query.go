// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TraceStreamQuery Test-only generated schema for the dashboard versioning proof of concept.
type TraceStreamQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Context *EventsContext `json:"context,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CrossOrgUuids []string `json:"cross_org_uuids,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource TraceStreamQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Indexes []string `json:"indexes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryString string `json:"query_string"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SemanticMode *FormulaSemanticMode `json:"semantic_mode,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort *TraceStreamQuery20270101Sort `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTraceStreamQuery instantiates a new TraceStreamQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTraceStreamQuery(dataSource TraceStreamQueryDataSource, queryString string) *TraceStreamQuery {
	this := TraceStreamQuery{}
	this.DataSource = dataSource
	this.QueryString = queryString
	return &this
}

// NewTraceStreamQueryWithDefaults instantiates a new TraceStreamQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTraceStreamQueryWithDefaults() *TraceStreamQuery {
	this := TraceStreamQuery{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *TraceStreamQuery) GetContext() EventsContext {
	if o == nil || o.Context == nil {
		var ret EventsContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceStreamQuery) GetContextOk() (*EventsContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *TraceStreamQuery) HasContext() bool {
	return o != nil && o.Context != nil
}

// SetContext gets a reference to the given EventsContext and assigns it to the Context field.
func (o *TraceStreamQuery) SetContext(v EventsContext) {
	o.Context = &v
}

// GetCrossOrgUuids returns the CrossOrgUuids field value if set, zero value otherwise.
func (o *TraceStreamQuery) GetCrossOrgUuids() []string {
	if o == nil || o.CrossOrgUuids == nil {
		var ret []string
		return ret
	}
	return o.CrossOrgUuids
}

// GetCrossOrgUuidsOk returns a tuple with the CrossOrgUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceStreamQuery) GetCrossOrgUuidsOk() (*[]string, bool) {
	if o == nil || o.CrossOrgUuids == nil {
		return nil, false
	}
	return &o.CrossOrgUuids, true
}

// HasCrossOrgUuids returns a boolean if a field has been set.
func (o *TraceStreamQuery) HasCrossOrgUuids() bool {
	return o != nil && o.CrossOrgUuids != nil
}

// SetCrossOrgUuids gets a reference to the given []string and assigns it to the CrossOrgUuids field.
func (o *TraceStreamQuery) SetCrossOrgUuids(v []string) {
	o.CrossOrgUuids = v
}

// GetDataSource returns the DataSource field value.
func (o *TraceStreamQuery) GetDataSource() TraceStreamQueryDataSource {
	if o == nil {
		var ret TraceStreamQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *TraceStreamQuery) GetDataSourceOk() (*TraceStreamQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *TraceStreamQuery) SetDataSource(v TraceStreamQueryDataSource) {
	o.DataSource = v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *TraceStreamQuery) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceStreamQuery) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *TraceStreamQuery) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *TraceStreamQuery) SetIndexes(v []string) {
	o.Indexes = v
}

// GetQueryString returns the QueryString field value.
func (o *TraceStreamQuery) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *TraceStreamQuery) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value.
func (o *TraceStreamQuery) SetQueryString(v string) {
	o.QueryString = v
}

// GetSemanticMode returns the SemanticMode field value if set, zero value otherwise.
func (o *TraceStreamQuery) GetSemanticMode() FormulaSemanticMode {
	if o == nil || o.SemanticMode == nil {
		var ret FormulaSemanticMode
		return ret
	}
	return *o.SemanticMode
}

// GetSemanticModeOk returns a tuple with the SemanticMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceStreamQuery) GetSemanticModeOk() (*FormulaSemanticMode, bool) {
	if o == nil || o.SemanticMode == nil {
		return nil, false
	}
	return o.SemanticMode, true
}

// HasSemanticMode returns a boolean if a field has been set.
func (o *TraceStreamQuery) HasSemanticMode() bool {
	return o != nil && o.SemanticMode != nil
}

// SetSemanticMode gets a reference to the given FormulaSemanticMode and assigns it to the SemanticMode field.
func (o *TraceStreamQuery) SetSemanticMode(v FormulaSemanticMode) {
	o.SemanticMode = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *TraceStreamQuery) GetSort() TraceStreamQuery20270101Sort {
	if o == nil || o.Sort == nil {
		var ret TraceStreamQuery20270101Sort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceStreamQuery) GetSortOk() (*TraceStreamQuery20270101Sort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *TraceStreamQuery) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given TraceStreamQuery20270101Sort and assigns it to the Sort field.
func (o *TraceStreamQuery) SetSort(v TraceStreamQuery20270101Sort) {
	o.Sort = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TraceStreamQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.CrossOrgUuids != nil {
		toSerialize["cross_org_uuids"] = o.CrossOrgUuids
	}
	toSerialize["data_source"] = o.DataSource
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	toSerialize["query_string"] = o.QueryString
	if o.SemanticMode != nil {
		toSerialize["semantic_mode"] = o.SemanticMode
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TraceStreamQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Context       *EventsContext                `json:"context,omitempty"`
		CrossOrgUuids []string                      `json:"cross_org_uuids,omitempty"`
		DataSource    *TraceStreamQueryDataSource   `json:"data_source"`
		Indexes       []string                      `json:"indexes,omitempty"`
		QueryString   *string                       `json:"query_string"`
		SemanticMode  *FormulaSemanticMode          `json:"semantic_mode,omitempty"`
		Sort          *TraceStreamQuery20270101Sort `json:"sort,omitempty"`
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
	o.Indexes = all.Indexes
	o.QueryString = *all.QueryString
	if all.SemanticMode != nil && !all.SemanticMode.IsValid() {
		hasInvalidField = true
	} else {
		o.SemanticMode = all.SemanticMode
	}
	o.Sort = all.Sort

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
