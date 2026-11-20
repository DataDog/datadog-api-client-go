// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApmMetricsQueryFragment Test-only generated schema for the dashboard versioning proof of concept.
type ApmMetricsQueryFragment struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Alias *string `json:"alias,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ByocCluster *string `json:"byoc_cluster,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource ApmMetricsQueryFragmentDataSource `json:"data_source"`
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
	Service *string `json:"service,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SpanKind *SpanKind `json:"span_kind,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewApmMetricsQueryFragment instantiates a new ApmMetricsQueryFragment object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApmMetricsQueryFragment(dataSource ApmMetricsQueryFragmentDataSource, name string) *ApmMetricsQueryFragment {
	this := ApmMetricsQueryFragment{}
	this.DataSource = dataSource
	this.Name = name
	return &this
}

// NewApmMetricsQueryFragmentWithDefaults instantiates a new ApmMetricsQueryFragment object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApmMetricsQueryFragmentWithDefaults() *ApmMetricsQueryFragment {
	this := ApmMetricsQueryFragment{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *ApmMetricsQueryFragment) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmMetricsQueryFragment) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *ApmMetricsQueryFragment) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *ApmMetricsQueryFragment) SetAlias(v string) {
	o.Alias = &v
}

// GetByocCluster returns the ByocCluster field value if set, zero value otherwise.
func (o *ApmMetricsQueryFragment) GetByocCluster() string {
	if o == nil || o.ByocCluster == nil {
		var ret string
		return ret
	}
	return *o.ByocCluster
}

// GetByocClusterOk returns a tuple with the ByocCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmMetricsQueryFragment) GetByocClusterOk() (*string, bool) {
	if o == nil || o.ByocCluster == nil {
		return nil, false
	}
	return o.ByocCluster, true
}

// HasByocCluster returns a boolean if a field has been set.
func (o *ApmMetricsQueryFragment) HasByocCluster() bool {
	return o != nil && o.ByocCluster != nil
}

// SetByocCluster gets a reference to the given string and assigns it to the ByocCluster field.
func (o *ApmMetricsQueryFragment) SetByocCluster(v string) {
	o.ByocCluster = &v
}

// GetDataSource returns the DataSource field value.
func (o *ApmMetricsQueryFragment) GetDataSource() ApmMetricsQueryFragmentDataSource {
	if o == nil {
		var ret ApmMetricsQueryFragmentDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *ApmMetricsQueryFragment) GetDataSourceOk() (*ApmMetricsQueryFragmentDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *ApmMetricsQueryFragment) SetDataSource(v ApmMetricsQueryFragmentDataSource) {
	o.DataSource = v
}

// GetName returns the Name field value.
func (o *ApmMetricsQueryFragment) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApmMetricsQueryFragment) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ApmMetricsQueryFragment) SetName(v string) {
	o.Name = v
}

// GetOperationMode returns the OperationMode field value if set, zero value otherwise.
func (o *ApmMetricsQueryFragment) GetOperationMode() string {
	if o == nil || o.OperationMode == nil {
		var ret string
		return ret
	}
	return *o.OperationMode
}

// GetOperationModeOk returns a tuple with the OperationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmMetricsQueryFragment) GetOperationModeOk() (*string, bool) {
	if o == nil || o.OperationMode == nil {
		return nil, false
	}
	return o.OperationMode, true
}

// HasOperationMode returns a boolean if a field has been set.
func (o *ApmMetricsQueryFragment) HasOperationMode() bool {
	return o != nil && o.OperationMode != nil
}

// SetOperationMode gets a reference to the given string and assigns it to the OperationMode field.
func (o *ApmMetricsQueryFragment) SetOperationMode(v string) {
	o.OperationMode = &v
}

// GetOperationName returns the OperationName field value if set, zero value otherwise.
func (o *ApmMetricsQueryFragment) GetOperationName() string {
	if o == nil || o.OperationName == nil {
		var ret string
		return ret
	}
	return *o.OperationName
}

// GetOperationNameOk returns a tuple with the OperationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmMetricsQueryFragment) GetOperationNameOk() (*string, bool) {
	if o == nil || o.OperationName == nil {
		return nil, false
	}
	return o.OperationName, true
}

// HasOperationName returns a boolean if a field has been set.
func (o *ApmMetricsQueryFragment) HasOperationName() bool {
	return o != nil && o.OperationName != nil
}

// SetOperationName gets a reference to the given string and assigns it to the OperationName field.
func (o *ApmMetricsQueryFragment) SetOperationName(v string) {
	o.OperationName = &v
}

// GetPeerTags returns the PeerTags field value if set, zero value otherwise.
func (o *ApmMetricsQueryFragment) GetPeerTags() []string {
	if o == nil || o.PeerTags == nil {
		var ret []string
		return ret
	}
	return o.PeerTags
}

// GetPeerTagsOk returns a tuple with the PeerTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmMetricsQueryFragment) GetPeerTagsOk() (*[]string, bool) {
	if o == nil || o.PeerTags == nil {
		return nil, false
	}
	return &o.PeerTags, true
}

// HasPeerTags returns a boolean if a field has been set.
func (o *ApmMetricsQueryFragment) HasPeerTags() bool {
	return o != nil && o.PeerTags != nil
}

// SetPeerTags gets a reference to the given []string and assigns it to the PeerTags field.
func (o *ApmMetricsQueryFragment) SetPeerTags(v []string) {
	o.PeerTags = v
}

// GetQueryFilter returns the QueryFilter field value if set, zero value otherwise.
func (o *ApmMetricsQueryFragment) GetQueryFilter() string {
	if o == nil || o.QueryFilter == nil {
		var ret string
		return ret
	}
	return *o.QueryFilter
}

// GetQueryFilterOk returns a tuple with the QueryFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmMetricsQueryFragment) GetQueryFilterOk() (*string, bool) {
	if o == nil || o.QueryFilter == nil {
		return nil, false
	}
	return o.QueryFilter, true
}

// HasQueryFilter returns a boolean if a field has been set.
func (o *ApmMetricsQueryFragment) HasQueryFilter() bool {
	return o != nil && o.QueryFilter != nil
}

// SetQueryFilter gets a reference to the given string and assigns it to the QueryFilter field.
func (o *ApmMetricsQueryFragment) SetQueryFilter(v string) {
	o.QueryFilter = &v
}

// GetResourceHash returns the ResourceHash field value if set, zero value otherwise.
func (o *ApmMetricsQueryFragment) GetResourceHash() string {
	if o == nil || o.ResourceHash == nil {
		var ret string
		return ret
	}
	return *o.ResourceHash
}

// GetResourceHashOk returns a tuple with the ResourceHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmMetricsQueryFragment) GetResourceHashOk() (*string, bool) {
	if o == nil || o.ResourceHash == nil {
		return nil, false
	}
	return o.ResourceHash, true
}

// HasResourceHash returns a boolean if a field has been set.
func (o *ApmMetricsQueryFragment) HasResourceHash() bool {
	return o != nil && o.ResourceHash != nil
}

// SetResourceHash gets a reference to the given string and assigns it to the ResourceHash field.
func (o *ApmMetricsQueryFragment) SetResourceHash(v string) {
	o.ResourceHash = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *ApmMetricsQueryFragment) GetResourceName() string {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmMetricsQueryFragment) GetResourceNameOk() (*string, bool) {
	if o == nil || o.ResourceName == nil {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *ApmMetricsQueryFragment) HasResourceName() bool {
	return o != nil && o.ResourceName != nil
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *ApmMetricsQueryFragment) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ApmMetricsQueryFragment) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmMetricsQueryFragment) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ApmMetricsQueryFragment) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *ApmMetricsQueryFragment) SetService(v string) {
	o.Service = &v
}

// GetSpanKind returns the SpanKind field value if set, zero value otherwise.
func (o *ApmMetricsQueryFragment) GetSpanKind() SpanKind {
	if o == nil || o.SpanKind == nil {
		var ret SpanKind
		return ret
	}
	return *o.SpanKind
}

// GetSpanKindOk returns a tuple with the SpanKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmMetricsQueryFragment) GetSpanKindOk() (*SpanKind, bool) {
	if o == nil || o.SpanKind == nil {
		return nil, false
	}
	return o.SpanKind, true
}

// HasSpanKind returns a boolean if a field has been set.
func (o *ApmMetricsQueryFragment) HasSpanKind() bool {
	return o != nil && o.SpanKind != nil
}

// SetSpanKind gets a reference to the given SpanKind and assigns it to the SpanKind field.
func (o *ApmMetricsQueryFragment) SetSpanKind(v SpanKind) {
	o.SpanKind = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApmMetricsQueryFragment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	if o.ByocCluster != nil {
		toSerialize["byoc_cluster"] = o.ByocCluster
	}
	toSerialize["data_source"] = o.DataSource
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
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.SpanKind != nil {
		toSerialize["span_kind"] = o.SpanKind
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApmMetricsQueryFragment) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias         *string                            `json:"alias,omitempty"`
		ByocCluster   *string                            `json:"byoc_cluster,omitempty"`
		DataSource    *ApmMetricsQueryFragmentDataSource `json:"data_source"`
		Name          *string                            `json:"name"`
		OperationMode *string                            `json:"operation_mode,omitempty"`
		OperationName *string                            `json:"operation_name,omitempty"`
		PeerTags      []string                           `json:"peer_tags,omitempty"`
		QueryFilter   *string                            `json:"query_filter,omitempty"`
		ResourceHash  *string                            `json:"resource_hash,omitempty"`
		ResourceName  *string                            `json:"resource_name,omitempty"`
		Service       *string                            `json:"service,omitempty"`
		SpanKind      *SpanKind                          `json:"span_kind,omitempty"`
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

	hasInvalidField := false
	o.Alias = all.Alias
	o.ByocCluster = all.ByocCluster
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Name = *all.Name
	o.OperationMode = all.OperationMode
	o.OperationName = all.OperationName
	o.PeerTags = all.PeerTags
	o.QueryFilter = all.QueryFilter
	o.ResourceHash = all.ResourceHash
	o.ResourceName = all.ResourceName
	o.Service = all.Service
	if all.SpanKind != nil && !all.SpanKind.IsValid() {
		hasInvalidField = true
	} else {
		o.SpanKind = all.SpanKind
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
