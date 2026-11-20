// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaApmMetricsQuery Test-only generated schema for the dashboard versioning proof of concept.
type FormulaApmMetricsQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ByocCluster *string `json:"byoc_cluster,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource FormulaApmMetricsQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy []string `json:"group_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	OperationMode *string `json:"operation_mode,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	OperationName *string `json:"operation_name,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	PeerTags []string `json:"peer_tags,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryFilter *string `json:"query_filter,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResourceHash *string `json:"resource_hash,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResourceName *string `json:"resource_name,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SemanticMode *FormulaSemanticMode `json:"semantic_mode,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Service *string `json:"service,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SpanKind *SpanKind `json:"span_kind,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Stat ApmMetricStat `json:"stat"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaApmMetricsQuery instantiates a new FormulaApmMetricsQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaApmMetricsQuery(dataSource FormulaApmMetricsQueryDataSource, name string, stat ApmMetricStat) *FormulaApmMetricsQuery {
	this := FormulaApmMetricsQuery{}
	this.DataSource = dataSource
	this.Name = name
	this.Stat = stat
	return &this
}

// NewFormulaApmMetricsQueryWithDefaults instantiates a new FormulaApmMetricsQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaApmMetricsQueryWithDefaults() *FormulaApmMetricsQuery {
	this := FormulaApmMetricsQuery{}
	return &this
}

// GetByocCluster returns the ByocCluster field value if set, zero value otherwise.
func (o *FormulaApmMetricsQuery) GetByocCluster() string {
	if o == nil || o.ByocCluster == nil {
		var ret string
		return ret
	}
	return *o.ByocCluster
}

// GetByocClusterOk returns a tuple with the ByocCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaApmMetricsQuery) GetByocClusterOk() (*string, bool) {
	if o == nil || o.ByocCluster == nil {
		return nil, false
	}
	return o.ByocCluster, true
}

// HasByocCluster returns a boolean if a field has been set.
func (o *FormulaApmMetricsQuery) HasByocCluster() bool {
	return o != nil && o.ByocCluster != nil
}

// SetByocCluster gets a reference to the given string and assigns it to the ByocCluster field.
func (o *FormulaApmMetricsQuery) SetByocCluster(v string) {
	o.ByocCluster = &v
}

// GetDataSource returns the DataSource field value.
func (o *FormulaApmMetricsQuery) GetDataSource() FormulaApmMetricsQueryDataSource {
	if o == nil {
		var ret FormulaApmMetricsQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaApmMetricsQuery) GetDataSourceOk() (*FormulaApmMetricsQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaApmMetricsQuery) SetDataSource(v FormulaApmMetricsQueryDataSource) {
	o.DataSource = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *FormulaApmMetricsQuery) GetGroupBy() []string {
	if o == nil || o.GroupBy == nil {
		var ret []string
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaApmMetricsQuery) GetGroupByOk() (*[]string, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *FormulaApmMetricsQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []string and assigns it to the GroupBy field.
func (o *FormulaApmMetricsQuery) SetGroupBy(v []string) {
	o.GroupBy = v
}

// GetName returns the Name field value.
func (o *FormulaApmMetricsQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaApmMetricsQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaApmMetricsQuery) SetName(v string) {
	o.Name = v
}

// GetOperationMode returns the OperationMode field value if set, zero value otherwise.
func (o *FormulaApmMetricsQuery) GetOperationMode() string {
	if o == nil || o.OperationMode == nil {
		var ret string
		return ret
	}
	return *o.OperationMode
}

// GetOperationModeOk returns a tuple with the OperationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaApmMetricsQuery) GetOperationModeOk() (*string, bool) {
	if o == nil || o.OperationMode == nil {
		return nil, false
	}
	return o.OperationMode, true
}

// HasOperationMode returns a boolean if a field has been set.
func (o *FormulaApmMetricsQuery) HasOperationMode() bool {
	return o != nil && o.OperationMode != nil
}

// SetOperationMode gets a reference to the given string and assigns it to the OperationMode field.
func (o *FormulaApmMetricsQuery) SetOperationMode(v string) {
	o.OperationMode = &v
}

// GetOperationName returns the OperationName field value if set, zero value otherwise.
func (o *FormulaApmMetricsQuery) GetOperationName() string {
	if o == nil || o.OperationName == nil {
		var ret string
		return ret
	}
	return *o.OperationName
}

// GetOperationNameOk returns a tuple with the OperationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaApmMetricsQuery) GetOperationNameOk() (*string, bool) {
	if o == nil || o.OperationName == nil {
		return nil, false
	}
	return o.OperationName, true
}

// HasOperationName returns a boolean if a field has been set.
func (o *FormulaApmMetricsQuery) HasOperationName() bool {
	return o != nil && o.OperationName != nil
}

// SetOperationName gets a reference to the given string and assigns it to the OperationName field.
func (o *FormulaApmMetricsQuery) SetOperationName(v string) {
	o.OperationName = &v
}

// GetPeerTags returns the PeerTags field value if set, zero value otherwise.
func (o *FormulaApmMetricsQuery) GetPeerTags() []string {
	if o == nil || o.PeerTags == nil {
		var ret []string
		return ret
	}
	return o.PeerTags
}

// GetPeerTagsOk returns a tuple with the PeerTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaApmMetricsQuery) GetPeerTagsOk() (*[]string, bool) {
	if o == nil || o.PeerTags == nil {
		return nil, false
	}
	return &o.PeerTags, true
}

// HasPeerTags returns a boolean if a field has been set.
func (o *FormulaApmMetricsQuery) HasPeerTags() bool {
	return o != nil && o.PeerTags != nil
}

// SetPeerTags gets a reference to the given []string and assigns it to the PeerTags field.
func (o *FormulaApmMetricsQuery) SetPeerTags(v []string) {
	o.PeerTags = v
}

// GetQueryFilter returns the QueryFilter field value if set, zero value otherwise.
func (o *FormulaApmMetricsQuery) GetQueryFilter() string {
	if o == nil || o.QueryFilter == nil {
		var ret string
		return ret
	}
	return *o.QueryFilter
}

// GetQueryFilterOk returns a tuple with the QueryFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaApmMetricsQuery) GetQueryFilterOk() (*string, bool) {
	if o == nil || o.QueryFilter == nil {
		return nil, false
	}
	return o.QueryFilter, true
}

// HasQueryFilter returns a boolean if a field has been set.
func (o *FormulaApmMetricsQuery) HasQueryFilter() bool {
	return o != nil && o.QueryFilter != nil
}

// SetQueryFilter gets a reference to the given string and assigns it to the QueryFilter field.
func (o *FormulaApmMetricsQuery) SetQueryFilter(v string) {
	o.QueryFilter = &v
}

// GetResourceHash returns the ResourceHash field value if set, zero value otherwise.
func (o *FormulaApmMetricsQuery) GetResourceHash() string {
	if o == nil || o.ResourceHash == nil {
		var ret string
		return ret
	}
	return *o.ResourceHash
}

// GetResourceHashOk returns a tuple with the ResourceHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaApmMetricsQuery) GetResourceHashOk() (*string, bool) {
	if o == nil || o.ResourceHash == nil {
		return nil, false
	}
	return o.ResourceHash, true
}

// HasResourceHash returns a boolean if a field has been set.
func (o *FormulaApmMetricsQuery) HasResourceHash() bool {
	return o != nil && o.ResourceHash != nil
}

// SetResourceHash gets a reference to the given string and assigns it to the ResourceHash field.
func (o *FormulaApmMetricsQuery) SetResourceHash(v string) {
	o.ResourceHash = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *FormulaApmMetricsQuery) GetResourceName() string {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaApmMetricsQuery) GetResourceNameOk() (*string, bool) {
	if o == nil || o.ResourceName == nil {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *FormulaApmMetricsQuery) HasResourceName() bool {
	return o != nil && o.ResourceName != nil
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *FormulaApmMetricsQuery) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetSemanticMode returns the SemanticMode field value if set, zero value otherwise.
func (o *FormulaApmMetricsQuery) GetSemanticMode() FormulaSemanticMode {
	if o == nil || o.SemanticMode == nil {
		var ret FormulaSemanticMode
		return ret
	}
	return *o.SemanticMode
}

// GetSemanticModeOk returns a tuple with the SemanticMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaApmMetricsQuery) GetSemanticModeOk() (*FormulaSemanticMode, bool) {
	if o == nil || o.SemanticMode == nil {
		return nil, false
	}
	return o.SemanticMode, true
}

// HasSemanticMode returns a boolean if a field has been set.
func (o *FormulaApmMetricsQuery) HasSemanticMode() bool {
	return o != nil && o.SemanticMode != nil
}

// SetSemanticMode gets a reference to the given FormulaSemanticMode and assigns it to the SemanticMode field.
func (o *FormulaApmMetricsQuery) SetSemanticMode(v FormulaSemanticMode) {
	o.SemanticMode = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *FormulaApmMetricsQuery) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaApmMetricsQuery) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *FormulaApmMetricsQuery) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *FormulaApmMetricsQuery) SetService(v string) {
	o.Service = &v
}

// GetSpanKind returns the SpanKind field value if set, zero value otherwise.
func (o *FormulaApmMetricsQuery) GetSpanKind() SpanKind {
	if o == nil || o.SpanKind == nil {
		var ret SpanKind
		return ret
	}
	return *o.SpanKind
}

// GetSpanKindOk returns a tuple with the SpanKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaApmMetricsQuery) GetSpanKindOk() (*SpanKind, bool) {
	if o == nil || o.SpanKind == nil {
		return nil, false
	}
	return o.SpanKind, true
}

// HasSpanKind returns a boolean if a field has been set.
func (o *FormulaApmMetricsQuery) HasSpanKind() bool {
	return o != nil && o.SpanKind != nil
}

// SetSpanKind gets a reference to the given SpanKind and assigns it to the SpanKind field.
func (o *FormulaApmMetricsQuery) SetSpanKind(v SpanKind) {
	o.SpanKind = &v
}

// GetStat returns the Stat field value.
func (o *FormulaApmMetricsQuery) GetStat() ApmMetricStat {
	if o == nil {
		var ret ApmMetricStat
		return ret
	}
	return o.Stat
}

// GetStatOk returns a tuple with the Stat field value
// and a boolean to check if the value has been set.
func (o *FormulaApmMetricsQuery) GetStatOk() (*ApmMetricStat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stat, true
}

// SetStat sets field value.
func (o *FormulaApmMetricsQuery) SetStat(v ApmMetricStat) {
	o.Stat = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaApmMetricsQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ByocCluster != nil {
		toSerialize["byoc_cluster"] = o.ByocCluster
	}
	toSerialize["data_source"] = o.DataSource
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	toSerialize["name"] = o.Name
	if o.OperationMode != nil {
		toSerialize["operation_mode"] = o.OperationMode
	}
	if o.OperationName != nil {
		toSerialize["operation_name"] = o.OperationName
	}
	if o.PeerTags != nil {
		toSerialize["peer_tags"] = o.PeerTags
	}
	if o.QueryFilter != nil {
		toSerialize["query_filter"] = o.QueryFilter
	}
	if o.ResourceHash != nil {
		toSerialize["resource_hash"] = o.ResourceHash
	}
	if o.ResourceName != nil {
		toSerialize["resource_name"] = o.ResourceName
	}
	if o.SemanticMode != nil {
		toSerialize["semantic_mode"] = o.SemanticMode
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.SpanKind != nil {
		toSerialize["span_kind"] = o.SpanKind
	}
	toSerialize["stat"] = o.Stat
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaApmMetricsQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ByocCluster   *string                           `json:"byoc_cluster,omitempty"`
		DataSource    *FormulaApmMetricsQueryDataSource `json:"data_source"`
		GroupBy       []string                          `json:"group_by,omitempty"`
		Name          *string                           `json:"name"`
		OperationMode *string                           `json:"operation_mode,omitempty"`
		OperationName *string                           `json:"operation_name,omitempty"`
		PeerTags      []string                          `json:"peer_tags,omitempty"`
		QueryFilter   *string                           `json:"query_filter,omitempty"`
		ResourceHash  *string                           `json:"resource_hash,omitempty"`
		ResourceName  *string                           `json:"resource_name,omitempty"`
		SemanticMode  *FormulaSemanticMode              `json:"semantic_mode,omitempty"`
		Service       *string                           `json:"service,omitempty"`
		SpanKind      *SpanKind                         `json:"span_kind,omitempty"`
		Stat          *ApmMetricStat                    `json:"stat"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Stat == nil {
		return fmt.Errorf("required field stat missing")
	}

	hasInvalidField := false
	o.ByocCluster = all.ByocCluster
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.GroupBy = all.GroupBy
	o.Name = *all.Name
	o.OperationMode = all.OperationMode
	o.OperationName = all.OperationName
	o.PeerTags = all.PeerTags
	o.QueryFilter = all.QueryFilter
	o.ResourceHash = all.ResourceHash
	o.ResourceName = all.ResourceName
	if all.SemanticMode != nil && !all.SemanticMode.IsValid() {
		hasInvalidField = true
	} else {
		o.SemanticMode = all.SemanticMode
	}
	o.Service = all.Service
	if all.SpanKind != nil && !all.SpanKind.IsValid() {
		hasInvalidField = true
	} else {
		o.SpanKind = all.SpanKind
	}
	if !all.Stat.IsValid() {
		hasInvalidField = true
	} else {
		o.Stat = *all.Stat
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
