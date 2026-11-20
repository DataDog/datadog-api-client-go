// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesWidgetLegacyEventRequest Test-only generated schema for the dashboard versioning proof of concept.
type TimeseriesWidgetLegacyEventRequest struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	DisplayType *WidgetDisplayType `json:"display_type,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	LegacyEventQuery LegacyEventManagementQuery `json:"legacy_event_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Metadata []TimeseriesWidgetExpressionAlias `json:"metadata,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	OnRightYaxis *bool `json:"on_right_yaxis,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *TimeseriesRequestStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTimeseriesWidgetLegacyEventRequest instantiates a new TimeseriesWidgetLegacyEventRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesWidgetLegacyEventRequest(legacyEventQuery LegacyEventManagementQuery) *TimeseriesWidgetLegacyEventRequest {
	this := TimeseriesWidgetLegacyEventRequest{}
	this.LegacyEventQuery = legacyEventQuery
	return &this
}

// NewTimeseriesWidgetLegacyEventRequestWithDefaults instantiates a new TimeseriesWidgetLegacyEventRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesWidgetLegacyEventRequestWithDefaults() *TimeseriesWidgetLegacyEventRequest {
	this := TimeseriesWidgetLegacyEventRequest{}
	return &this
}

// GetDisplayType returns the DisplayType field value if set, zero value otherwise.
func (o *TimeseriesWidgetLegacyEventRequest) GetDisplayType() WidgetDisplayType {
	if o == nil || o.DisplayType == nil {
		var ret WidgetDisplayType
		return ret
	}
	return *o.DisplayType
}

// GetDisplayTypeOk returns a tuple with the DisplayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetLegacyEventRequest) GetDisplayTypeOk() (*WidgetDisplayType, bool) {
	if o == nil || o.DisplayType == nil {
		return nil, false
	}
	return o.DisplayType, true
}

// HasDisplayType returns a boolean if a field has been set.
func (o *TimeseriesWidgetLegacyEventRequest) HasDisplayType() bool {
	return o != nil && o.DisplayType != nil
}

// SetDisplayType gets a reference to the given WidgetDisplayType and assigns it to the DisplayType field.
func (o *TimeseriesWidgetLegacyEventRequest) SetDisplayType(v WidgetDisplayType) {
	o.DisplayType = &v
}

// GetLegacyEventQuery returns the LegacyEventQuery field value.
func (o *TimeseriesWidgetLegacyEventRequest) GetLegacyEventQuery() LegacyEventManagementQuery {
	if o == nil {
		var ret LegacyEventManagementQuery
		return ret
	}
	return o.LegacyEventQuery
}

// GetLegacyEventQueryOk returns a tuple with the LegacyEventQuery field value
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetLegacyEventRequest) GetLegacyEventQueryOk() (*LegacyEventManagementQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LegacyEventQuery, true
}

// SetLegacyEventQuery sets field value.
func (o *TimeseriesWidgetLegacyEventRequest) SetLegacyEventQuery(v LegacyEventManagementQuery) {
	o.LegacyEventQuery = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *TimeseriesWidgetLegacyEventRequest) GetMetadata() []TimeseriesWidgetExpressionAlias {
	if o == nil || o.Metadata == nil {
		var ret []TimeseriesWidgetExpressionAlias
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetLegacyEventRequest) GetMetadataOk() (*[]TimeseriesWidgetExpressionAlias, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TimeseriesWidgetLegacyEventRequest) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given []TimeseriesWidgetExpressionAlias and assigns it to the Metadata field.
func (o *TimeseriesWidgetLegacyEventRequest) SetMetadata(v []TimeseriesWidgetExpressionAlias) {
	o.Metadata = v
}

// GetOnRightYaxis returns the OnRightYaxis field value if set, zero value otherwise.
func (o *TimeseriesWidgetLegacyEventRequest) GetOnRightYaxis() bool {
	if o == nil || o.OnRightYaxis == nil {
		var ret bool
		return ret
	}
	return *o.OnRightYaxis
}

// GetOnRightYaxisOk returns a tuple with the OnRightYaxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetLegacyEventRequest) GetOnRightYaxisOk() (*bool, bool) {
	if o == nil || o.OnRightYaxis == nil {
		return nil, false
	}
	return o.OnRightYaxis, true
}

// HasOnRightYaxis returns a boolean if a field has been set.
func (o *TimeseriesWidgetLegacyEventRequest) HasOnRightYaxis() bool {
	return o != nil && o.OnRightYaxis != nil
}

// SetOnRightYaxis gets a reference to the given bool and assigns it to the OnRightYaxis field.
func (o *TimeseriesWidgetLegacyEventRequest) SetOnRightYaxis(v bool) {
	o.OnRightYaxis = &v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *TimeseriesWidgetLegacyEventRequest) GetStyle() TimeseriesRequestStyle {
	if o == nil || o.Style == nil {
		var ret TimeseriesRequestStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetLegacyEventRequest) GetStyleOk() (*TimeseriesRequestStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *TimeseriesWidgetLegacyEventRequest) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given TimeseriesRequestStyle and assigns it to the Style field.
func (o *TimeseriesWidgetLegacyEventRequest) SetStyle(v TimeseriesRequestStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesWidgetLegacyEventRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DisplayType != nil {
		toSerialize["display_type"] = o.DisplayType
	}
	toSerialize["legacy_event_query"] = o.LegacyEventQuery
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
func (o *TimeseriesWidgetLegacyEventRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayType      *WidgetDisplayType                `json:"display_type,omitempty"`
		LegacyEventQuery *LegacyEventManagementQuery       `json:"legacy_event_query"`
		Metadata         []TimeseriesWidgetExpressionAlias `json:"metadata,omitempty"`
		OnRightYaxis     *bool                             `json:"on_right_yaxis,omitempty"`
		Style            *TimeseriesRequestStyle           `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.LegacyEventQuery == nil {
		return fmt.Errorf("required field legacy_event_query missing")
	}

	hasInvalidField := false
	if all.DisplayType != nil && !all.DisplayType.IsValid() {
		hasInvalidField = true
	} else {
		o.DisplayType = all.DisplayType
	}
	if all.LegacyEventQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LegacyEventQuery = *all.LegacyEventQuery
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
