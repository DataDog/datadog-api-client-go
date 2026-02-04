// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardSearchAggregations Aggregations of dashboard search results.
type DashboardSearchAggregations struct {
	// Aggregation by author.
	Author []DashboardSearchAggregationBucketMultiKey `json:"author,omitempty"`
	// Aggregation by share status.
	IsShared []DashboardSearchAggregationBucketKey `json:"is_shared,omitempty"`
	// Aggregation by share type.
	ShareType []DashboardSearchAggregationBucketKey `json:"share_type,omitempty"`
	// Aggregation by who shared the dashboard.
	SharedByHandle []DashboardSearchAggregationBucketKey `json:"shared_by.handle,omitempty"`
	// Aggregation by status.
	Status []DashboardSearchAggregationBucketKey `json:"status,omitempty"`
	// Aggregation by tags.
	Tags []DashboardSearchAggregationBucketKey `json:"tags,omitempty"`
	// Aggregation by template variable names.
	TemplateVariablesName []DashboardSearchAggregationBucketKey `json:"template_variables.name,omitempty"`
	// Aggregation by dashboard type.
	Type []DashboardSearchAggregationBucketKey `json:"type,omitempty"`
	// Aggregation by widget metrics.
	WidgetsMetrics []DashboardSearchAggregationBucketKey `json:"widgets.metrics,omitempty"`
	// Aggregation by widget type.
	WidgetsType []DashboardSearchAggregationBucketKey `json:"widgets.type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDashboardSearchAggregations instantiates a new DashboardSearchAggregations object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardSearchAggregations() *DashboardSearchAggregations {
	this := DashboardSearchAggregations{}
	return &this
}

// NewDashboardSearchAggregationsWithDefaults instantiates a new DashboardSearchAggregations object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardSearchAggregationsWithDefaults() *DashboardSearchAggregations {
	this := DashboardSearchAggregations{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *DashboardSearchAggregations) GetAuthor() []DashboardSearchAggregationBucketMultiKey {
	if o == nil || o.Author == nil {
		var ret []DashboardSearchAggregationBucketMultiKey
		return ret
	}
	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSearchAggregations) GetAuthorOk() (*[]DashboardSearchAggregationBucketMultiKey, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return &o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *DashboardSearchAggregations) HasAuthor() bool {
	return o != nil && o.Author != nil
}

// SetAuthor gets a reference to the given []DashboardSearchAggregationBucketMultiKey and assigns it to the Author field.
func (o *DashboardSearchAggregations) SetAuthor(v []DashboardSearchAggregationBucketMultiKey) {
	o.Author = v
}

// GetIsShared returns the IsShared field value if set, zero value otherwise.
func (o *DashboardSearchAggregations) GetIsShared() []DashboardSearchAggregationBucketKey {
	if o == nil || o.IsShared == nil {
		var ret []DashboardSearchAggregationBucketKey
		return ret
	}
	return o.IsShared
}

// GetIsSharedOk returns a tuple with the IsShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSearchAggregations) GetIsSharedOk() (*[]DashboardSearchAggregationBucketKey, bool) {
	if o == nil || o.IsShared == nil {
		return nil, false
	}
	return &o.IsShared, true
}

// HasIsShared returns a boolean if a field has been set.
func (o *DashboardSearchAggregations) HasIsShared() bool {
	return o != nil && o.IsShared != nil
}

// SetIsShared gets a reference to the given []DashboardSearchAggregationBucketKey and assigns it to the IsShared field.
func (o *DashboardSearchAggregations) SetIsShared(v []DashboardSearchAggregationBucketKey) {
	o.IsShared = v
}

// GetShareType returns the ShareType field value if set, zero value otherwise.
func (o *DashboardSearchAggregations) GetShareType() []DashboardSearchAggregationBucketKey {
	if o == nil || o.ShareType == nil {
		var ret []DashboardSearchAggregationBucketKey
		return ret
	}
	return o.ShareType
}

// GetShareTypeOk returns a tuple with the ShareType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSearchAggregations) GetShareTypeOk() (*[]DashboardSearchAggregationBucketKey, bool) {
	if o == nil || o.ShareType == nil {
		return nil, false
	}
	return &o.ShareType, true
}

// HasShareType returns a boolean if a field has been set.
func (o *DashboardSearchAggregations) HasShareType() bool {
	return o != nil && o.ShareType != nil
}

// SetShareType gets a reference to the given []DashboardSearchAggregationBucketKey and assigns it to the ShareType field.
func (o *DashboardSearchAggregations) SetShareType(v []DashboardSearchAggregationBucketKey) {
	o.ShareType = v
}

// GetSharedByHandle returns the SharedByHandle field value if set, zero value otherwise.
func (o *DashboardSearchAggregations) GetSharedByHandle() []DashboardSearchAggregationBucketKey {
	if o == nil || o.SharedByHandle == nil {
		var ret []DashboardSearchAggregationBucketKey
		return ret
	}
	return o.SharedByHandle
}

// GetSharedByHandleOk returns a tuple with the SharedByHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSearchAggregations) GetSharedByHandleOk() (*[]DashboardSearchAggregationBucketKey, bool) {
	if o == nil || o.SharedByHandle == nil {
		return nil, false
	}
	return &o.SharedByHandle, true
}

// HasSharedByHandle returns a boolean if a field has been set.
func (o *DashboardSearchAggregations) HasSharedByHandle() bool {
	return o != nil && o.SharedByHandle != nil
}

// SetSharedByHandle gets a reference to the given []DashboardSearchAggregationBucketKey and assigns it to the SharedByHandle field.
func (o *DashboardSearchAggregations) SetSharedByHandle(v []DashboardSearchAggregationBucketKey) {
	o.SharedByHandle = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DashboardSearchAggregations) GetStatus() []DashboardSearchAggregationBucketKey {
	if o == nil || o.Status == nil {
		var ret []DashboardSearchAggregationBucketKey
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSearchAggregations) GetStatusOk() (*[]DashboardSearchAggregationBucketKey, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DashboardSearchAggregations) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given []DashboardSearchAggregationBucketKey and assigns it to the Status field.
func (o *DashboardSearchAggregations) SetStatus(v []DashboardSearchAggregationBucketKey) {
	o.Status = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DashboardSearchAggregations) GetTags() []DashboardSearchAggregationBucketKey {
	if o == nil || o.Tags == nil {
		var ret []DashboardSearchAggregationBucketKey
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSearchAggregations) GetTagsOk() (*[]DashboardSearchAggregationBucketKey, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DashboardSearchAggregations) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []DashboardSearchAggregationBucketKey and assigns it to the Tags field.
func (o *DashboardSearchAggregations) SetTags(v []DashboardSearchAggregationBucketKey) {
	o.Tags = v
}

// GetTemplateVariablesName returns the TemplateVariablesName field value if set, zero value otherwise.
func (o *DashboardSearchAggregations) GetTemplateVariablesName() []DashboardSearchAggregationBucketKey {
	if o == nil || o.TemplateVariablesName == nil {
		var ret []DashboardSearchAggregationBucketKey
		return ret
	}
	return o.TemplateVariablesName
}

// GetTemplateVariablesNameOk returns a tuple with the TemplateVariablesName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSearchAggregations) GetTemplateVariablesNameOk() (*[]DashboardSearchAggregationBucketKey, bool) {
	if o == nil || o.TemplateVariablesName == nil {
		return nil, false
	}
	return &o.TemplateVariablesName, true
}

// HasTemplateVariablesName returns a boolean if a field has been set.
func (o *DashboardSearchAggregations) HasTemplateVariablesName() bool {
	return o != nil && o.TemplateVariablesName != nil
}

// SetTemplateVariablesName gets a reference to the given []DashboardSearchAggregationBucketKey and assigns it to the TemplateVariablesName field.
func (o *DashboardSearchAggregations) SetTemplateVariablesName(v []DashboardSearchAggregationBucketKey) {
	o.TemplateVariablesName = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DashboardSearchAggregations) GetType() []DashboardSearchAggregationBucketKey {
	if o == nil || o.Type == nil {
		var ret []DashboardSearchAggregationBucketKey
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSearchAggregations) GetTypeOk() (*[]DashboardSearchAggregationBucketKey, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DashboardSearchAggregations) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given []DashboardSearchAggregationBucketKey and assigns it to the Type field.
func (o *DashboardSearchAggregations) SetType(v []DashboardSearchAggregationBucketKey) {
	o.Type = v
}

// GetWidgetsMetrics returns the WidgetsMetrics field value if set, zero value otherwise.
func (o *DashboardSearchAggregations) GetWidgetsMetrics() []DashboardSearchAggregationBucketKey {
	if o == nil || o.WidgetsMetrics == nil {
		var ret []DashboardSearchAggregationBucketKey
		return ret
	}
	return o.WidgetsMetrics
}

// GetWidgetsMetricsOk returns a tuple with the WidgetsMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSearchAggregations) GetWidgetsMetricsOk() (*[]DashboardSearchAggregationBucketKey, bool) {
	if o == nil || o.WidgetsMetrics == nil {
		return nil, false
	}
	return &o.WidgetsMetrics, true
}

// HasWidgetsMetrics returns a boolean if a field has been set.
func (o *DashboardSearchAggregations) HasWidgetsMetrics() bool {
	return o != nil && o.WidgetsMetrics != nil
}

// SetWidgetsMetrics gets a reference to the given []DashboardSearchAggregationBucketKey and assigns it to the WidgetsMetrics field.
func (o *DashboardSearchAggregations) SetWidgetsMetrics(v []DashboardSearchAggregationBucketKey) {
	o.WidgetsMetrics = v
}

// GetWidgetsType returns the WidgetsType field value if set, zero value otherwise.
func (o *DashboardSearchAggregations) GetWidgetsType() []DashboardSearchAggregationBucketKey {
	if o == nil || o.WidgetsType == nil {
		var ret []DashboardSearchAggregationBucketKey
		return ret
	}
	return o.WidgetsType
}

// GetWidgetsTypeOk returns a tuple with the WidgetsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSearchAggregations) GetWidgetsTypeOk() (*[]DashboardSearchAggregationBucketKey, bool) {
	if o == nil || o.WidgetsType == nil {
		return nil, false
	}
	return &o.WidgetsType, true
}

// HasWidgetsType returns a boolean if a field has been set.
func (o *DashboardSearchAggregations) HasWidgetsType() bool {
	return o != nil && o.WidgetsType != nil
}

// SetWidgetsType gets a reference to the given []DashboardSearchAggregationBucketKey and assigns it to the WidgetsType field.
func (o *DashboardSearchAggregations) SetWidgetsType(v []DashboardSearchAggregationBucketKey) {
	o.WidgetsType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardSearchAggregations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Author != nil {
		toSerialize["author"] = o.Author
	}
	if o.IsShared != nil {
		toSerialize["is_shared"] = o.IsShared
	}
	if o.ShareType != nil {
		toSerialize["share_type"] = o.ShareType
	}
	if o.SharedByHandle != nil {
		toSerialize["shared_by.handle"] = o.SharedByHandle
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TemplateVariablesName != nil {
		toSerialize["template_variables.name"] = o.TemplateVariablesName
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.WidgetsMetrics != nil {
		toSerialize["widgets.metrics"] = o.WidgetsMetrics
	}
	if o.WidgetsType != nil {
		toSerialize["widgets.type"] = o.WidgetsType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardSearchAggregations) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Author                []DashboardSearchAggregationBucketMultiKey `json:"author,omitempty"`
		IsShared              []DashboardSearchAggregationBucketKey      `json:"is_shared,omitempty"`
		ShareType             []DashboardSearchAggregationBucketKey      `json:"share_type,omitempty"`
		SharedByHandle        []DashboardSearchAggregationBucketKey      `json:"shared_by.handle,omitempty"`
		Status                []DashboardSearchAggregationBucketKey      `json:"status,omitempty"`
		Tags                  []DashboardSearchAggregationBucketKey      `json:"tags,omitempty"`
		TemplateVariablesName []DashboardSearchAggregationBucketKey      `json:"template_variables.name,omitempty"`
		Type                  []DashboardSearchAggregationBucketKey      `json:"type,omitempty"`
		WidgetsMetrics        []DashboardSearchAggregationBucketKey      `json:"widgets.metrics,omitempty"`
		WidgetsType           []DashboardSearchAggregationBucketKey      `json:"widgets.type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"author", "is_shared", "share_type", "shared_by.handle", "status", "tags", "template_variables.name", "type", "widgets.metrics", "widgets.type"})
	} else {
		return err
	}
	o.Author = all.Author
	o.IsShared = all.IsShared
	o.ShareType = all.ShareType
	o.SharedByHandle = all.SharedByHandle
	o.Status = all.Status
	o.Tags = all.Tags
	o.TemplateVariablesName = all.TemplateVariablesName
	o.Type = all.Type
	o.WidgetsMetrics = all.WidgetsMetrics
	o.WidgetsType = all.WidgetsType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
