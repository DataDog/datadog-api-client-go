// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesWidgetMetricRequest Test-only generated schema for the dashboard versioning proof of concept.
type TimeseriesWidgetMetricRequest struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryOptions map[string]DashboardInlineSchema50 `json:"_query_options,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DisplayType *WidgetDisplayType `json:"display_type,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Metadata []TimeseriesWidgetExpressionAlias `json:"metadata,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	OnRightYaxis *bool `json:"on_right_yaxis,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Q string `json:"q"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *TimeseriesRequestStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTimeseriesWidgetMetricRequest instantiates a new TimeseriesWidgetMetricRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesWidgetMetricRequest(q string) *TimeseriesWidgetMetricRequest {
	this := TimeseriesWidgetMetricRequest{}
	this.Q = q
	return &this
}

// NewTimeseriesWidgetMetricRequestWithDefaults instantiates a new TimeseriesWidgetMetricRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesWidgetMetricRequestWithDefaults() *TimeseriesWidgetMetricRequest {
	this := TimeseriesWidgetMetricRequest{}
	return &this
}

// GetQueryOptions returns the QueryOptions field value if set, zero value otherwise.
func (o *TimeseriesWidgetMetricRequest) GetQueryOptions() map[string]DashboardInlineSchema50 {
	if o == nil || o.QueryOptions == nil {
		var ret map[string]DashboardInlineSchema50
		return ret
	}
	return o.QueryOptions
}

// GetQueryOptionsOk returns a tuple with the QueryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetMetricRequest) GetQueryOptionsOk() (*map[string]DashboardInlineSchema50, bool) {
	if o == nil || o.QueryOptions == nil {
		return nil, false
	}
	return &o.QueryOptions, true
}

// HasQueryOptions returns a boolean if a field has been set.
func (o *TimeseriesWidgetMetricRequest) HasQueryOptions() bool {
	return o != nil && o.QueryOptions != nil
}

// SetQueryOptions gets a reference to the given map[string]DashboardInlineSchema50 and assigns it to the QueryOptions field.
func (o *TimeseriesWidgetMetricRequest) SetQueryOptions(v map[string]DashboardInlineSchema50) {
	o.QueryOptions = v
}

// GetDisplayType returns the DisplayType field value if set, zero value otherwise.
func (o *TimeseriesWidgetMetricRequest) GetDisplayType() WidgetDisplayType {
	if o == nil || o.DisplayType == nil {
		var ret WidgetDisplayType
		return ret
	}
	return *o.DisplayType
}

// GetDisplayTypeOk returns a tuple with the DisplayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetMetricRequest) GetDisplayTypeOk() (*WidgetDisplayType, bool) {
	if o == nil || o.DisplayType == nil {
		return nil, false
	}
	return o.DisplayType, true
}

// HasDisplayType returns a boolean if a field has been set.
func (o *TimeseriesWidgetMetricRequest) HasDisplayType() bool {
	return o != nil && o.DisplayType != nil
}

// SetDisplayType gets a reference to the given WidgetDisplayType and assigns it to the DisplayType field.
func (o *TimeseriesWidgetMetricRequest) SetDisplayType(v WidgetDisplayType) {
	o.DisplayType = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *TimeseriesWidgetMetricRequest) GetMetadata() []TimeseriesWidgetExpressionAlias {
	if o == nil || o.Metadata == nil {
		var ret []TimeseriesWidgetExpressionAlias
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetMetricRequest) GetMetadataOk() (*[]TimeseriesWidgetExpressionAlias, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TimeseriesWidgetMetricRequest) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given []TimeseriesWidgetExpressionAlias and assigns it to the Metadata field.
func (o *TimeseriesWidgetMetricRequest) SetMetadata(v []TimeseriesWidgetExpressionAlias) {
	o.Metadata = v
}

// GetOnRightYaxis returns the OnRightYaxis field value if set, zero value otherwise.
func (o *TimeseriesWidgetMetricRequest) GetOnRightYaxis() bool {
	if o == nil || o.OnRightYaxis == nil {
		var ret bool
		return ret
	}
	return *o.OnRightYaxis
}

// GetOnRightYaxisOk returns a tuple with the OnRightYaxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetMetricRequest) GetOnRightYaxisOk() (*bool, bool) {
	if o == nil || o.OnRightYaxis == nil {
		return nil, false
	}
	return o.OnRightYaxis, true
}

// HasOnRightYaxis returns a boolean if a field has been set.
func (o *TimeseriesWidgetMetricRequest) HasOnRightYaxis() bool {
	return o != nil && o.OnRightYaxis != nil
}

// SetOnRightYaxis gets a reference to the given bool and assigns it to the OnRightYaxis field.
func (o *TimeseriesWidgetMetricRequest) SetOnRightYaxis(v bool) {
	o.OnRightYaxis = &v
}

// GetQ returns the Q field value.
func (o *TimeseriesWidgetMetricRequest) GetQ() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Q
}

// GetQOk returns a tuple with the Q field value
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetMetricRequest) GetQOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Q, true
}

// SetQ sets field value.
func (o *TimeseriesWidgetMetricRequest) SetQ(v string) {
	o.Q = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *TimeseriesWidgetMetricRequest) GetStyle() TimeseriesRequestStyle {
	if o == nil || o.Style == nil {
		var ret TimeseriesRequestStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetMetricRequest) GetStyleOk() (*TimeseriesRequestStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *TimeseriesWidgetMetricRequest) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given TimeseriesRequestStyle and assigns it to the Style field.
func (o *TimeseriesWidgetMetricRequest) SetStyle(v TimeseriesRequestStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesWidgetMetricRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.QueryOptions != nil {
		toSerialize["_query_options"] = o.QueryOptions
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
	toSerialize["q"] = o.Q
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesWidgetMetricRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		QueryOptions map[string]DashboardInlineSchema50 `json:"_query_options,omitempty"`
		DisplayType  *WidgetDisplayType                 `json:"display_type,omitempty"`
		Metadata     []TimeseriesWidgetExpressionAlias  `json:"metadata,omitempty"`
		OnRightYaxis *bool                              `json:"on_right_yaxis,omitempty"`
		Q            *string                            `json:"q"`
		Style        *TimeseriesRequestStyle            `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Q == nil {
		return fmt.Errorf("required field q missing")
	}

	hasInvalidField := false
	o.QueryOptions = all.QueryOptions
	if all.DisplayType != nil && !all.DisplayType.IsValid() {
		hasInvalidField = true
	} else {
		o.DisplayType = all.DisplayType
	}
	o.Metadata = all.Metadata
	o.OnRightYaxis = all.OnRightYaxis
	o.Q = *all.Q
	if all.Style != nil && all.Style.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Style = all.Style

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
