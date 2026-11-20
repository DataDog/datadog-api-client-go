// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatmapBackendRequest20270101OneOf4458140576 Test-only generated schema for the dashboard versioning proof of concept.
type HeatmapBackendRequest20270101OneOf4458140576 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	NumberFormat *WidgetNumberFormat_20270101 `json:"number_format,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query HeatmapBackendRequestAnyOfAnyOfItemQuery_20270101 `json:"query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType HeatmapBackendRequestAnyOfAnyOfItemRequestType_20270101 `json:"request_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *RequestStyle_20270101 `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatmapBackendRequest20270101OneOf4458140576 instantiates a new HeatmapBackendRequest20270101OneOf4458140576 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatmapBackendRequest20270101OneOf4458140576(query HeatmapBackendRequestAnyOfAnyOfItemQuery_20270101, requestType HeatmapBackendRequestAnyOfAnyOfItemRequestType_20270101) *HeatmapBackendRequest20270101OneOf4458140576 {
	this := HeatmapBackendRequest20270101OneOf4458140576{}
	this.Query = query
	this.RequestType = requestType
	return &this
}

// NewHeatmapBackendRequest20270101OneOf4458140576WithDefaults instantiates a new HeatmapBackendRequest20270101OneOf4458140576 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatmapBackendRequest20270101OneOf4458140576WithDefaults() *HeatmapBackendRequest20270101OneOf4458140576 {
	this := HeatmapBackendRequest20270101OneOf4458140576{}
	return &this
}

// GetNumberFormat returns the NumberFormat field value if set, zero value otherwise.
func (o *HeatmapBackendRequest20270101OneOf4458140576) GetNumberFormat() WidgetNumberFormat_20270101 {
	if o == nil || o.NumberFormat == nil {
		var ret WidgetNumberFormat_20270101
		return ret
	}
	return *o.NumberFormat
}

// GetNumberFormatOk returns a tuple with the NumberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatmapBackendRequest20270101OneOf4458140576) GetNumberFormatOk() (*WidgetNumberFormat_20270101, bool) {
	if o == nil || o.NumberFormat == nil {
		return nil, false
	}
	return o.NumberFormat, true
}

// HasNumberFormat returns a boolean if a field has been set.
func (o *HeatmapBackendRequest20270101OneOf4458140576) HasNumberFormat() bool {
	return o != nil && o.NumberFormat != nil
}

// SetNumberFormat gets a reference to the given WidgetNumberFormat_20270101 and assigns it to the NumberFormat field.
func (o *HeatmapBackendRequest20270101OneOf4458140576) SetNumberFormat(v WidgetNumberFormat_20270101) {
	o.NumberFormat = &v
}

// GetQuery returns the Query field value.
func (o *HeatmapBackendRequest20270101OneOf4458140576) GetQuery() HeatmapBackendRequestAnyOfAnyOfItemQuery_20270101 {
	if o == nil {
		var ret HeatmapBackendRequestAnyOfAnyOfItemQuery_20270101
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *HeatmapBackendRequest20270101OneOf4458140576) GetQueryOk() (*HeatmapBackendRequestAnyOfAnyOfItemQuery_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *HeatmapBackendRequest20270101OneOf4458140576) SetQuery(v HeatmapBackendRequestAnyOfAnyOfItemQuery_20270101) {
	o.Query = v
}

// GetRequestType returns the RequestType field value.
func (o *HeatmapBackendRequest20270101OneOf4458140576) GetRequestType() HeatmapBackendRequestAnyOfAnyOfItemRequestType_20270101 {
	if o == nil {
		var ret HeatmapBackendRequestAnyOfAnyOfItemRequestType_20270101
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *HeatmapBackendRequest20270101OneOf4458140576) GetRequestTypeOk() (*HeatmapBackendRequestAnyOfAnyOfItemRequestType_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *HeatmapBackendRequest20270101OneOf4458140576) SetRequestType(v HeatmapBackendRequestAnyOfAnyOfItemRequestType_20270101) {
	o.RequestType = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *HeatmapBackendRequest20270101OneOf4458140576) GetStyle() RequestStyle_20270101 {
	if o == nil || o.Style == nil {
		var ret RequestStyle_20270101
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatmapBackendRequest20270101OneOf4458140576) GetStyleOk() (*RequestStyle_20270101, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *HeatmapBackendRequest20270101OneOf4458140576) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given RequestStyle_20270101 and assigns it to the Style field.
func (o *HeatmapBackendRequest20270101OneOf4458140576) SetStyle(v RequestStyle_20270101) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatmapBackendRequest20270101OneOf4458140576) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.NumberFormat != nil {
		toSerialize["number_format"] = o.NumberFormat
	}
	toSerialize["query"] = o.Query
	toSerialize["request_type"] = o.RequestType
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HeatmapBackendRequest20270101OneOf4458140576) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NumberFormat *WidgetNumberFormat_20270101                             `json:"number_format,omitempty"`
		Query        *HeatmapBackendRequestAnyOfAnyOfItemQuery_20270101       `json:"query"`
		RequestType  *HeatmapBackendRequestAnyOfAnyOfItemRequestType_20270101 `json:"request_type"`
		Style        *RequestStyle_20270101                                   `json:"style,omitempty"`
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

	hasInvalidField := false
	o.NumberFormat = all.NumberFormat
	o.Query = *all.Query
	if !all.RequestType.IsValid() {
		hasInvalidField = true
	} else {
		o.RequestType = *all.RequestType
	}
	o.Style = all.Style

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
