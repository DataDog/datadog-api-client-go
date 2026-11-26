// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesWidgetAdvancedQueryRequest Test-only generated schema for the dashboard versioning proof of concept.
type TimeseriesWidgetAdvancedQueryRequest struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	DatasetToQuery string `json:"dataset_to_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Datasets []AdvancedQueryDataset `json:"datasets"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DisplayType *WidgetDisplayType `json:"display_type,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit *float64 `json:"limit,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Metadata []TimeseriesWidgetExpressionAlias `json:"metadata,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	OnRightYaxis *bool `json:"on_right_yaxis,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Projection *TimeseriesProjection `json:"projection,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType TimeseriesWidgetAdvancedQueryRequestRequestType `json:"request_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat TimeseriesWidgetAdvancedQueryRequestResponseFormat `json:"response_format"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *TimeseriesRequestStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTimeseriesWidgetAdvancedQueryRequest instantiates a new TimeseriesWidgetAdvancedQueryRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesWidgetAdvancedQueryRequest(datasetToQuery string, datasets []AdvancedQueryDataset, requestType TimeseriesWidgetAdvancedQueryRequestRequestType, responseFormat TimeseriesWidgetAdvancedQueryRequestResponseFormat) *TimeseriesWidgetAdvancedQueryRequest {
	this := TimeseriesWidgetAdvancedQueryRequest{}
	this.DatasetToQuery = datasetToQuery
	this.Datasets = datasets
	this.RequestType = requestType
	this.ResponseFormat = responseFormat
	return &this
}

// NewTimeseriesWidgetAdvancedQueryRequestWithDefaults instantiates a new TimeseriesWidgetAdvancedQueryRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesWidgetAdvancedQueryRequestWithDefaults() *TimeseriesWidgetAdvancedQueryRequest {
	this := TimeseriesWidgetAdvancedQueryRequest{}
	return &this
}

// GetDatasetToQuery returns the DatasetToQuery field value.
func (o *TimeseriesWidgetAdvancedQueryRequest) GetDatasetToQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DatasetToQuery
}

// GetDatasetToQueryOk returns a tuple with the DatasetToQuery field value
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetAdvancedQueryRequest) GetDatasetToQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetToQuery, true
}

// SetDatasetToQuery sets field value.
func (o *TimeseriesWidgetAdvancedQueryRequest) SetDatasetToQuery(v string) {
	o.DatasetToQuery = v
}

// GetDatasets returns the Datasets field value.
func (o *TimeseriesWidgetAdvancedQueryRequest) GetDatasets() []AdvancedQueryDataset {
	if o == nil {
		var ret []AdvancedQueryDataset
		return ret
	}
	return o.Datasets
}

// GetDatasetsOk returns a tuple with the Datasets field value
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetAdvancedQueryRequest) GetDatasetsOk() (*[]AdvancedQueryDataset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datasets, true
}

// SetDatasets sets field value.
func (o *TimeseriesWidgetAdvancedQueryRequest) SetDatasets(v []AdvancedQueryDataset) {
	o.Datasets = v
}

// GetDisplayType returns the DisplayType field value if set, zero value otherwise.
func (o *TimeseriesWidgetAdvancedQueryRequest) GetDisplayType() WidgetDisplayType {
	if o == nil || o.DisplayType == nil {
		var ret WidgetDisplayType
		return ret
	}
	return *o.DisplayType
}

// GetDisplayTypeOk returns a tuple with the DisplayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetAdvancedQueryRequest) GetDisplayTypeOk() (*WidgetDisplayType, bool) {
	if o == nil || o.DisplayType == nil {
		return nil, false
	}
	return o.DisplayType, true
}

// HasDisplayType returns a boolean if a field has been set.
func (o *TimeseriesWidgetAdvancedQueryRequest) HasDisplayType() bool {
	return o != nil && o.DisplayType != nil
}

// SetDisplayType gets a reference to the given WidgetDisplayType and assigns it to the DisplayType field.
func (o *TimeseriesWidgetAdvancedQueryRequest) SetDisplayType(v WidgetDisplayType) {
	o.DisplayType = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *TimeseriesWidgetAdvancedQueryRequest) GetLimit() float64 {
	if o == nil || o.Limit == nil {
		var ret float64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetAdvancedQueryRequest) GetLimitOk() (*float64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *TimeseriesWidgetAdvancedQueryRequest) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given float64 and assigns it to the Limit field.
func (o *TimeseriesWidgetAdvancedQueryRequest) SetLimit(v float64) {
	o.Limit = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *TimeseriesWidgetAdvancedQueryRequest) GetMetadata() []TimeseriesWidgetExpressionAlias {
	if o == nil || o.Metadata == nil {
		var ret []TimeseriesWidgetExpressionAlias
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetAdvancedQueryRequest) GetMetadataOk() (*[]TimeseriesWidgetExpressionAlias, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TimeseriesWidgetAdvancedQueryRequest) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given []TimeseriesWidgetExpressionAlias and assigns it to the Metadata field.
func (o *TimeseriesWidgetAdvancedQueryRequest) SetMetadata(v []TimeseriesWidgetExpressionAlias) {
	o.Metadata = v
}

// GetOnRightYaxis returns the OnRightYaxis field value if set, zero value otherwise.
func (o *TimeseriesWidgetAdvancedQueryRequest) GetOnRightYaxis() bool {
	if o == nil || o.OnRightYaxis == nil {
		var ret bool
		return ret
	}
	return *o.OnRightYaxis
}

// GetOnRightYaxisOk returns a tuple with the OnRightYaxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetAdvancedQueryRequest) GetOnRightYaxisOk() (*bool, bool) {
	if o == nil || o.OnRightYaxis == nil {
		return nil, false
	}
	return o.OnRightYaxis, true
}

// HasOnRightYaxis returns a boolean if a field has been set.
func (o *TimeseriesWidgetAdvancedQueryRequest) HasOnRightYaxis() bool {
	return o != nil && o.OnRightYaxis != nil
}

// SetOnRightYaxis gets a reference to the given bool and assigns it to the OnRightYaxis field.
func (o *TimeseriesWidgetAdvancedQueryRequest) SetOnRightYaxis(v bool) {
	o.OnRightYaxis = &v
}

// GetProjection returns the Projection field value if set, zero value otherwise.
func (o *TimeseriesWidgetAdvancedQueryRequest) GetProjection() TimeseriesProjection {
	if o == nil || o.Projection == nil {
		var ret TimeseriesProjection
		return ret
	}
	return *o.Projection
}

// GetProjectionOk returns a tuple with the Projection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetAdvancedQueryRequest) GetProjectionOk() (*TimeseriesProjection, bool) {
	if o == nil || o.Projection == nil {
		return nil, false
	}
	return o.Projection, true
}

// HasProjection returns a boolean if a field has been set.
func (o *TimeseriesWidgetAdvancedQueryRequest) HasProjection() bool {
	return o != nil && o.Projection != nil
}

// SetProjection gets a reference to the given TimeseriesProjection and assigns it to the Projection field.
func (o *TimeseriesWidgetAdvancedQueryRequest) SetProjection(v TimeseriesProjection) {
	o.Projection = &v
}

// GetRequestType returns the RequestType field value.
func (o *TimeseriesWidgetAdvancedQueryRequest) GetRequestType() TimeseriesWidgetAdvancedQueryRequestRequestType {
	if o == nil {
		var ret TimeseriesWidgetAdvancedQueryRequestRequestType
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetAdvancedQueryRequest) GetRequestTypeOk() (*TimeseriesWidgetAdvancedQueryRequestRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *TimeseriesWidgetAdvancedQueryRequest) SetRequestType(v TimeseriesWidgetAdvancedQueryRequestRequestType) {
	o.RequestType = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *TimeseriesWidgetAdvancedQueryRequest) GetResponseFormat() TimeseriesWidgetAdvancedQueryRequestResponseFormat {
	if o == nil {
		var ret TimeseriesWidgetAdvancedQueryRequestResponseFormat
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetAdvancedQueryRequest) GetResponseFormatOk() (*TimeseriesWidgetAdvancedQueryRequestResponseFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *TimeseriesWidgetAdvancedQueryRequest) SetResponseFormat(v TimeseriesWidgetAdvancedQueryRequestResponseFormat) {
	o.ResponseFormat = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *TimeseriesWidgetAdvancedQueryRequest) GetStyle() TimeseriesRequestStyle {
	if o == nil || o.Style == nil {
		var ret TimeseriesRequestStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetAdvancedQueryRequest) GetStyleOk() (*TimeseriesRequestStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *TimeseriesWidgetAdvancedQueryRequest) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given TimeseriesRequestStyle and assigns it to the Style field.
func (o *TimeseriesWidgetAdvancedQueryRequest) SetStyle(v TimeseriesRequestStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesWidgetAdvancedQueryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["dataset_to_query"] = o.DatasetToQuery
	toSerialize["datasets"] = o.Datasets
	if o.DisplayType != nil {
		toSerialize["display_type"] = o.DisplayType
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.OnRightYaxis != nil {
		toSerialize["on_right_yaxis"] = o.OnRightYaxis
	}
	if o.Projection != nil {
		toSerialize["projection"] = o.Projection
	}
	toSerialize["request_type"] = o.RequestType
	toSerialize["response_format"] = o.ResponseFormat
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesWidgetAdvancedQueryRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatasetToQuery *string                                             `json:"dataset_to_query"`
		Datasets       *[]AdvancedQueryDataset                             `json:"datasets"`
		DisplayType    *WidgetDisplayType                                  `json:"display_type,omitempty"`
		Limit          *float64                                            `json:"limit,omitempty"`
		Metadata       []TimeseriesWidgetExpressionAlias                   `json:"metadata,omitempty"`
		OnRightYaxis   *bool                                               `json:"on_right_yaxis,omitempty"`
		Projection     *TimeseriesProjection                               `json:"projection,omitempty"`
		RequestType    *TimeseriesWidgetAdvancedQueryRequestRequestType    `json:"request_type"`
		ResponseFormat *TimeseriesWidgetAdvancedQueryRequestResponseFormat `json:"response_format"`
		Style          *TimeseriesRequestStyle                             `json:"style,omitempty"`
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
	o.DatasetToQuery = *all.DatasetToQuery
	o.Datasets = *all.Datasets
	if all.DisplayType != nil && !all.DisplayType.IsValid() {
		hasInvalidField = true
	} else {
		o.DisplayType = all.DisplayType
	}
	o.Limit = all.Limit
	o.Metadata = all.Metadata
	o.OnRightYaxis = all.OnRightYaxis
	if all.Projection != nil && all.Projection.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
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
	if all.Style != nil && all.Style.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Style = all.Style

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
