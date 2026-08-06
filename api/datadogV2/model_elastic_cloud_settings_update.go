// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudSettingsUpdate Partial Elastic Cloud interface settings for updates.
type ElasticCloudSettingsUpdate struct {
	// Enable to collect shard allocation metrics.
	CatAllocationStatsEnabled *bool `json:"cat_allocation_stats_enabled,omitempty"`
	// Enable to collect index-specific stats.
	DetailedIndexStatsEnabled *bool `json:"detailed_index_stats_enabled,omitempty"`
	// Enable to collect metrics about the indices in your cluster.
	IndexStatsEnabled *bool `json:"index_stats_enabled,omitempty"`
	// Enable to collect metrics about pending tasks.
	PendingTaskStatsEnabled *bool `json:"pending_task_stats_enabled,omitempty"`
	// Enable to collect all metrics even if primary shard metric collection times out.
	PshardGracefulToEnabled *bool `json:"pshard_graceful_to_enabled,omitempty"`
	// Enable to collect metrics over primary shards.
	PshardStatsEnabled *bool `json:"pshard_stats_enabled,omitempty"`
	// Enable to collect snapshot lifecycle management metrics.
	SlmStatsEnabled *bool `json:"slm_stats_enabled,omitempty"`
	// Custom tags for this deployment.
	Tags []string `json:"tags,omitempty"`
	// Deployment URL.
	Url *string `json:"url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticCloudSettingsUpdate instantiates a new ElasticCloudSettingsUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticCloudSettingsUpdate() *ElasticCloudSettingsUpdate {
	this := ElasticCloudSettingsUpdate{}
	return &this
}

// NewElasticCloudSettingsUpdateWithDefaults instantiates a new ElasticCloudSettingsUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticCloudSettingsUpdateWithDefaults() *ElasticCloudSettingsUpdate {
	this := ElasticCloudSettingsUpdate{}
	return &this
}

// GetCatAllocationStatsEnabled returns the CatAllocationStatsEnabled field value if set, zero value otherwise.
func (o *ElasticCloudSettingsUpdate) GetCatAllocationStatsEnabled() bool {
	if o == nil || o.CatAllocationStatsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CatAllocationStatsEnabled
}

// GetCatAllocationStatsEnabledOk returns a tuple with the CatAllocationStatsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudSettingsUpdate) GetCatAllocationStatsEnabledOk() (*bool, bool) {
	if o == nil || o.CatAllocationStatsEnabled == nil {
		return nil, false
	}
	return o.CatAllocationStatsEnabled, true
}

// HasCatAllocationStatsEnabled returns a boolean if a field has been set.
func (o *ElasticCloudSettingsUpdate) HasCatAllocationStatsEnabled() bool {
	return o != nil && o.CatAllocationStatsEnabled != nil
}

// SetCatAllocationStatsEnabled gets a reference to the given bool and assigns it to the CatAllocationStatsEnabled field.
func (o *ElasticCloudSettingsUpdate) SetCatAllocationStatsEnabled(v bool) {
	o.CatAllocationStatsEnabled = &v
}

// GetDetailedIndexStatsEnabled returns the DetailedIndexStatsEnabled field value if set, zero value otherwise.
func (o *ElasticCloudSettingsUpdate) GetDetailedIndexStatsEnabled() bool {
	if o == nil || o.DetailedIndexStatsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DetailedIndexStatsEnabled
}

// GetDetailedIndexStatsEnabledOk returns a tuple with the DetailedIndexStatsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudSettingsUpdate) GetDetailedIndexStatsEnabledOk() (*bool, bool) {
	if o == nil || o.DetailedIndexStatsEnabled == nil {
		return nil, false
	}
	return o.DetailedIndexStatsEnabled, true
}

// HasDetailedIndexStatsEnabled returns a boolean if a field has been set.
func (o *ElasticCloudSettingsUpdate) HasDetailedIndexStatsEnabled() bool {
	return o != nil && o.DetailedIndexStatsEnabled != nil
}

// SetDetailedIndexStatsEnabled gets a reference to the given bool and assigns it to the DetailedIndexStatsEnabled field.
func (o *ElasticCloudSettingsUpdate) SetDetailedIndexStatsEnabled(v bool) {
	o.DetailedIndexStatsEnabled = &v
}

// GetIndexStatsEnabled returns the IndexStatsEnabled field value if set, zero value otherwise.
func (o *ElasticCloudSettingsUpdate) GetIndexStatsEnabled() bool {
	if o == nil || o.IndexStatsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IndexStatsEnabled
}

// GetIndexStatsEnabledOk returns a tuple with the IndexStatsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudSettingsUpdate) GetIndexStatsEnabledOk() (*bool, bool) {
	if o == nil || o.IndexStatsEnabled == nil {
		return nil, false
	}
	return o.IndexStatsEnabled, true
}

// HasIndexStatsEnabled returns a boolean if a field has been set.
func (o *ElasticCloudSettingsUpdate) HasIndexStatsEnabled() bool {
	return o != nil && o.IndexStatsEnabled != nil
}

// SetIndexStatsEnabled gets a reference to the given bool and assigns it to the IndexStatsEnabled field.
func (o *ElasticCloudSettingsUpdate) SetIndexStatsEnabled(v bool) {
	o.IndexStatsEnabled = &v
}

// GetPendingTaskStatsEnabled returns the PendingTaskStatsEnabled field value if set, zero value otherwise.
func (o *ElasticCloudSettingsUpdate) GetPendingTaskStatsEnabled() bool {
	if o == nil || o.PendingTaskStatsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.PendingTaskStatsEnabled
}

// GetPendingTaskStatsEnabledOk returns a tuple with the PendingTaskStatsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudSettingsUpdate) GetPendingTaskStatsEnabledOk() (*bool, bool) {
	if o == nil || o.PendingTaskStatsEnabled == nil {
		return nil, false
	}
	return o.PendingTaskStatsEnabled, true
}

// HasPendingTaskStatsEnabled returns a boolean if a field has been set.
func (o *ElasticCloudSettingsUpdate) HasPendingTaskStatsEnabled() bool {
	return o != nil && o.PendingTaskStatsEnabled != nil
}

// SetPendingTaskStatsEnabled gets a reference to the given bool and assigns it to the PendingTaskStatsEnabled field.
func (o *ElasticCloudSettingsUpdate) SetPendingTaskStatsEnabled(v bool) {
	o.PendingTaskStatsEnabled = &v
}

// GetPshardGracefulToEnabled returns the PshardGracefulToEnabled field value if set, zero value otherwise.
func (o *ElasticCloudSettingsUpdate) GetPshardGracefulToEnabled() bool {
	if o == nil || o.PshardGracefulToEnabled == nil {
		var ret bool
		return ret
	}
	return *o.PshardGracefulToEnabled
}

// GetPshardGracefulToEnabledOk returns a tuple with the PshardGracefulToEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudSettingsUpdate) GetPshardGracefulToEnabledOk() (*bool, bool) {
	if o == nil || o.PshardGracefulToEnabled == nil {
		return nil, false
	}
	return o.PshardGracefulToEnabled, true
}

// HasPshardGracefulToEnabled returns a boolean if a field has been set.
func (o *ElasticCloudSettingsUpdate) HasPshardGracefulToEnabled() bool {
	return o != nil && o.PshardGracefulToEnabled != nil
}

// SetPshardGracefulToEnabled gets a reference to the given bool and assigns it to the PshardGracefulToEnabled field.
func (o *ElasticCloudSettingsUpdate) SetPshardGracefulToEnabled(v bool) {
	o.PshardGracefulToEnabled = &v
}

// GetPshardStatsEnabled returns the PshardStatsEnabled field value if set, zero value otherwise.
func (o *ElasticCloudSettingsUpdate) GetPshardStatsEnabled() bool {
	if o == nil || o.PshardStatsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.PshardStatsEnabled
}

// GetPshardStatsEnabledOk returns a tuple with the PshardStatsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudSettingsUpdate) GetPshardStatsEnabledOk() (*bool, bool) {
	if o == nil || o.PshardStatsEnabled == nil {
		return nil, false
	}
	return o.PshardStatsEnabled, true
}

// HasPshardStatsEnabled returns a boolean if a field has been set.
func (o *ElasticCloudSettingsUpdate) HasPshardStatsEnabled() bool {
	return o != nil && o.PshardStatsEnabled != nil
}

// SetPshardStatsEnabled gets a reference to the given bool and assigns it to the PshardStatsEnabled field.
func (o *ElasticCloudSettingsUpdate) SetPshardStatsEnabled(v bool) {
	o.PshardStatsEnabled = &v
}

// GetSlmStatsEnabled returns the SlmStatsEnabled field value if set, zero value otherwise.
func (o *ElasticCloudSettingsUpdate) GetSlmStatsEnabled() bool {
	if o == nil || o.SlmStatsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SlmStatsEnabled
}

// GetSlmStatsEnabledOk returns a tuple with the SlmStatsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudSettingsUpdate) GetSlmStatsEnabledOk() (*bool, bool) {
	if o == nil || o.SlmStatsEnabled == nil {
		return nil, false
	}
	return o.SlmStatsEnabled, true
}

// HasSlmStatsEnabled returns a boolean if a field has been set.
func (o *ElasticCloudSettingsUpdate) HasSlmStatsEnabled() bool {
	return o != nil && o.SlmStatsEnabled != nil
}

// SetSlmStatsEnabled gets a reference to the given bool and assigns it to the SlmStatsEnabled field.
func (o *ElasticCloudSettingsUpdate) SetSlmStatsEnabled(v bool) {
	o.SlmStatsEnabled = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ElasticCloudSettingsUpdate) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudSettingsUpdate) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ElasticCloudSettingsUpdate) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ElasticCloudSettingsUpdate) SetTags(v []string) {
	o.Tags = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ElasticCloudSettingsUpdate) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudSettingsUpdate) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ElasticCloudSettingsUpdate) HasUrl() bool {
	return o != nil && o.Url != nil
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ElasticCloudSettingsUpdate) SetUrl(v string) {
	o.Url = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticCloudSettingsUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CatAllocationStatsEnabled != nil {
		toSerialize["cat_allocation_stats_enabled"] = o.CatAllocationStatsEnabled
	}
	if o.DetailedIndexStatsEnabled != nil {
		toSerialize["detailed_index_stats_enabled"] = o.DetailedIndexStatsEnabled
	}
	if o.IndexStatsEnabled != nil {
		toSerialize["index_stats_enabled"] = o.IndexStatsEnabled
	}
	if o.PendingTaskStatsEnabled != nil {
		toSerialize["pending_task_stats_enabled"] = o.PendingTaskStatsEnabled
	}
	if o.PshardGracefulToEnabled != nil {
		toSerialize["pshard_graceful_to_enabled"] = o.PshardGracefulToEnabled
	}
	if o.PshardStatsEnabled != nil {
		toSerialize["pshard_stats_enabled"] = o.PshardStatsEnabled
	}
	if o.SlmStatsEnabled != nil {
		toSerialize["slm_stats_enabled"] = o.SlmStatsEnabled
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticCloudSettingsUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CatAllocationStatsEnabled *bool    `json:"cat_allocation_stats_enabled,omitempty"`
		DetailedIndexStatsEnabled *bool    `json:"detailed_index_stats_enabled,omitempty"`
		IndexStatsEnabled         *bool    `json:"index_stats_enabled,omitempty"`
		PendingTaskStatsEnabled   *bool    `json:"pending_task_stats_enabled,omitempty"`
		PshardGracefulToEnabled   *bool    `json:"pshard_graceful_to_enabled,omitempty"`
		PshardStatsEnabled        *bool    `json:"pshard_stats_enabled,omitempty"`
		SlmStatsEnabled           *bool    `json:"slm_stats_enabled,omitempty"`
		Tags                      []string `json:"tags,omitempty"`
		Url                       *string  `json:"url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cat_allocation_stats_enabled", "detailed_index_stats_enabled", "index_stats_enabled", "pending_task_stats_enabled", "pshard_graceful_to_enabled", "pshard_stats_enabled", "slm_stats_enabled", "tags", "url"})
	} else {
		return err
	}
	o.CatAllocationStatsEnabled = all.CatAllocationStatsEnabled
	o.DetailedIndexStatsEnabled = all.DetailedIndexStatsEnabled
	o.IndexStatsEnabled = all.IndexStatsEnabled
	o.PendingTaskStatsEnabled = all.PendingTaskStatsEnabled
	o.PshardGracefulToEnabled = all.PshardGracefulToEnabled
	o.PshardStatsEnabled = all.PshardStatsEnabled
	o.SlmStatsEnabled = all.SlmStatsEnabled
	o.Tags = all.Tags
	o.Url = all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
