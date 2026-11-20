// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ToplistBackendRequest20270101OneOf Test-only generated schema for the dashboard versioning proof of concept.
type ToplistBackendRequest20270101OneOf struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats []ThresholdFormattingRule `json:"conditional_formats,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DatasetToQuery string `json:"dataset_to_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Datasets []AdvancedQueryDataset `json:"datasets"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit *float64 `json:"limit,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Projection *TabularProjection `json:"projection,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType ToplistBackendRequestAnyOfAnyOfItemRequestType `json:"request_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat ToplistBackendRequestAnyOfAnyOfItemResponseFormat `json:"response_format"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewToplistBackendRequest20270101OneOf instantiates a new ToplistBackendRequest20270101OneOf object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewToplistBackendRequest20270101OneOf(datasetToQuery string, datasets []AdvancedQueryDataset, requestType ToplistBackendRequestAnyOfAnyOfItemRequestType, responseFormat ToplistBackendRequestAnyOfAnyOfItemResponseFormat) *ToplistBackendRequest20270101OneOf {
	this := ToplistBackendRequest20270101OneOf{}
	this.DatasetToQuery = datasetToQuery
	this.Datasets = datasets
	this.RequestType = requestType
	this.ResponseFormat = responseFormat
	return &this
}

// NewToplistBackendRequest20270101OneOfWithDefaults instantiates a new ToplistBackendRequest20270101OneOf object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewToplistBackendRequest20270101OneOfWithDefaults() *ToplistBackendRequest20270101OneOf {
	this := ToplistBackendRequest20270101OneOf{}
	return &this
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *ToplistBackendRequest20270101OneOf) GetConditionalFormats() []ThresholdFormattingRule {
	if o == nil || o.ConditionalFormats == nil {
		var ret []ThresholdFormattingRule
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToplistBackendRequest20270101OneOf) GetConditionalFormatsOk() (*[]ThresholdFormattingRule, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *ToplistBackendRequest20270101OneOf) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []ThresholdFormattingRule and assigns it to the ConditionalFormats field.
func (o *ToplistBackendRequest20270101OneOf) SetConditionalFormats(v []ThresholdFormattingRule) {
	o.ConditionalFormats = v
}

// GetDatasetToQuery returns the DatasetToQuery field value.
func (o *ToplistBackendRequest20270101OneOf) GetDatasetToQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DatasetToQuery
}

// GetDatasetToQueryOk returns a tuple with the DatasetToQuery field value
// and a boolean to check if the value has been set.
func (o *ToplistBackendRequest20270101OneOf) GetDatasetToQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetToQuery, true
}

// SetDatasetToQuery sets field value.
func (o *ToplistBackendRequest20270101OneOf) SetDatasetToQuery(v string) {
	o.DatasetToQuery = v
}

// GetDatasets returns the Datasets field value.
func (o *ToplistBackendRequest20270101OneOf) GetDatasets() []AdvancedQueryDataset {
	if o == nil {
		var ret []AdvancedQueryDataset
		return ret
	}
	return o.Datasets
}

// GetDatasetsOk returns a tuple with the Datasets field value
// and a boolean to check if the value has been set.
func (o *ToplistBackendRequest20270101OneOf) GetDatasetsOk() (*[]AdvancedQueryDataset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datasets, true
}

// SetDatasets sets field value.
func (o *ToplistBackendRequest20270101OneOf) SetDatasets(v []AdvancedQueryDataset) {
	o.Datasets = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ToplistBackendRequest20270101OneOf) GetLimit() float64 {
	if o == nil || o.Limit == nil {
		var ret float64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToplistBackendRequest20270101OneOf) GetLimitOk() (*float64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ToplistBackendRequest20270101OneOf) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given float64 and assigns it to the Limit field.
func (o *ToplistBackendRequest20270101OneOf) SetLimit(v float64) {
	o.Limit = &v
}

// GetProjection returns the Projection field value if set, zero value otherwise.
func (o *ToplistBackendRequest20270101OneOf) GetProjection() TabularProjection {
	if o == nil || o.Projection == nil {
		var ret TabularProjection
		return ret
	}
	return *o.Projection
}

// GetProjectionOk returns a tuple with the Projection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToplistBackendRequest20270101OneOf) GetProjectionOk() (*TabularProjection, bool) {
	if o == nil || o.Projection == nil {
		return nil, false
	}
	return o.Projection, true
}

// HasProjection returns a boolean if a field has been set.
func (o *ToplistBackendRequest20270101OneOf) HasProjection() bool {
	return o != nil && o.Projection != nil
}

// SetProjection gets a reference to the given TabularProjection and assigns it to the Projection field.
func (o *ToplistBackendRequest20270101OneOf) SetProjection(v TabularProjection) {
	o.Projection = &v
}

// GetRequestType returns the RequestType field value.
func (o *ToplistBackendRequest20270101OneOf) GetRequestType() ToplistBackendRequestAnyOfAnyOfItemRequestType {
	if o == nil {
		var ret ToplistBackendRequestAnyOfAnyOfItemRequestType
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *ToplistBackendRequest20270101OneOf) GetRequestTypeOk() (*ToplistBackendRequestAnyOfAnyOfItemRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *ToplistBackendRequest20270101OneOf) SetRequestType(v ToplistBackendRequestAnyOfAnyOfItemRequestType) {
	o.RequestType = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *ToplistBackendRequest20270101OneOf) GetResponseFormat() ToplistBackendRequestAnyOfAnyOfItemResponseFormat {
	if o == nil {
		var ret ToplistBackendRequestAnyOfAnyOfItemResponseFormat
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *ToplistBackendRequest20270101OneOf) GetResponseFormatOk() (*ToplistBackendRequestAnyOfAnyOfItemResponseFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *ToplistBackendRequest20270101OneOf) SetResponseFormat(v ToplistBackendRequestAnyOfAnyOfItemResponseFormat) {
	o.ResponseFormat = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ToplistBackendRequest20270101OneOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	toSerialize["dataset_to_query"] = o.DatasetToQuery
	toSerialize["datasets"] = o.Datasets
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Projection != nil {
		toSerialize["projection"] = o.Projection
	}
	toSerialize["request_type"] = o.RequestType
	toSerialize["response_format"] = o.ResponseFormat
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ToplistBackendRequest20270101OneOf) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConditionalFormats []ThresholdFormattingRule                          `json:"conditional_formats,omitempty"`
		DatasetToQuery     *string                                            `json:"dataset_to_query"`
		Datasets           *[]AdvancedQueryDataset                            `json:"datasets"`
		Limit              *float64                                           `json:"limit,omitempty"`
		Projection         *TabularProjection                                 `json:"projection,omitempty"`
		RequestType        *ToplistBackendRequestAnyOfAnyOfItemRequestType    `json:"request_type"`
		ResponseFormat     *ToplistBackendRequestAnyOfAnyOfItemResponseFormat `json:"response_format"`
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
	o.ConditionalFormats = all.ConditionalFormats
	o.DatasetToQuery = *all.DatasetToQuery
	o.Datasets = *all.Datasets
	o.Limit = all.Limit
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

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
