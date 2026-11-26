// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesWidgetAppsecSpanRequest Test-only generated schema for the dashboard versioning proof of concept.
type TimeseriesWidgetAppsecSpanRequest struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	AppsecspanQuery LogQueryDefinition `json:"appsecspan_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DisplayType *WidgetDisplayType `json:"display_type,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Metadata []TimeseriesWidgetExpressionAlias `json:"metadata,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	OnRightYaxis *bool `json:"on_right_yaxis,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *TimeseriesRequestStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTimeseriesWidgetAppsecSpanRequest instantiates a new TimeseriesWidgetAppsecSpanRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesWidgetAppsecSpanRequest(appsecspanQuery LogQueryDefinition) *TimeseriesWidgetAppsecSpanRequest {
	this := TimeseriesWidgetAppsecSpanRequest{}
	this.AppsecspanQuery = appsecspanQuery
	return &this
}

// NewTimeseriesWidgetAppsecSpanRequestWithDefaults instantiates a new TimeseriesWidgetAppsecSpanRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesWidgetAppsecSpanRequestWithDefaults() *TimeseriesWidgetAppsecSpanRequest {
	this := TimeseriesWidgetAppsecSpanRequest{}
	return &this
}

// GetAppsecspanQuery returns the AppsecspanQuery field value.
func (o *TimeseriesWidgetAppsecSpanRequest) GetAppsecspanQuery() LogQueryDefinition {
	if o == nil {
		var ret LogQueryDefinition
		return ret
	}
	return o.AppsecspanQuery
}

// GetAppsecspanQueryOk returns a tuple with the AppsecspanQuery field value
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetAppsecSpanRequest) GetAppsecspanQueryOk() (*LogQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppsecspanQuery, true
}

// SetAppsecspanQuery sets field value.
func (o *TimeseriesWidgetAppsecSpanRequest) SetAppsecspanQuery(v LogQueryDefinition) {
	o.AppsecspanQuery = v
}

// GetDisplayType returns the DisplayType field value if set, zero value otherwise.
func (o *TimeseriesWidgetAppsecSpanRequest) GetDisplayType() WidgetDisplayType {
	if o == nil || o.DisplayType == nil {
		var ret WidgetDisplayType
		return ret
	}
	return *o.DisplayType
}

// GetDisplayTypeOk returns a tuple with the DisplayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetAppsecSpanRequest) GetDisplayTypeOk() (*WidgetDisplayType, bool) {
	if o == nil || o.DisplayType == nil {
		return nil, false
	}
	return o.DisplayType, true
}

// HasDisplayType returns a boolean if a field has been set.
func (o *TimeseriesWidgetAppsecSpanRequest) HasDisplayType() bool {
	return o != nil && o.DisplayType != nil
}

// SetDisplayType gets a reference to the given WidgetDisplayType and assigns it to the DisplayType field.
func (o *TimeseriesWidgetAppsecSpanRequest) SetDisplayType(v WidgetDisplayType) {
	o.DisplayType = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *TimeseriesWidgetAppsecSpanRequest) GetMetadata() []TimeseriesWidgetExpressionAlias {
	if o == nil || o.Metadata == nil {
		var ret []TimeseriesWidgetExpressionAlias
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetAppsecSpanRequest) GetMetadataOk() (*[]TimeseriesWidgetExpressionAlias, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TimeseriesWidgetAppsecSpanRequest) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given []TimeseriesWidgetExpressionAlias and assigns it to the Metadata field.
func (o *TimeseriesWidgetAppsecSpanRequest) SetMetadata(v []TimeseriesWidgetExpressionAlias) {
	o.Metadata = v
}

// GetOnRightYaxis returns the OnRightYaxis field value if set, zero value otherwise.
func (o *TimeseriesWidgetAppsecSpanRequest) GetOnRightYaxis() bool {
	if o == nil || o.OnRightYaxis == nil {
		var ret bool
		return ret
	}
	return *o.OnRightYaxis
}

// GetOnRightYaxisOk returns a tuple with the OnRightYaxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetAppsecSpanRequest) GetOnRightYaxisOk() (*bool, bool) {
	if o == nil || o.OnRightYaxis == nil {
		return nil, false
	}
	return o.OnRightYaxis, true
}

// HasOnRightYaxis returns a boolean if a field has been set.
func (o *TimeseriesWidgetAppsecSpanRequest) HasOnRightYaxis() bool {
	return o != nil && o.OnRightYaxis != nil
}

// SetOnRightYaxis gets a reference to the given bool and assigns it to the OnRightYaxis field.
func (o *TimeseriesWidgetAppsecSpanRequest) SetOnRightYaxis(v bool) {
	o.OnRightYaxis = &v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *TimeseriesWidgetAppsecSpanRequest) GetStyle() TimeseriesRequestStyle {
	if o == nil || o.Style == nil {
		var ret TimeseriesRequestStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetAppsecSpanRequest) GetStyleOk() (*TimeseriesRequestStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *TimeseriesWidgetAppsecSpanRequest) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given TimeseriesRequestStyle and assigns it to the Style field.
func (o *TimeseriesWidgetAppsecSpanRequest) SetStyle(v TimeseriesRequestStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesWidgetAppsecSpanRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["appsecspan_query"] = o.AppsecspanQuery
	if o.DisplayType != nil {
		toSerialize["display_type"] = o.DisplayType
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.OnRightYaxis != nil {
		toSerialize["on_right_yaxis"] = o.OnRightYaxis
	}
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesWidgetAppsecSpanRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AppsecspanQuery *LogQueryDefinition               `json:"appsecspan_query"`
		DisplayType     *WidgetDisplayType                `json:"display_type,omitempty"`
		Metadata        []TimeseriesWidgetExpressionAlias `json:"metadata,omitempty"`
		OnRightYaxis    *bool                             `json:"on_right_yaxis,omitempty"`
		Style           *TimeseriesRequestStyle           `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AppsecspanQuery == nil {
		return fmt.Errorf("required field appsecspan_query missing")
	}

	hasInvalidField := false
	if all.AppsecspanQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AppsecspanQuery = *all.AppsecspanQuery
	if all.DisplayType != nil && !all.DisplayType.IsValid() {
		hasInvalidField = true
	} else {
		o.DisplayType = all.DisplayType
	}
	o.Metadata = all.Metadata
	o.OnRightYaxis = all.OnRightYaxis
	if all.Style != nil && all.Style.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Style = all.Style

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
