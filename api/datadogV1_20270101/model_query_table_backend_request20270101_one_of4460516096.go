// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryTableBackendRequest20270101OneOf4460516096 Test-only generated schema for the dashboard versioning proof of concept.
type QueryTableBackendRequest20270101OneOf4460516096 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Aggregator *ScalarAggregator `json:"aggregator,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Alias *string `json:"alias,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CellDisplayMode []CellDisplayModeOption `json:"cell_display_mode,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats []ThresholdFormattingRule `json:"conditional_formats,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DatasetToQuery string `json:"dataset_to_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Datasets []AdvancedQueryDataset `json:"datasets"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit *float64 `json:"limit,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Order *QuerySortOrder `json:"order,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Projection *TabularProjection `json:"projection,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType QueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840 `json:"request_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640 `json:"response_format"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TextFormats [][]TextFormattingRule `json:"text_formats,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewQueryTableBackendRequest20270101OneOf4460516096 instantiates a new QueryTableBackendRequest20270101OneOf4460516096 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQueryTableBackendRequest20270101OneOf4460516096(datasetToQuery string, datasets []AdvancedQueryDataset, requestType QueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840, responseFormat QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640) *QueryTableBackendRequest20270101OneOf4460516096 {
	this := QueryTableBackendRequest20270101OneOf4460516096{}
	this.DatasetToQuery = datasetToQuery
	this.Datasets = datasets
	this.RequestType = requestType
	this.ResponseFormat = responseFormat
	return &this
}

// NewQueryTableBackendRequest20270101OneOf4460516096WithDefaults instantiates a new QueryTableBackendRequest20270101OneOf4460516096 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueryTableBackendRequest20270101OneOf4460516096WithDefaults() *QueryTableBackendRequest20270101OneOf4460516096 {
	this := QueryTableBackendRequest20270101OneOf4460516096{}
	return &this
}

// GetAggregator returns the Aggregator field value if set, zero value otherwise.
func (o *QueryTableBackendRequest20270101OneOf4460516096) GetAggregator() ScalarAggregator {
	if o == nil || o.Aggregator == nil {
		var ret ScalarAggregator
		return ret
	}
	return *o.Aggregator
}

// GetAggregatorOk returns a tuple with the Aggregator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460516096) GetAggregatorOk() (*ScalarAggregator, bool) {
	if o == nil || o.Aggregator == nil {
		return nil, false
	}
	return o.Aggregator, true
}

// HasAggregator returns a boolean if a field has been set.
func (o *QueryTableBackendRequest20270101OneOf4460516096) HasAggregator() bool {
	return o != nil && o.Aggregator != nil
}

// SetAggregator gets a reference to the given ScalarAggregator and assigns it to the Aggregator field.
func (o *QueryTableBackendRequest20270101OneOf4460516096) SetAggregator(v ScalarAggregator) {
	o.Aggregator = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *QueryTableBackendRequest20270101OneOf4460516096) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460516096) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *QueryTableBackendRequest20270101OneOf4460516096) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *QueryTableBackendRequest20270101OneOf4460516096) SetAlias(v string) {
	o.Alias = &v
}

// GetCellDisplayMode returns the CellDisplayMode field value if set, zero value otherwise.
func (o *QueryTableBackendRequest20270101OneOf4460516096) GetCellDisplayMode() []CellDisplayModeOption {
	if o == nil || o.CellDisplayMode == nil {
		var ret []CellDisplayModeOption
		return ret
	}
	return o.CellDisplayMode
}

// GetCellDisplayModeOk returns a tuple with the CellDisplayMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460516096) GetCellDisplayModeOk() (*[]CellDisplayModeOption, bool) {
	if o == nil || o.CellDisplayMode == nil {
		return nil, false
	}
	return &o.CellDisplayMode, true
}

// HasCellDisplayMode returns a boolean if a field has been set.
func (o *QueryTableBackendRequest20270101OneOf4460516096) HasCellDisplayMode() bool {
	return o != nil && o.CellDisplayMode != nil
}

// SetCellDisplayMode gets a reference to the given []CellDisplayModeOption and assigns it to the CellDisplayMode field.
func (o *QueryTableBackendRequest20270101OneOf4460516096) SetCellDisplayMode(v []CellDisplayModeOption) {
	o.CellDisplayMode = v
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *QueryTableBackendRequest20270101OneOf4460516096) GetConditionalFormats() []ThresholdFormattingRule {
	if o == nil || o.ConditionalFormats == nil {
		var ret []ThresholdFormattingRule
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460516096) GetConditionalFormatsOk() (*[]ThresholdFormattingRule, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *QueryTableBackendRequest20270101OneOf4460516096) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []ThresholdFormattingRule and assigns it to the ConditionalFormats field.
func (o *QueryTableBackendRequest20270101OneOf4460516096) SetConditionalFormats(v []ThresholdFormattingRule) {
	o.ConditionalFormats = v
}

// GetDatasetToQuery returns the DatasetToQuery field value.
func (o *QueryTableBackendRequest20270101OneOf4460516096) GetDatasetToQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DatasetToQuery
}

// GetDatasetToQueryOk returns a tuple with the DatasetToQuery field value
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460516096) GetDatasetToQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetToQuery, true
}

// SetDatasetToQuery sets field value.
func (o *QueryTableBackendRequest20270101OneOf4460516096) SetDatasetToQuery(v string) {
	o.DatasetToQuery = v
}

// GetDatasets returns the Datasets field value.
func (o *QueryTableBackendRequest20270101OneOf4460516096) GetDatasets() []AdvancedQueryDataset {
	if o == nil {
		var ret []AdvancedQueryDataset
		return ret
	}
	return o.Datasets
}

// GetDatasetsOk returns a tuple with the Datasets field value
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460516096) GetDatasetsOk() (*[]AdvancedQueryDataset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datasets, true
}

// SetDatasets sets field value.
func (o *QueryTableBackendRequest20270101OneOf4460516096) SetDatasets(v []AdvancedQueryDataset) {
	o.Datasets = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *QueryTableBackendRequest20270101OneOf4460516096) GetLimit() float64 {
	if o == nil || o.Limit == nil {
		var ret float64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460516096) GetLimitOk() (*float64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *QueryTableBackendRequest20270101OneOf4460516096) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given float64 and assigns it to the Limit field.
func (o *QueryTableBackendRequest20270101OneOf4460516096) SetLimit(v float64) {
	o.Limit = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *QueryTableBackendRequest20270101OneOf4460516096) GetOrder() QuerySortOrder {
	if o == nil || o.Order == nil {
		var ret QuerySortOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460516096) GetOrderOk() (*QuerySortOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *QueryTableBackendRequest20270101OneOf4460516096) HasOrder() bool {
	return o != nil && o.Order != nil
}

// SetOrder gets a reference to the given QuerySortOrder and assigns it to the Order field.
func (o *QueryTableBackendRequest20270101OneOf4460516096) SetOrder(v QuerySortOrder) {
	o.Order = &v
}

// GetProjection returns the Projection field value if set, zero value otherwise.
func (o *QueryTableBackendRequest20270101OneOf4460516096) GetProjection() TabularProjection {
	if o == nil || o.Projection == nil {
		var ret TabularProjection
		return ret
	}
	return *o.Projection
}

// GetProjectionOk returns a tuple with the Projection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460516096) GetProjectionOk() (*TabularProjection, bool) {
	if o == nil || o.Projection == nil {
		return nil, false
	}
	return o.Projection, true
}

// HasProjection returns a boolean if a field has been set.
func (o *QueryTableBackendRequest20270101OneOf4460516096) HasProjection() bool {
	return o != nil && o.Projection != nil
}

// SetProjection gets a reference to the given TabularProjection and assigns it to the Projection field.
func (o *QueryTableBackendRequest20270101OneOf4460516096) SetProjection(v TabularProjection) {
	o.Projection = &v
}

// GetRequestType returns the RequestType field value.
func (o *QueryTableBackendRequest20270101OneOf4460516096) GetRequestType() QueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840 {
	if o == nil {
		var ret QueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460516096) GetRequestTypeOk() (*QueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *QueryTableBackendRequest20270101OneOf4460516096) SetRequestType(v QueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840) {
	o.RequestType = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *QueryTableBackendRequest20270101OneOf4460516096) GetResponseFormat() QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640 {
	if o == nil {
		var ret QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460516096) GetResponseFormatOk() (*QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *QueryTableBackendRequest20270101OneOf4460516096) SetResponseFormat(v QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640) {
	o.ResponseFormat = v
}

// GetTextFormats returns the TextFormats field value if set, zero value otherwise.
func (o *QueryTableBackendRequest20270101OneOf4460516096) GetTextFormats() [][]TextFormattingRule {
	if o == nil || o.TextFormats == nil {
		var ret [][]TextFormattingRule
		return ret
	}
	return o.TextFormats
}

// GetTextFormatsOk returns a tuple with the TextFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460516096) GetTextFormatsOk() (*[][]TextFormattingRule, bool) {
	if o == nil || o.TextFormats == nil {
		return nil, false
	}
	return &o.TextFormats, true
}

// HasTextFormats returns a boolean if a field has been set.
func (o *QueryTableBackendRequest20270101OneOf4460516096) HasTextFormats() bool {
	return o != nil && o.TextFormats != nil
}

// SetTextFormats gets a reference to the given [][]TextFormattingRule and assigns it to the TextFormats field.
func (o *QueryTableBackendRequest20270101OneOf4460516096) SetTextFormats(v [][]TextFormattingRule) {
	o.TextFormats = v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueryTableBackendRequest20270101OneOf4460516096) MarshalJSON() ([]byte, error) {
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
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	toSerialize["dataset_to_query"] = o.DatasetToQuery
	toSerialize["datasets"] = o.Datasets
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.Projection != nil {
		toSerialize["projection"] = o.Projection
	}
	toSerialize["request_type"] = o.RequestType
	toSerialize["response_format"] = o.ResponseFormat
	if o.TextFormats != nil {
		toSerialize["text_formats"] = o.TextFormats
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *QueryTableBackendRequest20270101OneOf4460516096) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregator         *ScalarAggregator                                               `json:"aggregator,omitempty"`
		Alias              *string                                                         `json:"alias,omitempty"`
		CellDisplayMode    []CellDisplayModeOption                                         `json:"cell_display_mode,omitempty"`
		ConditionalFormats []ThresholdFormattingRule                                       `json:"conditional_formats,omitempty"`
		DatasetToQuery     *string                                                         `json:"dataset_to_query"`
		Datasets           *[]AdvancedQueryDataset                                         `json:"datasets"`
		Limit              *float64                                                        `json:"limit,omitempty"`
		Order              *QuerySortOrder                                                 `json:"order,omitempty"`
		Projection         *TabularProjection                                              `json:"projection,omitempty"`
		RequestType        *QueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840    `json:"request_type"`
		ResponseFormat     *QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640 `json:"response_format"`
		TextFormats        [][]TextFormattingRule                                          `json:"text_formats,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DatasetToQuery == nil {
		return fmt.Errorf("required field dataset_to_query missing")
	}
	if all.Datasets == nil {
		return fmt.Errorf("required field datasets missing")
	}
	if all.RequestType == nil {
		return fmt.Errorf("required field request_type missing")
	}
	if all.ResponseFormat == nil {
		return fmt.Errorf("required field response_format missing")
	}

	hasInvalidField := false
	if all.Aggregator != nil && !all.Aggregator.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregator = all.Aggregator
	}
	o.Alias = all.Alias
	o.CellDisplayMode = all.CellDisplayMode
	o.ConditionalFormats = all.ConditionalFormats
	o.DatasetToQuery = *all.DatasetToQuery
	o.Datasets = *all.Datasets
	o.Limit = all.Limit
	if all.Order != nil && !all.Order.IsValid() {
		hasInvalidField = true
	} else {
		o.Order = all.Order
	}
	o.Projection = all.Projection
	if !all.RequestType.IsValid() {
		hasInvalidField = true
	} else {
		o.RequestType = *all.RequestType
	}
	if !all.ResponseFormat.IsValid() {
		hasInvalidField = true
	} else {
		o.ResponseFormat = *all.ResponseFormat
	}
	o.TextFormats = all.TextFormats

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
