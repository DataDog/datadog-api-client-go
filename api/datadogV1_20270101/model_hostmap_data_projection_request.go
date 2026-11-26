// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostmapDataProjectionRequest Test-only generated schema for the dashboard versioning proof of concept.
type HostmapDataProjectionRequest struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats []ThresholdFormattingRule `json:"conditional_formats,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit *float64 `json:"limit,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	NoGroupHosts *bool `json:"no_group_hosts,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	NoMetricHosts *bool `json:"no_metric_hosts,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Projection HostmapProjection `json:"projection"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query DatasetListQuery `json:"query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType HostmapDataProjectionRequestRequestType `json:"request_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *HostmapStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHostmapDataProjectionRequest instantiates a new HostmapDataProjectionRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHostmapDataProjectionRequest(projection HostmapProjection, query DatasetListQuery, requestType HostmapDataProjectionRequestRequestType) *HostmapDataProjectionRequest {
	this := HostmapDataProjectionRequest{}
	this.Projection = projection
	this.Query = query
	this.RequestType = requestType
	return &this
}

// NewHostmapDataProjectionRequestWithDefaults instantiates a new HostmapDataProjectionRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHostmapDataProjectionRequestWithDefaults() *HostmapDataProjectionRequest {
	this := HostmapDataProjectionRequest{}
	return &this
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *HostmapDataProjectionRequest) GetConditionalFormats() []ThresholdFormattingRule {
	if o == nil || o.ConditionalFormats == nil {
		var ret []ThresholdFormattingRule
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapDataProjectionRequest) GetConditionalFormatsOk() (*[]ThresholdFormattingRule, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *HostmapDataProjectionRequest) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []ThresholdFormattingRule and assigns it to the ConditionalFormats field.
func (o *HostmapDataProjectionRequest) SetConditionalFormats(v []ThresholdFormattingRule) {
	o.ConditionalFormats = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *HostmapDataProjectionRequest) GetLimit() float64 {
	if o == nil || o.Limit == nil {
		var ret float64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapDataProjectionRequest) GetLimitOk() (*float64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *HostmapDataProjectionRequest) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given float64 and assigns it to the Limit field.
func (o *HostmapDataProjectionRequest) SetLimit(v float64) {
	o.Limit = &v
}

// GetNoGroupHosts returns the NoGroupHosts field value if set, zero value otherwise.
func (o *HostmapDataProjectionRequest) GetNoGroupHosts() bool {
	if o == nil || o.NoGroupHosts == nil {
		var ret bool
		return ret
	}
	return *o.NoGroupHosts
}

// GetNoGroupHostsOk returns a tuple with the NoGroupHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapDataProjectionRequest) GetNoGroupHostsOk() (*bool, bool) {
	if o == nil || o.NoGroupHosts == nil {
		return nil, false
	}
	return o.NoGroupHosts, true
}

// HasNoGroupHosts returns a boolean if a field has been set.
func (o *HostmapDataProjectionRequest) HasNoGroupHosts() bool {
	return o != nil && o.NoGroupHosts != nil
}

// SetNoGroupHosts gets a reference to the given bool and assigns it to the NoGroupHosts field.
func (o *HostmapDataProjectionRequest) SetNoGroupHosts(v bool) {
	o.NoGroupHosts = &v
}

// GetNoMetricHosts returns the NoMetricHosts field value if set, zero value otherwise.
func (o *HostmapDataProjectionRequest) GetNoMetricHosts() bool {
	if o == nil || o.NoMetricHosts == nil {
		var ret bool
		return ret
	}
	return *o.NoMetricHosts
}

// GetNoMetricHostsOk returns a tuple with the NoMetricHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapDataProjectionRequest) GetNoMetricHostsOk() (*bool, bool) {
	if o == nil || o.NoMetricHosts == nil {
		return nil, false
	}
	return o.NoMetricHosts, true
}

// HasNoMetricHosts returns a boolean if a field has been set.
func (o *HostmapDataProjectionRequest) HasNoMetricHosts() bool {
	return o != nil && o.NoMetricHosts != nil
}

// SetNoMetricHosts gets a reference to the given bool and assigns it to the NoMetricHosts field.
func (o *HostmapDataProjectionRequest) SetNoMetricHosts(v bool) {
	o.NoMetricHosts = &v
}

// GetProjection returns the Projection field value.
func (o *HostmapDataProjectionRequest) GetProjection() HostmapProjection {
	if o == nil {
		var ret HostmapProjection
		return ret
	}
	return o.Projection
}

// GetProjectionOk returns a tuple with the Projection field value
// and a boolean to check if the value has been set.
func (o *HostmapDataProjectionRequest) GetProjectionOk() (*HostmapProjection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Projection, true
}

// SetProjection sets field value.
func (o *HostmapDataProjectionRequest) SetProjection(v HostmapProjection) {
	o.Projection = v
}

// GetQuery returns the Query field value.
func (o *HostmapDataProjectionRequest) GetQuery() DatasetListQuery {
	if o == nil {
		var ret DatasetListQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *HostmapDataProjectionRequest) GetQueryOk() (*DatasetListQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *HostmapDataProjectionRequest) SetQuery(v DatasetListQuery) {
	o.Query = v
}

// GetRequestType returns the RequestType field value.
func (o *HostmapDataProjectionRequest) GetRequestType() HostmapDataProjectionRequestRequestType {
	if o == nil {
		var ret HostmapDataProjectionRequestRequestType
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *HostmapDataProjectionRequest) GetRequestTypeOk() (*HostmapDataProjectionRequestRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *HostmapDataProjectionRequest) SetRequestType(v HostmapDataProjectionRequestRequestType) {
	o.RequestType = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *HostmapDataProjectionRequest) GetStyle() HostmapStyle {
	if o == nil || o.Style == nil {
		var ret HostmapStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapDataProjectionRequest) GetStyleOk() (*HostmapStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *HostmapDataProjectionRequest) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given HostmapStyle and assigns it to the Style field.
func (o *HostmapDataProjectionRequest) SetStyle(v HostmapStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HostmapDataProjectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.NoGroupHosts != nil {
		toSerialize["no_group_hosts"] = o.NoGroupHosts
	}
	if o.NoMetricHosts != nil {
		toSerialize["no_metric_hosts"] = o.NoMetricHosts
	}
	toSerialize["projection"] = o.Projection
	toSerialize["query"] = o.Query
	toSerialize["request_type"] = o.RequestType
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HostmapDataProjectionRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConditionalFormats []ThresholdFormattingRule                `json:"conditional_formats,omitempty"`
		Limit              *float64                                 `json:"limit,omitempty"`
		NoGroupHosts       *bool                                    `json:"no_group_hosts,omitempty"`
		NoMetricHosts      *bool                                    `json:"no_metric_hosts,omitempty"`
		Projection         *HostmapProjection                       `json:"projection"`
		Query              *DatasetListQuery                        `json:"query"`
		RequestType        *HostmapDataProjectionRequestRequestType `json:"request_type"`
		Style              *HostmapStyle                            `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Projection == nil {
		return fmt.Errorf("required field projection missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.RequestType == nil {
		return fmt.Errorf("required field request_type missing")
	}

	hasInvalidField := false
	o.ConditionalFormats = all.ConditionalFormats
	o.Limit = all.Limit
	o.NoGroupHosts = all.NoGroupHosts
	o.NoMetricHosts = all.NoMetricHosts
	if all.Projection.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Projection = *all.Projection
	if all.Query.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Query = *all.Query
	if !all.RequestType.IsValid() {
		hasInvalidField = true
	} else {
		o.RequestType = *all.RequestType
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
