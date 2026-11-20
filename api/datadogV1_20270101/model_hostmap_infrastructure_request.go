// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostmapInfrastructureRequest Test-only generated schema for the dashboard versioning proof of concept.
type HostmapInfrastructureRequest struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Child *HostmapInfrastructureRequest `json:"child,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats []ThresholdFormattingRule `json:"conditional_formats,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Enrichments []HostmapScalarFormulaRequest `json:"enrichments"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Filter *string `json:"filter,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy []HostmapGroupBy `json:"group_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	NoGroupHosts *bool `json:"no_group_hosts,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	NoMetricHosts *bool `json:"no_metric_hosts,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	NodeType HostmapNodeType `json:"node_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType HostmapInfrastructureRequestRequestType `json:"request_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *HostmapStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHostmapInfrastructureRequest instantiates a new HostmapInfrastructureRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHostmapInfrastructureRequest(enrichments []HostmapScalarFormulaRequest, nodeType HostmapNodeType, requestType HostmapInfrastructureRequestRequestType) *HostmapInfrastructureRequest {
	this := HostmapInfrastructureRequest{}
	this.Enrichments = enrichments
	this.NodeType = nodeType
	this.RequestType = requestType
	return &this
}

// NewHostmapInfrastructureRequestWithDefaults instantiates a new HostmapInfrastructureRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHostmapInfrastructureRequestWithDefaults() *HostmapInfrastructureRequest {
	this := HostmapInfrastructureRequest{}
	return &this
}

// GetChild returns the Child field value if set, zero value otherwise.
func (o *HostmapInfrastructureRequest) GetChild() HostmapInfrastructureRequest {
	if o == nil || o.Child == nil {
		var ret HostmapInfrastructureRequest
		return ret
	}
	return *o.Child
}

// GetChildOk returns a tuple with the Child field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapInfrastructureRequest) GetChildOk() (*HostmapInfrastructureRequest, bool) {
	if o == nil || o.Child == nil {
		return nil, false
	}
	return o.Child, true
}

// HasChild returns a boolean if a field has been set.
func (o *HostmapInfrastructureRequest) HasChild() bool {
	return o != nil && o.Child != nil
}

// SetChild gets a reference to the given HostmapInfrastructureRequest and assigns it to the Child field.
func (o *HostmapInfrastructureRequest) SetChild(v HostmapInfrastructureRequest) {
	o.Child = &v
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *HostmapInfrastructureRequest) GetConditionalFormats() []ThresholdFormattingRule {
	if o == nil || o.ConditionalFormats == nil {
		var ret []ThresholdFormattingRule
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapInfrastructureRequest) GetConditionalFormatsOk() (*[]ThresholdFormattingRule, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *HostmapInfrastructureRequest) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []ThresholdFormattingRule and assigns it to the ConditionalFormats field.
func (o *HostmapInfrastructureRequest) SetConditionalFormats(v []ThresholdFormattingRule) {
	o.ConditionalFormats = v
}

// GetEnrichments returns the Enrichments field value.
func (o *HostmapInfrastructureRequest) GetEnrichments() []HostmapScalarFormulaRequest {
	if o == nil {
		var ret []HostmapScalarFormulaRequest
		return ret
	}
	return o.Enrichments
}

// GetEnrichmentsOk returns a tuple with the Enrichments field value
// and a boolean to check if the value has been set.
func (o *HostmapInfrastructureRequest) GetEnrichmentsOk() (*[]HostmapScalarFormulaRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enrichments, true
}

// SetEnrichments sets field value.
func (o *HostmapInfrastructureRequest) SetEnrichments(v []HostmapScalarFormulaRequest) {
	o.Enrichments = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *HostmapInfrastructureRequest) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapInfrastructureRequest) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *HostmapInfrastructureRequest) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *HostmapInfrastructureRequest) SetFilter(v string) {
	o.Filter = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *HostmapInfrastructureRequest) GetGroupBy() []HostmapGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []HostmapGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapInfrastructureRequest) GetGroupByOk() (*[]HostmapGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *HostmapInfrastructureRequest) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []HostmapGroupBy and assigns it to the GroupBy field.
func (o *HostmapInfrastructureRequest) SetGroupBy(v []HostmapGroupBy) {
	o.GroupBy = v
}

// GetNoGroupHosts returns the NoGroupHosts field value if set, zero value otherwise.
func (o *HostmapInfrastructureRequest) GetNoGroupHosts() bool {
	if o == nil || o.NoGroupHosts == nil {
		var ret bool
		return ret
	}
	return *o.NoGroupHosts
}

// GetNoGroupHostsOk returns a tuple with the NoGroupHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapInfrastructureRequest) GetNoGroupHostsOk() (*bool, bool) {
	if o == nil || o.NoGroupHosts == nil {
		return nil, false
	}
	return o.NoGroupHosts, true
}

// HasNoGroupHosts returns a boolean if a field has been set.
func (o *HostmapInfrastructureRequest) HasNoGroupHosts() bool {
	return o != nil && o.NoGroupHosts != nil
}

// SetNoGroupHosts gets a reference to the given bool and assigns it to the NoGroupHosts field.
func (o *HostmapInfrastructureRequest) SetNoGroupHosts(v bool) {
	o.NoGroupHosts = &v
}

// GetNoMetricHosts returns the NoMetricHosts field value if set, zero value otherwise.
func (o *HostmapInfrastructureRequest) GetNoMetricHosts() bool {
	if o == nil || o.NoMetricHosts == nil {
		var ret bool
		return ret
	}
	return *o.NoMetricHosts
}

// GetNoMetricHostsOk returns a tuple with the NoMetricHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapInfrastructureRequest) GetNoMetricHostsOk() (*bool, bool) {
	if o == nil || o.NoMetricHosts == nil {
		return nil, false
	}
	return o.NoMetricHosts, true
}

// HasNoMetricHosts returns a boolean if a field has been set.
func (o *HostmapInfrastructureRequest) HasNoMetricHosts() bool {
	return o != nil && o.NoMetricHosts != nil
}

// SetNoMetricHosts gets a reference to the given bool and assigns it to the NoMetricHosts field.
func (o *HostmapInfrastructureRequest) SetNoMetricHosts(v bool) {
	o.NoMetricHosts = &v
}

// GetNodeType returns the NodeType field value.
func (o *HostmapInfrastructureRequest) GetNodeType() HostmapNodeType {
	if o == nil {
		var ret HostmapNodeType
		return ret
	}
	return o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value
// and a boolean to check if the value has been set.
func (o *HostmapInfrastructureRequest) GetNodeTypeOk() (*HostmapNodeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeType, true
}

// SetNodeType sets field value.
func (o *HostmapInfrastructureRequest) SetNodeType(v HostmapNodeType) {
	o.NodeType = v
}

// GetRequestType returns the RequestType field value.
func (o *HostmapInfrastructureRequest) GetRequestType() HostmapInfrastructureRequestRequestType {
	if o == nil {
		var ret HostmapInfrastructureRequestRequestType
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *HostmapInfrastructureRequest) GetRequestTypeOk() (*HostmapInfrastructureRequestRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *HostmapInfrastructureRequest) SetRequestType(v HostmapInfrastructureRequestRequestType) {
	o.RequestType = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *HostmapInfrastructureRequest) GetStyle() HostmapStyle {
	if o == nil || o.Style == nil {
		var ret HostmapStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapInfrastructureRequest) GetStyleOk() (*HostmapStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *HostmapInfrastructureRequest) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given HostmapStyle and assigns it to the Style field.
func (o *HostmapInfrastructureRequest) SetStyle(v HostmapStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HostmapInfrastructureRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Child != nil {
		toSerialize["child"] = o.Child
	}
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	toSerialize["enrichments"] = o.Enrichments
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	if o.NoGroupHosts != nil {
		toSerialize["no_group_hosts"] = o.NoGroupHosts
	}
	if o.NoMetricHosts != nil {
		toSerialize["no_metric_hosts"] = o.NoMetricHosts
	}
	toSerialize["node_type"] = o.NodeType
	toSerialize["request_type"] = o.RequestType
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HostmapInfrastructureRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Child              *HostmapInfrastructureRequest            `json:"child,omitempty"`
		ConditionalFormats []ThresholdFormattingRule                `json:"conditional_formats,omitempty"`
		Enrichments        *[]HostmapScalarFormulaRequest           `json:"enrichments"`
		Filter             *string                                  `json:"filter,omitempty"`
		GroupBy            []HostmapGroupBy                         `json:"group_by,omitempty"`
		NoGroupHosts       *bool                                    `json:"no_group_hosts,omitempty"`
		NoMetricHosts      *bool                                    `json:"no_metric_hosts,omitempty"`
		NodeType           *HostmapNodeType                         `json:"node_type"`
		RequestType        *HostmapInfrastructureRequestRequestType `json:"request_type"`
		Style              *HostmapStyle                            `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enrichments == nil {
		return fmt.Errorf("required field enrichments missing")
	}
	if all.NodeType == nil {
		return fmt.Errorf("required field node_type missing")
	}
	if all.RequestType == nil {
		return fmt.Errorf("required field request_type missing")
	}

	hasInvalidField := false
	if all.Child != nil && all.Child.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Child = all.Child
	o.ConditionalFormats = all.ConditionalFormats
	o.Enrichments = *all.Enrichments
	o.Filter = all.Filter
	o.GroupBy = all.GroupBy
	o.NoGroupHosts = all.NoGroupHosts
	o.NoMetricHosts = all.NoMetricHosts
	if !all.NodeType.IsValid() {
		hasInvalidField = true
	} else {
		o.NodeType = *all.NodeType
	}
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
