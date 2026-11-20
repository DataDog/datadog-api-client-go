// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryTableBackendRequest20270101OneOf4460313504 Test-only generated schema for the dashboard versioning proof of concept.
type QueryTableBackendRequest20270101OneOf4460313504 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Aggregator *ScalarAggregator `json:"aggregator,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Alias *string `json:"alias,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CellDisplayMode []CellDisplayModeOption `json:"cell_display_mode,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Columns []TableColumn `json:"columns"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats []ThresholdFormattingRule `json:"conditional_formats,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit *float64 `json:"limit,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Order *QuerySortOrder `json:"order,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Queries []GuidedTableQuery `json:"queries"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType QueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280 `json:"request_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Rows TableRows `json:"rows"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TextFormats [][]TextFormattingRule `json:"text_formats,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewQueryTableBackendRequest20270101OneOf4460313504 instantiates a new QueryTableBackendRequest20270101OneOf4460313504 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQueryTableBackendRequest20270101OneOf4460313504(columns []TableColumn, queries []GuidedTableQuery, requestType QueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280, rows TableRows) *QueryTableBackendRequest20270101OneOf4460313504 {
	this := QueryTableBackendRequest20270101OneOf4460313504{}
	this.Columns = columns
	this.Queries = queries
	this.RequestType = requestType
	this.Rows = rows
	return &this
}

// NewQueryTableBackendRequest20270101OneOf4460313504WithDefaults instantiates a new QueryTableBackendRequest20270101OneOf4460313504 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueryTableBackendRequest20270101OneOf4460313504WithDefaults() *QueryTableBackendRequest20270101OneOf4460313504 {
	this := QueryTableBackendRequest20270101OneOf4460313504{}
	return &this
}

// GetAggregator returns the Aggregator field value if set, zero value otherwise.
func (o *QueryTableBackendRequest20270101OneOf4460313504) GetAggregator() ScalarAggregator {
	if o == nil || o.Aggregator == nil {
		var ret ScalarAggregator
		return ret
	}
	return *o.Aggregator
}

// GetAggregatorOk returns a tuple with the Aggregator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460313504) GetAggregatorOk() (*ScalarAggregator, bool) {
	if o == nil || o.Aggregator == nil {
		return nil, false
	}
	return o.Aggregator, true
}

// HasAggregator returns a boolean if a field has been set.
func (o *QueryTableBackendRequest20270101OneOf4460313504) HasAggregator() bool {
	return o != nil && o.Aggregator != nil
}

// SetAggregator gets a reference to the given ScalarAggregator and assigns it to the Aggregator field.
func (o *QueryTableBackendRequest20270101OneOf4460313504) SetAggregator(v ScalarAggregator) {
	o.Aggregator = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *QueryTableBackendRequest20270101OneOf4460313504) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460313504) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *QueryTableBackendRequest20270101OneOf4460313504) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *QueryTableBackendRequest20270101OneOf4460313504) SetAlias(v string) {
	o.Alias = &v
}

// GetCellDisplayMode returns the CellDisplayMode field value if set, zero value otherwise.
func (o *QueryTableBackendRequest20270101OneOf4460313504) GetCellDisplayMode() []CellDisplayModeOption {
	if o == nil || o.CellDisplayMode == nil {
		var ret []CellDisplayModeOption
		return ret
	}
	return o.CellDisplayMode
}

// GetCellDisplayModeOk returns a tuple with the CellDisplayMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460313504) GetCellDisplayModeOk() (*[]CellDisplayModeOption, bool) {
	if o == nil || o.CellDisplayMode == nil {
		return nil, false
	}
	return &o.CellDisplayMode, true
}

// HasCellDisplayMode returns a boolean if a field has been set.
func (o *QueryTableBackendRequest20270101OneOf4460313504) HasCellDisplayMode() bool {
	return o != nil && o.CellDisplayMode != nil
}

// SetCellDisplayMode gets a reference to the given []CellDisplayModeOption and assigns it to the CellDisplayMode field.
func (o *QueryTableBackendRequest20270101OneOf4460313504) SetCellDisplayMode(v []CellDisplayModeOption) {
	o.CellDisplayMode = v
}

// GetColumns returns the Columns field value.
func (o *QueryTableBackendRequest20270101OneOf4460313504) GetColumns() []TableColumn {
	if o == nil {
		var ret []TableColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460313504) GetColumnsOk() (*[]TableColumn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value.
func (o *QueryTableBackendRequest20270101OneOf4460313504) SetColumns(v []TableColumn) {
	o.Columns = v
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *QueryTableBackendRequest20270101OneOf4460313504) GetConditionalFormats() []ThresholdFormattingRule {
	if o == nil || o.ConditionalFormats == nil {
		var ret []ThresholdFormattingRule
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460313504) GetConditionalFormatsOk() (*[]ThresholdFormattingRule, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *QueryTableBackendRequest20270101OneOf4460313504) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []ThresholdFormattingRule and assigns it to the ConditionalFormats field.
func (o *QueryTableBackendRequest20270101OneOf4460313504) SetConditionalFormats(v []ThresholdFormattingRule) {
	o.ConditionalFormats = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *QueryTableBackendRequest20270101OneOf4460313504) GetLimit() float64 {
	if o == nil || o.Limit == nil {
		var ret float64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460313504) GetLimitOk() (*float64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *QueryTableBackendRequest20270101OneOf4460313504) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given float64 and assigns it to the Limit field.
func (o *QueryTableBackendRequest20270101OneOf4460313504) SetLimit(v float64) {
	o.Limit = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *QueryTableBackendRequest20270101OneOf4460313504) GetOrder() QuerySortOrder {
	if o == nil || o.Order == nil {
		var ret QuerySortOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460313504) GetOrderOk() (*QuerySortOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *QueryTableBackendRequest20270101OneOf4460313504) HasOrder() bool {
	return o != nil && o.Order != nil
}

// SetOrder gets a reference to the given QuerySortOrder and assigns it to the Order field.
func (o *QueryTableBackendRequest20270101OneOf4460313504) SetOrder(v QuerySortOrder) {
	o.Order = &v
}

// GetQueries returns the Queries field value.
func (o *QueryTableBackendRequest20270101OneOf4460313504) GetQueries() []GuidedTableQuery {
	if o == nil {
		var ret []GuidedTableQuery
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460313504) GetQueriesOk() (*[]GuidedTableQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *QueryTableBackendRequest20270101OneOf4460313504) SetQueries(v []GuidedTableQuery) {
	o.Queries = v
}

// GetRequestType returns the RequestType field value.
func (o *QueryTableBackendRequest20270101OneOf4460313504) GetRequestType() QueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280 {
	if o == nil {
		var ret QueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460313504) GetRequestTypeOk() (*QueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *QueryTableBackendRequest20270101OneOf4460313504) SetRequestType(v QueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280) {
	o.RequestType = v
}

// GetRows returns the Rows field value.
func (o *QueryTableBackendRequest20270101OneOf4460313504) GetRows() TableRows {
	if o == nil {
		var ret TableRows
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460313504) GetRowsOk() (*TableRows, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rows, true
}

// SetRows sets field value.
func (o *QueryTableBackendRequest20270101OneOf4460313504) SetRows(v TableRows) {
	o.Rows = v
}

// GetTextFormats returns the TextFormats field value if set, zero value otherwise.
func (o *QueryTableBackendRequest20270101OneOf4460313504) GetTextFormats() [][]TextFormattingRule {
	if o == nil || o.TextFormats == nil {
		var ret [][]TextFormattingRule
		return ret
	}
	return o.TextFormats
}

// GetTextFormatsOk returns a tuple with the TextFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460313504) GetTextFormatsOk() (*[][]TextFormattingRule, bool) {
	if o == nil || o.TextFormats == nil {
		return nil, false
	}
	return &o.TextFormats, true
}

// HasTextFormats returns a boolean if a field has been set.
func (o *QueryTableBackendRequest20270101OneOf4460313504) HasTextFormats() bool {
	return o != nil && o.TextFormats != nil
}

// SetTextFormats gets a reference to the given [][]TextFormattingRule and assigns it to the TextFormats field.
func (o *QueryTableBackendRequest20270101OneOf4460313504) SetTextFormats(v [][]TextFormattingRule) {
	o.TextFormats = v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueryTableBackendRequest20270101OneOf4460313504) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Aggregator != nil {
		toSerialize["aggregator"] = o.Aggregator
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	if o.CellDisplayMode != nil {
		toSerialize["cell_display_mode"] = o.CellDisplayMode
	}
	toSerialize["columns"] = o.Columns
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	toSerialize["queries"] = o.Queries
	toSerialize["request_type"] = o.RequestType
	toSerialize["rows"] = o.Rows
	if o.TextFormats != nil {
		toSerialize["text_formats"] = o.TextFormats
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *QueryTableBackendRequest20270101OneOf4460313504) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregator         *ScalarAggregator                                            `json:"aggregator,omitempty"`
		Alias              *string                                                      `json:"alias,omitempty"`
		CellDisplayMode    []CellDisplayModeOption                                      `json:"cell_display_mode,omitempty"`
		Columns            *[]TableColumn                                               `json:"columns"`
		ConditionalFormats []ThresholdFormattingRule                                    `json:"conditional_formats,omitempty"`
		Limit              *float64                                                     `json:"limit,omitempty"`
		Order              *QuerySortOrder                                              `json:"order,omitempty"`
		Queries            *[]GuidedTableQuery                                          `json:"queries"`
		RequestType        *QueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280 `json:"request_type"`
		Rows               *TableRows                                                   `json:"rows"`
		TextFormats        [][]TextFormattingRule                                       `json:"text_formats,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Columns == nil {
		return fmt.Errorf("required field columns missing")
	}
	if all.Queries == nil {
		return fmt.Errorf("required field queries missing")
	}
	if all.RequestType == nil {
		return fmt.Errorf("required field request_type missing")
	}
	if all.Rows == nil {
		return fmt.Errorf("required field rows missing")
	}

	hasInvalidField := false
	if all.Aggregator != nil && !all.Aggregator.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregator = all.Aggregator
	}
	o.Alias = all.Alias
	o.CellDisplayMode = all.CellDisplayMode
	o.Columns = *all.Columns
	o.ConditionalFormats = all.ConditionalFormats
	o.Limit = all.Limit
	if all.Order != nil && !all.Order.IsValid() {
		hasInvalidField = true
	} else {
		o.Order = all.Order
	}
	o.Queries = *all.Queries
	if !all.RequestType.IsValid() {
		hasInvalidField = true
	} else {
		o.RequestType = *all.RequestType
	}
	if all.Rows.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Rows = *all.Rows
	o.TextFormats = all.TextFormats

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
