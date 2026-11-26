// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesWidgetLocalDatasetRequest Test-only generated schema for the dashboard versioning proof of concept.
type TimeseriesWidgetLocalDatasetRequest struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	DisplayType *WidgetDisplayType `json:"display_type,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Metadata []TimeseriesWidgetExpressionAlias `json:"metadata,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	OnRightYaxis *bool `json:"on_right_yaxis,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Projection *TimeseriesProjection `json:"projection,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query StructuredAnalysisQuery `json:"query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType TimeseriesWidgetLocalDatasetRequestRequestType `json:"request_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat TimeseriesWidgetLocalDatasetRequestResponseFormat `json:"response_format"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *TimeseriesRequestStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTimeseriesWidgetLocalDatasetRequest instantiates a new TimeseriesWidgetLocalDatasetRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesWidgetLocalDatasetRequest(query StructuredAnalysisQuery, requestType TimeseriesWidgetLocalDatasetRequestRequestType, responseFormat TimeseriesWidgetLocalDatasetRequestResponseFormat) *TimeseriesWidgetLocalDatasetRequest {
	this := TimeseriesWidgetLocalDatasetRequest{}
	this.Query = query
	this.RequestType = requestType
	this.ResponseFormat = responseFormat
	return &this
}

// NewTimeseriesWidgetLocalDatasetRequestWithDefaults instantiates a new TimeseriesWidgetLocalDatasetRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesWidgetLocalDatasetRequestWithDefaults() *TimeseriesWidgetLocalDatasetRequest {
	this := TimeseriesWidgetLocalDatasetRequest{}
	return &this
}

// GetDisplayType returns the DisplayType field value if set, zero value otherwise.
func (o *TimeseriesWidgetLocalDatasetRequest) GetDisplayType() WidgetDisplayType {
	if o == nil || o.DisplayType == nil {
		var ret WidgetDisplayType
		return ret
	}
	return *o.DisplayType
}

// GetDisplayTypeOk returns a tuple with the DisplayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetLocalDatasetRequest) GetDisplayTypeOk() (*WidgetDisplayType, bool) {
	if o == nil || o.DisplayType == nil {
		return nil, false
	}
	return o.DisplayType, true
}

// HasDisplayType returns a boolean if a field has been set.
func (o *TimeseriesWidgetLocalDatasetRequest) HasDisplayType() bool {
	return o != nil && o.DisplayType != nil
}

// SetDisplayType gets a reference to the given WidgetDisplayType and assigns it to the DisplayType field.
func (o *TimeseriesWidgetLocalDatasetRequest) SetDisplayType(v WidgetDisplayType) {
	o.DisplayType = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *TimeseriesWidgetLocalDatasetRequest) GetMetadata() []TimeseriesWidgetExpressionAlias {
	if o == nil || o.Metadata == nil {
		var ret []TimeseriesWidgetExpressionAlias
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetLocalDatasetRequest) GetMetadataOk() (*[]TimeseriesWidgetExpressionAlias, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TimeseriesWidgetLocalDatasetRequest) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given []TimeseriesWidgetExpressionAlias and assigns it to the Metadata field.
func (o *TimeseriesWidgetLocalDatasetRequest) SetMetadata(v []TimeseriesWidgetExpressionAlias) {
	o.Metadata = v
}

// GetOnRightYaxis returns the OnRightYaxis field value if set, zero value otherwise.
func (o *TimeseriesWidgetLocalDatasetRequest) GetOnRightYaxis() bool {
	if o == nil || o.OnRightYaxis == nil {
		var ret bool
		return ret
	}
	return *o.OnRightYaxis
}

// GetOnRightYaxisOk returns a tuple with the OnRightYaxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetLocalDatasetRequest) GetOnRightYaxisOk() (*bool, bool) {
	if o == nil || o.OnRightYaxis == nil {
		return nil, false
	}
	return o.OnRightYaxis, true
}

// HasOnRightYaxis returns a boolean if a field has been set.
func (o *TimeseriesWidgetLocalDatasetRequest) HasOnRightYaxis() bool {
	return o != nil && o.OnRightYaxis != nil
}

// SetOnRightYaxis gets a reference to the given bool and assigns it to the OnRightYaxis field.
func (o *TimeseriesWidgetLocalDatasetRequest) SetOnRightYaxis(v bool) {
	o.OnRightYaxis = &v
}

// GetProjection returns the Projection field value if set, zero value otherwise.
func (o *TimeseriesWidgetLocalDatasetRequest) GetProjection() TimeseriesProjection {
	if o == nil || o.Projection == nil {
		var ret TimeseriesProjection
		return ret
	}
	return *o.Projection
}

// GetProjectionOk returns a tuple with the Projection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetLocalDatasetRequest) GetProjectionOk() (*TimeseriesProjection, bool) {
	if o == nil || o.Projection == nil {
		return nil, false
	}
	return o.Projection, true
}

// HasProjection returns a boolean if a field has been set.
func (o *TimeseriesWidgetLocalDatasetRequest) HasProjection() bool {
	return o != nil && o.Projection != nil
}

// SetProjection gets a reference to the given TimeseriesProjection and assigns it to the Projection field.
func (o *TimeseriesWidgetLocalDatasetRequest) SetProjection(v TimeseriesProjection) {
	o.Projection = &v
}

// GetQuery returns the Query field value.
func (o *TimeseriesWidgetLocalDatasetRequest) GetQuery() StructuredAnalysisQuery {
	if o == nil {
		var ret StructuredAnalysisQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetLocalDatasetRequest) GetQueryOk() (*StructuredAnalysisQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *TimeseriesWidgetLocalDatasetRequest) SetQuery(v StructuredAnalysisQuery) {
	o.Query = v
}

// GetRequestType returns the RequestType field value.
func (o *TimeseriesWidgetLocalDatasetRequest) GetRequestType() TimeseriesWidgetLocalDatasetRequestRequestType {
	if o == nil {
		var ret TimeseriesWidgetLocalDatasetRequestRequestType
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetLocalDatasetRequest) GetRequestTypeOk() (*TimeseriesWidgetLocalDatasetRequestRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *TimeseriesWidgetLocalDatasetRequest) SetRequestType(v TimeseriesWidgetLocalDatasetRequestRequestType) {
	o.RequestType = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *TimeseriesWidgetLocalDatasetRequest) GetResponseFormat() TimeseriesWidgetLocalDatasetRequestResponseFormat {
	if o == nil {
		var ret TimeseriesWidgetLocalDatasetRequestResponseFormat
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetLocalDatasetRequest) GetResponseFormatOk() (*TimeseriesWidgetLocalDatasetRequestResponseFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *TimeseriesWidgetLocalDatasetRequest) SetResponseFormat(v TimeseriesWidgetLocalDatasetRequestResponseFormat) {
	o.ResponseFormat = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *TimeseriesWidgetLocalDatasetRequest) GetStyle() TimeseriesRequestStyle {
	if o == nil || o.Style == nil {
		var ret TimeseriesRequestStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetLocalDatasetRequest) GetStyleOk() (*TimeseriesRequestStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *TimeseriesWidgetLocalDatasetRequest) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given TimeseriesRequestStyle and assigns it to the Style field.
func (o *TimeseriesWidgetLocalDatasetRequest) SetStyle(v TimeseriesRequestStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesWidgetLocalDatasetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DisplayType != nil {
		toSerialize["display_type"] = o.DisplayType
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
	toSerialize["query"] = o.Query
	toSerialize["request_type"] = o.RequestType
	toSerialize["response_format"] = o.ResponseFormat
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesWidgetLocalDatasetRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayType    *WidgetDisplayType                                 `json:"display_type,omitempty"`
		Metadata       []TimeseriesWidgetExpressionAlias                  `json:"metadata,omitempty"`
		OnRightYaxis   *bool                                              `json:"on_right_yaxis,omitempty"`
		Projection     *TimeseriesProjection                              `json:"projection,omitempty"`
		Query          *StructuredAnalysisQuery                           `json:"query"`
		RequestType    *TimeseriesWidgetLocalDatasetRequestRequestType    `json:"request_type"`
		ResponseFormat *TimeseriesWidgetLocalDatasetRequestResponseFormat `json:"response_format"`
		Style          *TimeseriesRequestStyle                            `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.RequestType == nil {
		return fmt.Errorf("required field request_type missing")
	}
	if all.ResponseFormat == nil {
		return fmt.Errorf("required field response_format missing")
	}

	hasInvalidField := false
	if all.DisplayType != nil && !all.DisplayType.IsValid() {
		hasInvalidField = true
	} else {
		o.DisplayType = all.DisplayType
	}
	o.Metadata = all.Metadata
	o.OnRightYaxis = all.OnRightYaxis
	if all.Projection != nil && all.Projection.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Projection = all.Projection
	if all.Query.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Query = *all.Query
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
