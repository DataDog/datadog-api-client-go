// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatmapBackendRequest20270101OneOf4458138496 Test-only generated schema for the dashboard versioning proof of concept.
type HeatmapBackendRequest20270101OneOf4458138496 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryOptions map[string]DashboardInlineSchema5 `json:"_query_options,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Q string `json:"q"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *RequestStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatmapBackendRequest20270101OneOf4458138496 instantiates a new HeatmapBackendRequest20270101OneOf4458138496 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatmapBackendRequest20270101OneOf4458138496(q string) *HeatmapBackendRequest20270101OneOf4458138496 {
	this := HeatmapBackendRequest20270101OneOf4458138496{}
	this.Q = q
	return &this
}

// NewHeatmapBackendRequest20270101OneOf4458138496WithDefaults instantiates a new HeatmapBackendRequest20270101OneOf4458138496 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatmapBackendRequest20270101OneOf4458138496WithDefaults() *HeatmapBackendRequest20270101OneOf4458138496 {
	this := HeatmapBackendRequest20270101OneOf4458138496{}
	return &this
}

// GetQueryOptions returns the QueryOptions field value if set, zero value otherwise.
func (o *HeatmapBackendRequest20270101OneOf4458138496) GetQueryOptions() map[string]DashboardInlineSchema5 {
	if o == nil || o.QueryOptions == nil {
		var ret map[string]DashboardInlineSchema5
		return ret
	}
	return o.QueryOptions
}

// GetQueryOptionsOk returns a tuple with the QueryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatmapBackendRequest20270101OneOf4458138496) GetQueryOptionsOk() (*map[string]DashboardInlineSchema5, bool) {
	if o == nil || o.QueryOptions == nil {
		return nil, false
	}
	return &o.QueryOptions, true
}

// HasQueryOptions returns a boolean if a field has been set.
func (o *HeatmapBackendRequest20270101OneOf4458138496) HasQueryOptions() bool {
	return o != nil && o.QueryOptions != nil
}

// SetQueryOptions gets a reference to the given map[string]DashboardInlineSchema5 and assigns it to the QueryOptions field.
func (o *HeatmapBackendRequest20270101OneOf4458138496) SetQueryOptions(v map[string]DashboardInlineSchema5) {
	o.QueryOptions = v
}

// GetQ returns the Q field value.
func (o *HeatmapBackendRequest20270101OneOf4458138496) GetQ() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Q
}

// GetQOk returns a tuple with the Q field value
// and a boolean to check if the value has been set.
func (o *HeatmapBackendRequest20270101OneOf4458138496) GetQOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Q, true
}

// SetQ sets field value.
func (o *HeatmapBackendRequest20270101OneOf4458138496) SetQ(v string) {
	o.Q = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *HeatmapBackendRequest20270101OneOf4458138496) GetStyle() RequestStyle {
	if o == nil || o.Style == nil {
		var ret RequestStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatmapBackendRequest20270101OneOf4458138496) GetStyleOk() (*RequestStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *HeatmapBackendRequest20270101OneOf4458138496) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given RequestStyle and assigns it to the Style field.
func (o *HeatmapBackendRequest20270101OneOf4458138496) SetStyle(v RequestStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatmapBackendRequest20270101OneOf4458138496) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.QueryOptions != nil {
		toSerialize["_query_options"] = o.QueryOptions
	}
	toSerialize["q"] = o.Q
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HeatmapBackendRequest20270101OneOf4458138496) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		QueryOptions map[string]DashboardInlineSchema5 `json:"_query_options,omitempty"`
		Q            *string                           `json:"q"`
		Style        *RequestStyle                     `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Q == nil {
		return fmt.Errorf("required field q missing")
	}

	hasInvalidField := false
	o.QueryOptions = all.QueryOptions
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
